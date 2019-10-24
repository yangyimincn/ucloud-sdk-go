//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UDB CheckRecoverUDBInstance

package udb

import (
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/request"
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/response"
)

// CheckRecoverUDBInstanceRequest is request schema for CheckRecoverUDBInstance action
type CheckRecoverUDBInstanceRequest struct {
	request.CommonBase

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"false"`

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// 源实例的Id(只支持普通版DB不支持高可用)
	SrcDBId *string `required:"true"`
}

// CheckRecoverUDBInstanceResponse is response schema for CheckRecoverUDBInstance action
type CheckRecoverUDBInstanceResponse struct {
	response.CommonBase

	// 核查成功返回值为可以回档到的最近时刻,核查失败不返回
	LastestTime int
}

// NewCheckRecoverUDBInstanceRequest will create request of CheckRecoverUDBInstance action.
func (c *UDBClient) NewCheckRecoverUDBInstanceRequest() *CheckRecoverUDBInstanceRequest {
	req := &CheckRecoverUDBInstanceRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// CheckRecoverUDBInstance - 核查db是否可以使用回档功能
func (c *UDBClient) CheckRecoverUDBInstance(req *CheckRecoverUDBInstanceRequest) (*CheckRecoverUDBInstanceResponse, error) {
	var err error
	var res CheckRecoverUDBInstanceResponse

	err = c.Client.InvokeAction("CheckRecoverUDBInstance", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
