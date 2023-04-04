SHELL:=/bin/bash

model/mock/mock_user_repository.go:
	mockgen -destination=model/mock/mock_user_repository.go -package=mock cbupnvj/model UserRepository
model/mock/mock_session_repository.go:
	mockgen -destination=model/mock/mock_session_repository.go -package=mock cbupnvj/model SessionRepository
model/mock/mock_intent_repository.go:
	mockgen -destination=model/mock/mock_intent_repository.go -package=mock cbupnvj/model IntentRepository
model/mock/mock_user_service.go:
	mockgen -destination=model/mock/mock_user_service.go -package=mock cbupnvj/model UserService
model/mock/mock_auth_service.go:
	mockgen -destination=model/mock/mock_auth_service.go -package=mock cbupnvj/model AuthService
model/mock/mock_intent_service.go:
	mockgen -destination=model/mock/mock_intent_service.go -package=mock cbupnvj/model IntentService
model/mock/mock_user_controller.go:
	mockgen -destination=model/mock/mock_user_controller.go -package=mock cbupnvj/model UserController
model/mock/mock_auth_controller.go:
	mockgen -destination=model/mock/mock_auth_controller.go -package=mock cbupnvj/model AuthController
model/mock/mock_intent_controller.go:
	mockgen -destination=model/mock/mock_intent_controller.go -package=mock cbupnvj/model IntentController

mockgen: model/mock/mock_user_repository.go \
	model/mock/mock_session_repository.go \
	model/mock/mock_intent_repository.go \
	model/mock/mock_user_service.go \
	model/mock/mock_auth_service.go \
	model/mock/mock_intent_service.go \
	model/mock/mock_user_controller.go \
	model/mock/mock_auth_controller.go \
	model/mock/mock_intent_controller.go \

clean:
	rm -v model/mock/mock_*.go