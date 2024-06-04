package response

type HotelResponseDTO struct {
	AuditData AuditData `json:"auditData"`
	Hotels    *Hotels   `json:"hotels"`
}

type AuditData struct {
	ProcessTime string `json:"processTime"`
	Timestamp   string `json:"timestamp"`
	RequestHost string `json:"requestHost"`
	ServerID    string `json:"serverId"`
	Environment string `json:"environment"`
	Release     string `json:"release"`
	Token       string `json:"token"`
	Internal    string `json:"internal"`
}

type Hotels struct {
	Hotels   *[]Hotel `json:"hotels"`
	CheckIn  string   `json:"checkIn"`
	Total    int      `json:"total"`
	CheckOut string   `json:"checkOut"`
}

type Hotel struct {
	Code            int    `json:"code"`
	Name            string `json:"name"`
	CategoryCode    string `json:"categoryCode"`
	CategoryName    string `json:"categoryName"`
	DestinationCode string `json:"destinationCode"`
	DestinationName string `json:"destinationName"`
	ZoneCode        int    `json:"zoneCode"`
	ZoneName        string `json:"zoneName"`
	Latitude        string `json:"latitude"`
	Longitude       string `json:"longitude"`
	Rooms           []Room `json:"rooms"`
	MinRate         string `json:"minRate"`
	MaxRate         string `json:"maxRate"`
	Currency        string `json:"currency"`
}

type Room struct {
	Code  string `json:"code"`
	Name  string `json:"name"`
	Rates []Rate `json:"rates"`
}

type Rate struct {
	RateKey              string               `json:"rateKey"`
	RateClass            string               `json:"rateClass"`
	RateType             string               `json:"rateType"`
	Net                  string               `json:"net"`
	Allotment            int                  `json:"allotment"`
	PaymentType          string               `json:"paymentType"`
	Packaging            bool                 `json:"packaging"`
	BoardCode            string               `json:"boardCode"`
	BoardName            string               `json:"boardName"`
	CancellationPolicies []CancellationPolicy `json:"cancellationPolicies"`
	Taxes                Taxes                `json:"taxes"`
	Rooms                int                  `json:"rooms"`
	Adults               int                  `json:"adults"`
	Children             int                  `json:"children"`
	Promotions           []Promotion          `json:"promotions,omitempty"`
}

type CancellationPolicy struct {
	Amount string `json:"amount"`
	From   string `json:"from"`
}

type Taxes struct {
	Taxes       []Tax `json:"taxes"`
	AllIncluded bool  `json:"allIncluded"`
}

type Tax struct {
	Included       bool   `json:"included"`
	Amount         string `json:"amount"`
	Currency       string `json:"currency"`
	Type           string `json:"type"`
	SubType        string `json:"subType"`
	ClientAmount   string `json:"clientAmount,omitempty"`
	ClientCurrency string `json:"clientCurrency,omitempty"`
	Percent        string `json:"percent,omitempty"`
}

type Promotion struct {
	Code   string `json:"code"`
	Name   string `json:"name"`
	Remark string `json:"remark,omitempty"`
}
