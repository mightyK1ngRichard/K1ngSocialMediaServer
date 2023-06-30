package repository

import "K1ngSochialMediaServer/internal/app/ds"

func (r *Repository) GetPostsOfUser(userID string) (*[]ds.Post, error) {
	rows, err := r.db.Query(`SELECT * FROM posts WHERE user_id = $1;`, userID)
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
		); err != nil {
			r.logger.Error(err)
			continue
		}
		posts = append(posts, p)
	}

	return &posts, nil
}
