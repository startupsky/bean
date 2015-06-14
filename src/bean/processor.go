package main

import (
	"protocol"
)

type handlerProc func(args []string) (data []string, err error)

type RequestProcessor struct {
	handler          *RequestHandler
	proceduerMapping map[string]handlerProc
}

func NewRequestProcessor(handler *RequestHandler) *RequestProcessor {
	this := new(RequestProcessor)
	this.handler = handler
	this.proceduerMapping = make(map[string]handlerProc)

	//Add More proceduere there
	this.proceduerMapping["ListAuthenticator"] = this.handler.ListAuthenticator
	this.proceduerMapping["IP"] = this.handler.IP
	this.proceduerMapping["ScoreIP"] = this.handler.ScoreIP
	this.proceduerMapping["Timeout"] = this.handler.Timeout
	//end adding proceduere

	return this
}

func (this *RequestProcessor) Process(in *protocol.Request, out *protocol.Response) {
	if in == nil {
		out.State = protocol.StateUnknownCmd
		return
	}

	pro, ok := this.proceduerMapping[in.Procedure]
	if !ok {
		out.State = protocol.StateUnknownCmd
		return
	}

	var err error
	out.Data, err = pro(in.Arguments)
	if err != nil {
		out.State = protocol.StateProcedureException
	}
}
