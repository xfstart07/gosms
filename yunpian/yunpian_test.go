package yunpian

import (
	"testing"

	"fmt"

	"github.com/xfstart07/gosms/yunpian"
)

var (
	service = yunpian.New("apikey")
)

func TestQuery(t *testing.T) {
	result, err := service.UserInfo()

	fmt.Println(result)
	fmt.Println(err)

	if err != nil {
		t.Errorf("%v", err)
	}

	if result.Code != 200 {
		t.Errorf("http error")
	}
}
