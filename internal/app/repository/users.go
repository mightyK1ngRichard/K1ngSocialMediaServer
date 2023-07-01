package repository

import (
	"K1ngSochialMediaServer/internal/app/ds"
	"K1ngSochialMediaServer/internal/utils"
	"strconv"
)

func (r *Repository) GetAllUsers() (*[]ds.Users, error) {
	rows, err := r.db.Query(`SELECT * FROM users;`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var users []ds.Users
	for rows.Next() {
		u := ds.Users{}
		if err := rows.Scan(&u.ID,
			&u.Nickname,
			&u.Description,
			&u.Location,
			&u.University,
			&u.HeaderImage,
			&u.Avatar,
			&u.CountOfFriends); err != nil {
			r.logger.Error(err)
			continue
		}

		// Получаем посты пользователя.
		posts, err2 := r.GetPostsOfUser(strconv.Itoa(int(u.ID)))
		if err2 != nil {
			r.logger.Error(err2)
			users = append(users, u)
			continue
		}
		u.Posts = posts

		// Получаем фотографии юзера.
		files, err3 := r.GetUserImages(strconv.Itoa(int(u.ID)))
		if err3 != nil {
			r.logger.Error(err3)
			users = append(users, u)
			continue
		}
		u.Images = files

		users = append(users, u)
	}

	return &users, nil
}

func (r *Repository) GetUserById(userId string) (*ds.Users, error) {
	u := ds.Users{}
	// TODO: Сделать async.
	row := r.db.QueryRow(`SELECT * from users WHERE id = $1`, userId)
	images, err2 := r.GetUserImages(userId)

	if err := row.Scan(&u.ID,
		&u.Nickname,
		&u.Description,
		&u.Location,
		&u.University,
		&u.HeaderImage,
		&u.Avatar,
		&u.CountOfFriends,
	); err != nil {
		return nil, err
	}

	if err2 != nil {
		r.logger.Error(err2)
		return &u, nil
	}
	u.Images = images
	return &u, nil
}

func (r *Repository) GetUserImages(userId string) (*[]ds.UserImages, error) {
	rows, err := r.db.Query(`SELECT * FROM user_images WHERE user_id = $1;`, userId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var images []ds.UserImages
	for rows.Next() {
		i := ds.UserImages{}
		if err := rows.Scan(&i.ID,
			&i.DatePublic,
			&i.ImageURL,
			&i.CountOfLikes,
			&i.CountOfComments,
			&i.UserID,
		); err != nil {
			r.logger.Error(err)
			continue
		}
		images = append(images, i)

	}

	return &images, nil
}

func (r *Repository) AddUserImage(expansion, userId string) (string, error) {
	imageName := utils.GenerateUniqueFileName(expansion)
	_, err := r.db.Exec(`INSERT INTO user_images (image_name, user_id) VALUES ($1, $2);`, imageName, userId)
	if err != nil {
		return "", err
	}
	return imageName, nil
}
