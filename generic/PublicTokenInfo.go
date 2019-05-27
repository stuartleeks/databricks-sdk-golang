package generic

type PublicTokenInfo struct {
	TokenID      string `json:"token_id,omitempty"`
	CreationTime int64  `json:"creation_time,omitempty"`
	ExpiryTime   int64  `json:"expiry_time,omitempty"`
	Comment      string `json:"comment,omitempty"`
}
