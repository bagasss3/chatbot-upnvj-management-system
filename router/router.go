package router

import (
	"cbupnvj/model"

	"github.com/labstack/echo/v4"
)

type router struct {
	group                     *echo.Group
	userController            model.UserController
	authController            model.AuthController
	intentController          model.IntentController
	utteranceController       model.UtteranceController
	exampleController         model.ExampleController
	actionHttpController      model.ActionHttpController
	reqBodyController         model.ReqBodyController
	krsActionController       model.KrsActionController
	entityController          model.EntityController
	ruleController            model.RuleController
	storyController           model.StoryController
	stepController            model.StepController
	configurationController   model.ConfigurationController
	trainingHistoryController model.TrainingHistoryController
	logIntentController       model.LogIntentController
	workerController          model.WorkerController
	fallbackChatLogController model.FallbackChatLogController
	conversationController    model.ConversationController
	majorController           model.MajorController
	facultyController         model.FacultyController
}

func NewRouter(group *echo.Group, userController model.UserController,
	authController model.AuthController,
	intentController model.IntentController,
	utteranceController model.UtteranceController,
	exampleController model.ExampleController,
	actionHttpController model.ActionHttpController,
	reqBodyController model.ReqBodyController,
	krsActionController model.KrsActionController,
	entityController model.EntityController,
	ruleController model.RuleController,
	storyController model.StoryController,
	stepController model.StepController,
	configurationController model.ConfigurationController,
	trainingHistoryController model.TrainingHistoryController,
	logIntentController model.LogIntentController,
	workerController model.WorkerController,
	fallbackChatLogController model.FallbackChatLogController,
	conversationController model.ConversationController,
	majorController model.MajorController,
	facultyController model.FacultyController) {
	rt := &router{
		group:                     group,
		userController:            userController,
		authController:            authController,
		intentController:          intentController,
		utteranceController:       utteranceController,
		exampleController:         exampleController,
		actionHttpController:      actionHttpController,
		reqBodyController:         reqBodyController,
		krsActionController:       krsActionController,
		entityController:          entityController,
		ruleController:            ruleController,
		storyController:           storyController,
		stepController:            stepController,
		configurationController:   configurationController,
		trainingHistoryController: trainingHistoryController,
		logIntentController:       logIntentController,
		workerController:          workerController,
		fallbackChatLogController: fallbackChatLogController,
		conversationController:    conversationController,
		majorController:           majorController,
		facultyController:         facultyController,
	}

	rt.RouterInit()
}

func (r *router) RouterInit() {
	r.userRouter()
	r.authRouter()
	r.intentRouter()
	r.utteranceRouter()
	r.exampleRouter()
	r.actionRouter()
	r.conversationRouter()
	r.configurationRouter()
	r.workerRouter()
	r.majorFacultyRouter()
}
