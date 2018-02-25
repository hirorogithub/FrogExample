package Service

import (
	"github.com/nats-io/go-nats"

	util "../../util"
	echo "../../util/echo"
	"context"
	"encoding/json"
	"github.com/golang/protobuf/proto"
	"github.com/yplusplus/frog"
	"log"
	"time"
	"github.com/pkg/errors"
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

func (impl * MyEchoServiceImpl)errReply(subj string,args util.RPCHeader,err error){
	log.Println(err)
	args.StatusCode = util.CallLogicErr
	args.Msg = nil
	impl.conn.Publish(subj,args.ToByte())
	return
}

func (impl *MyEchoServiceImpl) listenRPCCall() {

	//registe listener by service name
	subj := *echo.EchoService_ServiceDesc.Name
	impl.conn.Subscribe(subj, func(msg *nats.Msg) {

		//get args
		args := util.RPCHeader{}
		err := json.Unmarshal(msg.Data, &args)
		if err != nil {
			impl.errReply(msg.Reply,args,err)
			return
		}

		//get rpcmeth
		rpcMeth, exist := methodTable[(args.MethodName)]
		if !exist {
			impl.errReply(msg.Reply,args,err)
			return
		}


		//reflect response/request to real type
		response := rpcMeth.NewResponse()
		request := rpcMeth.NewRequest()
		if !args.GetAsRequest(request){
			impl.errReply(msg.Reply,args,errors.New("reflect to request faild"))
			return
		}

		log.Printf("Request \t%s: %+v",rpcMeth.Name(),request)

		//!!here I want to pass ctx as nil,but some panic happend
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		err = frog.CallMethod(rpcMeth, ctx, request, response)
		if err != nil {
			impl.errReply(msg.Reply,args,err)
			return
		}

		args.SetMsg(response)
		args.StatusCode= util.CallOK
		res, _ := json.Marshal(args)
		impl.conn.Publish(msg.Reply, res)

		log.Printf("Response\t%s: %+v",rpcMeth.Name(),response)


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
