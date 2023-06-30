package repository

import "K1ngSochialMediaServer/internal/app/ds"

func (r *Repository) GetPostsOfUser(userID string) (*[]ds.Post, error) {
	rows, err := r.db.Query(`
	SELECT p.id, p.date_public, p.content, p.count_of_likes, p.count_of_comments, p.user_id, u.avatar, u.nickname
	FROM posts p
	LEFT JOIN users u on p.user_id = u.id
	WHERE user_id = $1
	ORDER BY p.date_public DESC;	
	`, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var posts []ds.Post
	for rows.Next() {
		p := ds.Post{}
		if err := rows.Scan(&p.ID,
			&p.DatePublic,
			&p.Content,
			&p.CountOfLikes,
			&p.CountOfComments,
			&p.UserID,
			&p.UserAvatar,
			&p.Nickname,
		); err != nil {
			r.logger.Error(err)
			continue
		}
		posts = append(posts, p)
	}

	return &posts, nil
}
