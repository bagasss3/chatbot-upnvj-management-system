SHELL:=/bin/bash

model/mock/mock_user_repository.go:
	mockgen -destination=model/mock/mock_user_repository.go -package=mock cbupnvj/model UserRepository
model/mock/mock_session_repository.go:
	mockgen -destination=model/mock/mock_session_repository.go -package=mock cbupnvj/model SessionRepository
model/mock/mock_intent_repository.go:
	mockgen -destination=model/mock/mock_intent_repository.go -package=mock cbupnvj/model IntentRepository
model/mock/mock_action_http_repository.go:
	mockgen -destination=model/mock/mock_action_http_repository.go -package=mock cbupnvj/model ActionHttpRepository
model/mock/mock_configuration_repository.go:
	mockgen -destination=model/mock/mock_configuration_repository.go -package=mock cbupnvj/model ConfigurationRepository	
model/mock/mock_entity_repository.go:
	mockgen -destination=model/mock/mock_entity_repository.go -package=mock cbupnvj/model EntityRepository
model/mock/mock_example_repository.go:
	mockgen -destination=model/mock/mock_example_repository.go -package=mock cbupnvj/model ExampleRepository
model/mock/mock_faculty_repository.go:
	mockgen -destination=model/mock/mock_faculty_repository.go -package=mock cbupnvj/model FacultyRepository
model/mock/mock_fallback_chat_log_repository.go:
	mockgen -destination=model/mock/mock_fallback_chat_log_repository.go -package=mock cbupnvj/model FallbackChatLogRepository
model/mock/mock_log_intent_repository.go:
	mockgen -destination=model/mock/mock_log_intent_repository.go -package=mock cbupnvj/model LogIntentRepository
model/mock/mock_major_repository.go:
	mockgen -destination=model/mock/mock_major_repository.go -package=mock cbupnvj/model MajorRepository
model/mock/mock_req_body_repository.go:
	mockgen -destination=model/mock/mock_req_body_repository.go -package=mock cbupnvj/model ReqBodyRepository
model/mock/mock_rule_repository.go:
	mockgen -destination=model/mock/mock_rule_repository.go -package=mock cbupnvj/model RuleRepository
model/mock/mock_step_repository.go:
	mockgen -destination=model/mock/mock_step_repository.go -package=mock cbupnvj/model StepRepository
model/mock/mock_story_repository.go:
	mockgen -destination=model/mock/mock_story_repository.go -package=mock cbupnvj/model StoryRepository
model/mock/mock_training_history_repository.go:
	mockgen -destination=model/mock/mock_training_history_repository.go -package=mock cbupnvj/model TrainingHistoryRepository
model/mock/mock_utterance_repository.go:
	mockgen -destination=model/mock/mock_utterance_repository.go -package=mock cbupnvj/model UtteranceRepository

model/mock/mock_user_service.go:
	mockgen -destination=model/mock/mock_user_service.go -package=mock cbupnvj/model UserService
model/mock/mock_auth_service.go:
	mockgen -destination=model/mock/mock_auth_service.go -package=mock cbupnvj/model AuthService
model/mock/mock_intent_service.go:
	mockgen -destination=model/mock/mock_intent_service.go -package=mock cbupnvj/model IntentService
model/mock/mock_action_http_service.go:
	mockgen -destination=model/mock/mock_action_http_service.go -package=mock cbupnvj/model ActionHttpService
model/mock/mock_configuration_service.go:
	mockgen -destination=model/mock/mock_configuration_service.go -package=mock cbupnvj/model ConfigurationService
model/mock/mock_conversation_service.go:
	mockgen -destination=model/mock/mock_conversation_service.go -package=mock cbupnvj/model ConversationService
model/mock/mock_entity_service.go:
	mockgen -destination=model/mock/mock_entity_service.go -package=mock cbupnvj/model EntityService
model/mock/mock_example_service.go:
	mockgen -destination=model/mock/mock_example_service.go -package=mock cbupnvj/model ExampleService
model/mock/mock_faculty_service.go:
	mockgen -destination=model/mock/mock_faculty_service.go -package=mock cbupnvj/model FacultyService
model/mock/mock_fallback_chat_log_service.go:
	mockgen -destination=model/mock/mock_fallback_chat_log_service.go -package=mock cbupnvj/model FallbackChatLogService
model/mock/mock_log_intent_service.go:
	mockgen -destination=model/mock/mock_log_intent_service.go -package=mock cbupnvj/model LogIntentService
model/mock/mock_major_service.go:
	mockgen -destination=model/mock/mock_major_service.go -package=mock cbupnvj/model MajorService
model/mock/mock_req_body_service.go:
	mockgen -destination=model/mock/mock_req_body_service.go -package=mock cbupnvj/model ReqBodyService
model/mock/mock_rule_service.go:
	mockgen -destination=model/mock/mock_rule_service.go -package=mock cbupnvj/model RuleService
model/mock/mock_step_service.go:
	mockgen -destination=model/mock/mock_step_service.go -package=mock cbupnvj/model StepService
model/mock/mock_story_service.go:
	mockgen -destination=model/mock/mock_story_service.go -package=mock cbupnvj/model StoryService
model/mock/mock_training_history_service.go:
	mockgen -destination=model/mock/mock_training_history_service.go -package=mock cbupnvj/model TrainingHistoryService
model/mock/mock_utterance_service.go:
	mockgen -destination=model/mock/mock_utterance_service.go -package=mock cbupnvj/model UtteranceService



mockgen: model/mock/mock_user_repository.go \
	model/mock/mock_session_repository.go \
	model/mock/mock_intent_repository.go \
	model/mock/mock_user_service.go \
	model/mock/mock_auth_service.go \
	model/mock/mock_intent_service.go \
	model/mock/mock_action_http_repository.go \
	model/mock/mock_configuration_repository.go \
	model/mock/mock_entity_repository.go \
	model/mock/mock_example_repository.go \
	model/mock/mock_faculty_repository.go \
	model/mock/mock_fallback_chat_log_repository.go \
	model/mock/mock_log_intent_repository.go \
	model/mock/mock_major_repository.go \
	model/mock/mock_req_body_repository.go \
	model/mock/mock_rule_repository.go \
	model/mock/mock_step_repository.go \
	model/mock/mock_story_repository.go \
	model/mock/mock_training_history_repository.go \
	model/mock/mock_utterance_repository.go \
	model/mock/mock_action_http_service.go \
	model/mock/mock_configuration_service.go \
	model/mock/mock_conversation_service.go \
	model/mock/mock_entity_service.go \
	model/mock/mock_example_service.go \
	model/mock/mock_faculty_service.go \
	model/mock/mock_fallback_chat_log_service.go \
	model/mock/mock_log_intent_service.go \
	model/mock/mock_major_service.go \
	model/mock/mock_req_body_service.go \
	model/mock/mock_rule_service.go \
	model/mock/mock_step_service.go \
	model/mock/mock_story_service.go \
	model/mock/mock_training_history_service.go \
	model/mock/mock_utterance_service.go \

clean:
	rm -v model/mock/mock_*.go