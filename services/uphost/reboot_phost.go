//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UPHost RebootPHost

package uphost

import (
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/request"
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/response"
)

// RebootPHostRequest is request schema for RebootPHost action
type RebootPHostRequest struct {
	request.CommonBase

	// 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"false"`

	// PHost资源ID
	PHostId *string `required:"true"`
}

// RebootPHostResponse is response schema for RebootPHost action
type RebootPHostResponse struct {
	response.CommonBase

	// PHost 的资源ID
	PHostId string
}

// NewRebootPHostRequest will create request of RebootPHost action.
func (c *UPHostClient) NewRebootPHostRequest() *RebootPHostRequest {
	req := &RebootPHostRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// RebootPHost - 重启物理机
func (c *UPHostClient) RebootPHost(req *RebootPHostRequest) (*RebootPHostResponse, error) {
	var err error
	var res RebootPHostResponse

	err = c.Client.InvokeAction("RebootPHost", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
