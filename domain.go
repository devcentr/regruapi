package regruapi

import (
	"encoding/json"
	"github.com/pkg/errors"
)

func (c *Client) DomainGetPrice() (*DomainGetPriceResponse, error) {
	body, err := c.request("domain/get_prices", map[string]string{
		"show_renew_data": "1",
	})
	if err != nil {
		return nil, errors.Wrap(err, "send request")
	}

	response := DomainGetPriceResponse{}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, errors.Wrap(err, "response unmarshal")
	}

	return &response, nil
}

type DomainGetPriceResponse struct {
	Answer struct {
		Currency   string                                  `json:"currency"`
		PriceGroup string                                  `json:"price_group"`
		Prices     map[string]DomainGetPriceResponseEntity `json:"prices"`
	} `json:"answer"`
	Charset      string      `json:"charset"`
	Messagestore interface{} `json:"messagestore"`
	Result       string      `json:"result"`
}

type DomainGetPriceResponseEntity struct {
	ExtcreatePriceEqRenew int     `json:"extcreate_price_eq_renew"`
	Idn                   int     `json:"idn"`
	RegMaxPeriod          int     `json:"reg_max_period"`
	RegMinPeriod          int     `json:"reg_min_period"`
	RegPrice              float64 `json:"reg_price,string"`
	RenewMaxPeriod        int     `json:"renew_max_period"`
	RenewMinPeriod        int     `json:"renew_min_period"`
	RenewPrice            float64 `json:"renew_price,string"`
	RetailRegPrice        float64 `json:"retail_reg_price,string"`
	RetailRenewPrice      float64 `json:"retail_renew_price,string"`
}
