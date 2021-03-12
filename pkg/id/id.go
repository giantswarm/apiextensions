package id

import (
	"math/rand"
	"regexp"
	"time"
)

const (
	// IDChars represents the character set used to generate cluster IDs.
	// (does not contain 1 and l, to avoid confusion)
	IDChars = "023456789abcdefghijkmnopqrstuvwxyz"
	// IDLength represents the number of characters used to create a cluster ID.
	IDLength = 5
)

func Generate() string {
	idRegexp := regexp.MustCompile("^[a-z]([a-z][0-9]|[0-9][a-z])+$")
	letterRunes := []rune(IDChars)

	b := make([]rune, IDLength)
	for {
		rand.Seed(time.Now().UnixNano())
		for i := range b {
			b[i] = letterRunes[rand.Intn(len(letterRunes))]
		}

		id := string(b)
		matched := idRegexp.MatchString(id)

		if !matched {
			continue
		}

		return id
	}
}
