package models

import (
	"errors"
	"html"
)

type Post struct {
	ID         uint64 `json:"id,omitempty"`
	Title      string `json:"title" validate:"required,min=1,max=100"`
	Content    string `json:"content" validate:"required,min=1,max=1000"`
	AuthorID   uint64 `json:"authorID,omitempty"`
	AuthorNick string `json:"authorNick,omitempty"`
	CreatedAt  string `json:"createdAt,omitempty"`
	Likes      int    `json:"likes"`
}

func (post *Post) Prepare() error {
	if error := post.validate(); error != nil {
		return error
	}
	post.format()
	return nil
}

func (post *Post) validate() error {
	if post.Title == "" {
		return errors.New("the title is required")
	}

	if post.Content == "" {
		return errors.New("the content is required")
	}

	return nil
}

func (post *Post) format() {
	post.Title = html.EscapeString(post.Title)
	post.Content = html.EscapeString(post.Content)
	post.AuthorNick = html.EscapeString(post.AuthorNick)
}
