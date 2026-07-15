package protocol

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProtocolNonce(t *testing.T) {

	nonce := 45

	solution := BuildSolution(nonce)
	s, err := ParseSolution(solution)

	assert.Nil(t, err)
	assert.Equal(t, 45, s)
}

func TestProtocolChallenge(t *testing.T) {

	hexValue := "1a52bc32"
	challengeDifficulty := 4

	challenge := BuildChallenge(hexValue, challengeDifficulty)

	ch, difficulty, err := ParseChallenge(challenge)
	assert.Nil(t, err)
	assert.Equal(t, challengeDifficulty, difficulty)
	assert.Equal(t, hexValue, ch)
}

func TestProtocolQuote(t *testing.T) {

	quote := "some quote"

	buildQuote := BuildQuote(quote)

	q, err := ParseQuote(buildQuote)
	assert.Nil(t, err)
	assert.Equal(t, quote, q)
}

func TestProtocolError(t *testing.T) {

	var err string = "some error"

	buildErr := BuildError(err)
	getError, e := ParseError(buildErr)
	assert.Equal(t, err, getError)
	assert.Nil(t, e)
}
