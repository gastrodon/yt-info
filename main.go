package main

import (
	"github.com/gastrodon/yt-info/listener"

	"github.com/turnage/graw/reddit"

	"log"
)

func main() {
	var client reddit.Bot
	var err error
	client, err = reddit.NewBotFromAgentFile("agent", 0)
	if err != nil {
		log.Fatal(err)
	}

	var comment_kill chan bool = make(chan bool)
	var err_chan chan error = make(chan error)
	go listener.CommentsOn("all", client, comment_kill, err_chan)
	go listener.LogErrs(err_chan)

	select {
	case <-comment_kill:
		return
	}
}
