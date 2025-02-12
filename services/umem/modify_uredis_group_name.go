//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UMem ModifyURedisGroupName

package umem

import (
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/request"
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/response"
)

// ModifyURedisGroupNameRequest is request schema for ModifyURedisGroupName action
type ModifyURedisGroupNameRequest struct {
	request.CommonBase

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"false"`

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// 组的ID
	GroupId *string `required:"true"`

	// Redis组名称 (范围[6-63],只能包含英文、数字以及符号-和_)
	Name *string `required:"true"`
}

// ModifyURedisGroupNameResponse is response schema for ModifyURedisGroupName action
type ModifyURedisGroupNameResponse struct {
	response.CommonBase
}

// NewModifyURedisGroupNameRequest will create request of ModifyURedisGroupName action.
func (c *UMemClient) NewModifyURedisGroupNameRequest() *ModifyURedisGroupNameRequest {
	req := &ModifyURedisGroupNameRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// ModifyURedisGroupName - 修改主备redis名称
func (c *UMemClient) ModifyURedisGroupName(req *ModifyURedisGroupNameRequest) (*ModifyURedisGroupNameResponse, error) {
	var err error
	var res ModifyURedisGroupNameResponse

	err = c.Client.InvokeAction("ModifyURedisGroupName", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
