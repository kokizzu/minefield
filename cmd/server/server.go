package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	connectcors "connectrpc.com/cors"
	service "github.com/bitbomdev/minefield/api/v1"
	"github.com/bitbomdev/minefield/gen/api/v1/apiv1connect"
	"github.com/bitbomdev/minefield/pkg/graph"
	"github.com/bitbomdev/minefield/pkg/storages"
	"github.com/rs/cors"
	"github.com/spf13/cobra"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type options struct {
	storage     graph.Storage
	concurrency int32
	addr        string
	StorageType string
	StorageAddr string
	StoragePath string
	UseInMemory bool
	CORS        []string
}

const (
	defaultConcurrency = 10
	defaultAddr        = "localhost:8089"
	redisStorageType   = "redis"
	sqliteStorageType  = "sqlite"
)

func (o *options) AddFlags(cmd *cobra.Command) {
	cmd.Flags().Int32Var(&o.concurrency, "concurrency", defaultConcurrency, "Maximum number of concurrent operations for leaderboard operations")
	cmd.Flags().StringVar(&o.addr, "addr", defaultAddr, "Network address and port for the server (e.g. localhost:8089)")
	cmd.Flags().StringVar(&o.StorageType, "storage-type", sqliteStorageType, "Type of storage to use (e.g., redis, sqlite)")
	cmd.Flags().StringVar(&o.StorageAddr, "storage-addr", "localhost:6379", "Address for redis storage backend")
	cmd.Flags().StringVar(&o.StoragePath, "storage-path", "", "Path to the SQLite database file")
	cmd.Flags().BoolVar(&o.UseInMemory, "use-in-memory", true, "Use in-memory SQLite database")
	cmd.Flags().StringSliceVar(
		&o.CORS,
		"cors",
		[]string{"http://localhost:8089"},
		"Allowed origins for CORS (e.g., 'https://app.bitbom.dev')",
	)
}

func (o *options) ProvideStorage() (graph.Storage, error) {
	switch o.StorageType {
	case redisStorageType:
		return storages.NewRedisStorage(o.StorageAddr)
	case sqliteStorageType:
		return storages.NewSQLStorage(o.StoragePath, o.UseInMemory)
	default:
		return nil, fmt.Errorf("unknown storage type: %s", o.StorageType)
	}
}

func (o *options) Run(_ *cobra.Command, _ []string) error {
	var err error
	o.storage, err = o.ProvideStorage()
	if err != nil {
		return fmt.Errorf("failed to initialize storage: %w", err)
	}

	server, err := o.setupServer()
	if err != nil {
		return fmt.Errorf("failed to setup server: %w", err)
	}
	return o.startServer(server)
}

func (o *options) PersistentPreRunE(_ *cobra.Command, _ []string) error {
	if o.StorageType != redisStorageType && o.StorageType != sqliteStorageType {
		return fmt.Errorf("invalid storage-type %q: must be one of [redis, sqlite]", o.StorageType)
	}

	if o.StorageType == sqliteStorageType && o.StoragePath == "" {
		if !o.UseInMemory {
			return fmt.Errorf("storage-path is required when using SQLite with file-based storage")
		}
	}

	if o.StorageType == redisStorageType && o.StorageAddr == "" {
		return fmt.Errorf("storage-addr is required when using Redis (format: host:port)")
	}

	return nil
}

func (o *options) setupServer() (*http.Server, error) {
	if o.concurrency <= 0 {
		return nil, fmt.Errorf("concurrency must be greater than zero")
	}

	serviceAddr := o.addr
	if serviceAddr == "" {
		serviceAddr = defaultAddr
	}

	newService := service.NewService(o.storage, o.concurrency)
	mux := http.NewServeMux()
	path, handler := apiv1connect.NewQueryServiceHandler(newService)
	mux.Handle(path, handler)
	path, handler = apiv1connect.NewLeaderboardServiceHandler(newService)
	mux.Handle(path, handler)
	path, handler = apiv1connect.NewCacheServiceHandler(newService)
	mux.Handle(path, handler)
	path, handler = apiv1connect.NewGraphServiceHandler(newService)
	mux.Handle(path, handler)
	path, handler = apiv1connect.NewHealthServiceHandler(newService)
	mux.Handle(path, handler)
	path, handler = apiv1connect.NewIngestServiceHandler(newService)
	mux.Handle(path, handler)

	server := &http.Server{
		Addr:    serviceAddr,
		Handler: h2c.NewHandler(withCORS(mux, o), &http2.Server{}),
	}

	return server, nil
}

func (o *options) startServer(server *http.Server) error {
	// Handle graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Printf("Server is starting on %s\n", server.Addr)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("ListenAndServe(): %s\n", err)
		}
	}()

	<-stop
	log.Println("Shutting down the server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		return fmt.Errorf("server shutdown failed: %w", err)
	}

	log.Println("Server gracefully stopped")
	return nil
}

// New returns a new cobra command for the server.
func New() *cobra.Command {
	o := &options{}
	cmd := &cobra.Command{
		Use:               "server",
		Short:             "Start the minefield server for graph operations and queries",
		Args:              cobra.ExactArgs(0),
		PersistentPreRunE: o.PersistentPreRunE,
		RunE:              o.Run,
		DisableAutoGenTag: true,
	}
	o.AddFlags(cmd)
	return cmd
}

func NewServerCommand(storage graph.Storage, o *options) (*cobra.Command, error) {
	o.storage = storage
	cmd := &cobra.Command{
		Use:               "server",
		Short:             "Start the minefield server for graph operations and queries",
		Args:              cobra.ExactArgs(0),
		RunE:              o.Run,
		DisableAutoGenTag: true,
	}
	o.AddFlags(cmd)
	return cmd, nil
}

// withCORS adds CORS support to a Connect HTTP handler.
func withCORS(h http.Handler, o *options) http.Handler {
	middleware := cors.New(cors.Options{
		AllowedOrigins:   o.CORS,
		AllowedMethods:   connectcors.AllowedMethods(),
		AllowedHeaders:   connectcors.AllowedHeaders(),
		ExposedHeaders:   connectcors.ExposedHeaders(),
		AllowCredentials: true,
		MaxAge:           3600,
	})
	return middleware.Handler(h)
}
