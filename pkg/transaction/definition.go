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
	transmissionId string
	instanceId     string
	tenantId       string
	unitName       string
	clientId       string

	elapsed         string
	from            string
	routeAddress    string
	routeName       string
	created         string
	rxNumbers       []interface{}
	ncpdp           string
	callOrigin      int
	transactionData struct {
		ncpdpData string
	}
	replyTo  string
	timeRcvd time.Time
}
type Response struct {
	transmissionId string
	instanceId     string
	tenantId       string
	unitName       string
	clientId       string
	elapsed        string
	from           string
	routeAddress   string
	routeName      string
	created        string
	rxNumbers      []interface{}
	ncpdp          string
	callOrigin     int
	returnCode     string
	statusCode     string
	reader         string
	contextData    string

	transactionData struct {
		ncpdpData string
	}
	errorData *ErrorInfo
}
