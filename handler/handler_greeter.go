package handler

import (
	"context"
	"interview/gen-go/Sample"
)

type GreeterServiceHandler struct {

}

func (g *GreeterServiceHandler) SayHello(ctx context.Context, user *Sample.User) (_r *Sample.Response, _err error) {
	return
}

func (g *GreeterServiceHandler) GetUser(ctx context.Context, uid int32) (_r *Sample.Response, _err error) {
	return
}