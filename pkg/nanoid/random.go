package nanoid

import (
	"math/rand"
	"time"
)

const (
	AlphabetLower       = "abcdefghijklmnopqrstuvwxyz"
	AlphabetUpper       = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	AlphaNum            = "0123456789"
	AlphabetLowerAndNum = AlphabetLower + AlphaNum
	AlphabetAndNum      = AlphabetLower + AlphaNum + AlphabetUpper
)

func Random(length int) string {
	alphabetLength := len(AlphabetAndNum)
	haystacks := make([]byte, length)
	for i := 0; i < length; i++ {
		rand.Seed(time.Now().UnixNano())
		idx := rand.Intn(alphabetLength - 1)
		haystacks[i] = AlphabetLowerAndNum[idx]
	}

	return string(haystacks)
}

func RandomMixed(length int) string {
	alphabetLength := len(AlphabetAndNum)
	haystacks := make([]byte, length)
	for i := 0; i < length; i++ {
		rand.Seed(time.Now().UnixNano())
		idx := rand.Intn(alphabetLength - 1)
		haystacks[i] = AlphabetLower[idx]
	}

	return string(haystacks)
}

func RandomAlphabet(alphabet string, length int) string {
	alphabetLength := len(alphabet)
	haystacks := make([]byte, length)
	for i := 0; i < length; i++ {
		rand.Seed(time.Now().UnixNano())
		idx := rand.Intn(alphabetLength - 1)
		haystacks[i] = alphabet[idx]
	}

	return string(haystacks)
}
