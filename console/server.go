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
	gormTransationer := repository.NewGormTransactioner(MysqlDB)

	userRepository := repository.NewUserRepository(MysqlDB)
	sessionRepository := repository.NewSessionRepository(MysqlDB)
	intentRepository := repository.NewIntentRepository(MysqlDB)
	utteranceRepository := repository.NewUtteranceRepository(MysqlDB)
	exampleRepository := repository.NewExampleRepository(MysqlDB)
	actionHttpRepository := repository.NewActionHttpRepository(MysqlDB)
	reqBodyRepository := repository.NewReqBodyRepository(MysqlDB)
	krsActionRepository := repository.NewKrsActionRepository(MysqlDB)
	entityRepository := repository.NewEntityRepository(MysqlDB)
	ruleRepository := repository.NewRuleRepository(MysqlDB)
	storyRepository := repository.NewStoryRepository(MysqlDB)
	stepRepository := repository.NewStepRepository(MysqlDB)
	configurationRepository := repository.NewConfigurationRepository(MysqlDB)
	trainingHistoryRepository := repository.NewTrainingHistoryRepository(MysqlDB)
	logIntentRepository := repository.NewLogIntentRepository(MysqlDB)
	fallbackChatLogRepository := repository.NewFallbackChatLogRepository(MysqlDB)
	majorRepository := repository.NewMajorRepository(MysqlDB)
	facultyRepository := repository.NewFacultyRepository(MysqlDB)

	userService := service.NewUserService(userRepository)
	authService := service.NewAuthService(userRepository, sessionRepository)
	intentService := service.NewIntentService(intentRepository, gormTransationer, logIntentRepository, exampleRepository)
	utteranceService := service.NewUtteranceService(utteranceRepository)
	exampleService := service.NewExampleService(exampleRepository, intentRepository)
	actionHttpService := service.NewActionHttpService(actionHttpRepository, reqBodyRepository)
	reqBodyService := service.NewReqBodyService(reqBodyRepository, actionHttpRepository, gormTransationer)
	krsActionService := service.NewKrsActionService(krsActionRepository)
	entityService := service.NewEntityService(entityRepository, intentRepository)
	ruleService := service.NewRuleService(ruleRepository, intentRepository, actionHttpRepository, utteranceRepository)
	storyService := service.NewStoryService(storyRepository, stepRepository, gormTransationer)
	stepService := service.NewStepService(storyRepository, stepRepository, intentRepository, utteranceRepository, actionHttpRepository, gormTransationer)
	configurationService := service.NewConfigurationService(configurationRepository, utteranceRepository)
	trainingHistoryService := service.NewTrainingHistoryService(trainingHistoryRepository, userRepository)
	logIntentService := service.NewLogIntentService(logIntentRepository, intentRepository)
	workerService := service.NewWorkerService(trainingHistoryRepository, intentRepository, utteranceRepository, actionHttpRepository, entityRepository, exampleRepository, ruleRepository, storyRepository, stepRepository, configurationRepository, userRepository)
	fallbackChatLogService := service.NewFallbackChatLogService(fallbackChatLogRepository)
	conversationService := service.NewConversationService(ruleRepository, storyRepository)
	majorService := service.NewMajorService(majorRepository)
	facultyService := service.NewFacultyService(facultyRepository)

	userController := controller.NewUserController(userService)
	authController := controller.NewAuthController(authService)
	intentController := controller.NewIntentController(intentService)
	utteranceController := controller.NewUtteranceController(utteranceService)
	exampleController := controller.NewExampleController(exampleService)
	actionHttpController := controller.NewActionHttpController(actionHttpService)
	reqBodyController := controller.NewReqBodyController(reqBodyService)
	krsActionController := controller.NewKrsActionController(krsActionService)
	entityController := controller.NewEntityController(entityService)
	ruleController := controller.NewRuleController(ruleService)
	storyController := controller.NewStoryController(storyService)
	stepController := controller.NewStepController(stepService)
	configurationController := controller.NewConfigurationController(configurationService)
	trainingHistoryController := controller.NewTrainingHistoryController(trainingHistoryService)
	logIntentController := controller.NewLogIntentController(logIntentService)
	workerController := controller.NewWorkerController(workerService)
	fallbackChatLogController := controller.NewFallbackChatLogController(fallbackChatLogService)
	conversationController := controller.NewConversationController(conversationService)
	majorController := controller.NewMajorController(majorService)
	facultyController := controller.NewFacultyController(facultyService)

	router.NewRouter(httpServer.Group("/api"),
		userController,
		authController,
		intentController,
		utteranceController,
		exampleController,
		actionHttpController,
		reqBodyController,
		krsActionController,
		entityController,
		ruleController,
		storyController,
		stepController,
		configurationController,
		trainingHistoryController,
		logIntentController,
		workerController,
		fallbackChatLogController,
		conversationController,
		majorController,
		facultyController)

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
