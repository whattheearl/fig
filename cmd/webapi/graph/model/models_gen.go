// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type NewProfile struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type Post struct {
	ID      string   `json:"id"`
	Author  *Profile `json:"author"`
	Message string   `json:"message"`
	Image   string   `json:"image"`
}

type Posts struct {
	Result []*Post `json:"result"`
	Count  int     `json:"count"`
	Cursor string  `json:"cursor"`
}

type Profile struct {
	Name     string  `json:"name"`
	Username string  `json:"username"`
	Summary  string  `json:"summary"`
	Posts    []*Post `json:"posts"`
}
