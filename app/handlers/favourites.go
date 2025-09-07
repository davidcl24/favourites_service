package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/davidcl24/favourites_service/app/models"
	"github.com/go-chi/chi/v5"
)

type FavouriteHandler struct {
	DB *models.DB
}

func (f *FavouriteHandler) ListUserFavourites(w http.ResponseWriter, r *http.Request) {
	userId, _ := strconv.Atoi(chi.URLParam(r, "user_id"))

	favs, err := f.DB.GetAllUserFavourites(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(favs)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (f *FavouriteHandler) GetUserMovieFavourite(w http.ResponseWriter, r *http.Request) {
	userId, _ := strconv.Atoi(chi.URLParam(r, "user_id"))
	movieId, _ := strconv.Atoi(chi.URLParam(r, "movie_id"))

	fav, err := f.DB.GetMovieFavouriteFromUser(userId, movieId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if fav == nil {
		http.Error(w, "Element not found", http.StatusNotFound)
	} else {
		err = json.NewEncoder(w).Encode(fav)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
}

func (f *FavouriteHandler) GetUserShowFavourite(w http.ResponseWriter, r *http.Request) {
	userId, _ := strconv.Atoi(chi.URLParam(r, "user_id"))
	showId, _ := strconv.Atoi(chi.URLParam(r, "show_id"))

	fav, err := f.DB.GetShowFavouriteFromUser(userId, showId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if fav == nil {
		http.Error(w, "Element not found", http.StatusNotFound)
	} else {
		err = json.NewEncoder(w).Encode(fav)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
}

func (f *FavouriteHandler) GetFavourite(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	fav, err := f.DB.GetFavouriteByID(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if fav == nil {
		http.Error(w, "Element not found", http.StatusNotFound)
	} else {
		err = json.NewEncoder(w).Encode(fav)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
}

func (f *FavouriteHandler) CreateFavourite(w http.ResponseWriter, r *http.Request) {
	var favourite models.Favourite

	err := json.NewDecoder(r.Body).Decode(&favourite)
	if err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	createdFav, err := f.DB.AddFavourite(&favourite)
	if err != nil {
		http.Error(w, "Failed to insert element: "+err.Error(), http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(createdFav)
	if err != nil {
		http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func (f *FavouriteHandler) DeleteFavourite(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	fav, err := f.DB.DeleteFavourite(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if fav == nil {
		http.Error(w, "Element not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (f *FavouriteHandler) ClearUserFavourites(w http.ResponseWriter, r *http.Request) {
	userId, _ := strconv.Atoi(chi.URLParam(r, "user_id"))
	historyElements, err := f.DB.ClearUserFavourites(userId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if historyElements == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
