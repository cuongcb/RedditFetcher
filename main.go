package main

import (
	"fmt"

	"github.com/turnage/graw/reddit"
)

// where is the sorting kind of the posts/comments/messages
type where string

func (w where) toString() string {
	return string(w)
}

const (
	best          where = "best"
	top           where = "top"
	new           where = "new"
	controversial where = "controversial"
	old           where = "old"
)

const (
	defaultLimit string = "25"
	defaultDepth string = "10"
)

func main() {
	bot, err := reddit.NewBotFromAgentFile(".env", 0)
	if err != nil {
		panic(err)
	}

	options := map[string]string{
		"limit": defaultLimit,
		"sort":  best.toString(),
		"depth": defaultDepth,
	}

	harvest, err := bot.ListingWithParams("/r/golang",
		options)
	if err != nil {
		panic(err)
	}

	showPosts(harvest.Posts)
}

func showPosts(posts []*reddit.Post) {
	for _, post := range posts {
		newline := false
		titleLen := len(post.Title)
		if titleLen > 80 {
			titleLen = 80
			newline = true
		}
		fmt.Printf("> %s \n", post.Title[:titleLen])
		if newline {
			fmt.Println("...")
		}
	}
}
