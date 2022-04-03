package api

import (
	"k8s-dashboard-server/internal/code"
)

type BusinessResponse struct {
	Code int         `json:"code"` // 业务码
	Data interface{} `json:"data"` // 返回数据
}

func Success(data interface{}) BusinessResponse {
	return BusinessResponse{
		Code: code.Success,
		Data: data,
	}
}

func Response(code int, data interface{}) BusinessResponse {
	return BusinessResponse{
		Code: code,
		Data: data,
	}
}
