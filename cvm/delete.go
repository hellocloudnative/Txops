package cvm

import (
	"Txops/config"
	"Txops/image"
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
)

func Deletecvm()  {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "cvm.tencentcloudapi.com"
	client, _ := cvm.NewClient(config.Credential, fmt.Sprintf("ap-%s",Zone), cpf)

	request := cvm.NewTerminateInstancesRequest()
	cvmid:=image.GetCvmID(Ipaddress)
	params :=  fmt.Sprintf("{\"InstanceIds\":[\"%s\"]}",cvmid)
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.TerminateInstances(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", response.ToJsonString())
	fmt.Print("[%s]主机退还成功",Ipaddress)
}

