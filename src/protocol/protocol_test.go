package protocol

import (
	"bufio"
	"bytes"
	"io"
	"testing"
)

func Test_ReadCommand(t *testing.T) {
	p := &Protocol{}
	in_str := bytes.NewBufferString("procedure 1\r\narg1\r\n")

	ioReader := bufio.NewReader(in_str)

	req, err := p.ReadCommand(ioReader)
	if err != nil {
		t.Error("Read Command should be non error")
	}

	if req.CommandID != "procedure" || len(req.Arguments) != 1 || req.Arguments[0] != "arg1" {
		t.Error("wrong read request:", *req)
	}
	///////////////////////////////////////////////////////////////////////////////////////////
	in_str = bytes.NewBufferString("procedure\r\nfdafa\r\narg1\r\n")

	ioReader = bufio.NewReader(in_str)

	req, err = p.ReadCommand(ioReader)
	if err != WrongFmtError {
		t.Error("should WrongFmtError here")
	}
	/////////////////////////////////////////////////////////////////////////////////////////////
	in_str = bytes.NewBufferString("")

	ioReader = bufio.NewReader(in_str)

	req, err = p.ReadCommand(ioReader)
	if err != io.EOF {
		t.Error("Should io.EOF here")
	}
}

func Test_ReadResponse(t *testing.T) {
	p := &Protocol{}
	in_str := bytes.NewBufferString("0 1\r\ndata1\r\n")

	ioReader := bufio.NewReader(in_str)

	response, err := p.ReadResponse(ioReader)
	if err != nil {
		t.Error("Read response should be non error")
	}

	if response.ErrNo != StateOK || len(response.Data) != 1 || response.Data[0] != "data1" {
		t.Error("wrong read response:", response)
	}
	///////////////////////////////////////////////////////////////////////////////////////////
	in_str = bytes.NewBufferString("1\r\n")

	ioReader = bufio.NewReader(in_str)

	response, err = p.ReadResponse(ioReader)
	if err != nil {
		t.Error("Read response should be non error")
	}

	if response.ErrNo != StateUnknownCmd || len(response.Data) != 0 {
		t.Error("wrong read response")
	}
	///////////////////////////////////////////////////////////////////////////////////////////////
	in_str = bytes.NewBufferString("")

	ioReader = bufio.NewReader(in_str)

	response, err = p.ReadResponse(ioReader)
	if err != io.EOF {
		t.Error("Should be EOF Error here")
	}
	////////////////////////////////////////////////////////////////////////////////////////////////
	in_str = bytes.NewBufferString("xxxxx")

	ioReader = bufio.NewReader(in_str)

	response, err = p.ReadResponse(ioReader)
	if err != WrongFmtError {
		t.Error("Should be WrongFmtError Error here")
	}

}

func Test_ResponseSerialize(t *testing.T) {
	p := &Protocol{}
	resp := p.CreateResponse()
	resp.ErrNo = StateOK
	resp.Data = []string{"1", "abc"}

	res := resp.Serialize()
	if string(res) != "0 2\r\n1\r\nabc\r\n" {
		t.Error("Wrong Serialize data:", string(res))
	}
}

func Test_CommandSerialize(t *testing.T) {
	p := &Protocol{}
	req := p.CreateCommand()
	req.CommandID = "test_procedure"
	req.Arguments = []string{"1", "abc"}

	res := req.Serialize()
	if string(res) != "test_procedure 2\r\n1\r\nabc\r\n" {
		t.Error("Wrong Serialize data:", string(res))
	}
}
