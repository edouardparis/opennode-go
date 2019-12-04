package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/edouardparis/opennode-go/opennode"
)

type Client struct {
	client   *http.Client
	apikey   string
	endpoint string
}

func New(apiKey string, env opennode.Env) *Client {
	return &Client{
		apikey:   apiKey,
		client:   http.DefaultClient,
		endpoint: string(env),
	}
}

func (c *Client) CreateCharge(payload *opennode.ChargePayload) (*opennode.Charge, error) {
	url := fmt.Sprintf("%s/v1/charges", c.endpoint)

	p, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(p))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", c.apikey)
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("error during charge creation: status: %d, body %s",
			resp.StatusCode, string(body))
	}

	resource := struct {
		Data opennode.Charge `json:"data"`
	}{}
	err = json.Unmarshal(body, &resource)
	if err != nil {
		return nil, err
	}
	return &resource.Data, nil
}
