package service

import (
	"cbupnvj/config"
	"cbupnvj/constant"
	"cbupnvj/helper"
	"cbupnvj/middleware"
	"cbupnvj/model"
	"context"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

type workerService struct {
	trainingHistoryRepository model.TrainingHistoryRepository
	intentRepository          model.IntentRepository
	utteranceRepository       model.UtteranceRepository
	actionHttpRepository      model.ActionHttpRepository
	entityRepository          model.EntityRepository
	exampleRepository         model.ExampleRepository
	ruleRepository            model.RuleRepository
	storyRepository           model.StoryRepository
	stepRepository            model.StepRepository
	configurationRepository   model.ConfigurationRepository
}

func NewWorkerService(
	trainingHistoryRepository model.TrainingHistoryRepository,
	intentRepository model.IntentRepository,
	utteranceRepository model.UtteranceRepository,
	actionHttpRepository model.ActionHttpRepository,
	entityRepository model.EntityRepository,
	exampleRepository model.ExampleRepository,
	ruleRepository model.RuleRepository,
	storyRepository model.StoryRepository,
	stepRepository model.StepRepository,
	configurationRepository model.ConfigurationRepository,
) model.WorkerService {
	return &workerService{
		trainingHistoryRepository: trainingHistoryRepository,
		intentRepository:          intentRepository,
		utteranceRepository:       utteranceRepository,
		actionHttpRepository:      actionHttpRepository,
		entityRepository:          entityRepository,
		exampleRepository:         exampleRepository,
		ruleRepository:            ruleRepository,
		storyRepository:           storyRepository,
		stepRepository:            stepRepository,
		configurationRepository:   configurationRepository,
	}
}

func (w *workerService) StartTrainingModel(ctx context.Context) (*model.TrainingHistory, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
	})
	// start time
	startTime := time.Now()

	ctxUser := middleware.GetUserFromCtx(ctx)
	if ctxUser == nil {
		log.Error(constant.ErrUnauthorized)
		return nil, constant.ErrUnauthorized
	}

	// initiate strings builder for yml
	var sb strings.Builder

	// write configuration
	configModel, err := w.configurationRepository.FindAll(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	findUtteranceConfig, err := w.utteranceRepository.FindByID(ctx, configModel[0].FallbackUtteranceId)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	// write pipeline
	sb.WriteString("pipeline:\n")
	sb.WriteString("  - name: WhitespaceTokenizer\n")
	sb.WriteString("  - name: RegexFeaturizer\n")
	sb.WriteString("  - name: LexicalSyntacticFeaturizer\n")
	sb.WriteString("  - name: CountVectorsFeaturizer\n")
	sb.WriteString("  - name: CountVectorsFeaturizer\n")
	sb.WriteString("    analyzer: char_wb\n")
	sb.WriteString("    min_ngram: 1\n")
	sb.WriteString("    max_ngram: 4\n")
	sb.WriteString("  - name: DIETClassifier\n")
	sb.WriteString(fmt.Sprintf("    epochs: %d\n", configModel[0].DietClassifierEpoch))
	sb.WriteString("    constrain_similarities: true\n")
	sb.WriteString("  - name: EntitySynonymMapper\n")
	sb.WriteString("  - name: ResponseSelector\n")
	sb.WriteString(fmt.Sprintf("    epochs: %d\n", configModel[0].ResponseSelectorEpoch))
	sb.WriteString("    constrain_similarities: true\n")
	sb.WriteString("  - name: FallbackClassifier\n")
	sb.WriteString(fmt.Sprintf("    threshold: %0.1f\n", configModel[0].FallbackClassifierTreshold))
	sb.WriteString("    ambiguity_threshold: 0.1\n")

	// write policy
	sb.WriteString("\npolicy:\n")
	sb.WriteString("  - name: MemoizationPolicy\n")
	sb.WriteString("  - name: TEDPolicy\n")
	sb.WriteString("    max_history: 5\n")
	sb.WriteString(fmt.Sprintf("    epochs: %d\n", configModel[0].TedPolicyEpoch))
	sb.WriteString("    constrain_similarities: true\n")
	sb.WriteString("  - name: UnexpecTEDIntentPolicy\n")
	sb.WriteString("    max_history: 5\n")
	sb.WriteString(fmt.Sprintf("    epochs: %d\n", configModel[0].UnexpecTEDIntentPolicyEpoch))
	sb.WriteString("  - name: RulePolicy\n")
	sb.WriteString(fmt.Sprintf("    core_fallback_threshold: %0.1f\n", configModel[0].FallbackTreshold))
	sb.WriteString(fmt.Sprintf("    core_fallback_action_name: %s\n", findUtteranceConfig.Name))

	// write intents
	intents, err := w.intentRepository.FindAll(ctx, "")
	if err != nil {
		log.Error(err)
		return nil, err
	}

	// store intents in map
	intentsMap := make(map[string]*model.Intent)
	for _, intent := range intents {
		intentsMap[intent.Id] = intent
	}

	sb.WriteString("\nintents:\n")
	for i := range intents {
		sb.WriteString(fmt.Sprintf("  - %s\n", intents[i].Name))
	}

	// write entities
	entities, err := w.entityRepository.FindAllWithNoIntentId(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if len(entities) > 0 {
		sb.WriteString("\nentities:\n")
		for i := range entities {
			sb.WriteString(fmt.Sprintf("  - %s\n", entities[i].Name))
		}
	} else {
		sb.WriteString("\nentities: []\n")
	}

	// write action
	actionHttp, err := w.actionHttpRepository.FindAll(ctx, "")
	if err != nil {
		log.Error(err)
		return nil, err
	}

	// store action in map
	actionHttpMap := make(map[string]*model.ActionHttp)
	for _, action := range actionHttp {
		actionHttpMap[action.Id] = action
	}

	if len(actionHttp) > 0 {
		sb.WriteString("\nactions:\n")
		for i := range actionHttp {
			sb.WriteString(fmt.Sprintf("  - %s\n", actionHttp[i].Name))
		}
	} else {
		sb.WriteString("\nactions: []\n")
	}

	// write utterances
	utterances, err := w.utteranceRepository.FindAll(ctx, "")
	if err != nil {
		log.Error(err)
		return nil, err
	}

	// store utterances in map
	utteranceMap := make(map[string]*model.Utterance)
	for _, utterance := range utterances {
		utteranceMap[utterance.Id] = utterance
	}

	sb.WriteString("\nresponses:\n")
	for i := range utterances {
		sb.WriteString(fmt.Sprintf("  %s:\n", utterances[i].Name))
		sb.WriteString(fmt.Sprintf("  - text: \"%s\"\n\n", utterances[i].Response))
	}

	// write session
	sb.WriteString("session_config:\n")
	sb.WriteString("  session_expiration_time: 60\n")
	sb.WriteString("  carry_over_slots_to_new_session: true\n")

	// write nlu
	sb.WriteString("\nnlu:\n")
	for i := range intents {
		examples, err := w.exampleRepository.FindAllByIntentID(ctx, intents[i].Id)
		if err != nil {
			log.Error(err)
			return nil, err
		}

		sb.WriteString(fmt.Sprintf("- intent: %s\n", intents[i].Name))
		if len(examples) > 0 {
			sb.WriteString("  examples: |\n")

			for i := range examples {
				sb.WriteString(fmt.Sprintf("    - %s\n", examples[i].Example))
			}
		} else {
			sb.WriteString("  examples: []\n")
		}
		sb.WriteString("\n")
	}

	// write rules
	rules, err := w.ruleRepository.FindAll(ctx, "")
	if err != nil {
		log.Error(err)
		return nil, err
	}

	sb.WriteString("rules:\n")
	for i := range rules {
		sb.WriteString(fmt.Sprintf("\n- rule: %s\n", rules[i].RuleTitle))
		sb.WriteString("  steps:\n")

		intent, ok := intentsMap[rules[i].IntentId]
		if !ok {
			// Intent not found
			return nil, fmt.Errorf("[rules] intent not found for ID: %s", rules[i].IntentId)
		}
		sb.WriteString(fmt.Sprintf("  - intent: %s\n", intent.Name))

		var actionName string
		if rules[i].Type == model.RuleUtterance {
			// Find utterance from the utteranceMap
			utterance, ok := utteranceMap[rules[i].ResponseId]
			if !ok {
				// Utterance not found
				return nil, fmt.Errorf("[rules] utterance not found for ID: %s", rules[i].ResponseId)
			}
			actionName = utterance.Name
		} else {
			// Find actionHttp from the actionHttpMap
			actionHttp, ok := actionHttpMap[rules[i].ResponseId]
			if !ok {
				// ActionHttp not found
				return nil, fmt.Errorf("[rules] actionHttp not found for ID: %s", rules[i].ResponseId)
			}
			actionName = actionHttp.Name
		}
		sb.WriteString(fmt.Sprintf("  - action: %s\n", actionName))
	}

	// write stories
	stories, err := w.storyRepository.FindAll(ctx, "")
	if err != nil {
		log.Error(err)
		return nil, err
	}

	sb.WriteString("\nstories:\n")
	for i := range stories {
		sb.WriteString(fmt.Sprintf("\n- story: %s\n", stories[i].StoryTitle))
		sb.WriteString("  steps:\n")

		steps, err := w.stepRepository.FindAll(ctx, stories[i].Id)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		for j := range steps {
			switch steps[j].Type {
			case model.StepIntent:
				intent, ok := intentsMap[steps[j].ResponseId]
				if !ok {
					// ActionHttp not found
					return nil, fmt.Errorf("[stories] intent not found for ID: %s", steps[j].ResponseId)
				}
				sb.WriteString(fmt.Sprintf("  - intent: %s\n", intent.Name))

			case model.StepAction:
				action, ok := actionHttpMap[steps[j].ResponseId]
				if !ok {
					// ActionHttp not found
					return nil, fmt.Errorf("[stories] actionHttp not found for ID: %s", steps[j].ResponseId)
				}
				sb.WriteString(fmt.Sprintf("  - action: %s\n", action.Name))
			case model.StepUtterance:
				// Find utterance from the utteranceMap
				utterance, ok := utteranceMap[steps[j].ResponseId]
				if !ok {
					// ActionHttp not found
					return nil, fmt.Errorf("[stories] utterance not found for ID: %s", steps[j].ResponseId)
				}
				sb.WriteString(fmt.Sprintf("  - action: %s\n", utterance.Name))
			}
		}
	}

	id := helper.GenerateID()
	err = ioutil.WriteFile(fmt.Sprintf("%s/%s.yml", config.GeneratedPath(), id), []byte(sb.String()), 0644)
	if err != nil {
		log.Fatal(err)
	}

	// Calculate the elapsed time
	elapsedTime := time.Since(startTime)
	elapsedSeconds := int(elapsedTime.Seconds())

	create := &model.TrainingHistory{
		Id:        id,
		UserId:    ctxUser.UserID,
		TotalTime: elapsedSeconds,
		Status:    "DONE",
	}

	err = w.trainingHistoryRepository.Create(ctx, create)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return create, nil
}
