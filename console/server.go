package console

import (
	"cbupnvj/config"
	"cbupnvj/controller"
	"cbupnvj/database"
	"cbupnvj/repository"
	"cbupnvj/router"
	"cbupnvj/service"
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "run server",
	Long:  "Start running the server",
	Run:   server,
}

func init() {
	RootCmd.AddCommand(serverCmd)
}

func server(cmd *cobra.Command, args []string) {
	// Initiate Connection
	MysqlDB := database.InitDB()
	db, err := MysqlDB.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create Echo instance
	httpServer := echo.New()
	httpServer.Use(middleware.Logger())
	httpServer.Use(middleware.Recover())
	httpServer.Use(middleware.CORS())

	// Initiate Depedency
	userRepository := repository.NewUserRepository(MysqlDB)
	sessionRepository := repository.NewSessionRepository(MysqlDB)

	userService := service.NewUserService(userRepository)
	authService := service.NewAuthService(userRepository, sessionRepository)

	userController := controller.NewUserController(userService)
	authController := controller.NewAuthController(authService)

	router.NewRouter(httpServer.Group("/api"), userController, authController)

	// Graceful Shutdown
	// Catch Signal
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		defer signal.Stop(sigChan)
		defer close(sigChan)

		<-sigChan
		log.Info("Received termination signal, initiating graceful shutdown...")
		cancel()
	}()

	// Start http server
	go func() {
		log.Info("Starting server...")
		if err := httpServer.Start(":" + config.Port()); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting the server: %v", err)
		}
	}()

	// Shutting down any connection and server
	<-ctx.Done()
	log.Info("Shutting down server...")
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatalf("Error shutting down server: %v", err)
	}

	log.Info("Server gracefully shut down")
}
