package listener

import (
	"github.com/gastrodon/yt-info/handler"

	"github.com/turnage/graw/reddit"
	"github.com/turnage/graw/streams"

	"log"
	"regexp"
)

var early_pattern *regexp.Regexp = regexp.MustCompile(`youtu\.?be`)
var patterns map[string]*regexp.Regexp = map[string]*regexp.Regexp{
	"video_long":  regexp.MustCompile(`https?://(www\.)?youtube\.com/watch\?v=[\w_-]+`),
	"video_short": regexp.MustCompile(`https?://youtu\.be/[\w_-]+`),
}

func CommentsOn(subs string, client reddit.Bot, kill chan bool, err_chan chan<- error) {
	var stream <-chan *reddit.Comment
	var err error
	stream, err = streams.SubredditComments(client, kill, err_chan, subs)
	if err != nil {
		panic(err)
	}

	var match_size int
	var matches map[string][]string
	var key string
	var value *regexp.Regexp

	var comment *reddit.Comment
	for comment = range stream {
		match_size = 1 + len(comment.Body)
		if early_pattern.FindAllString(comment.Body, match_size) == nil {
			continue
		}

		matches = map[string][]string{}
		for key, value = range patterns {
			matches[key] = value.FindAllString(comment.Body, match_size)
		}

		err = handler.DispatchForMap(client, comment.Name, matches)
		if err != nil {
			err_chan <- err
		}
	}

}

func LogErrs(err_chan chan error) {
	var err error
	for err = range err_chan {
		log.Println(err)
	}
}
