package handlers

import (
	"encoding/json"
	"net/http"
	"server/helpers"
	"server/models"
	"server/repository"
	"server/server"

	"github.com/segmentio/ksuid"
)

type CreatePostRequest struct {
	PostContent string `json:"post_content"`
}

type PostResponse struct {
	Id string `json:"id"`
	PostContent string `json:"post_content"`
}

func CreatePostHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := helpers.GetTokenFromRequest(s, w, r)

		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(*models.AppClaims); ok && token.Valid {
			postRequest := CreatePostRequest{}

			if err := json.NewDecoder(r.Body).Decode(&postRequest); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			id, err := ksuid.NewRandom()

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			post := models.Post{
				Id: id.String(),
				PostContent: postRequest.PostContent,
				UserId: claims.UserId,
			}

			err = repository.CreatePost(r.Context(), &post)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(PostResponse{
				Id: post.Id,
				PostContent: post.PostContent,
			})
		} else {
			http.Error(w, err.Error(), http.StatusUnauthorized)
		}
	}
}