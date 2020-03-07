package tapas

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

// Episode (or chapter) of a comic.
type Episode struct {
	EpisodeMeta
	CommentCount int         `json:"comment_cnt"`
	ContentSize  int         `json:"content_size"`
	Contents     []Image     `json:"contents"`
	NextEpisode  EpisodeMeta `json:"next_episode"`
}

// Episode fetches an chapter of a comic.
func (c *Client) Episode(series, episode int) (Episode, error) {
	raw, err := c.get(fmt.Sprintf("/series/%d/episodes/%d", series, episode))
	if err != nil {
		return Episode{}, err
	}
	var res Episode
	if err := json.Unmarshal(raw, &res); err != nil {
		return Episode{}, errors.Wrap(err, "could not unmarshal episode")
	}
	return res, nil
}
