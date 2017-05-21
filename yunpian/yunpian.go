package yunpian

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	//GetURL get info
	GetURL = "https://sms.yunpian.com/v2/user/get.json"

	//SingleSendURL send sms
	SingleSendURL = "https://sms.yunpian.com/v2/sms/single_send.json"

	//BatchSendURL send sms
	BatchSendURL = "https://sms.yunpian.com/v2/sms/batch_send.json"

	//TplSendURL send template sms
	TplSendURL = "https://sms.yunpian.com/v2/sms/tpl_batch_send.json"

	//VoiceSendURL send voice sms
	VoiceSendURL = "https://voice.yunpian.com/v2/voice/send.json"
)

// Config 云片 apikey
type Config struct {
	apikey string // 不对外调用
}

// Result  返回的结果
type Result struct {
	Code    int
	Message string
}

// New sms service
func New(apikey string) *Config {
	return &Config{
		apikey: apikey,
	}
}

// UserInfo 获得用户的信息，例如费用
func (cfg *Config) UserInfo() (*Result, error) {
	query := url.Values{}
	query.Add("apikey", cfg.apikey)

	req, err := http.PostForm(GetURL, query)

	if err != nil {
		return &Result{Code: req.StatusCode, Message: "请求失败"}, err
	}

	fmt.Println(req.StatusCode)

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return &Result{Code: req.StatusCode, Message: "读取失败"}, err
	}

	fmt.Println(string(body))

	return &Result{Code: req.StatusCode, Message: string(body)}, nil
}

// SingleSend 单条发送
func (cfg *Config) SingleSend(mobile, content string) (Result, error) {

	query := url.Values{}
	query.Add("mobile", mobile)
	query.Add("apikey", cfg.apikey)
	query.Add("text", content)

	req, err := http.PostForm(SingleSendURL, query)

	if err != nil {
		return Result{Code: req.StatusCode, Message: "请求失败"}, err
	}

	fmt.Println(req.StatusCode)
	fmt.Println(req.Body)

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return Result{Code: req.StatusCode, Message: "读取失败"}, err
	}
	fmt.Println(string(body))

	return Result{Code: req.StatusCode, Message: string(body)}, nil
}

// BatchSend 批量发送
func (cfg *Config) BatchSend(mobile, content string) (Result, error) {

	query := url.Values{}
	query.Add("mobile", mobile)
	query.Add("apikey", cfg.apikey)
	query.Add("text", content)

	req, err := http.PostForm(BatchSendURL, query)

	if err != nil {
		return Result{Code: req.StatusCode, Message: "请求失败"}, err
	}

	fmt.Println(req.StatusCode)
	fmt.Println(req.Body)

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return Result{Code: req.StatusCode, Message: "读取失败"}, err
	}
	fmt.Println(string(body))

	return Result{Code: req.StatusCode, Message: string(body)}, nil
}

// TplSend 通过模版发送短信
// Deprecated 官方不推荐使用了，请使用单条或批量发送的借口
func (cfg *Config) TplSend(mobile, tplID string, options map[string]string) (Result, error) {

	tplValue := url.Values{}
	for key, value := range options {
		tplValue.Add("#"+key+"#", value)
	}

	query := url.Values{}
	query.Add("mobile", mobile)
	query.Add("apikey", cfg.apikey)
	query.Add("tpl_id", tplID)
	query.Add("tpl_value", tplValue.Encode())

	req, err := http.PostForm(TplSendURL, query)

	if err != nil {
		return Result{Code: req.StatusCode, Message: "请求失败"}, err
	}

	fmt.Println(req.StatusCode)
	fmt.Println(req.Body)

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return Result{Code: req.StatusCode, Message: "读取失败"}, err
	}
	fmt.Println(string(body))

	return Result{Code: req.StatusCode, Message: string(body)}, nil
}

// Voice voice
func (cfg *Config) Voice(mobile, code string, options map[string]string) (Result, error) {

	query := url.Values{}
	query.Add("mobile", mobile)
	query.Add("apikey", cfg.apikey)
	query.Add("code", code)

	// 添加其他参数，例如 callback_url
	if options != nil {
		for key, value := range options {
			query.Add(key, value)
		}
	}

	req, err := http.PostForm(VoiceSendURL, query)
	if err != nil {
		return Result{req.StatusCode, "请求失败"}, err
	}

	fmt.Println(req.StatusCode)
	fmt.Println(req.Body)

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return Result{req.StatusCode, "读取失败"}, err
	}
	fmt.Println(string(body))

	return Result{req.StatusCode, string(body)}, nil
}
