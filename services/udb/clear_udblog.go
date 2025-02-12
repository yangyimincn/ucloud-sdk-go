//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UDB ClearUDBLog

package udb

import (
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/request"
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/response"
)

// ClearUDBLogRequest is request schema for ClearUDBLog action
type ClearUDBLogRequest struct {
	request.CommonBase

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"false"`

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// DB实例的id,该值可以通过DescribeUDBInstance获取
	DBId *string `required:"true"`

	// 日志类型，10-error（暂不支持）、20-slow（暂不支持 ）、30-binlog
	LogType *int `required:"true"`

	// 删除时间点(至少前一天)之前log，采用时间戳(秒)，默认当 前时间点前一天
	BeforeTime *int `required:"false"`
}

// ClearUDBLogResponse is response schema for ClearUDBLog action
type ClearUDBLogResponse struct {
	response.CommonBase
}

// NewClearUDBLogRequest will create request of ClearUDBLog action.
func (c *UDBClient) NewClearUDBLogRequest() *ClearUDBLogRequest {
	req := &ClearUDBLogRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// ClearUDBLog - 清除UDB实例的log
func (c *UDBClient) ClearUDBLog(req *ClearUDBLogRequest) (*ClearUDBLogResponse, error) {
	var err error
	var res ClearUDBLogResponse

	err = c.Client.InvokeAction("ClearUDBLog", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
