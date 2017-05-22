package yunpian

import (
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

// 返回值说明
// https://www.yunpian.com/api2.0/api-recode.html

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

// Query 获得用户的信息，例如费用
func (cfg *Config) Query() (Result, error) {
	query := url.Values{}
	query.Add("apikey", cfg.apikey)

	return queryByURL(GetURL, query)
}

// SingleSend 单条发送
func (cfg *Config) SingleSend(mobile, content string) (Result, error) {

	query := url.Values{}
	query.Add("mobile", mobile)
	query.Add("apikey", cfg.apikey)
	query.Add("text", content)

	return queryByURL(SingleSendURL, query)
}

// BatchSend 批量发送
func (cfg *Config) BatchSend(mobile, content string) (Result, error) {

	query := url.Values{}
	query.Add("mobile", mobile)
	query.Add("apikey", cfg.apikey)
	query.Add("text", content)

	return queryByURL(BatchSendURL, query)
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

	return queryByURL(TplSendURL, query)
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

	return queryByURL(VoiceSendURL, query)
}

// private method

// queryByURL 通过 URL 查询
func queryByURL(url string, query url.Values) (Result, error) {
	req, err := http.PostForm(url, query)

	if err != nil {
		return Result{Code: req.StatusCode, Message: "请求失败"}, err
	}

	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return Result{Code: req.StatusCode, Message: "读取失败"}, err
	}

	// TODO 还没有解析结果
	bodyString := string(body)

	return Result{Code: req.StatusCode, Message: bodyString}, nil
}
