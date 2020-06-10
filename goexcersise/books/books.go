package books

import (
	"encoding/json"
	"strings"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  int    `json:"pages"`
}

func (b *Book) CategoryByLength() string {

	if b.Pages >= 300 {
		return "NOVEL"
	}

	return "SHORT STORY"
}

func NewBookFromJSON(str string) (Book, error) {
	var book Book
	if err := json.NewDecoder(strings.NewReader(str)).Decode(&book); err != nil {
		return Book{}, err
	}
	return book, nil
}
