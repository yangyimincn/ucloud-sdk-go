//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UHost DescribeImage

package uhost

import (
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/request"
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/response"
)

// DescribeImageRequest is request schema for DescribeImage action
type DescribeImageRequest struct {
	request.CommonBase

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"false"`

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// 镜像类型。标准镜像：Base，镜像市场：Business， 自定义镜像：Custom，默认返回所有类型
	ImageType *string `required:"false"`

	// 操作系统类型：Linux， Windows 默认返回所有类型
	OsType *string `required:"false"`

	// 镜像Id
	ImageId *string `required:"false"`

	// 列表起始位置偏移量，默认为0
	Offset *int `required:"false"`

	// 返回数据长度，默认为20
	Limit *int `required:"false"`

	// 是否返回价格：1返回，0不返回；默认不返回
	PriceSet *int `required:"false"`
}

// DescribeImageResponse is response schema for DescribeImage action
type DescribeImageResponse struct {
	response.CommonBase

	// 满足条件的镜像总数
	TotalCount int

	// 镜像列表详见 UHostImageSet
	ImageSet []UHostImageSet
}

// NewDescribeImageRequest will create request of DescribeImage action.
func (c *UHostClient) NewDescribeImageRequest() *DescribeImageRequest {
	req := &DescribeImageRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// DescribeImage - 获取指定数据中心镜像列表，用户可通过指定操作系统类型，镜像Id进行过滤。
func (c *UHostClient) DescribeImage(req *DescribeImageRequest) (*DescribeImageResponse, error) {
	var err error
	var res DescribeImageResponse

	err = c.Client.InvokeAction("DescribeImage", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
