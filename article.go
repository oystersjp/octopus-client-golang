package octopus

type Article struct {
	Id            int    `json:"id"`
	Title         string `json:"title"`
	Summary       string `json:"summary"`
	Url           string `json:"url"`
	PublishedAt   string `json:"publishedAt"`
	PublisherName string `json:"publisherName"`
	PublisherUrl  string `json:"publisherUrl"`
}
