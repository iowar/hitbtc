package hitbtc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	BaseURL = "https://api.hitbtc.com"
)

var (
	spr      = fmt.Sprintf
	throttle = time.Tick(time.Second / 9)
)

type HitBtc struct {
	api_key    string
	api_secret string
	httpClient *http.Client
}

func NewClient(api_key, api_secret string) (client *HitBtc, err error) {

	client = &HitBtc{
		api_key:    api_key,
		api_secret: api_secret,
		httpClient: &http.Client{Timeout: time.Second * 15},
	}

	return
}

func (h *HitBtc) publicRequest(action string, respch chan<- []byte, errch chan<- error) {

	<-throttle

	defer close(respch)
	defer close(errch)

	rawurl := BaseURL + action

	req, err := http.NewRequest("GET", rawurl, nil)
	if err != nil {
		respch <- nil
		errch <- Error(RequestError)
		return
	}

	req.Header.Add("Accept", "application/json")

	resp, err := h.httpClient.Do(req)
	if err != nil {
		respch <- nil
		errch <- Error(ConnectError)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {

		respch <- body
		errch <- err

		return
	}

	respch <- body
	errch <- nil
}

type serverError struct {
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

func checkServerError(response []byte) error {
	var check serverError

	err := json.Unmarshal(response, &check)
	if err != nil {
		return nil
	}
	if check.Error.Message != "" {
		return Error(ServerError, check.Error.Message)
	} else {
		return nil
	}
}

func (h *HitBtc) tradeRequest(
	method, action string,
	parameters map[string]string,
	respch chan<- []byte,
	errch chan<- error,
) {

	<-throttle

	defer close(respch)
	defer close(errch)

	if len(h.api_key) == 0 || len(h.api_secret) == 0 {
		respch <- nil
		errch <- Error(SetApiError)
	}

	rawurl := spr("%s%s", BaseURL, action)
	method = strings.ToUpper(method)

	formValues := url.Values{}
	for k, v := range parameters {
		formValues.Set(k, v)
	}

	formData := formValues.Encode()

	if method == "GET" && parameters != nil {
		var URL *url.URL
		URL, err := url.Parse(rawurl)
		if err != nil {
			respch <- nil
			errch <- Error(UrlParseError)
		}

		URL.RawQuery = formData
		rawurl = URL.String()
	}

	req, err := http.NewRequest(method, rawurl,
		strings.NewReader(formData))
	if err != nil {
		respch <- nil
		errch <- Error(RequestError)
		return
	}

	req.Header.Add("Accept", "application/json")
	req.SetBasicAuth(h.api_key, h.api_secret)

	resp, err := h.httpClient.Do(req)
	if err != nil {
		respch <- nil
		errch <- Error(ConnectError)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		respch <- body
		errch <- err
		return
	}

	err = checkServerError(body)
	if err != nil {
		respch <- nil
		errch <- err
	}

	respch <- body
	errch <- nil
}
