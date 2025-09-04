package models

import (
	"errors"
	"net"
)

type Favourite struct {
	ID      int  `json:"id"`
	UserID  int  `json:"user_id"`
	MovieID *int `json:"movie_id"`
	ShowID  *int `json:"show_id"`
}

func (db *DB) GetAllUserFavourites(userId int) ([]*Favourite, error) {
	query := `
		SELECT id, user_id, movie_id, show_id
		FROM favourites
		WHERE user_id = $1`
	rows, err := db.Conn.Query(query, userId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	favourites := []*Favourite{}

	for rows.Next() {
		fav := &Favourite{}
		err := rows.Scan(&fav.ID, &fav.UserID, &fav.MovieID, &fav.ShowID)
		if err != nil {
			continue
		}
		favourites = append(favourites, fav)
	}
	return favourites, nil
}

func (db *DB) GetMovieFavouriteFromUser(userId int, movieId int) (*Favourite, error) {
	query := `
		SELECT id, user_id, movie_id, show_id
		FROM favourites
		WHERE user_id = $1 AND movie_id = $2`
	fav := &Favourite{}

	err := db.Conn.QueryRow(query, userId, movieId).Scan(&fav.ID, &fav.UserID, &fav.MovieID, &fav.ShowID)
	var opErr *net.OpError
	if errors.As(err, &opErr) {
		return nil, err
	}
	if err != nil {
		return nil, nil
	}
	return fav, nil
}

func (db *DB) GetShowFavouriteFromUser(userId int, showId int) (*Favourite, error) {
	query := `
		SELECT id, user_id, movie_id, show_id
		FROM favourites
		WHERE user_id = $1 AND show_id = $2`
	fav := &Favourite{}

	err := db.Conn.QueryRow(query, userId, showId).Scan(&fav.ID, &fav.UserID, &fav.MovieID, &fav.ShowID)
	var opErr *net.OpError
	if errors.As(err, &opErr) {
		return nil, err
	}
	if err != nil {
		return nil, nil
	}
	return fav, nil
}

func (db *DB) GetFavouriteByID(id int) (*Favourite, error) {
	query := `
		SELECT id, user_id, movie_id, show_id
		FROM favourites
		WHERE id = $1`

	fav := &Favourite{}

	err := db.Conn.QueryRow(query, id).Scan(&fav.ID, &fav.UserID, &fav.MovieID, &fav.ShowID)
	var opErr *net.OpError
	if errors.As(err, &opErr) {
		return nil, err
	}
	if err != nil {
		return nil, nil
	}
	return fav, nil
}

func (db *DB) AddFavourite(favourite *Favourite) (*Favourite, error) {
	query := `
		INSERT INTO favourites (user_id, movie_id, show_id)
		VALUES ($1, $2, $3)
		RETURNING id`
	err := db.Conn.QueryRow(
		query,
		favourite.UserID,
		favourite.MovieID,
		favourite.ShowID,
	).Scan(&favourite.ID)

	if err != nil {
		return nil, err
	}
	return favourite, nil
}

func (db *DB) DeleteFavourite(id int) (*Favourite, error) {
	fav, err := db.GetFavouriteByID(id)

	if err != nil {
		return nil, err
	}

	if fav == nil {
		return nil, nil
	}

	_, err = db.Conn.Exec("DELETE FROM favourites WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	return fav, nil
}

func (db *DB) ClearUserFavourites(userId int) ([]*Favourite, error) {
	favourites, err := db.GetAllUserFavourites(userId)

	if err != nil {
		return nil, err
	}

	if len(favourites) == 0 {
		return nil, nil
	}

	_, err = db.Conn.Exec("DELETE FROM favourites WHERE user_id = $1", userId)
	if err != nil {
		return nil, err
	}
	return favourites, nil
}
