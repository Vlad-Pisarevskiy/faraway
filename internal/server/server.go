package server

import (
	"bufio"
	"crypto/rand"
	"encoding/hex"
	"log"
	"net"

	"github.com/Vlad-Pisarevskiy/faraway/internal/pow"
	"github.com/Vlad-Pisarevskiy/faraway/internal/protocol"
)

const baseDifficulty = 4

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Run() error {

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		return err
	}
	defer ln.Close()

	for {
		var conn net.Conn
		conn, err = ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go s.handle(conn)
	}
}

func (s *Server) handle(conn net.Conn) {

	defer conn.Close()

	writer := bufio.NewWriter(conn)
	var byteChallenge [16]byte
	_, _ = rand.Read(byteChallenge[:])

	challenge := hex.EncodeToString(byteChallenge[:])
	_, err := writer.Write([]byte(protocol.BuildChallenge(challenge, baseDifficulty)))
	if err != nil {
		log.Println(err)
		return
	}

	reader := bufio.NewReader(conn)
	data, err := reader.ReadString('\n')
	if err != nil {
		log.Println(err)
		return
	}

	solution, err := protocol.ParseSolution(data)
	if err != nil {
		log.Println(err)
		return
	}

	if ok := pow.Verify(challenge, solution, baseDifficulty); !ok {
		stringErr := protocol.BuildError("bad solution")
		_, err = writer.Write([]byte(stringErr))
		if err != nil {
			log.Println(err)
			return
		}

		log.Println(stringErr)
		return
	}

}
