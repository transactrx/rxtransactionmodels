package transaction

import "time"

//some comment.

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

	Elapsed         string `json:"elapsed"`
	From            string `json:"from"`
	RouteAddress    string `json:"routeAddress"`
	RouteName       string `json:"routeName"`
	Created         string `json:"created"`
	RxNumbers       []interface{} `json:"rxNumbers"`
	Ncpdp           string `json:"ncpdp"`
	CallOrigin      int `json:"callOrigin"`
	TransactionData struct {
		NcpdpData string `json:"ncpdpData"`
	} `json:"transactionData"`
	ReplyTo  string `json:"replyTo"`
	TimeRcvd time.Time `json:"timeRcvd"`
}


type Response struct {
	TransmissionId string `json:"transmissionId"`
	InstanceId     string `json:"instanceId"`
	TenantId       string `json:"tenantId"`
	UnitName       string `json:"unitName"`
	ClientId       string `json:"clientId"`
	Elapsed        string `json:"elapsed"`
	From           string `json:"from"`
	RouteAddress   string `json:"routeAddress"`
	RouteName      string `json:"routeName"`
	Created        string `json:"created"`
	RxNumbers      []interface{} `json:"rxNumbers"`
	Ncpdp          string `json:"ncpdp"`
	CallOrigin     int `json:"callOrigin"`
	ReturnCode     string `json:"returnCode"`
	StatusCode     string `json:"statusCode"`
	Reader         string `json:"reader"`
	ContextData    string `json:"contextData"`
	TransactionData struct {
		NcpdpData string `json:"ncpdpData"`
	} `json:"transactionData"`
	ErrorData *ErrorInfo `json:"errorData"`
}
