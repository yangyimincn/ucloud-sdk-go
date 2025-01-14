// Code is generated by ucloud-model, DO NOT EDIT IT.

package ucloudstack

import (
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/request"
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/response"
)

// DeleteVMInstanceRequest is request schema for DeleteVMInstance action
type DeleteVMInstanceRequest struct {
	request.CommonBase

	// [公共参数] 地域。 枚举值：cn，表示中国；
	// Region *string `required:"true"`

	// [公共参数] 可用区。枚举值：zone-01，表示中国；
	// Zone *string `required:"true"`

	// 虚拟机 ID。输入有效的虚拟机 ID。
	VMID *string `required:"true"`
}

// DeleteVMInstanceResponse is response schema for DeleteVMInstance action
type DeleteVMInstanceResponse struct {
	response.CommonBase

	// 返回信息描述。
	Message string
}

// NewDeleteVMInstanceRequest will create request of DeleteVMInstance action.
func (c *UCloudStackClient) NewDeleteVMInstanceRequest() *DeleteVMInstanceRequest {
	req := &DeleteVMInstanceRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// DeleteVMInstance - 删除UCloudStack虚拟机
func (c *UCloudStackClient) DeleteVMInstance(req *DeleteVMInstanceRequest) (*DeleteVMInstanceResponse, error) {
	var err error
	var res DeleteVMInstanceResponse

	reqCopier := *req

	err = c.Client.InvokeAction("DeleteVMInstance", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
