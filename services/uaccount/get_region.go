//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UAccount GetRegion

package uaccount

import (
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/request"
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/response"
)

// GetRegionRequest is request schema for GetRegion action
type GetRegionRequest struct {
	request.CommonBase
}

// GetRegionResponse is response schema for GetRegion action
type GetRegionResponse struct {
	response.CommonBase

	// 各数据中心信息
	Regions []RegionInfo
}

// NewGetRegionRequest will create request of GetRegion action.
func (c *UAccountClient) NewGetRegionRequest() *GetRegionRequest {
	req := &GetRegionRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// GetRegion - 获取用户在各数据中心的权限等信息
func (c *UAccountClient) GetRegion(req *GetRegionRequest) (*GetRegionResponse, error) {
	var err error
	var res GetRegionResponse

	err = c.Client.InvokeAction("GetRegion", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
