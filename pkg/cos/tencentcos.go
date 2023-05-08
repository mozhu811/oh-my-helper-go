package cos

import (
	"context"
	"github.com/tencentyun/cos-go-sdk-v5"
	"io"
	"net/http"
	"net/url"
	"os"
)

func Upload(file, name string) error {
	u, _ := url.Parse("https://bilibili-cruii-io-1251547651.cos.ap-chengdu.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  "", // 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
			SecretKey: "", // 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
		},
	})

	name = "avatars/" + name
	response, err := http.Get(file)
	if err != nil {
		return err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(response.Body)

	temp, err := os.CreateTemp("", "avatar-*.png")
	if err != nil {
		return err
	}

	_, err = io.Copy(temp, response.Body)
	if err != nil {
		return err
	}

	_, err = c.Object.PutFromFile(context.Background(), name, temp.Name(), nil)
	if err != nil {
		return err
	}
	return nil
}
