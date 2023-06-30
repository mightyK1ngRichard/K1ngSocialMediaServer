package repository

import "K1ngSochialMediaServer/internal/app/ds"

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
		users = append(users, u)
	}

	return &users, nil
}

func (r *Repository) GetUserById(userId string) (*ds.Users, error) {
	u := ds.Users{}
	row := r.db.QueryRow(`SELECT * from users WHERE id = $1`, userId)
	if err := row.Scan(&u.ID,
		&u.Nickname,
		&u.Description,
		&u.Location,
		&u.University,
		&u.HeaderImage,
		&u.Avatar,
		&u.CountOfFriends); err != nil {
		return nil, err
	}

	return &u, nil
}
