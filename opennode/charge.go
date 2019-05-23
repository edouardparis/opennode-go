package opennode

type Charge struct {
	ID string `json:"id"`
	// Charge Description
	Description string `json:"description"`
	// Charge price in satoshis
	Amount int64 `json:"amount"`
	// Charge status
	// unpaid/processing/paid
	Status string `json:"status"`
	// Charge fee in satoshis
	Fee int64 `json:"fee"`
	// Charge value at issue time
	FiatValue float64 `json:"fiat_value"`
	// Charge currency
	Currency string `json:"currency"`
	// timestamp
	CreatedAt int64 `json:"created_at"`
	// URL where user gets redirected after payment
	SuccessURL string `json:"success_url"`
	// Charge notes
	Notes string `json:"notes"`
	// Charge requested instant exchange
	AutoSettle bool `json:"auto_settle"`
	// External order ID
	OrderID string `json:"order_id"`
	// Charge payment details
	ChainInvoice ChainInvoice `json:"chain_invoice"`
	// lightning_invoice
	LightningInvoice LightningInvoice `json:"lightning_invoice"`
	// Hashed Order
	// OpenNode signs all charge related events it sends
	// to your endpoints by including a hashed_order field
	// on the event payload. This allows you to validate that
	// the events were sent by OpenNode and not by a third party.
	HashedOrder string `json:"hashed_order"`
}

type ChainInvoice struct {
	//  Bitcoin address
	Address string `json:"address"`
	// Charge settlement timestamp
	SettledAt int64 `json:"settle_at"`
	// Transaction ID on Bitcoin Blockchain
	Tx string `json:"tx"`
}

type LightningInvoice struct {
	// Payment Request creation timestamp
	CreatedAt int64 `json:"created_at"`
	// Charge settlement timestamp
	SettledAt int64 `json:"settle_at"`
	// Payment Request hash
	PayReq string `json:"payreq"`
}

type ChargePayload struct {
	Description string `json:"description"`
	// Charge price in satoshis
	Amount int64 `json:"amount"`
	// Charge currency
	Currency string `json:"currency,omitempty"`
	// URL where user gets redirected after payment
	CallbackURL string `json:"callback_url"`
	// URL where user gets redirected after payment
	SuccessURL string `json:"success_url"`
	// Charge requested instant exchange
	AutoSettle bool `json:"auto_settle"`
	// External order ID
	OrderID string `json:"order_id"`
}
