package tapas

// Image with URL and sizes.
type Image struct {
	Size   int    `json:"file_size"`
	URL    string `json:"file_url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}

func (i Image) String() string { return i.URL }
