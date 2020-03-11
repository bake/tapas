package tapas

// Tag of a comic.
type Tag struct {
	ID           int64   `json:"id"`
	Title        string  `json:"title"`
	Type         string  `json:"type"`
	SalesType    string  `json:"sale_type"`
	SubscribeCnt int     `json:"subscribe_cnt"`
	LikeCnt      int     `json:"like_cnt"`
	ViewCnt      int     `json:"view_cnt"`
	Description  string  `json:"description"`
	ThumbURL     string  `json:"thumb_url"`
	BookCoverURL string  `json:"book_cover_url"`
	ReviewRating int     `json:"review_rating"`
	ReviewCnt    int     `json:"review_cnt"`
	Creator      Creator `json:"creator"`
	Genre        Genre   `json:"genre"`
}

func (t Tag) String() string { return t.Title }
