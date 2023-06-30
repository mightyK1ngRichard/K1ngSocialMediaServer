package ds

type Users struct {
	ID             uint          `json:"id"`
	Nickname       string        `json:"nickname"`
	Description    string        `json:"description,omitempty"`
	Location       string        `json:"location,omitempty"`
	University     string        `json:"university,omitempty"`
	HeaderImage    string        `json:"header_image,omitempty"`
	Avatar         string        `json:"avatar,omitempty"`
	CountOfFriends int           `json:"count_of_friends"`
	Posts          *[]Post       `json:"posts,omitempty"`
	Images         *[]UserImages `json:"images"`
}

// UserImages Фотографии пользователя.
type UserImages struct {
	ID              uint   `json:"id"`
	DatePublic      string `json:"date_public"`
	ImageURL        string `json:"image_name"`
	CountOfLikes    int    `json:"count_of_likes"`
	CountOfComments int    `json:"count_of_comments"`
	UserID          int    `json:"user_id"`
}
