package sms

const unavailableErrMsg = "is not available! Please try again later."
const emptyMessageErrMsg = "Empty message! Please add message content."
const emptyReceiverErrMsg = "Empty receiver! Please add receiver."
const emptyRecipientsErrMsg = "Empty recipients! Please add least one receiver."

// SendSMS Sends a SMS to a specified receiver
func SendSMS(message string, receiver string) error {
	return nil
}

// SendBulkSMS Send a SMS to multiple recipients
func SendBulkSMS(message string, recipients []string) error {
	return nil
}
