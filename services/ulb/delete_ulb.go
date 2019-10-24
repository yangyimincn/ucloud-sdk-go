//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api ULB DeleteULB

package ulb

import (
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/request"
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/response"
)

// DeleteULBRequest is request schema for DeleteULB action
type DeleteULBRequest struct {
	request.CommonBase

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"true"`

	// 负载均衡实例的ID
	ULBId *string `required:"true"`

	// 删除ulb时是否释放绑定的EIP，false标识只解绑EIP，true表示会释放绑定的EIP，默认是false
	ReleaseEip *bool `required:"false"`
}

// DeleteULBResponse is response schema for DeleteULB action
type DeleteULBResponse struct {
	response.CommonBase
}

// NewDeleteULBRequest will create request of DeleteULB action.
func (c *ULBClient) NewDeleteULBRequest() *DeleteULBRequest {
	req := &DeleteULBRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// DeleteULB - 删除负载均衡实例
func (c *ULBClient) DeleteULB(req *DeleteULBRequest) (*DeleteULBResponse, error) {
	var err error
	var res DeleteULBResponse

	err = c.Client.InvokeAction("DeleteULB", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
