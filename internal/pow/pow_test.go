package pow

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSolveAndVerify(t *testing.T) {

	challenge := "4ad79ca77b123"
	difficulty := 4

	result := Solve(challenge, difficulty)

	require.NotEmpty(t, result)
	assert.True(t, Verify(challenge, result, difficulty))
}
