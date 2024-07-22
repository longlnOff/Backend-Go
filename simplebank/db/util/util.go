package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// Constants for all supported currencies
const (
    USD = "USD"
    EUR = "EUR"
CAD = "CAD"
)

// IsSupportedCurrency returns true if the currency is supported
func IsSupportedCurrency(currency string) bool {
    switch currency {
    case USD, EUR, CAD:
        return true
    }
    return false
}


// HashPassword returns the bcrypt has of the password
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to has password: %w", err)
	}
	return string(hashedPassword), nil
}

func CheckPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}