package models

type NaverLocalSearchResponse struct {
	LastBuildDate string `json:"lastBuildDate"`
	Total         int    `json:"total"`
	Start         int    `json:"start"`
	Display       int    `json:"display"`
	Items         []NaverLocalSearchItem `json:"items"`
}

type NaverLocalSearchItem struct {
	Title       string `json:"title"`
	Link        string `json:"link"`
	Category    string `json:"category"`
	Description string `json:"description"`
	Telephone   string `json:"telephone"`
	Address     string `json:"address"`
	RoadAddress string `json:"roadAddress"`
	MapX        string `json:"mapx"`
	MapY        string `json:"mapy"`
}
