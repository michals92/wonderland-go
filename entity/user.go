package entity

type User struct {
	Address     string   `json:"address"`
	Name        string   `json:"name"`
	VisitedNFTS []string `json:"visited_nfts"`
}
