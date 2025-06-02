package utils

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
