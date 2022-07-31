package qgproxy

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
)

var endpoint = `https://proxy.qg.net`

type ProxyInfo struct {
	Code int `json:"Code"`
	Data []struct {
		IP       string `json:"IP"`
		Port     string `json:"port"`
		Deadline string `json:"deadline"`
		Host     string `json:"host"`
	} `json:"Data"`
	Num    int    `json:"Num"`
	Msg    string    `json:"Msg"`
	TaskID string `json:"TaskID"`
}

type QGProxy struct {
	client *resty.Client
}

func NewQGProxy() *QGProxy {
	client := resty.New()
	client.SetBaseURL(endpoint)

	return &QGProxy{
		client: client,
	}
}

func (q *QGProxy) request(method string, uri string, query map[string]string, data map[string]string) (proxyInfo *ProxyInfo, err error) {
	req := q.client.R()
	if query != nil {
		req.SetQueryParams(query)
	}

	if data != nil {
		req.SetFormData(data)
	}

	var resp *resty.Response

	if method == "post" {
		resp, err = req.Post(uri)
	} else {
		resp, err = req.Get(uri)
	}

	if err != nil {
		return &ProxyInfo{}, err
	}

	err = json.Unmarshal(resp.Body(), &proxyInfo)
	if err != nil {
		return &ProxyInfo{}, err
	}

	return proxyInfo, err

}

func (q *QGProxy) Allocate(key string) (proxyInfo *ProxyInfo, err error) {

	query := map[string]string{
		"Key": key,
	}

	return q.request("get", "/allocate", query, nil)
}

func (q *QGProxy) Release(key, ip string) (proxyInfo *ProxyInfo, err error) {

	query := map[string]string{
		"Key": key,
		"IP": ip,
	}

	return q.request("get", "/release", query, nil)
}

func (q *QGProxy) Replace(key, ip string) (proxyInfo *ProxyInfo, err error) {

	query := map[string]string{
		"Key": key,
		"IP": ip,
	}

	return q.request("get", "/replace", query, nil)
}

func (q *QGProxy) Query(key string) (proxyInfo *ProxyInfo, err error) {

	query := map[string]string{
		"Key": key,
	}

	return q.request("get", "/query", query, nil)
}