package regruapi

import (
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

type Client struct {
	apiUri   string
	login    string
	password string
	client   *http.Client
}

func NewClient(login, password string) *Client {
	return &Client{
		apiUri:   "https://api.reg.ru/api/regru2/",
		login:    login,
		password: password,
		client:   &http.Client{},
	}
}

func (c *Client) request(method string, params map[string]string) ([]byte, error) {
	req, err := http.NewRequest("GET", c.apiUri+method, nil)
	if err != nil {
		return []byte(""), errors.Wrap(err, "create new request")
	}

	q := req.URL.Query()
	q.Add("username", c.login)
	q.Add("password", c.password)
	q.Add("output_format", "json")
	for paramName, paramValue := range params {
		q.Add(paramName, paramValue)
	}
	req.URL.RawQuery = q.Encode()

	req.Header.Add("Accept", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		fmt.Println("clientDo")
		fmt.Println(err)
		return []byte(""), err
	}

	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ReadAll body")
		fmt.Println(err)
		return []byte(""), err
	}

	fmt.Println(resp.Status)
	fmt.Println(string(respBody))
	return respBody, nil
}
