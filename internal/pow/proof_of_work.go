package pow

import (
	"crypto/sha256"
	"math/bits"
	"strconv"
)

func Verify(challenge string, nonce int, difficulty int) bool {

	data := challenge + strconv.Itoa(nonce)
	value := sha256.Sum256([]byte(data))

	total := 0

	for _, b := range value {

		if bits.LeadingZeros8(b) == 8 {
			total += 8
		} else {
			total += bits.LeadingZeros8(b)
			return total >= difficulty
		}
	}

	return total >= difficulty
}

func Solve(challenge string, difficulty int) int {

	i := 0

	for {
		if Verify(challenge, i, difficulty) {
			return i
		}
		i++
	}
}
