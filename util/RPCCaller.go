package util

import (
	"encoding/json"
	"errors"
	"github.com/gogo/protobuf/proto"
	"github.com/yplusplus/frog"
	"time"
)

type RPCCallNatsMsg struct {
	Method   *frog.MethodDesc
	Request  interface{}
	Response interface{}
}

const (
	CallTimeOut  = -1
	CallOK       = 0
	CallLogicErr = 1
)

type RPCHeader struct {
	ServiceName string
	MethodName  string
	Msg         []byte
	StatusCode  int
	TimeStamp   string
}

func NewRPCHeader(service, method string) *RPCHeader {
	return &RPCHeader{
		ServiceName: service, MethodName: method, Msg: nil,
		StatusCode: CallOK, TimeStamp: time.Now().Format(DefaultTimeFmt),
	}
}

func (this *RPCHeader) SetMsg(data interface{}) bool {
	bdata, err := json.Marshal(data)
	if err != nil {
		return false
	} else {
		this.Msg = bdata
		return true
	}
}

func (this *RPCHeader) GetAsRequest(request proto.Message) bool {
	err := json.Unmarshal(this.Msg, request)
	if err != nil {
		return false
	} else {
		return true
	}
}

func (this *RPCHeader) GetAsResponse(response proto.Message) bool {
	return this.GetAsRequest(response)
}

func (this RPCHeader) ToByte() []byte {
	res, err := json.Marshal(this)
	if err != nil {
		return nil
	} else {
		return res
	}
}

func (this *RPCHeader) CheckStatus() error {
	switch this.StatusCode {
	case CallOK:
		return nil
	case CallLogicErr:
		{
			err := errors.New("some logic err happend")
			return err
		}
	case CallTimeOut:
		{
			err := errors.New("rpccall timeout")
			return err
		}
	default:
		{
			err := errors.New("got unknown status code")
			return err
		}
	}
}
