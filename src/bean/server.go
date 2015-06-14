package main

import (
	"bufio"
	"net"
	"os"
	"os/signal"
	"protocol"
	"syscall"
)

type Server struct {
	proto     *protocol.Protocol
	processor *RequestProcessor
	unixSock  string
}

func NewServer(unixSock string, proto *protocol.Protocol, processor *RequestProcessor) *Server {
	this := new(Server)
	this.proto = proto
	this.processor = processor
	this.unixSock = unixSock

	return this
}

func (this *Server) Start() {
	os.Remove(this.unixSock)

	ls, err := net.Listen("unix", this.unixSock)
	if err != nil {
		log.Critical("unix sock listen error=[%s]", err.Error())
		return
	}

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, os.Kill, syscall.SIGTERM)
	go func(c chan os.Signal) {
		// Wait for a SIGINT or SIGKILL:
		<-c
		// Stop listening (and unlink the socket if unix type):
		ls.Close()
		// And we're done:
		os.Exit(0)
	}(sigc)

	log.Info("Listening unix_sock=[%s]", this.unixSock)

	for {
		conn, err := ls.Accept()
		if err != nil {
			log.Fatal("accept error:", err)
		}
		go this.handleConnection(conn)
	}

}

func (this *Server) handleConnection(conn net.Conn) {
	defer func() {
		if r := recover(); r != nil {
			log.Critical("UnExpected Fatal %v", r)
		}
		conn.Close()
		log.Debug("connection leaves")
	}()
	log.Debug("A new connection comes")
	io := bufio.NewReaderSize(conn, protocol.LINE_SIZE)
	for {
		in, err := this.proto.ReadRequest(io)
		if err != nil && err != protocol.WrongFmtError {
			break
		}

		out := this.proto.CreateResponse()
		this.processor.Process(in, out)

		_, err = conn.Write(out.Serialize())
		if err != nil {
			break
		}
	}
}
