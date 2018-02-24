package CommonInterface

import (
	"github.com/yplusplus/frog"
	echo "../CommonInterface/echo"
)

type RPCCallNatsMsg struct {
	Method   *frog.MethodDesc
	Request  echo.ProtoEchoRequest
	Response echo.ProtoEchoResponse
}
