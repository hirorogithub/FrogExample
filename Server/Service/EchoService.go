package Service

import (
	"../../CommonInterface/MethodsTable"
	"github.com/nats-io/go-nats"

	util "../../CommonInterface"
	echo "../../CommonInterface/echo"
	"context"
	"encoding/json"
	"github.com/golang/protobuf/proto"
	"github.com/yplusplus/frog"
	"log"
	"strings"
	"time"
)

type MyEchoServiceImpl struct {
	conn *nats.Conn
}

var instance *MyEchoServiceImpl

func RunMyEchoService() {
	instance = NewMyEchoServiceImpl()
}

func NewMyEchoServiceImpl() *MyEchoServiceImpl {

	tmpConn, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		//error!!exirintft??
		log.Printf("can not connect to" + "nats://localhost:4222")
		return nil
	}
	baseService := &MyEchoServiceImpl{tmpConn}
	log.Println("Link Start!")

	err = echo.RegisterEchoService(baseService, MethodsTable.RegisterMethods)
	if err != nil {
		log.Fatal(err)
	}

	baseService.listenRPCCall()
	return baseService
}

func (impl *MyEchoServiceImpl) listenRPCCall() {
	//impl.conn.Req
	impl.conn.Subscribe("frog", func(msg *nats.Msg) {

		args := util.RPCCallNatsMsg{}

		err := json.Unmarshal(msg.Data, &args)
		if err != nil {
			log.Println(err)
			return
		}

		//request:=args.Request.(echo.ProtoEchoRequest)

		var rpcMeth *frog.RpcMethod
		for _, meth := range MethodsTable.Methods {
			//should deep compare
			if strings.EqualFold(*meth.Descriptor().Name, *args.Method.Name) {
				rpcMeth = meth
				break
			}
		}
		if rpcMeth == nil {
			log.Println(err)
			return
		}
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		var response echo.ProtoEchoResponse
		err = frog.CallMethod(rpcMeth, ctx, &args.Request, &response)

		if err != nil {
			log.Println(err)
		}

		args.Response = response
		res, err := json.Marshal(args)
		impl.conn.Publish(msg.Reply, res)

	})
}

func (impl *MyEchoServiceImpl) Echo(ctx context.Context, in *echo.ProtoEchoRequest, out *echo.ProtoEchoResponse) error {
	out.Text = proto.String(*in.Text)
	return nil
}

func (impl *MyEchoServiceImpl) Echo2(ctx context.Context, in *echo.ProtoEchoRequest, out *echo.ProtoEchoResponse) error {
	out.Text = proto.String(*in.Text)
	return nil
}
