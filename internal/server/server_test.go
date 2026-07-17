package server

import (
	"bufio"
	"net"
	"testing"
	"time"

	"github.com/Vlad-Pisarevskiy/faraway/internal/pow"
	"github.com/Vlad-Pisarevskiy/faraway/internal/protocol"
	"github.com/Vlad-Pisarevskiy/faraway/internal/quotes"
	"github.com/stretchr/testify/require"
)

func TestServer_HappyPath(t *testing.T) {

	quoter := quotes.NewQuoter()
	server := NewServer(quoter)

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	defer ln.Close()
	require.NoError(t, err)

	go server.Serve(ln)

	conn, err := net.Dial("tcp", ln.Addr().String())
	defer conn.Close()
	require.NoError(t, err)

	err = conn.SetReadDeadline(time.Now().Add(time.Minute))
	require.NoError(t, err)

	reader := bufio.NewReader(conn)
	data, err := reader.ReadString('\n')
	require.NoError(t, err)
	challenge, difficulty, err := protocol.ParseChallenge(data)
	require.NoError(t, err)

	solution := pow.Solve(challenge, difficulty)

	_, err = conn.Write([]byte(protocol.BuildSolution(solution)))
	require.NoError(t, err)

	bytesData, err := reader.ReadString('\n')
	require.NoError(t, err)
	quote, err := protocol.ParseQuote(bytesData)
	require.NoError(t, err)
	t.Log(quote)
}
