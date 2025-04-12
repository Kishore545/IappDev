package handlers

import (
	"IAppDev/models"
	"encoding/json"
	"net/http"
)

var posts = []models.Post{}
var postID = 1

func CreatePost(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	json.NewDecoder(r.Body).Decode(&post)

	post.ID = postID
	postID++
	posts = append(posts, post)

	json.NewEncoder(w).Encode(post)
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(posts)
}
