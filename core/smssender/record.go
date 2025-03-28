package smssender

import (
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/sms"
)

func SendRecordOTP(app core.App, authRecord *core.Record, otpId string, pass string) error {
	smsClient := app.NewSMSClient()

	message := &sms.SMSMessage{
		To:      authRecord.GetString("phone"),
		Message: "Your OTP is " + pass,
	}

	return smsClient.Send(message)
}
