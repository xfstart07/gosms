package smsbao

import (
	"testing"

	"fmt"

	"github.com/xfstart07/gosms/smsbao"
)

var (
	service = smsbao.New("username", "password")
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
