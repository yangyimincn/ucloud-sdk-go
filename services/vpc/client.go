// Code is generated by ucloud-model, DO NOT EDIT IT.

package vpc

import (
	"github.com/yangyimincn/ucloud-sdk-go/ucloud"
	"github.com/yangyimincn/ucloud-sdk-go/ucloud/auth"
)

// VPCClient is the client of VPC
type VPCClient struct {
	*ucloud.Client
}

// NewClient will return a instance of VPCClient
func NewClient(config *ucloud.Config, credential *auth.Credential) *VPCClient {
	meta := ucloud.ClientMeta{Product: "VPC2.0"}
	client := ucloud.NewClientWithMeta(config, credential, meta)
	return &VPCClient{
		client,
	}
}
