package Service

import (
	"github.com/nats-io/go-nats"

	util "../../CommonInterface"
	echo "../../CommonInterface/echo"
	"context"
	"encoding/json"
	"github.com/golang/protobuf/proto"
	"github.com/yplusplus/frog"
	"log"
	"time"
)

var (
	methodTable = make(map[string]*frog.RpcMethod)
)

func registerMethods(methods []*frog.RpcMethod) error {
	for _, method := range methods {
		methodTable[*(method.Descriptor().Name)] = method
	}
	return nil

}

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

	err = echo.RegisterEchoService(baseService, registerMethods)
	if err != nil {
		log.Fatal(err)
	}

	baseService.listenRPCCall()
	return baseService
}

func (impl *MyEchoServiceImpl) listenRPCCall() {

	//registe listener by service name
	subj := *echo.EchoService_ServiceDesc.Name
	impl.conn.Subscribe(subj, func(msg *nats.Msg) {

		//get args
		args := util.RPCCallNatsMsg{}
		err := json.Unmarshal(msg.Data, &args)
		if err != nil {
			log.Println(err)
			return
		}

		//get rpcmeth
		rpcMeth, exist := methodTable[*(args.Method.Name)]
		if !exist {
			log.Println(err)
			return
		}

		//reflect response/request to real type
		args.Response = rpcMeth.NewResponse()

		request := rpcMeth.NewRequest()
		ok := util.MapToStruct(args.Request.(map[string]interface{}), &request)
		if !ok {
			log.Println("map to struct fail ")
		}
		args.Request = request

		//call the rpc func
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		err = frog.CallMethod(rpcMeth, ctx, args.Request.(proto.Message), args.Response.(proto.Message))

		if err != nil {
			log.Println(err)
		}

		//marshal the args and reply
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
