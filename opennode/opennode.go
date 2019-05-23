package opennode

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	DEV_API_ENDPOINT = "https://dev-api.opennode.co/v1"
	API_ENDPOINT     = "https://api.opennode.co/v1"
)

type Client struct {
	client   *http.Client
	APIKey   string
	Endpoint string
}

func (c *Client) CreateCharge(payload *ChargePayload) (*Charge, error) {
	url := fmt.Sprintf("%s/charges", c.Endpoint)

	p, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(p))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", c.APIKey)
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
		Data Charge `json:"data"`
	}{}
	err = json.Unmarshal(body, &resource)
	if err != nil {
		return nil, err
	}
	return &resource.Data, nil
}

func (c *Client) UpdateCharge(ch *Charge) error {
	if ch == nil {
		return nil
	}

	url := fmt.Sprintf("%s/charge/%s", c.Endpoint, ch.ID)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", c.APIKey)
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf(
			"error during charge update: status: %d, body: %s",
			resp.StatusCode, string(body))
	}

	resource := struct {
		Data Charge `json:"data"`
	}{}
	err = json.Unmarshal(body, &resource)
	if err != nil {
		return err
	}

	*ch = resource.Data

	return nil
}

type Config struct {
	Testnet bool   `json:"debug"`
	APIKey  string `json:"api_key"`
}

func NewClient(c *Config) *Client {
	client := &Client{
		APIKey:   c.APIKey,
		client:   http.DefaultClient,
		Endpoint: API_ENDPOINT,
	}
	if c.Testnet {
		client.Endpoint = DEV_API_ENDPOINT
	}
	return client
}
