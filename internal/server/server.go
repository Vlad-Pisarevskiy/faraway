package server

import (
	"bufio"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"
	"log"
	"net"
	"time"

	"github.com/Vlad-Pisarevskiy/faraway/internal/pow"
	"github.com/Vlad-Pisarevskiy/faraway/internal/protocol"
	"github.com/Vlad-Pisarevskiy/faraway/internal/quotes"
)

const deadlineSeconds = 30

type Server struct {
	quoter     *quotes.Quoter
	difficulty int
}

func NewServer(quoter *quotes.Quoter, difficulty int) *Server {

	return &Server{
		quoter:     quoter,
		difficulty: difficulty,
	}
}

func (s *Server) Run(addr string) error {

	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	defer ln.Close()

	return s.Serve(ln)
}

func (s *Server) Serve(ln net.Listener) error {

	for {

		conn, err := ln.Accept()

		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				return nil
			}

			log.Println(err)
			continue
		}

		go s.handle(conn)
	}
}

func (s *Server) handle(conn net.Conn) {

	err := conn.SetDeadline(time.Now().Add(deadlineSeconds * time.Second))
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	var byteChallenge [16]byte
	_, _ = rand.Read(byteChallenge[:])

	challenge := hex.EncodeToString(byteChallenge[:])
	_, err = conn.Write([]byte(protocol.BuildChallenge(challenge, s.difficulty)))
	if err != nil {
		log.Println(err)
		return
	}

	reader := bufio.NewReader(conn)
	data, err := reader.ReadString('\n')
	if errors.Is(err, io.EOF) {
		return
	}
	if err != nil {
		log.Println(err)
		return
	}

	solution, err := protocol.ParseSolution(data)
	if err != nil {
		log.Println(err)
		return
	}

	if ok := pow.Verify(challenge, solution, s.difficulty); !ok {
		stringErr := protocol.BuildError("bad solution")
		_, err = conn.Write([]byte(stringErr))
		if err != nil {
			log.Println(err)

			return
		}

		log.Println(stringErr)
		return
	}

	_, err = conn.Write([]byte(protocol.BuildQuote(s.quoter.Random())))
	if err != nil {
		log.Println(err)
		return
	}

}
