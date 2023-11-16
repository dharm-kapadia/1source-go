// Package models contains the type structures related to 1source-go
package models

type ContractInitiationResponse struct {
	Timestamp string `json:"timestamp"`
	Status    uint32 `json:"status"`
	Message   string `json:"message"`
	Path      string `json:"path"`
}

type ContractCancelReponse struct {
	Timestamp string `json:"timestamp"`
	Status    uint32 `json:"status"`
	Message   string `json:"message"`
	Path      string `json:"path"`
}

type ContractDeclineReponse struct {
	Timestamp string `json:"timestamp"`
	Status    uint32 `json:"status"`
	Message   string `json:"message"`
	Path      string `json:"path"`
}
