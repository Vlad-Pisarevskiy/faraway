package pow

import (
	"crypto/sha256"
	"math/bits"
	"strconv"
)

const zeroByte = 8
const startNonce = 0

func Verify(challenge string, nonce int, difficulty int) bool {

	data := challenge + strconv.Itoa(nonce)
	value := sha256.Sum256([]byte(data))

	total := countingLeadingZeroBits(value)

	return total >= difficulty
}

func Solve(challenge string, difficulty int) int {

	i := startNonce

	for {
		if Verify(challenge, i, difficulty) {
			return i
		}
		i++
	}
}

func countingLeadingZeroBits(value [32]byte) int {

	total := 0

	for _, b := range value {

		lz := bits.LeadingZeros8(b)
		if lz == zeroByte {
			total += zeroByte
		} else {
			total += lz
			return total
		}
	}

	return total
}
