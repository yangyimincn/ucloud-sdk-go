//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UDB PromoteUDBInstanceToHA

package udb

import (
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/request"
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/response"
)

// PromoteUDBInstanceToHARequest is request schema for PromoteUDBInstanceToHA action
type PromoteUDBInstanceToHARequest struct {
	request.CommonBase

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// 实例的Id,该值可以通过DescribeUDBInstance获取
	DBId *string `required:"true"`

	// 升级时是否锁库，默认为锁库。不锁定：false；锁定：true。现在统一为不锁库，即false
	IsLock *bool `required:"false"`
}

// PromoteUDBInstanceToHAResponse is response schema for PromoteUDBInstanceToHA action
type PromoteUDBInstanceToHAResponse struct {
	response.CommonBase
}

// NewPromoteUDBInstanceToHARequest will create request of PromoteUDBInstanceToHA action.
func (c *UDBClient) NewPromoteUDBInstanceToHARequest() *PromoteUDBInstanceToHARequest {
	req := &PromoteUDBInstanceToHARequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// PromoteUDBInstanceToHA - 普通db升级为高可用(只针对mysql5.5及以上版本)
func (c *UDBClient) PromoteUDBInstanceToHA(req *PromoteUDBInstanceToHARequest) (*PromoteUDBInstanceToHAResponse, error) {
	var err error
	var res PromoteUDBInstanceToHAResponse

	err = c.Client.InvokeAction("PromoteUDBInstanceToHA", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
