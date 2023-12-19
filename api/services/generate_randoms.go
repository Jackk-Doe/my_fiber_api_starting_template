package services

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Generate 6 random characters
*/
func GenerateSixRandomCharacters() string {
	// Create a seed for random number
	source := rand.NewSource(time.Now().UnixNano())
	myRand := rand.New(source)

	var result string
	characters := "ABCDEFGHJKMNPRTUVWXYZ23456789"
	// characters := "0123456789"
	charactersLength := len(characters)
	for i := 0; i < 6; i++ {
		result += string(characters[myRand.Intn(charactersLength)])
	}
	return result
}

/*
Generate 6 random digits OTP
*/
func GenerateSixRandomOTP() string {
	// Seed the random number generator with the current time
	rand.Seed(time.Now().UnixNano())

	// Generate a random 6-digit OTP
	otp := ""
	for i := 0; i < 6; i++ {
		digit := rand.Intn(10) // Generates a random digit from 0 to 9
		otp += fmt.Sprintf("%d", digit)
	}

	return otp
}
