package http

import (
	"bytes"
	"log"
	"net/http"
	"watchAlert/globals"
)

func Get(url string) (*http.Response, error) {

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println("err 1")
	}
	if err != nil {
		globals.Logger.Sugar().Error("请求建立失败", err)
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		globals.Logger.Sugar().Error("请求发送失败", err)
		return nil, err
	}

	return resp, nil
}

func Post(url string, bodyReader *bytes.Reader) (*http.Response, error) {

	request, err := http.NewRequest(http.MethodPost, url, bodyReader)
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		globals.Logger.Sugar().Error("请求建立失败", err)
		return nil, err
	}
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		globals.Logger.Sugar().Error("请求发送失败", err)
		return nil, err
	}

	return resp, nil

}

func PostReloadPrometheus() error {

	url := globals.Config.Prometheus.URL + "/-/reload"
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		globals.Logger.Sugar().Error("PostReloadPrometheus 请求建立失败 ->", err)
		return err
	}
	_, err = http.DefaultClient.Do(req)
	if err != nil {
		globals.Logger.Sugar().Error("PostReloadPrometheus 请求发送失败 ->", err)
		return err
	}

	return nil
}
