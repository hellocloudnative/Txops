package image

import (
	"Txops/config"
	"encoding/json"
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
  image	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
)

func Deleteimage()  {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "cvm.tencentcloudapi.com"
	client, _ := image.NewClient(config.Credential, fmt.Sprintf("ap-%s",Zone), cpf)

	request := image.NewDeleteImagesRequest()
	ImageId=GetImageID(ImageName)
	params := fmt.Sprintf("{\"ImageIds\":[\"%s\"]}",ImageId)
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DeleteImages(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", response.ToJsonString())
	fmt.Printf("[%s]镜像删除成功",ImageName)
}

func GetImageID(imagename  string) string {
	var images config.ImageList
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