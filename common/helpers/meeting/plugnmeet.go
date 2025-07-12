package meetingHelper

import (
	httpHelper "api/common/helpers/http"
	"api/config"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func GetPlugnMeetSignature(body string) (signature string) {
	mac := hmac.New(sha256.New, []byte(config.Env.MeetingApiSecret))
	mac.Write([]byte(body))
	signature = hex.EncodeToString(mac.Sum(nil))
	return
}

func GetPlugnMeetPostHeaders(signature string) []httpHelper.HttpHeader {
	var headers = make([]httpHelper.HttpHeader, 0, 2)
	headers = append(headers, httpHelper.HttpHeader{Label: "API-KEY", Value: config.Env.MeetingApiKey})
	headers = append(headers, httpHelper.HttpHeader{Label: "HASH-SIGNATURE", Value: signature})
	return headers
}
