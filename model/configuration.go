package model

import "time"

type Configuration struct {
	Id                         int64     `json:"id"`
	DietClassifierEpoch        int32     `json:"dietclassifier_epoch"`
	FallbackClassifierTreshold int32     `json:"fallbackclassifier_treshold"`
	ResponseSelectorEpoch      int32     `json:"responseselector_epoch"`
	TedPolicyEpoch             int32     `json:"tedpolicy_epoch"`
	FallbackUtteranceId        int64     `json:"fallback_utterance_id"`
	FallbackTreshold           int32     `json:"fallback_treshold"`
	CreatedAt                  time.Time `json:"created_at"`
	UpdatedAt                  time.Time `json:"updated_at"`
}
