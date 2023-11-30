package transaction

import "time"

// Transaction types
const (
	REQUEST  = "REQUEST"
	RESPONSE = "RESPONSE"
)

// Separators
const (
	SEGMENT_SEPARATOR = "\u001E"
	GROUP_SEPARATOR   = "\u001D"
	FIELD_SEPARATOR   = "\u001C"
	ETX               = "\u0003"
)

type Claim struct {
	TransmissionId string `json:"transmissionId"`
	InstanceId     string `json:"instanceId"`
	TenantId       string `json:"tenantId"`
	UnitName       string `json:"unitName"`
	ClientId       string `json:"clientId"`

	Elapsed         string        `json:"elapsed"`
	From            string        `json:"from"`
	RouteAddress    string        `json:"routeAddress"`
	RouteName       string        `json:"routeName"`
	Created         string        `json:"created"`
	RxNumbers       []interface{} `json:"rxNumbers"`
	Ncpdp           string        `json:"ncpdp"`
	CallOrigin      int           `json:"callOrigin"`
	TransactionData struct {
		NcpdpData string `json:"ncpdpData"`
	} `json:"transactionData"`
	ReplyTo  string    `json:"replyTo"`
	TimeRcvd time.Time `json:"timeRcvd"`
	Header   string    `json:"header"`
}

type Response struct {
	TransmissionId  string        `json:"transmissionId"`
	InstanceId      string        `json:"instanceId"`
	TenantId        string        `json:"tenantId"`
	UnitName        string        `json:"unitName"`
	ClientId        string        `json:"clientId"`
	Elapsed         string        `json:"elapsed"`
	From            string        `json:"from"`
	RouteAddress    string        `json:"routeAddress"`
	RouteName       string        `json:"routeName"`
	Created         string        `json:"created"`
	RxNumbers       []interface{} `json:"rxNumbers"`
	Ncpdp           string        `json:"ncpdp"`
	CallOrigin      int           `json:"callOrigin"`
	ReturnCode      string        `json:"returnCode"`
	StatusCode      string        `json:"statusCode"`
	Reader          string        `json:"reader"`
	ContextData     string        `json:"contextData"`
	TransactionData struct {
		NcpdpData string `json:"ncpdpData"`
	} `json:"transactionData"`
	ErrorData *ErrorInfo `json:"errorData"`
	PpeStatus *PpeStatus `json:"ppeStatus"`
}

type PpeStatus struct {
	ReplyToSwitch   string `json:"replyToSwitch"`
	Status          string `json:"status"`
	CallOrigin      int    `json:"callOrigin"`
	PpeRouteAddress string `json:"ppeRouteAddress"`
}

type MultiTripResponse struct {
	TransmissionId string        `json:"transmissionId"`
	InstanceId     string        `json:"instanceId"`
	TenantId       string        `json:"tenantId"`
	UnitName       string        `json:"unitName"`
	ClientId       string        `json:"clientId"`
	Elapsed        string        `json:"elapsed"`
	From           string        `json:"from"`
	RouteAddress   string        `json:"routeAddress"`
	RouteName      string        `json:"routeName"`
	Created        string        `json:"created"`
	RxNumbers      []interface{} `json:"rxNumbers"`
	Ncpdp          string        `json:"ncpdp"`
	CallOrigin     int           `json:"callOrigin"`
	ReturnCode     string        `json:"returnCode"`
	StatusCode     string        `json:"statusCode"`
	Reader         string        `json:"reader"`
	ContextData    string        `json:"contextData"`
	ErrorData      *ErrorInfo    `json:"errorData"`
	PpeStatus      *PpeStatus    `json:"ppeStatus"`

	//Final response
	TransactionData TransactionData `json:"transactionData"`

	//List of other transaction trips that took place between original request and final response.
	Trips []TransactionTrip `json:"trips"`
}

type TransactionData struct {
	NcpdpData string `json:"ncpdpData"`
}

type TransactionTrip struct {
	TransmissionId string      `json:"transmissionId"`
	TripType       string      `json:"type"` //Tran code: B1, B2, etc.
	Request        Transaction `json:"request"`
	Response       Transaction `json:"response"`
}

type Transaction struct {
	Created            string   `json:"created"`
	TransactionType    string   `json:"type"` //Request,Response
	RouteCode          string   `json:"routeCode"`
	RouteAddress       string   `json:"routeAddress"`
	ResponseStatusCode string   `json:"responseStatusCode"`
	RejectCodes        []string `json:"rejectCodes"`

	NcpdpData string `json:"ncpdpData"`
}
