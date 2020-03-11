package tapas

import (
	"encoding/json"
	"net/url"
)

// Search for books and comics.
func (c *Client) Search(query string) (books []Series, comics []Series, err error) {
	raw, err := c.get("/search", url.Values{"q": {query}})
	if err != nil {
		return nil, nil, err
	}
	var res struct {
		Books []struct {
			Series Series `json:"series"`
		} `json:"books"`
		Comics []struct {
			Series Series `json:"series"`
		} `json:"comics"`
	}
	if err := json.Unmarshal(raw, &res); err != nil {
		return nil, nil, err
	}

	books = make([]Series, len(res.Books))
	for i, b := range res.Books {
		books[i] = b.Series
	}
	comics = make([]Series, len(res.Comics))
	for i, b := range res.Comics {
		comics[i] = b.Series
	}
	return books, comics, nil
}
