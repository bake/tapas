package tapas

// Creator of a comic.
type Creator struct {
	ID               int    `json:"id"`
	Name             string `json:"uname"`
	DisplayName      string `json:"display_name"`
	JoinedCreatorTip bool   `json:"joined_creator_tip"`
	ProfilePic       string `json:"profile_pic_url"`
	Staff            bool   `json:"staff"`
	SupportBanner    Image  `json:"support_banner"`
	SubscriberCount  int    `json:"subscriber_cnt"`
}

func (c Creator) String() string { return c.DisplayName }
