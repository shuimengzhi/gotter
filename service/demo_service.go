package service

import "gotter/request_struct"

func DemoService(param request_struct.DemoRequest) ResultService {
	return ResultService{Code: SUCCESS, Data: "请求内容是" + param.Account}
}
