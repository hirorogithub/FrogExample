package Channel

import (
	util "../../util"
	"context"
	"encoding/json"
	"errors"
	"github.com/golang/protobuf/proto"
	"github.com/nats-io/go-nats"
	"github.com/yplusplus/frog"
	"log"
	"time"
)

type MyChannel struct {
	conn *nats.Conn
}

func NewMyChannel() *MyChannel {
	tmpConn, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		//error!!exirintft??
		log.Printf("can not connect to" + "nats://localhost:4222")
		return nil
	}

	channel := &MyChannel{tmpConn}
	return channel

}

func (this *MyChannel) Go(method *frog.MethodDesc, ctx context.Context, request proto.Message, response proto.Message) frog.RpcCall {
	call := frog.NewDefaultCall(request, response)

	serviceName := method.GetServiceDesc().GetName()
	methodName := method.GetName()

	header := util.NewRPCHeader(serviceName, methodName)
	header.SetMsg(request)

	//should async
	go func() {
		msg, err := this.conn.Request(serviceName, header.ToByte(), time.Second*5)
		if err != nil {
			log.Println(err)
			call.Close(err)
			return
		}

		//unmarshal the args
		args := util.RPCHeader{}
		err = json.Unmarshal(msg.Data, &args)
		if err != nil {
			log.Println(err)
			call.Close(err)
			return
		}

		//check the status
		if err = args.CheckStatus(); err != nil {
			log.Println(err)
			call.Close(err)
			return
		}

		//reflect to response
		if !args.GetAsResponse(response) {
			log.Println("Unmarshal fail ")
			call.Close(err)
			return
		}

		call.Close(nil)
	}()

	if _, ok := ctx.Deadline(); ok {
		go func() {
			select {
			case <-call.Done():
				// already Close, do nothing
			case <-ctx.Done():
				call.Close(errors.New("request timeout"))
			}
		}()
	}

	return call
}
