package Channel

import (
	util "../../CommonInterface"
	echo "../../CommonInterface/echo"
	"context"
	"encoding/json"
	"errors"
	"github.com/golang/protobuf/proto"
	"github.com/nats-io/go-nats"
	"github.com/yplusplus/frog"
	"log"
	"sync"
	"time"
)

type Call struct {
	mu       *sync.Mutex
	request  proto.Message
	response proto.Message
	ch       chan struct{}
	err      error
	hasDone  bool
}

func (c *Call) Request() proto.Message {
	return c.request
}

func (c *Call) Response() proto.Message {
	return c.response
}

func (c *Call) Error() error {
	return c.err
}

func (c *Call) Done() chan struct{} {
	return c.ch
}

func (c *Call) done(err error) {
	c.mu.Lock()
	if !c.hasDone {
		c.hasDone = true
		c.err = err
		close(c.ch)
	}
	c.mu.Unlock()
}

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
	call := &Call{
		new(sync.Mutex),
		request,
		response,
		make(chan struct{}),
		nil,
		false,
	}
	//var rpcMeth *frog.RpcMethod
	//for _, meth := range MethodsTable.Methods {
	//	if meth.Descriptor() == method {
	//		rpcMeth = meth
	//		break
	//	}
	//}
	//if rpcMeth == nil {
	//	call.done(errors.New("method not found"))
	//	return call
	//}

	jdata, _ := json.Marshal(util.RPCCallNatsMsg{
		Method:   method,
		Request:  *request.(*echo.ProtoEchoRequest),
		Response: *response.(*echo.ProtoEchoResponse),
	})

	msg, err := this.conn.Request("frog", jdata, time.Second*5)
	if err != nil {
		log.Println(err)
		call.done(errors.New("request faild"))
		return call
	}

	res := util.RPCCallNatsMsg{}
	err = json.Unmarshal(msg.Data, &res)
	if err != nil {
		log.Println(err)
		call.done(errors.New("request faild"))
		return call
	}

	//should deep copy
	//p := &(res.Response)
	response.(*echo.ProtoEchoResponse).Text = res.Response.Text

	call.done(nil)
	// async invoke
	//go func() {
	//
	//	//publish my request
	//	time.Sleep(time.Second * 5)
	//	err := frog.CallMethod(rpcMeth, ctx, request, response)
	//	call.done(err)
	//}()

	if _, ok := ctx.Deadline(); ok {
		go func() {
			select {
			case <-call.Done():
				// already done, do nothing
			case <-ctx.Done():
				call.done(errors.New("request timeout"))
			}
		}()
	}

	return call
}
