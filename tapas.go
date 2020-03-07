// Package tapas is a client for Tapas API at api.tapas.io.
package tapas

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

// Client implements a way to talk to Tapas' API.
type Client struct {
	base, path string
	client     *http.Client

	deviceType, deviceUUID string
}

// An OptionFunc can be used to modify the client.
type OptionFunc func(*Client)

// WithBase sets the API base.
func WithBase(base string) OptionFunc {
	return func(c *Client) { c.base = base }
}

// WithPath replaces the default path. Might be used on a new API version.
func WithPath(path string) OptionFunc {
	return func(c *Client) { c.path = path }
}

// WithDeviceType sets the device type. Defaults to "ANDROID".
func WithDeviceType(deviceType string) OptionFunc {
	return func(c *Client) { c.deviceType = deviceType }
}

// WithDeviceUUID sets the device UUID. Defaults to an empty string, please use
// 16 character long hex string if needed.
func WithDeviceUUID(deviceUUID string) OptionFunc {
	return func(c *Client) { c.deviceUUID = deviceUUID }
}

// WithHTTPClient makes the manga client use a given http.Client to make
// requests.
func WithHTTPClient(client *http.Client) OptionFunc {
	return func(c *Client) { c.client = client }
}

// New returns a new Tapas Client.
func New(options ...OptionFunc) *Client {
	c := &Client{
		base:   "https://api.tapas.io",
		path:   "/v3",
		client: http.DefaultClient,
	}
	for _, option := range options {
		option(c)
	}
	return c
}

// get sends a HTTP GET request.
func (c *Client) get(endpoint string) (json.RawMessage, error) {
	req, err := http.NewRequest(http.MethodGet, c.base+c.path+endpoint, nil)
	if err != nil {
		return nil, errors.Wrap(err, "could not create get request")
	}
	req.Header.Add("accept", "application/panda+json")
	req.Header.Add("x-device-type", c.deviceType)
	req.Header.Add("x-device-uuid", c.deviceUUID)
	res, err := c.client.Do(req)
	if err != nil {
		return nil, errors.Wrapf(err, "could not get %s", req.URL)
	}
	defer res.Body.Close()
	var raw json.RawMessage
	if err := json.NewDecoder(res.Body).Decode(&raw); err != nil {
		return nil, errors.Wrap(err, "could not decode response")
	}
	return raw, nil
}
