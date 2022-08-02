package interview_thrift_demo

import (
	"context"
	"fmt"
	"strconv"
)

type UserServiceHandler struct {

}

func (u *UserServiceHandler) Add(ctx context.Context, num1 int32, num2 string) (_r int32, _err error) {
	fmt.Println("调用Add的实现")
	num2int32, _ := strconv.Atoi(num2)
	return num1 + int32(num2int32), nil
}