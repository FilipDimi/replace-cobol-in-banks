package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "qazwsxedcrfvtgbyhnujmikolp"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Random int between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// Generate Random Owner
func RandomOwner() string {
	return RandomString(6)
}

// Generate Random Money Amount
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// Generate EUR, USD or CAD
func RandomCurrency() string {
	currencies := []string{EUR, USD, CAD}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

// RandomEmail generates a random email
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}
