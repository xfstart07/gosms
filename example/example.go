package example

import (
	"fmt"

	"github.com/xfstart07/gosms/luosimao"
)

func main() {
	service := luosimao.New("apikey")

	result, err := service.Send("you mobile", "你的验证码: 1231")
	if err != nil {
		fmt.Println("err")
	}

	fmt.Println(result.Code)
	fmt.Println(result.Message)
}
