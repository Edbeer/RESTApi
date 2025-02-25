package rest

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Edbeer/restapi/config"
	"github.com/Edbeer/restapi/internal/service"
	"github.com/Edbeer/restapi/internal/storage/psql"
	"github.com/Edbeer/restapi/internal/storage/redis"
	"github.com/Edbeer/restapi/internal/transport/rest/api"
	"github.com/Edbeer/restapi/pkg/logger"
	"github.com/go-redis/redis/v9"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type Server struct {
	echo        *echo.Echo
	config      *config.Config
	psqlClient  *sqlx.DB
	redisClient *redis.Client
	logger      logger.Logger
}

// New server constructor
func NewServer(config *config.Config, psqlClient *sqlx.DB, redisClient *redis.Client, logger logger.Logger) *Server {
	return &Server{
		echo:        echo.New(),
		psqlClient:  psqlClient,
		redisClient: redisClient,
		config:      config,
		logger:      logger,
	}
}

// Run server depends on config SSL option
func (s *Server) Run() error {
	if s.config.Server.SSL {
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"

		// Services, Repos & API Handlers
		cfg := config.GetConfig()

		psql := psql.NewStorage(s.psqlClient)
		redis := redisrepo.NewStorage(s.redisClient, s.config)
		service := service.NewService(service.Deps{
			Logger:       s.logger,
			Config:       s.config,
			PsqlStorage:  psql,
			RedisStorage: redis})
		handler := api.NewHandlers(api.Deps{
			AuthService:     service.Auth,
			NewsService:     service.News,
			CommentsService: service.Comments,
			Config:          cfg,
			Logger:          s.logger,
		})
		if err := handler.Init(s.echo); err != nil {
			s.logger.Fatal(err)
		}

		s.echo.Server.ReadTimeout = time.Second * time.Duration(s.config.Server.ReadTimeout)
		s.echo.Server.WriteTimeout = time.Second * time.Duration(s.config.Server.WriteTimeout)

		go func() {
			s.logger.Infof("Server listening on port: %s", s.config.Server.Port)
			if err := s.echo.StartTLS(s.config.Server.Port, certFile, keyFile); err != nil {
				s.logger.Fatalf("error starting TLS server %v", err)
			}
		}()

		// Graceful shutdown
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

		<-quit

		ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
		defer shutdown()

		s.logger.Info("Server Exited Properly")
		return s.echo.Shutdown(ctx)
	} else {
		e := echo.New()

		// Services, Repos & API Handlers
		cfg := config.GetConfig()

		psql := psql.NewStorage(s.psqlClient)
		redis := redisrepo.NewStorage(s.redisClient, s.config)
		service := service.NewService(service.Deps{
			Logger:       s.logger,
			Config:       s.config,
			PsqlStorage:  psql,
			RedisStorage: redis})
		handler := api.NewHandlers(api.Deps{
			AuthService:     service.Auth,
			NewsService:     service.News,
			CommentsService: service.Comments,
			Config:          cfg,
			Logger:          s.logger,
		})
		if err := handler.Init(e); err != nil {
			s.logger.Fatal(err)
		}

		server := &http.Server{
			Addr:           s.config.Server.Port,
			ReadTimeout:    time.Second * time.Duration(s.config.Server.ReadTimeout),
			WriteTimeout:   time.Second * time.Duration(s.config.Server.WriteTimeout),
			MaxHeaderBytes: 1 << 20,
		}

		go func() {
			s.logger.Infof("Server is listening on port: %s", s.config.Server.Port)
			if err := e.StartServer(server); err != nil {
				s.logger.Fatalf("Error starting Server: %v", err)
			}
		}()

		// Graceful shutdown
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

		<-quit

		ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
		defer shutdown()

		s.logger.Info("Server Exited Properly")
		return s.echo.Server.Shutdown(ctx)
	}
}
