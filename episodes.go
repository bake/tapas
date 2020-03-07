package tapas

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/pkg/errors"
)

// EpisodeMeta holds metadata of a chapter.
type EpisodeMeta struct {
	CreatedDate         time.Time `json:"created_date"`
	EarlyAccess         bool      `json:"early_access"`
	Free                bool      `json:"free"`
	ID                  int       `json:"id"`
	NSFW                bool      `json:"nsfw"`
	Nu                  bool      `json:"nu"`
	Read                bool      `json:"read"`
	Scene               int       `json:"scene"`
	ScheduledDate       time.Time `json:"scheduled_date"`
	SupportSupportingAd bool      `json:"support_supporting_ad"`
	Thumbnail           Image     `json:"thumb"`
	Title               string    `json:"title"`
	Unlocked            bool      `json:"unlocked"`
}

func (e EpisodeMeta) String() string { return e.Title }

func (c *Client) Episodes(series int) ([]EpisodeMeta, error) {
	raw, err := c.get(fmt.Sprintf("/series/%d/episodes", series))
	if err != nil {
		return nil, err
	}
	var res []EpisodeMeta
	if err := json.Unmarshal(raw, &res); err != nil {
		return nil, errors.Wrap(err, "could not unmarshal episodes")
	}
	return res, nil
}
