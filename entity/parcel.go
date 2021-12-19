package entity

type PinnedNFT struct {
	NFTaddress string `json:"nft_address,omitempty"`
	NFTid      int    `json:"nft_id,omitempty"`
	Blockchain string `json:"blockchain,omitempty"`
}

type Parcel struct {
	H3Index   int        `json:"h3index"`
	Name      string     `json:"name"`
	Type      string     `json:"type"`
	PinnedNFT *PinnedNFT `json:"pinned_nft,omitempty"`
	Owner     string     `json:"owner"`
}
