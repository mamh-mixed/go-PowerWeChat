package base

import (
	"github.com/ArtisanCloud/go-wechat/src/kernel"
)

type Client struct {
	*kernel.BaseClient
}

type ResponseGetCallBackIp struct {
	IPList []string `json:"ip_list"`
	*kernel.ResponseWX
}

func (comp *Client) GetCallbackIp() *ResponseGetCallBackIp {

	result := &ResponseGetCallBackIp{}

	comp.HttpGet("cgi-bin/getcallbackip", nil, result)

	return result
}
