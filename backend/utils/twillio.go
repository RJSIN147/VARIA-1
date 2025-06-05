package utils

import (
	"fmt"
	"os"

	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/verify/v2"
)

var (
    TWILIO_ACCOUNT_SID = os.Getenv("TWILIO_ACCOUNT_SID")
    TWILIO_AUTH_TOKEN = os.Getenv("TWILIO_AUTH_TOKEN")
    VERIFY_SERVICE_SID = os.Getenv("VERIFY_SERVICE_SID")
    client = twilio.NewRestClientWithParams(twilio.ClientParams{
        Username: TWILIO_ACCOUNT_SID,
        Password: TWILIO_AUTH_TOKEN,
    })
)

func SendOTP(to string) error {
    params := &openapi.CreateVerificationParams{}
    params.SetTo(to)
    params.SetChannel("sms")
    _, err := client.VerifyV2.CreateVerification(VERIFY_SERVICE_SID, params)
    if err != nil {
        return fmt.Errorf("failed to send OTP: %w", err)
    }
    return nil
}

func VerifyOTP(to, code string) (bool, error) {
    params := &openapi.CreateVerificationCheckParams{}
    params.SetTo(to)
    params.SetCode(code)
    resp, err := client.VerifyV2.CreateVerificationCheck(VERIFY_SERVICE_SID, params)
    if err != nil {
        return false, fmt.Errorf("failed to verify OTP: %w", err)
    }
    return resp.Status != nil && *resp.Status == "approved", nil
}
// import (
//     "fmt"
//     "math/rand"
//     "os"
//     "time"

//     "github.com/twilio/twilio-go"
//     "github.com/twilio/twilio-go/rest/api/v2010"
// )

// // generateOTP generates a 6-digit random OTP
// func generateOTP() string {
//     rand.Seed(time.Now().UnixNano())
//     return fmt.Sprintf("%06d", rand.Intn(1000000))
// }

// // SendOTP generates and sends an OTP via Twilio, and returns the OTP string
// func SendOTP(phone string) (string, error) {
//     otp := generateOTP()

//     client := twilio.NewRestClientWithParams(twilio.ClientParams{
//         Username: os.Getenv("TWILIO_ACCOUNT_SID"),
//         Password: os.Getenv("TWILIO_AUTH_TOKEN"),
//     })

//     body := fmt.Sprintf("Your OTP is: %s", otp)

//     messageParams := &v2010.CreateMessageParams{
//         To:   &phone,
//         From: twilio.String(os.Getenv("TWILIO_PHONE_NUMBER")),
//         Body: &body,
//     }

//     _, err := client.Api.CreateMessage(messageParams)
//     if err != nil {
//         return "", err
//     }

//     return otp, nil
// }
