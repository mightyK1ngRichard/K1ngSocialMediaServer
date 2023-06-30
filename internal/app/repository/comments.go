package repository

import (
	"K1ngSochialMediaServer/internal/app/ds"
	"database/sql"
)

func (r *Repository) GetAllCommentsOfPostsByUserID(userID string) (*[]ds.Comments, error) {
	rows, err := r.db.Query(`
	SELECT c.id, c.date_public, c.content, c.count_of_likes, c.user_id, c.post_id, u.avatar, u.nickname
	FROM comments_under_post c
	LEFT JOIN users u on u.id = c.user_id
	WHERE user_id = $1
	ORDER BY c.date_public DESC;
	`, userID)
	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		if err := rows.Close(); err != nil {
			r.logger.Errorf("error of the closure the connection of rows in GetAllCommentsByUserID: %s", err)
		}

	}(rows)

	var comments []ds.Comments
	for rows.Next() {
		c := ds.Comments{}
		if err := rows.Scan(&c.ID,
			&c.DatePublic,
			&c.Content,
			&c.CountOfLikes,
			&c.UserID,
			&c.PostID,
			&c.UserAvatar,
			&c.Nickname,
		); err != nil {
			r.logger.Error(err)
			continue
		}
		comments = append(comments, c)
	}

	return &comments, nil
}
