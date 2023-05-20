package model

type RasaResponse struct {
	Version string      `json:"version"`
	Status  Status      `json:"status"`
	Message string      `json:"message"`
	Reason  string      `json:"reason"`
	Details struct{}    `json:"details"`
	Help    interface{} `json:"help"`
	Code    int         `json:"code"`
}

type RasaChangeModelReq struct {
	ModelFile string `json:"model_file"`
}
