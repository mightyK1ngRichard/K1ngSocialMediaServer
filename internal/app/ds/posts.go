package ds

type Post struct {
	ID              uint   `json:"id"`
	DatePublic      string `json:"date_public"`
	Content         string `json:"content"`
	CountOfLikes    int    `json:"count_of_likes"`
	CountOfComments int    `json:"count_of_comments"`
	UserID          int    `json:"user_id"`
}
