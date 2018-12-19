package iris

const (
	ONLINE Source = iota + 1
	OFFLINE
)

const (
	USAGE RevenueClass = iota
	DAILY
	MONTHLY
)

const (
	ACTIVE Status = iota + 1
	INACTIVE
	PENDING
)

type (
	Amount struct {
		Currency string  `json:"currency"`
		Value    float32 `json:"value"`
	}

	Payment struct {
		Amount             Amount `json:"amount"`
		BusinessUnit       string `json:"businessUnit"`
		ExecutionTime      string `json:"executionTime"`
		PaymentIdentifier  string `json:"paymentIdentifier"`
		PaymentMethod      string `json:"paymentMethod"`
		SalesforceRecordId string `json:"salesforceRecordId"`
		PaymentType        string `json:"type"`
	}

	Category struct {
		Level1 string `json:"level1"`
		Level2 string `json:"level2"`
		Level3 string `json:"level3"`
	}

	RevenueClass int

	Product struct {
		ActivationTime string   `json:"activationTime"`
		Category       Category `json:"category"`
		Description    string   `json:"description"`
		ExpirationTime string   `json:"expirationTime"`
		Freebie        bool     `json:"freebie"`
		GrossPrice     float32  `json:"grossPrice"`
		Name           string   `json:"name"`
		PackageId      string   `json:"packageId"`
		SkuId          string   `json:"skuId"`
		RevenueClass   string   `json:"revenueClass"`
		ProductType    string   `json:"type"`
		UnitPrice      float32  `json:"unitPrice"`
		Units          int32    `json:"units"`
	}

	Order struct {
		Customer         Customer  `json:"customer"`
		PaymentReference Payment   `json:"paymentReference"`
		Products         []Product `json:"products"`
		OrderId          string    `json:"orderId"`
		Source           string    `json:"source"`
		Status           string    `json:"status"`
	}

	Customer struct {
		Address      Address `json:"address"`
		BusinessName string  `json:"businessName"`
		CustomerName string  `json:"customerName"`
		CustomerType string  `json:"type"`
		EmailAddress string  `json:"emailAddress"`
		FirstName    string  `json:"firstName"`
		Language     string  `json:"language"`
		LastName     string  `json:"lastName"`
		PhoneNumber  string  `json:"phoneNumber"`
		SalesforceId string  `json:"salesforceId"`
		VatNumber    string  `json:"vatNumber"`
	}

	Address struct {
		AddressLines []string `json:"addressLines"`
		City         string   `json:"city"`
		Country      string   `json:"country"`
		HouseNumber  string   `json:"houseNumber"`
		Region       string   `json:"region"`
		State        string   `json:"state"`
		StateCode    string   `json:"stateCode"`
		Street       string   `json:"street"`
		ZipCode      string   `json:"zipCode"`
	}

	Source int

	Status int
)
