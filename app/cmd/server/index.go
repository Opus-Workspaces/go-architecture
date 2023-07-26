package server

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-architecture/app/config"
	dbCfg "go-architecture/app/config/database"
	serverCfg "go-architecture/app/config/server"
	"go-architecture/app/db/mongo"
	"go-architecture/app/models"
	mongo2 "go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Server struct {
	Echo *echo.Echo

	ServerConfig serverCfg.ServerType
	DBConfig     dbCfg.DatabaseType

	DBMongo *mongo.MongoDB
	// DBRedis *redis.Client
}

func InitServer(cfg config.Config) *Server {

	serverCfg := cfg.Server
	dbCfg := cfg.DB

	mongoConn := mongo.ConnectMongo(dbCfg)
	// redisConn := redis.NewRedisClient(dbCfg)

	return &Server{
		Echo:         echo.New(),
		ServerConfig: serverCfg,
		DBConfig:     dbCfg,
		DBMongo:      mongoConn,
		// DBRedis:      redisConn,
	}
}

func Run(s *Server) {
	e := s.Echo
	rateLimiter := middleware.RateLimiterConfig{
		Skipper: middleware.DefaultSkipper,
		Store: middleware.NewRateLimiterMemoryStoreWithConfig(
			middleware.RateLimiterMemoryStoreConfig{
				Rate:      10,
				Burst:     30,
				ExpiresIn: 1 * time.Minute,
			},
		),
		IdentifierExtractor: func(ctx echo.Context) (string, error) {
			id := ctx.RealIP()
			return id, nil
		},
		ErrorHandler: func(context echo.Context, err error) error {
			return context.JSON(http.StatusForbidden, models.ResponseError{
				StatusCode: http.StatusForbidden,
				Type:       "server.go.rate_limiter.error_handler",
				Message:    err.Error(),
			})
		},
		DenyHandler: func(context echo.Context, identifier string, err error) error {
			return context.JSON(http.StatusTooManyRequests, models.ResponseError{
				StatusCode: http.StatusTooManyRequests,
				Type:       "server.go.rate_limiter.deny_handler",
				Message:    err.Error(),
			})
		},
	}

	e.Use(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAuthorization,
			echo.HeaderAccessControlAllowOrigin,
			echo.HeaderAccessControlAllowHeaders},
	}))
	e.Use(middleware.RateLimiterWithConfig(rateLimiter))
	s.Echo.GET("/", func(e echo.Context) error {
		time.Sleep(5 * time.Second) // Simulate long running task
		return e.JSON(http.StatusOK, "Hello World!")
	})

	serverConfig := s.ServerConfig

	go func() {
		if err := e.Start(serverConfig.Port); err != nil && err != http.ErrServerClosed {
			fmt.Println("error starting server: ", err)
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server with a timeout of 10 seconds.
	// use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	defer func(c *mongo2.Client) {
		err := mongo.DisconnectMongo(c)
		if err != nil {
			fmt.Println("error disconnecting from mongo: ", err)
		}
	}(s.DBMongo.Client)

	if err := s.Echo.Shutdown(ctx); err != nil {
		s.Echo.Logger.Fatal(err)
	}
}
