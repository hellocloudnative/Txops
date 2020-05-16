package image

import (
	"Txops/config"
	"encoding/json"
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	image "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
)

func MakeImage()  {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "cvm.tencentcloudapi.com"
	client, _ := image.NewClient(config.Credential, "ap-guangzhou", cpf)

	request := image.NewCreateImageRequest()
	cvmid:=GetCvmID(Ipaddress)
	params := fmt.Sprintf("{\"InstanceId\":\"%s\",\"ImageName\":\"%s\",\"DryRun\":%s}",cvmid,ImageName,DryRun)
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.CreateImage(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
	fmt.Println("镜像创建成功")
}

func  GetCvmID(ipaddress string) string {
	var cvms  config.CvmList
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "cvm.tencentcloudapi.com"
	client, _ := image.NewClient(config.Credential, fmt.Sprintf("ap-%s",Zone), cpf)

	request := image.NewDescribeInstancesRequest()

	params := "{}"
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DescribeInstances(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
	}
	if err != nil {
		panic(err)
	}
	jsonStr := fmt.Sprintf("%s", response.ToJsonString())
	if err := json.Unmarshal([]byte(jsonStr), &cvms); err == nil {
		cvmInfos := cvms.Response.InstanceSet
		for _, i := range (cvmInfos) {
			if ipaddress == i.PrivateIPAddresses[0] {
				return i.InstanceID
			}

		}
	}
	return ""
}