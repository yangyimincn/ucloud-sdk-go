//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UHost ModifyUHostInstanceName

package uhost

import (
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/request"
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/response"
)

// ModifyUHostInstanceNameRequest is request schema for ModifyUHostInstanceName action
type ModifyUHostInstanceNameRequest struct {
	request.CommonBase

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"false"`

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// UHost实例ID 参见 [DescribeUHostInstance](describe_uhost_instance.html)
	UHostId *string `required:"true"`

	// UHost实例名称
	Name *string `required:"false"`
}

// ModifyUHostInstanceNameResponse is response schema for ModifyUHostInstanceName action
type ModifyUHostInstanceNameResponse struct {
	response.CommonBase

	// UHost实例ID
	UhostId string
}

// NewModifyUHostInstanceNameRequest will create request of ModifyUHostInstanceName action.
func (c *UHostClient) NewModifyUHostInstanceNameRequest() *ModifyUHostInstanceNameRequest {
	req := &ModifyUHostInstanceNameRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// ModifyUHostInstanceName - 修改指定UHost实例名称，需要给出数据中心，UHostId，及新的实例名称。
func (c *UHostClient) ModifyUHostInstanceName(req *ModifyUHostInstanceNameRequest) (*ModifyUHostInstanceNameResponse, error) {
	var err error
	var res ModifyUHostInstanceNameResponse

	err = c.Client.InvokeAction("ModifyUHostInstanceName", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
