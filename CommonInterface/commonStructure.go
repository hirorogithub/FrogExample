package CommonInterface

import (
	"github.com/yplusplus/frog"
)

type RPCCallNatsMsg struct {
	Method   *frog.MethodDesc
	Request  interface{}
	Response interface{}
}
