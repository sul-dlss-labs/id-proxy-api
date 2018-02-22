package druid

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const truncatedAlpha string = "bcdfghjkmnpqrstvwxyz"

// Generate a random DRUID
func Generate() Druid {
	rand.Seed(time.Now().UTC().UnixNano())
	idArr := []string{randomChar(),
		randomChar(),
		randomInt(),
		randomInt(),
		randomInt(),
		randomChar(),
		randomChar(),
		randomInt(),
		randomInt(),
		randomInt(),
		randomInt(),
	}
	return Druid(strings.Join(idArr, ""))
}

func randomChar() string {
	rand := rand.Intn(len(truncatedAlpha))
	return string(truncatedAlpha[rand])
}

func randomInt() string {
	return strconv.Itoa(rand.Intn(9))
}
