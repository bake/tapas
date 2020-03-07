package tapas

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/pkg/errors"
)

// Series of a comic.
type Series struct {
	AgeRating                AgeRating `json:"age_rating"`
	AvailableImpression      bool      `json:"available_impression"`
	BackgroundURL            string    `json:"background_url"`
	Blurb                    string    `json:"blurb"`
	BookCoverURL             string    `json:"book_cover_url"`
	Bookmarked               bool      `json:"bookmarked"`
	Claimed                  bool      `json:"claimed"`
	Colophon                 string    `json:"colophon"`
	CommentCount             int       `json:"comment_cnt"`
	Completed                bool      `json:"completed"`
	Creators                 []Creator `json:"creators"`
	DescOrder                bool      `json:"desc_order"`
	Description              string    `json:"description"`
	DiscountRate             int       `json:"discount_rate"`
	DisplayAd                bool      `json:"display_ad"`
	EarlyAccessEpCount       int       `json:"early_access_ep_cnt"`
	EpisodeCount             int       `json:"episode_cnt"`
	Genre                    Genre     `json:"genre"`
	HasNewEpisode            bool      `json:"has_new_episode"`
	HumanURL                 string    `json:"human_url"`
	ID                       int       `json:"id"`
	LastEpisodeScheduledDate time.Time `json:"last_episode_scheduled_date"`
	LastEpisodeUpdatedDate   time.Time `json:"last_episode_updated_date"`
	LikeCount                int       `json:"like_cnt"`
	MasterKeyBanner          bool      `json:"master_key_banner"`
	MustPayCount             int       `json:"must_pay_cnt"`
	NewEpisodeCount          int       `json:"new_episode_cnt"`
	NotificationOn           bool      `json:"notification_on"`
	OnSale                   bool      `json:"on_sale"`
	Original                 bool      `json:"original"`
	PrivateReading           bool      `json:"private_reading"`
	PublishDays              []string  `json:"publish_days"`
	RectBannerURL            string    `json:"rect_banner_url"`
	Restricted               bool      `json:"restricted"`
	ReviewRating             Rating    `json:"review_rating"`
	RgbHex                   string    `json:"rgb_hex"`
	SaleEndDate              time.Time `json:"sale_end_date"`
	SaleStartDate            time.Time `json:"sale_start_date"`
	SaleType                 string    `json:"sale_type"`
	SpLikeCount              int       `json:"sp_like_cnt"`
	SubTitle                 string    `json:"sub_title"`
	SubscribeCount           int       `json:"subscribe_cnt"`
	SupportingAd             Image     `json:"supporting_ad"`
	SupportingAdLink         string    `json:"supporting_ad_link"`
	Tags                     []string  `json:"tags"`
	Thumbnail                Image     `json:"thumb"`
	Title                    string    `json:"title"`
	Type                     string    `json:"type"`
	UnusedKeyCount           int       `json:"unused_key_cnt"`
	Up                       bool      `json:"up"`
	UpdatedDate              time.Time `json:"updated_date"`
	ViewCount                int       `json:"view_cnt"`
	WopInterval              int       `json:"wop_interval"`
}

func (s Series) String() string { return s.Title }

// Creator of a comic.
type Creator struct {
	DisplayName      string `json:"display_name"`
	ID               int    `json:"id"`
	JoinedCreatorTip bool   `json:"joined_creator_tip"`
	ProfilePicURL    string `json:"profile_pic_url"`
	Staff            bool   `json:"staff"`
	SupportBanner    Image  `json:"support_banner"`
	Uname            string `json:"uname"`
}

func (c Creator) String() string { return c.DisplayName }

// Genre of a comic.
type Genre struct {
	Abbr  string `json:"abbr"`
	Books bool   `json:"books"`
	ID    int    `json:"id"`
	Name  string `json:"name"`
}

func (g Genre) String() string { return g.Name }

// AgeRating of a comic.
type AgeRating struct {
	IconPath string `json:"icon_path"`
	Name     string `json:"name"`
}

func (a AgeRating) String() string { return a.Name }

// Rating of a comic.
type Rating struct {
	ID     int     `json:"id"`
	Count  int     `json:"cnt"`
	Rating float64 `json:"rating"`
}

// Series fetches a series.
func (c *Client) Series(series int) (Series, error) {
	raw, err := c.get(fmt.Sprintf("/series/%d", series))
	if err != nil {
		return Series{}, err
	}
	var res Series
	if err := json.Unmarshal(raw, &res); err != nil {
		return Series{}, errors.Wrap(err, "could not unmarshal series")
	}
	return res, nil
}
