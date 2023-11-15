package models

type (
	AppConfig struct {
		General        general
		Endpoints      endpoints
		Authentication authentication
	}

	general struct {
		Auth_URL   string
		Realm_Name string
	}

	endpoints struct {
		Base       string
		Parties    string
		Events     string
		Agreements string
		Contracts  string
		Rerates    string
		Returns    string
		Recalls    string
		Buyins     string
	}

	authentication struct {
		Auth_Type     string
		Grant_Type    string
		Client_Id     string
		Username      string
		Password      string
		Client_Secret string
	}
)
