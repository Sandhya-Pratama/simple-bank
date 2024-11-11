package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().Unix())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// random owner generator
func RandomOwner() string {
	return RandomString(8)
}

// random money generator
func RandomMoney() int64 {
	return RandomInt(0, 100)
}

// random currency generator
func RandomCurrency() string {
	currencies := []string{"IDR", "USD", "EUR"}
	n := len(currencies)

	return currencies[rand.Intn(n)]
}

func RandomEmail() string {
	return RandomString(6) + "@gmail.com"
}
