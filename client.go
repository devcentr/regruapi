package regruapi

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Client struct {
	apiUri    string
	login     string
	password  string
	transport *http.Client
}

func NewClient(login, password string) *Client {
	return &Client{
		apiUri:    "https://api.reg.ru/api/regru2/",
		login:     login,
		password:  password,
		transport: &http.Client{},
	}
}

func (c *Client) request(method string, params map[string]string, responseInterface interface{}) (interface{}, error) {
	req, err := http.NewRequest("GET", c.apiUri+method, nil)
	if err != nil {
		return nil, errors.Wrap(err, "create new request")
	}

	req.URL.RawQuery = c.prepareParams(params)
	req.Header.Add("Accept", "application/json")

	resp, err := c.transport.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "transport do")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "read response body")
	}

	if err := json.Unmarshal(body, responseInterface); err != nil {
		return nil, errors.Wrap(err, "response body unmarshal")
	}

	return responseInterface, nil
}

func (c *Client) prepareParams(params map[string]string) string {
	q := url.Values{}
	q.Add("username", c.login)
	q.Add("password", c.password)
	q.Add("output_format", "json")
	for paramName, paramValue := range params {
		q.Add(paramName, paramValue)
	}
	return q.Encode()
}
