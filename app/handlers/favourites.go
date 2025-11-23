// It handles HTTP requests and makes a call to the desired CRUD operation
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

// Handles the specific GET HTTP request and returns a list of all favourites from the desired user in a JSON list.
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

// Handles the specific GET HTTP request and returns a favourite movie from the desired user in a JSON.
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

// Handles the specific GET HTTP request and returns a favourite show from the desired user in a JSON.
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

// Handles the specific GET HTTP request and returns a particular favourite element in a JSON.
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

// Handles the specific POST HTTP request to create a new favourite element and returns it in JSON format if successful.
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

// Handles the specific DELETE HTTP request to remove a specific favourite element
// and returns the desired HTTP status code if successful.
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

// Handles the specific DELETE HTTP request to remove all favourite elements from a specific user
// and returns the desired HTTP status code if successful.
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
