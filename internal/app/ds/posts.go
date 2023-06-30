package ds

type Post struct {
	ID              uint         `json:"id"`
	DatePublic      string       `json:"date_public"`
	Content         string       `json:"content,omitempty"`
	CountOfLikes    int          `json:"count_of_likes"`
	CountOfComments int          `json:"count_of_comments"`
	UserID          int          `json:"user_id"`
	UserAvatar      string       `json:"avatar,omitempty"`
	Nickname        string       `json:"nickname"`
	Files           *[]PostFiles `json:"files,omitempty"`
}

type PostFiles struct {
	ID     uint   `json:"id"`
	URL    string `json:"file_name"`
	PostID int    `json:"post_id"`
}
