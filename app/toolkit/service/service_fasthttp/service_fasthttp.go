package service_fasthttp

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"time"
)

var (
	SCU_Client *fasthttp.Client
)

func GetClient() *fasthttp.Client {
	return SCU_Client
}

func Get(url string, headers map[string]string, query map[string]string) (*fasthttp.ResponseHeader, []byte, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.Header.SetContentType("application/x-www-form-urlencoded")
	req.Header.SetMethod("Get")
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	if len(query) > 0 {
		url = url + "?"
		for k, v := range query {
			temp := k + "=" + v
			url = url + "&" + temp
		}
	}
	req.SetRequestURI(url)
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)
	if err := fasthttp.Do(req, resp); err != nil {
		return nil, nil, err
	}
	data_Header := new(fasthttp.ResponseHeader)
	*data_Header = resp.Header
	data_Body := resp.Body()
	return data_Header, data_Body, nil
}

func Post(url string, headers map[string]string, form map[string]string) (*fasthttp.ResponseHeader, []byte, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.Header.SetContentType("application/x-www-form-urlencoded")
	req.Header.SetMethod("POST")
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	if len(form) != 0 {
		param := ""
		for k, v := range form {
			temp := k + "=" + v
			param = param + "&" + temp
		}
		param = param[1:len(param)]
		req.SetBodyString(param)
	}
	req.SetRequestURI(url)
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)
	if err := fasthttp.Do(req, resp); err != nil {
		return nil, nil, err
	}
	data_Header := new(fasthttp.ResponseHeader)
	*data_Header = resp.Header
	data_Body := resp.Body()
	return data_Header, data_Body, nil
}

func DeadlineGet(url string, headers map[string]string, query map[string]string, Deadtime time.Time) (*fasthttp.ResponseHeader, []byte, error) {
	req := fasthttp.AcquireRequest()
	req.Header.SetContentType("application/x-www-form-urlencoded")
	req.Header.SetMethod("Get")
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	if len(query) > 0 {
		url = url + "?"
		for k, v := range query {
			temp := k + "=" + v
			url = url + "&" + temp
		}
	}
	req.SetRequestURI(url)
	resp := fasthttp.AcquireResponse()
	if err := fasthttp.DoDeadline(req, resp, Deadtime); err != nil {
		return nil, nil, err
	}
	data_Header := new(fasthttp.ResponseHeader)
	*data_Header = resp.Header
	data_Body := resp.Body()
	return data_Header, data_Body, nil
}

func DeadlinePost(url string, headers map[string]string, form map[string]string, Deadtime time.Time) (*fasthttp.ResponseHeader, []byte, error) {
	req := fasthttp.AcquireRequest()
	req.Header.SetContentType("application/x-www-form-urlencoded")
	req.Header.SetMethod("POST")
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	if len(form) != 0 {
		param := ""
		for k, v := range form {
			temp := k + "=" + v
			param = param + "&" + temp
		}
		param = param[1:len(param)]
		req.SetBodyString(param)
	}
	fmt.Println(string(req.Header.Peek("Cookie")))
	req.SetRequestURI(url)
	resp := fasthttp.AcquireResponse()
	if err := fasthttp.DoDeadline(req, resp, Deadtime); err != nil {
		return nil, nil, err
	}
	data_Header := new(fasthttp.ResponseHeader)
	*data_Header = resp.Header
	data_Body := resp.Body()
	return data_Header, data_Body, nil
}

func ClientGet(c *fasthttp.Client, url string, headers map[string]string, query map[string]string) (*fasthttp.ResponseHeader, []byte, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.Header.SetContentType("application/x-www-form-urlencoded")
	req.Header.SetMethod("Get")
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	if len(query) > 0 {
		url = url + "?"
		for k, v := range query {
			temp := k + "=" + v
			url = url + "&" + temp
		}
	}
	req.SetRequestURI(url)
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)
	if err := c.Do(req, resp); err != nil {
		return nil, nil, err
	}
	data_Header := new(fasthttp.ResponseHeader)
	*data_Header = resp.Header
	data_Body := resp.Body()
	return data_Header, data_Body, nil
}

func ClientPost(c *fasthttp.Client, url string, headers map[string]string, form map[string]string) (*fasthttp.ResponseHeader, []byte, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.Header.SetContentType("application/x-www-form-urlencoded")
	req.Header.SetMethod("POST")
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	if len(form) != 0 {
		param := ""
		for k, v := range form {
			temp := k + "=" + v
			param = param + "&" + temp
		}
		param = param[1:len(param)]
		req.SetBodyString(param)
	}
	req.SetRequestURI(url)
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)
	if err := c.Do(req, resp); err != nil {
		return nil, nil, err
	}
	data_Header := new(fasthttp.ResponseHeader)
	*data_Header = resp.Header
	data_Body := resp.Body()
	return data_Header, data_Body, nil
}

func ClientDeadlineGet(c *fasthttp.Client, url string, headers map[string]string, query map[string]string, Deadtime time.Time) (*fasthttp.ResponseHeader, []byte, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.Header.SetContentType("application/x-www-form-urlencoded")
	req.Header.SetMethod("Get")
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	if len(query) > 0 {
		url = url + "?"
		for k, v := range query {
			temp := k + "=" + v
			url = url + "&" + temp
		}
	}
	req.SetRequestURI(url)
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)
	if err := c.DoDeadline(req, resp, Deadtime); err != nil {
		return nil, nil, err
	}
	data_Header := new(fasthttp.ResponseHeader)
	*data_Header = resp.Header
	data_Body := resp.Body()
	return data_Header, data_Body, nil
}

func ClientDeadlinePost(c *fasthttp.Client, url string, headers map[string]string, form map[string]string, Deadtime time.Time) (*fasthttp.ResponseHeader, []byte, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.Header.SetContentType("application/x-www-form-urlencoded")
	req.Header.SetMethod("POST")
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	if len(form) != 0 {
		param := ""
		for k, v := range form {
			temp := k + "=" + v
			param = param + "&" + temp
		}
		param = param[1:len(param)]
		req.SetBodyString(param)
	}
	fmt.Println(string(req.Header.Peek("Cookie")))
	req.SetRequestURI(url)
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)
	if err := c.DoDeadline(req, resp, Deadtime); err != nil {
		return nil, nil, err
	}
	data_Header := new(fasthttp.ResponseHeader)
	*data_Header = resp.Header
	data_Body := resp.Body()
	return data_Header, data_Body, nil
}
