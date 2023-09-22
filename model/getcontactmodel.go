package model

type Pagination struct {
	Limit     int    `json:"limit"`
	Page      int    `json:"page"`
	Total     int    `json:"total"`
	Next_Page string `json:"next_page"`
}
type Content struct {
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	Property_id string `json:"property_id"`
	Message     string `json:"message"`
	Source      string `json:"source"`
	Happened_at string `json:"happened_at"`
}
type Contact struct {
	Pagination Pagination `json:"pagination"`
	Content    []Content  `json:"content"`
}
