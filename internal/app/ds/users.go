package ds

type Users struct {
	ID             uint   `json:"id"`
	Nickname       string `json:"nickname"`
	Description    string `json:"description"`
	Location       string `json:"location"`
	University     string `json:"university"`
	HeaderImage    string `json:"header_image"`
	Avatar         string `json:"avatar"`
	CountOfFriends int    `json:"count_of_friends"`
}
