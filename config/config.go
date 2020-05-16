package config

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"time"
)

var(
	Credential = common.NewCredential(
		"*",
		"*",
	)
)
type CvmList struct {
	Response struct {
		TotalCount  int `json:"TotalCount"`
		InstanceSet []struct {
			Placement struct {
				Zone      string      `json:"Zone"`
				HostID    interface{} `json:"HostId"`
				ProjectID int         `json:"ProjectId"`
			} `json:"Placement"`
			InstanceID         string `json:"InstanceId"`
			UUID               string `json:"Uuid"`
			InstanceState      string `json:"InstanceState"`
			RestrictState      string `json:"RestrictState"`
			InstanceType       string `json:"InstanceType"`
			CPU                int    `json:"CPU"`
			Memory             int    `json:"Memory"`
			InstanceName       string `json:"InstanceName"`
			InstanceChargeType string `json:"InstanceChargeType"`
			SystemDisk         struct {
				DiskType string `json:"DiskType"`
				DiskID   string `json:"DiskId"`
				DiskSize int    `json:"DiskSize"`
			} `json:"SystemDisk"`
			DataDisks          interface{} `json:"DataDisks"`
			PrivateIPAddresses []string    `json:"PrivateIpAddresses"`
			PublicIPAddresses  interface{} `json:"PublicIpAddresses"`
			IPv6Addresses      interface{} `json:"IPv6Addresses"`
			InternetAccessible struct {
				InternetMaxBandwidthOut int         `json:"InternetMaxBandwidthOut"`
				InternetChargeType      interface{} `json:"InternetChargeType"`
			} `json:"InternetAccessible"`
			VirtualPrivateCloud struct {
				VpcID        string `json:"VpcId"`
				SubnetID     string `json:"SubnetId"`
				AsVpcGateway bool   `json:"AsVpcGateway"`
			} `json:"VirtualPrivateCloud"`
			SecurityGroupIds []string `json:"SecurityGroupIds"`
			LoginSettings    struct {
				KeyIds interface{} `json:"KeyIds"`
			} `json:"LoginSettings"`
			ImageID                  string        `json:"ImageId"`
			OsName                   string        `json:"OsName"`
			RenewFlag                interface{}   `json:"RenewFlag"`
			CreatedTime              time.Time     `json:"CreatedTime"`
			ExpiredTime              interface{}   `json:"ExpiredTime"`
			Tags                     []interface{} `json:"Tags"`
			DisasterRecoverGroupID   string        `json:"DisasterRecoverGroupId"`
			CamRoleName              string        `json:"CamRoleName"`
			LatestOperation          interface{}   `json:"LatestOperation"`
			LatestOperationState     interface{}   `json:"LatestOperationState"`
			LatestOperationRequestID interface{}   `json:"LatestOperationRequestId"`
			IsolatedSource           string        `json:"IsolatedSource"`
			StopChargingMode         string        `json:"StopChargingMode"`
		} `json:"InstanceSet"`
		RequestID string `json:"RequestId"`
	} `json:"Response"`
}

type ImageList struct {
	Response struct {
		TotalCount int `json:"TotalCount"`
		ImageSet   []struct {
			OsName                        string      `json:"OsName"`
			ImageSize                     int         `json:"ImageSize"`
			ImageType                     string      `json:"ImageType"`
			CreatedTime                   time.Time   `json:"CreatedTime"`
			ImageDescription              string      `json:"ImageDescription"`
			ImageSource                   string      `json:"ImageSource"`
			ImageID                       string      `json:"ImageId"`
			ImageName                     string      `json:"ImageName"`
			ImageCreator                  string      `json:"ImageCreator"`
			ImageState                    string      `json:"ImageState"`
			SyncPercent                   interface{} `json:"SyncPercent"`
			IsSupportAutoInstallGPUDriver bool        `json:"IsSupportAutoInstallGPUDriver"`
			SnapshotSet                   []struct {
				SnapshotID string `json:"SnapshotId"`
				DiskUsage  string `json:"DiskUsage"`
				DiskSize   int    `json:"DiskSize"`
			} `json:"SnapshotSet,omitempty"`
			ImageClass         string `json:"_ImageClass"`
			Architecture       string `json:"Architecture"`
			Platform           string `json:"Platform"`
			IsSupportCloudinit bool   `json:"IsSupportCloudinit"`
		} `json:"ImageSet"`
		RequestID string `json:"RequestId"`
	} `json:"Response"`
}

