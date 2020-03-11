package tapas

// Genre of a comic.
type Genre struct {
	Abbr  string `json:"abbr"`
	Books bool   `json:"books"`
	ID    int    `json:"id"`
	Name  string `json:"name"`
}

func (g Genre) String() string { return g.Name }
