//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UAccount TerminateProject

package uaccount

import (
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/request"
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/response"
)

// TerminateProjectRequest is request schema for TerminateProject action
type TerminateProjectRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

}

// TerminateProjectResponse is response schema for TerminateProject action
type TerminateProjectResponse struct {
	response.CommonBase
}

// NewTerminateProjectRequest will create request of TerminateProject action.
func (c *UAccountClient) NewTerminateProjectRequest() *TerminateProjectRequest {
	req := &TerminateProjectRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// TerminateProject - 删除项目
func (c *UAccountClient) TerminateProject(req *TerminateProjectRequest) (*TerminateProjectResponse, error) {
	var err error
	var res TerminateProjectResponse

	err = c.Client.InvokeAction("TerminateProject", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
