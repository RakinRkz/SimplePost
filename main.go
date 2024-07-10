package main

import (
	"fmt"
	"net/http"
)

// BlogPost represents a single blog post.
type BlogPost struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt string    `json:"created_at"`
}

var posts []BlogPost

func main() {
	http.HandleFunc("/posts", handlePosts)
	http.HandleFunc("/posts/", handlePost)
	http.ListenAndServe(":8080", nil)
}

func handlePosts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getAllPosts(w, r)
	case "POST":
		createPost(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func getAllPosts(w http.ResponseWriter, _ *http.Request) {
	json.NewEncoder(w).Encode(posts)
}

func createPost(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var p BlogPost
	err := decodeJSON(r.Body, &p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	posts = append(posts, p)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Created post %v\n", p.Title)
}

func decodeJSON(body io.Reader, post *BlogPost) error {
	return json.NewDecoder(body).Decode(post)
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/posts/")
	if id == "" || id == "favicon.ico" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	for _, post := range posts {
		if post.ID == atoi(id) {
			json.NewEncoder(w).Encode(post)
			return
		}
	}
	http.Error(w, "Not found", http.StatusNotFound)
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
