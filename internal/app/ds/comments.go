package ds

type Comments struct {
	ID           uint   `json:"id"`
	DatePublic   string `json:"date_public"`
	Content      string `json:"content"`
	CountOfLikes int    `json:"count_of_likes"`
	UserID       int    `json:"user_id"`
	PostID       int    `json:"post_id"`
	UserAvatar   string `json:"avatar"`
	Nickname     string `json:"nickname"`
}
