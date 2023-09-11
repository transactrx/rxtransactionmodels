package transaction

import "time"

//some comment.

const (
	SEGMENT_SEPARATOR = "\u001E"
	GROUP_SEPARATOR   = "\u001D"
	FIELD_SEPARATOR   = "\u001C"
	ETX               = "\u0003"
)

type Claim structs {
	TransmissionID  string
	Elapsed         string
	From            string
	RouteAddress    string
	RouteName       string
	Created         string
	RxNumbers       []interface{}
	Ncpdp           string
	CallOrigin      int
	TransactionData struct {
		NcpdpData string
	}
	ReplyTo  string
	TimeRcvd time.Time
}
type Response struct {
	TransmissionID string
	Elapsed        string
	From           string
	RouteAddress   string
	RouteName      string
	Created        string
	RxNumbers      []interface{}
	Ncpdp          string
	CallOrigin     int
	ReturnCode     string
	StatusCode     string

	TransactionData struct {
		NcpdpData string
	}
	//ErrorData *ErrorCode
}
