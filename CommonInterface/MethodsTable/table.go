package MethodsTable

import "github.com/yplusplus/frog"

var (
    Methods = make([]*frog.RpcMethod, 0)
)

func RegisterMethods(method []*frog.RpcMethod) error {
    Methods = append(Methods, method...)
    return nil
}
