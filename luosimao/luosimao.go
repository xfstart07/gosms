package luosimao

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	//SendURL 发送
	SendURL = "https://sms-api.luosimao.com/v1/send.json"

	//BatchSendURL 发送
	BatchSendURL = "https://sms-api.luosimao.com/v1/send_batch.json"

	//QueryURL 查询信息
	QueryURL = "http://sms-api.luosimao.com/v1/status.json"
)

// error 0 成功
// https://luosimao.com/docs/api/20#status

// Config 用户名, 密码
type Config struct {
	username string
	password string
}

// Result  返回的结果
type Result struct {
	Code    int
	Message string
}

// New sms service
func New(apikey string) *Config {
	return &Config{
		username: "api",
		password: "key-" + apikey,
	}
}

// Send 发送短信
func (cfg *Config) Send(mobile, content string) (Result, error) {
	query := url.Values{}
	query.Add("mobile", mobile)
	query.Add("message", content)

	return cfg.queryByURL(SendURL, query)
}

// BatchSend 批量发送短信
func (cfg *Config) BatchSend(mobile []string, content string) (Result, error) {
	query := url.Values{}
	query.Add("mobile_list", strings.Join(mobile, ","))
	query.Add("message", content)

	return cfg.queryByURL(SendURL, query)
}

// Query 查询信息
func (cfg *Config) Query() (Result, error) {
	query := url.Values{}

	return cfg.queryByURL(SendURL, query)
}

// private method

// queryByURL 通过 URL 查询
func (cfg *Config) queryByURL(url string, query url.Values) (Result, error) {
	// http client
	client := &http.Client{}

	// 压缩成 string io Reader 类型
	queryEncode := strings.NewReader(query.Encode())

	// http request
	request, err := http.NewRequest("POST", QueryURL, queryEncode)
	request.SetBasicAuth(cfg.username, cfg.password)

	// 请求
	req, err := client.Do(request)

	if err != nil {
		return Result{Code: req.StatusCode, Message: "请求失败"}, err
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return Result{Code: req.StatusCode, Message: "读取失败"}, err
	}

	// TODO 还没有解析结果
	bodyString := string(body)

	return Result{Code: req.StatusCode, Message: bodyString}, nil
}
