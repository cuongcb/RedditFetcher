package main

import (
	"fmt"

	"github.com/turnage/graw/reddit"
)

func main() {
	bot, err := reddit.NewBotFromAgentFile("rf.env", 0)
	if err != nil {
		fmt.Println("Failed to create bot handle: ", err)
		return
	}

	options := map[string]string{
		"limit": "10",
		"sort":  "best",
		"depth": "10",
	}

	harvest, err := bot.ListingWithParams("/r/golang",
		options)
	if err != nil {
		fmt.Println("Failed to fetch /r/golang: ", err)
		return
	}

	for index, post := range harvest.Posts {
		fmt.Printf("Post [%d] *<%s>*\n", index, post.Title)
		for _, comment := range post.Replies {
			fmt.Printf("*** [%s] commented --> %s\n", comment.Author, comment.Body)
		}
	}
}
