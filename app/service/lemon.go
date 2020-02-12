package service

import (
	"context"
	"fmt"
	"goformat/app/global"
	"goformat/app/global/helper"
	"goformat/library/errorcode"
	"goformat/library/rpc/rpclemon"
	"sync"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// LemonSer lemon API
type LemonSer struct {
}

var lemonSingleton *LemonSer
var lemonOnce sync.Once

// LemonIns 獲得Rotate對象
func LemonIns() *LemonSer {
	lemonOnce.Do(func() {
		lemonSingleton = &LemonSer{}
	})
	return lemonSingleton
}

// GateWayServiceRegister gateway 服務註冊
func GateWayServiceRegister(project string) (output string, apiErr errorcode.Error) {
	// grpe server connect
	conn, apiErr := helper.GrpcServerConnect(global.Config.API.LemonGrpcServer)
	if apiErr != nil {
		return
	}
	defer conn.Close()

	// 連線超時5秒deadline
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Duration(5*time.Second)))
	defer cancel()

	// 呼叫 lemon grpc 初始化
	c := rpclemon.NewServiceClient(conn)
	resp, e := c.ServiceLists(ctx, &rpclemon.ListRequest{})
	if e != nil {
		statusErr, ok := status.FromError(e)
		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				apiErr = helper.ErrorHandle(global.FatalLog, "GOFORMAT_GET_GRPC_RESPONSE_DEADLINE", e.Error())
				return
			}
		}
		apiErr = helper.ErrorHandle(global.FatalLog, "GOFORMAT_GET_GRPC_RESPONSE_FAILED", e.Error())
		return
	}

	var out []string
	for _, val := range resp.Lists {
		out = append(out, val.Name)
	}
	// validate project name req is in array
	errArr := helper.InArray(project, out)
	// 檢查服務是否在 不存在註冊
	if errArr {
		output = fmt.Sprintf(">>---------------service already register---------------<<")

	} else {
		// register pitaya service
		str := global.Config.GrpcSetting.Name + ":service"
		md5Token := helper.Md5Encryption(str)
		req := rpclemon.RegisterRequest{
			Name:  project,
			Token: md5Token,
			Path:  global.Config.GrpcSetting.Path,
		}

		resp, err := c.ServiceRegister(context.Background(), &req)
		if err != nil {
			apiErr = helper.ErrorHandle(global.WarnLog, "GOFORMAT_GET_GRPC_RESPONSE_FAILED", err.Error())
			return
		}
		output = fmt.Sprintf(">>---------------service register successful!---------------<<:%v", resp)

	}

	return
}
