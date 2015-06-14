package main

import (
	"flag"
	"fmt"
	"github.com/nporsche/np-golang-logging"
	"log/syslog"
	"os"
	"protocol"
)

var log = logging.MustGetLogger("main")

func main() {
	pPort := flag.Int("port", 9090, "The listen port")
	logLevelStr := flag.String("log", "WARNING", "The log level CRITICAL,ERROR, WARNING,NOTICE,INFO,DEBUG")

	flag.Parse()

	backend, err := logging.NewSyslogBackendPriority("bean", syslog.LOG_LOCAL3)
	if err != nil {
		fmt.Printf("logging init error=[%s]", err.Error())
		os.Exit(1)
	}
	format := logging.MustStringFormatter("%{color}[%{module}.%{shortfunc}][%{level:.4s}]%{color:reset}%{message}")
	logging.SetBackend(logging.NewBackendFormatter(backend, format))

	logLevel, err := logging.LogLevel(*logLevelStr)
	if err == nil {
		logging.SetLevel(logLevel, "main")
		logging.SetLevel(logLevel, "protocol")
	}

	processor := NewRequestProcessor(NewRequestHandler(*path))
	svr := NewServer(*pUnixSock, &protocol.Protocol{}, processor)
	svr.Start()
	return
}
