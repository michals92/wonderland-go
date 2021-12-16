package entity

type Parcel struct {
	H3Index   string `json:"h3index"`
	Name      string `json:"name"`
	PinnedNFT string `json:"pinned_nft"`
	Owner     string `json:"owner"`
}
