package client

import (
	"bufio"
	"log"
	"net"

	"github.com/Vlad-Pisarevskiy/faraway/internal/pow"
	"github.com/Vlad-Pisarevskiy/faraway/internal/protocol"
)

type Client struct{}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) Run(addr string) error {

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}
	defer conn.Close()

	reader := bufio.NewReader(conn)
	data, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	challenge, difficulty, err := protocol.ParseChallenge(data)
	if err != nil {
		return err
	}

	solution := pow.Solve(challenge, difficulty)

	_, err = conn.Write([]byte(protocol.BuildSolution(solution)))
	if err != nil {
		return err
	}

	data, err = reader.ReadString('\n')
	if err != nil {
		return err
	}

	quote, err := protocol.ParseQuote(data)
	if err != nil {
		return err
	}

	log.Printf("Цитата получена успешно: %s", quote)
	return nil
}
