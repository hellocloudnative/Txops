package image

import (
	"Txops/config"
	"encoding/json"
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	image "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
)

var(
	images  config.ImageList
)

func ShowImage()  {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "cvm.tencentcloudapi.com"
	client, _ := image.NewClient(config.Credential, fmt.Sprintf("ap-%s",Zone), cpf)

	request := image.NewDescribeImagesRequest()

	params := "{\"Filters\":[{\"Name\":\"image-type\",\"Values\":[\"PRIVATE_IMAGE\"]}]}"
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DescribeImages(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	jsonStr:=fmt.Sprintf("%s", response.ToJsonString())
	if All == "images"{
		fmt.Println(jsonStr)
	}

	if err := json.Unmarshal([]byte(jsonStr), &images); err == nil {
		imageInfos := images.Response.ImageSet
		for _, i := range (imageInfos) {
			if ImageName == i.ImageName {
				fmt.Printf("镜像ID:[%s]\n",i.ImageID)
				fmt.Printf("镜像名称:[%s]\n",i.ImageName)
			}
			if ImageId == i.ImageID {
				fmt.Printf("镜像ID:[%s]\n",i.ImageID)
				fmt.Printf("镜像名称:[%s]\n",i.ImageName)
			}
		}
	}
}
