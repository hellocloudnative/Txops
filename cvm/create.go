package cvm

import (
	"Txops/config"
	"encoding/json"
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
)

func Createcvm()  {
		cpf := profile.NewClientProfile()
		cpf.HttpProfile.Endpoint = "cvm.tencentcloudapi.com"
		client, _ := cvm.NewClient(config.Credential, fmt.Sprintf("ap-%s",Zone), cpf)

		request := cvm.NewRunInstancesRequest()
		ImageId=GetImageID(ImageName)
		params :=fmt.Sprintf( "{\"Placement\":{\"Zone\":\"ap-guangzhou-4\"," +
			"\"ProjectId\":1112410},\"InstanceType\":\"S5.8XLARGE64\"," +
			"\"ImageId\":\"%s\",\"SystemDisk\":{\"DiskType\":\"CLOUD_PREMIUM\"," +
			"\"DiskSize\":500},\"VirtualPrivateCloud\":{\"VpcId\":\"vpc-nn4i3pj9\"," +
			"\"SubnetId\":\"subnet-7yuw15ss\",\"PrivateIpAddresses\":[\"%s\"]}," +
			"\"InstanceName\":\"%s\",\"LoginSettings\":{\"Password\":\"nEtbenTY2o20\"}," +
			"\"SecurityGroupIds\":[\"sg-p0nv061x\"],\"HostName\":\"%s\",\"DryRun\":%s}", ImageId, Ipaddress,HostName, HostName, DryRun)
		err := request.FromJsonString(params)
		if err != nil {
			panic(err)
		}
		response, err := client.RunInstances(request)
		if _, ok := err.(*errors.TencentCloudSDKError); ok {
			fmt.Printf("An API error has returned: %s", err)
			return
		}
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", response.ToJsonString())
		fmt.Println("cvm 主机创建成功")
}

func GetImageID(imagename  string) string {
	var images config.ImageList
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "cvm.tencentcloudapi.com"
	client, _ := cvm.NewClient(config.Credential, fmt.Sprintf("ap-%s",Zone), cpf)

	request := cvm.NewDescribeImagesRequest()

	params := "{\"Filters\":[{\"Name\":\"image-type\",\"Values\":[\"PRIVATE_IMAGE\"]}]}"
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DescribeImages(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
	}
	if err != nil {
		panic(err)
	}
	jsonStr := fmt.Sprintf("%s", response.ToJsonString())
	if err := json.Unmarshal([]byte(jsonStr), &images); err == nil {
		imageInfos := images.Response.ImageSet
		for _, i := range (imageInfos) {
			if imagename == i.ImageName {
				return i.ImageID
			}

		}
	}

	return ""
}