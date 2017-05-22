package luosimao

import (
	"testing"

	"fmt"

	"github.com/xfstart07/gosms/luosimao"
)

var (
	service = luosimao.New("apikey")
)

func TestQuery(t *testing.T) {
	result, err := service.Query()

	fmt.Println(result)
	fmt.Println(err)

	if err != nil {
		t.Errorf("%v", err)
	}

	if result.Code != 200 {
		t.Errorf("http error")
	}

}
