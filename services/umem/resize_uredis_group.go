//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UMem ResizeURedisGroup

package umem

import (
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/request"
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/response"
)

// ResizeURedisGroupRequest is request schema for ResizeURedisGroup action
type ResizeURedisGroupRequest struct {
	request.CommonBase

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"false"`

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// 组ID
	GroupId *string `required:"true"`

	// 内存大小, 单位:GB (需要大于原size,且小于等于32) 目前仅支持1/2/4/8/16/32 G 六种容量规格
	Size *int `required:"true"`

	//
	ChargeType *string `required:"false"`

	// 空间类型:single(无热备),double(热备)(默认: double)
	Type *string `required:"false"`

	// 代金券ID 请参考DescribeCoupon接口
	CouponId *int `required:"false"`
}

// ResizeURedisGroupResponse is response schema for ResizeURedisGroup action
type ResizeURedisGroupResponse struct {
	response.CommonBase
}

// NewResizeURedisGroupRequest will create request of ResizeURedisGroup action.
func (c *UMemClient) NewResizeURedisGroupRequest() *ResizeURedisGroupRequest {
	req := &ResizeURedisGroupRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// ResizeURedisGroup - 调整主备redis容量
func (c *UMemClient) ResizeURedisGroup(req *ResizeURedisGroupRequest) (*ResizeURedisGroupResponse, error) {
	var err error
	var res ResizeURedisGroupResponse

	err = c.Client.InvokeAction("ResizeURedisGroup", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
