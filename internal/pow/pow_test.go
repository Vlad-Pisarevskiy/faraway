package pow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolveAndVerify(t *testing.T) {

	challenge := "4ad79ca77b123"
	difficulty := 4

	result := Solve(challenge, difficulty)

	assert.True(t, Verify(challenge, result, difficulty))
	assert.False(t, Verify(challenge, 0, 256))
}

func TestTableDrivenSolve(t *testing.T) {

	values := []struct {
		name string
		hash [32]byte
		want int
	}{
		{"4 нуля", [32]byte{0x0A}, 4},
		{"0 нулей", [32]byte{0x80}, 0},
		{"8 нулей", [32]byte{0x00, 0xA0}, 8},
	}

	for _, v := range values {

		t.Run(v.name, func(t *testing.T) {
			assert.Equal(t, countingLeadingZeroBits(v.hash), v.want)
		})
	}
}
