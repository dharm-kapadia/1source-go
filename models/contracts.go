// Package models contains the type structures related to 1source-go
package models

type (
	Contract struct {
		ContractId         string
		LastEventId        uint32
		ContractStatus     string
		SettlementStatus   string
		LastUpdatePartyId  string
		LastUpdateDateTime string
		Trade              trade
		Settlement         settlement
	}

	trade struct {
		ExecutionVenue     executionvenue
		Instrument         instrument
		Rate               rate
		Quantity           uint32
		BillingCurrency    string
		DividentRatePct    float32
		TradeDate          string
		SettlementType     string
		Collateral         collateral
		TransactingParties []transactingparties
	}

	executionvenue struct {
		Evtype       string
		Platform     platform
		VenueParties []venueparties
	}

	venueparties struct {
		PartyRole string
	}

	platform struct {
		GleifLei   string
		LegalName  string
		VenueName  string
		VenueRefId string
	}

	instrument struct {
		Ticker      string
		Cusip       string
		Isin        string
		Sedol       string
		Figi        string
		Description string
	}

	rate struct {
		Rebate rebate
	}

	rebate struct {
		Fixed fixed
	}

	fixed struct {
		BaseRate      float32
		EffectiveDate string
		EffectiveRate float32
	}

	collateral struct {
		ContractValue   float64
		CollateralValue float64
		Currency        string
		RoundingRule    uint32
		RoundingMode    string
		Margin          uint32
	}

	transactingparties struct {
		PartyRole string
		Party     party
	}

	party struct {
		PartyId         string
		PartyName       string
		GleifLei        string
		InternalPartyId string
	}

	settlement struct {
		PartyRole   string
		Instruction instruction
	}

	instruction struct {
		SettlementBic     string
		LocalAgentBic     string
		LocalAgentName    string
		LocalAgentAcct    string
		LocalMarketFields localmarketfields
	}

	localmarketfields struct {
		LocalFieldName  string
		LocalFieldValue string
	}
)

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
