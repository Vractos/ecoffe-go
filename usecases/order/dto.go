package order

type CreateOrderDtoInput struct {
	Client      string  `json:"client"`
	Item        string  `json:"item"`
	Observation *string `json:"observation"`
	Quantity    int     `json:"quantity"`
}
