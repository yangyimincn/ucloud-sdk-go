//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UDB RestartUDBInstance

package udb

import (
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/request"
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/response"
)

// RestartUDBInstanceRequest is request schema for RestartUDBInstance action
type RestartUDBInstanceRequest struct {
	request.CommonBase

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"false"`

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// 实例的Id,该值可以通过DescribeUDBInstance获取
	DBId *string `required:"true"`
}

// RestartUDBInstanceResponse is response schema for RestartUDBInstance action
type RestartUDBInstanceResponse struct {
	response.CommonBase
}

// NewRestartUDBInstanceRequest will create request of RestartUDBInstance action.
func (c *UDBClient) NewRestartUDBInstanceRequest() *RestartUDBInstanceRequest {
	req := &RestartUDBInstanceRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// RestartUDBInstance - 重启UDB实例
func (c *UDBClient) RestartUDBInstance(req *RestartUDBInstanceRequest) (*RestartUDBInstanceResponse, error) {
	var err error
	var res RestartUDBInstanceResponse

	err = c.Client.InvokeAction("RestartUDBInstance", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
