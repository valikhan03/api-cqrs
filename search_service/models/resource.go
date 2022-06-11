package models

type Resource struct {
	ID string `json:"id"`
	Title     string   `json:"title"`
	Author    string   `json:"author"`
	Content   string   `json:"content"`
	CreatedAt string   `json:"created_at"`
	Tags      []string `json:"tags"`
}

type ResourceData struct {
	Title     string   `json:"title"`
	Author    string   `json:"author"`
	Content   string   `json:"content"`
	Tags      []string `json:"tags"`
}

type Filter struct {
	Title   string   `json:"title"`
	Author  string   `json:"author"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}
