package presenter

type Announcement struct {
	ID          string  `json:"id"`
	Client      string  `json:"client"`
	Item        string  `json:"item"`
	Observation *string `json:"observation"`
	Quantity    int     `json:"quantity"`
	Status      string  `json:"status"`
	DateCreated string  `json:"date_created"`
}
