//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UPHost DescribePHostTags

package uphost

import (
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/request"
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/response"
)

// DescribePHostTagsRequest is request schema for DescribePHostTags action
type DescribePHostTagsRequest struct {
	request.CommonBase

	// 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"false"`
}

// DescribePHostTagsResponse is response schema for DescribePHostTags action
type DescribePHostTagsResponse struct {
	response.CommonBase

	// Tag的个数
	TotalCount int

	// 具体参见 PHostTagSet
	TagSet []PHostTagSet
}

// NewDescribePHostTagsRequest will create request of DescribePHostTags action.
func (c *UPHostClient) NewDescribePHostTagsRequest() *DescribePHostTagsRequest {
	req := &DescribePHostTagsRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// DescribePHostTags - 获取物理机tag列表（业务组）
func (c *UPHostClient) DescribePHostTags(req *DescribePHostTagsRequest) (*DescribePHostTagsResponse, error) {
	var err error
	var res DescribePHostTagsResponse

	err = c.Client.InvokeAction("DescribePHostTags", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
