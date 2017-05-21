package smsbao

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	//SendURL 发送
	SendURL = "http://api.smsbao.com/sms"

	// VoiceURL 查询
	VoiceURL = "http://www.smsbao.com/voice"

	// QueryURL 查询
	QueryURL = "http://www.smsbao.com/query"
)

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
func New(username, password string) *Config {
	return &Config{
		username: username,
		password: password,
	}
}

// Send 发送短信
func (cfg *Config) Send(mobile, content string) (Result, error) {
	md5Password := hexToMd5(cfg.password)

	query := url.Values{}
	query.Add("u", cfg.username)
	query.Add("p", md5Password)
	query.Add("m", mobile)
	query.Add("c", content)

	return queryByURL(SendURL, query)
}

//Voice 语音
func (cfg *Config) Voice(mobile, code string) (Result, error) {
	md5Password := hexToMd5(cfg.password)

	query := url.Values{}
	query.Add("u", cfg.username)
	query.Add("p", md5Password)
	query.Add("m", mobile)
	query.Add("c", code)

	return queryByURL(VoiceURL, query)
}

//Query 查询短信剩余信息
func (cfg *Config) Query() (Result, error) {
	md5Password := hexToMd5(cfg.password)

	query := url.Values{}
	query.Add("u", cfg.username)
	query.Add("p", md5Password)

	return queryByURL(QueryURL, query)
}

// private method

// queryByURL 通过 URL 查询
func queryByURL(url string, query url.Values) (Result, error) {
	req, err := http.PostForm(url, query)

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

// hexToMd5 md5 加密
func hexToMd5(content string) string {
	hash := md5.New()
	hash.Write([]byte(content))
	return hex.EncodeToString(hash.Sum(nil))
}
