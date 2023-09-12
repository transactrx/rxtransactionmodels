package transaction

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func (response *Response) BuildResponseSuccess(claim Claim, startTime time.Time, ncpdpResponse []byte) {

	response.transmissionId = claim.transmissionId
	response.instanceId = claim.instanceId
	response.tenantId = claim.tenantId
	response.elapsed = fmt.Sprintf("%f", (time.Now().Sub(startTime).Seconds()))
	now := time.Now().UTC()
	formattedTime := now.Format("2006-01-02T15:04:05.9999999Z")
	response.created = formattedTime
	response.routeAddress = claim.routeAddress
	response.routeName = claim.routeName
	response.rxNumbers = claim.rxNumbers
	response.ncpdp = claim.ncpdp
	response.callOrigin = claim.callOrigin
	response.from, _ = os.Hostname()
	response.returnCode = "1" // ok status
	response.statusCode = "200"
	response.transactionData.ncpdpData = string(ncpdpResponse)
}

func (response *Response) BuildResponseError(claim Claim, errorCode ErrorInfo, startTime time.Time) {

	response.transmissionId = claim.transmissionId
	response.instanceId = claim.instanceId
	response.elapsed = fmt.Sprintf("%f", (time.Now().Sub(startTime).Seconds()))
	now := time.Now().UTC()
	formattedTime := now.Format("2006-01-02T15:04:05.9999999Z")
	response.created = formattedTime
	response.routeAddress = claim.routeAddress
	response.routeName = claim.routeName
	response.rxNumbers = claim.rxNumbers
	response.ncpdp = claim.ncpdp
	response.callOrigin = claim.callOrigin
	response.from, _ = os.Hostname()
	response.returnCode = "3" // failure
	response.statusCode = "200"
	response.transactionData.ncpdpData = GenerateError([]byte(claim.transactionData.ncpdpData), errorCode)
}

// func GenerateError(responseBuffer []byte,errorCode string) string {
func GenerateError(responseBuffer []byte, errorCode ErrorInfo) string {

	ncpdpClaim := string(responseBuffer)
	errorMessage := errorCode.Code + ": " + errorCode.Message
	response := "D0" + ncpdpClaim[8:8+2] + ncpdpClaim[20:20+1] + "R" + ncpdpClaim[21:21+2] + ncpdpClaim[23:23+15] + time.Now().Format("20060102")
	//Build Response Message Segment AM20
	response += SEGMENT_SEPARATOR
	response += FIELD_SEPARATOR + "AM20"
	response += FIELD_SEPARATOR + "F4" + errorMessage
	tmp := ncpdpClaim[20 : 20+1]
	numberOfClaims, err := strconv.Atoi(tmp)
	//Build Response Status Segment for each claim AM21
	if err != nil {
		numberOfClaims = 1
	}
	for i := 0; i < numberOfClaims; i++ {
		response += GROUP_SEPARATOR
		response += SEGMENT_SEPARATOR
		response += FIELD_SEPARATOR + "AM21"
		response += FIELD_SEPARATOR + "ANR"
		response += FIELD_SEPARATOR + "FA2"
		response += FIELD_SEPARATOR + "FB85"
		response += FIELD_SEPARATOR + "FBNN"
	}
	response += ETX
	return response
}
