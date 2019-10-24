//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UPHost StartPHost

package uphost

import (
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/request"
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/response"
)

// StartPHostRequest is request schema for StartPHost action
type StartPHostRequest struct {
	request.CommonBase

	// 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"false"`

	// PHost资源ID
	PHostId *string `required:"true"`
}

// StartPHostResponse is response schema for StartPHost action
type StartPHostResponse struct {
	response.CommonBase

	// PHost 的资源ID
	PHostId string
}

// NewStartPHostRequest will create request of StartPHost action.
func (c *UPHostClient) NewStartPHostRequest() *StartPHostRequest {
	req := &StartPHostRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// StartPHost - 启动物理机
func (c *UPHostClient) StartPHost(req *StartPHostRequest) (*StartPHostResponse, error) {
	var err error
	var res StartPHostResponse

	err = c.Client.InvokeAction("StartPHost", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
