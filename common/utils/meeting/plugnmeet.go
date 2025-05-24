package meeting

import (
	"api/common/utils"
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

func GetPlugnMeetPostHeaders(signature string) []utils.HttpHeader {
	var headers = make([]utils.HttpHeader, 0, 2)
	headers = append(headers, utils.HttpHeader{Label: "API-KEY", Value: config.Env.MeetingApiKey})
	headers = append(headers, utils.HttpHeader{Label: "HASH-SIGNATURE", Value: signature})
	return headers
}
