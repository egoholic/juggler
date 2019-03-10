package uuid

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var randomizer *rand.Rand

func init() {
	source := rand.NewSource(time.Now().Unix())
	randomizer = rand.New(source)
}

var chars = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
var length = len(chars)

func New() string {
	uuid := &strings.Builder{}
	uuid.WriteString(strconv.Itoa(int(time.Now().Unix())))
	uuid.WriteRune('-')
	for i := 0; i < 16; i++ {
		uuid.WriteRune(chars[randomizer.Intn(length)])
	}

	return uuid.String()
}
