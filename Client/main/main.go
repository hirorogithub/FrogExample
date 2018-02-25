package main

import ("time"
    "../Channel"
    echo "../../util/echo"
    "github.com/gogo/protobuf/proto"
    "log"
    "context"
)



func main(){
    // create stub
    channel := Channel.NewMyChannel()
    stub := echo.NewEchoServiceStub(channel)
    var request echo.ProtoEchoRequest
    var response echo.ProtoEchoResponse
    request.Text = proto.String("Hello, world!")

    // make a sync rpc
    log.Println("begin a sync rpc")
    ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
    err := stub.Echo(ctx, &request, &response)
    if err != nil {
        log.Println(err)
    } else if request.GetText() != response.GetText() {
        log.Println("Text not match:", request.GetText(), "vs", response.GetText())
    } else {
        log.Println(request.GetText(), response.GetText())
    }
    log.Println("end a sync rpc")
}