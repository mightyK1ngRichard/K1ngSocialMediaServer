package repository

import (
	"K1ngSochialMediaServer/internal/app/ds"
	"strconv"
)

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

		// Получаем все файлы поста.
		files, err2 := r.GetPostFiles(strconv.Itoa(int(p.ID)))
		if err2 != nil {
			r.logger.Error(err2)
			posts = append(posts, p)
			continue
		}
		p.Files = files

		posts = append(posts, p)
	}

	return &posts, nil
}

func (r *Repository) GetPostFiles(postID string) (*[]ds.PostFiles, error) {
	rows, err := r.db.Query(`SELECT * FROM files_in_post WHERE post_id = $1;`, postID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var files []ds.PostFiles
	for rows.Next() {
		f := ds.PostFiles{}
		if err := rows.Scan(&f.ID,
			&f.URL,
			&f.PostID,
		); err != nil {
			r.logger.Error(err)
			continue
		}
		files = append(files, f)
	}

	return &files, nil
}
