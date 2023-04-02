package helper

import (
	"errors"
	"math/rand"
	"time"

	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func ParseTimeDuration(t string, defaultt time.Duration) time.Duration {
	timeDurr, err := time.ParseDuration(t)
	if err != nil {
		return defaultt
	}
	return timeDurr
}

// GenerateID based on curr time
func GenerateID() int64 {
	return time.Now().UnixNano() + int64(rand.Intn(10000))
}

func GeneratePassword() string {
	rand.Seed(time.Now().UnixNano())

	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const length = 10

	randBytes := make([]byte, length)
	for i := 0; i < length; i++ {
		randBytes[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	randStr := string(randBytes)

	return randStr
}

func StringToSlice(s string) []string {
	slice := make([]string, len(s))

	for i, c := range s {
		slice[i] = string(c)
	}

	return slice
}

func HashString(txt string) (string, error) {
	bt, err := bcrypt.GenerateFromPassword([]byte(txt), 10)
	if err != nil {
		return "", err
	}

	return string(bt), nil
}

func IsHashStringMatch(plain, cipher []byte) bool {
	err := bcrypt.CompareHashAndPassword(cipher, plain)
	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return false
	}

	if err != nil {
		log.Error(err)
		return false
	}

	return true
}
