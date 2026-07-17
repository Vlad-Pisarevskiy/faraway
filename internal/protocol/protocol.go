package protocol

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	CHALLENGE          = "CHALLENGE"
	challengeArguments = 3
	SOLUTION           = "SOLUTION"
	solutionArguments  = 2
	QUOTE              = "QUOTE"
	quoteArguments     = 2
	ERROR              = "ERROR"
	errorArguments     = 2
)

func BuildChallenge(challenge string, difficulty int) string {

	return fmt.Sprintf("%s %s %d\n", CHALLENGE, challenge, difficulty)
}

func BuildSolution(nonce int) string {

	return fmt.Sprintf("%s %d\n", SOLUTION, nonce)
}

func BuildQuote(quote string) string {

	return fmt.Sprintf("%s %s\n", QUOTE, quote)
}

func BuildError(err string) string {

	return fmt.Sprintf("%s %s\n", ERROR, err)
}

func ParseSolution(solution string) (int, error) {

	solution = strings.TrimSpace(solution)
	values := strings.SplitN(solution, " ", solutionArguments)
	if values[0] != SOLUTION {
		return 0, errors.New("uncorrected request")
	}

	if len(values) != solutionArguments {
		return 0, errors.New("get bad solution")
	}

	nonce := values[1]
	result, err := strconv.Atoi(nonce)
	if err != nil {
		return 0, err
	}

	return result, nil
}

func ParseChallenge(challenge string) (string, int, error) {

	challenge = strings.TrimSpace(challenge)
	values := strings.SplitN(challenge, " ", challengeArguments)
	if len(values) != challengeArguments {
		return "", 0, errors.New("get bad challenge")
	}

	difficulty, err := strconv.Atoi(values[2])
	if err != nil {
		return "", 0, err
	}

	return values[1], difficulty, nil

}

func ParseQuote(quote string) (string, error) {

	quote = strings.TrimSpace(quote)
	values := strings.SplitN(quote, " ", quoteArguments)

	if len(values) != quoteArguments {
		return "", errors.New("error: get bad quote")
	}

	return values[1], nil
}

func ParseError(err string) (string, error) {

	err = strings.TrimSpace(err)
	values := strings.SplitN(err, " ", errorArguments)
	if len(values) != errorArguments {
		return "", errors.New("error: get bad error")
	}

	return values[1], nil
}
