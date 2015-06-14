package protocol

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
)

const (
	LINE_SIZE = 512
)

const (
	StateOK           = 0
	StateUnknownCmd   = 1
	StateCmdException = 2
)

var WrongFmtError = errors.New("Invalid Format")

//=========================================================================
/*
command argCount\r\n
arg1\r\n
arg2\r\n
*/
type Command struct {
	CommandID string
	Arguments []string
}

/*
state replyCount\r\n
reply1\r\n
reply2\r\n
*/

type Response struct {
	ErrNo int16
	Data  []string
}

type Event struct {
	EventID string
	Data    []string
}

type Protocol struct {
}

func (this *Protocol) CreateResponse() *Response {
	return new(Response)
}

func (this *Protocol) CreateCommand() *Command {
	return new(Command)
}

func (this *Protocol) CreateEvent() *Event {
	return new(Event)
}

func (this *Event) Serialize() []byte {
	buf := new(bytes.Buffer)
	buf.WriteString(fmt.Sprintf("%s %d\r\n", this.EventID, len(this.Data)))
	for _, d := range this.Data {
		buf.WriteString(fmt.Sprintf("%s\r\n", d))
	}
	return buf.Bytes()
}

func (this *Command) Serialize() []byte {
	buf := new(bytes.Buffer)
	buf.WriteString(fmt.Sprintf("%s %d\r\n", this.CommandID, len(this.Arguments)))
	for _, d := range this.Arguments {
		buf.WriteString(fmt.Sprintf("%s\r\n", d))
	}
	return buf.Bytes()
}

func (this *Response) Serialize() []byte {
	buf := new(bytes.Buffer)
	buf.WriteString(fmt.Sprintf("%d %d\r\n", this.ErrNo, len(this.Data)))
	for _, d := range this.Data {
		buf.WriteString(fmt.Sprintf("%s\r\n", d))
	}
	return buf.Bytes()
}

func (this *Protocol) ReadResponse(reader *bufio.Reader) (*Response, error) {
	resp := &Response{StateOK, []string{}}
	if n, err := fmt.Fscanf(reader, "%d\r\n", &resp.ErrNo); err == io.EOF {
		return nil, err
	} else if n != 1 {
		return nil, WrongFmtError
	}

	if resp.ErrNo == StateOK {
		var dataCount int
		if n, err := fmt.Fscanf(reader, "%d\r\n", &dataCount); err == io.EOF {
			return nil, err
		} else if n != 1 {
			return nil, WrongFmtError
		}

		for i := 0; i < dataCount; i++ {
			line, _, err := reader.ReadLine()
			if err != nil {
				return nil, WrongFmtError
			}
			resp.Data = append(resp.Data, string(line))
		}
	}
	return resp, nil
}

func (this *Protocol) ReadCommand(reader *bufio.Reader) (*Command, error) {
	req := &Command{"", []string{}}
	if n, err := fmt.Fscanf(reader, "%s\r\n", &req.CommandID); err == io.EOF {
		return nil, err
	} else if n != 1 {
		return nil, WrongFmtError
	}

	var argCount int
	if n, err := fmt.Fscanf(reader, "%d\r\n", &argCount); err == io.EOF {
		return nil, err
	} else if n != 1 {
		return nil, WrongFmtError
	}

	for i := 0; i < argCount; i++ {
		line, _, err := reader.ReadLine()
		if err != nil {
			return nil, WrongFmtError
		}
		req.Arguments = append(req.Arguments, string(line))
	}
	return req, nil
}
