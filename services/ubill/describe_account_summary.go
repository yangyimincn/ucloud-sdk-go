package ubill

import (
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/request"
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/response"
)

type DescribeAccountSummaryRequest struct {
	request.CommonBase
}

type DescribeAccountSummaryResponse struct {
	response.CommonBase

	// JSON格式的订单信息
	AccountBalance	AccountBalance
}

func (c *UBillClient) NewDescribeAccountSummaryRequest() *DescribeAccountSummaryRequest {
	req := &DescribeAccountSummaryRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// 获取账户余额信息
func (c *UBillClient) DescribeAccountSummary(req *DescribeAccountSummaryRequest) (*DescribeAccountSummaryResponse, error) {
	var err error
	var res DescribeAccountSummaryResponse

	err = c.Client.InvokeAction("DescribeAccountSummary", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
