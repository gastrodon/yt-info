package handler

import (
	"github.com/turnage/graw/reddit"

	"fmt"
	"log"
	"strings"
)

var pattern_handlers map[string]func(...string) (string, error) = map[string]func(...string) (string, error){
	"video_long":  video_long,
	"video_short": video_short,
}

func DispatchForMap(client reddit.Bot, parent_name string, matches map[string][]string) (err error) {
	var blocks []string = []string{}

	var key, block string
	var value []string
	for key, value = range matches {
		if len(value) == 0 {
			continue
		}

		block, err = pattern_handlers[key](value...)
		if err != nil {
			return
		}
		blocks = append(
			blocks,
			block,
		)
	}

	if len(blocks) == 0 {
		return
	}

	var reply_text string = fmt.Sprintf("%s[.](https://github.com/gastrodon/yt-info)", strings.Join(blocks, "\n\n"))
	log.Println(reply_text)
	err = client.Reply(parent_name, reply_text)
	return
}

func ids_after(delim string, links ...string) (ids []string) {
	ids = []string{}

	var split []string
	var link string
	for _, link = range links {
		split = strings.Split(link, delim)
		ids = append(ids, split[len(split)-1])
	}
	return
}

func video_long(matches ...string) (comment string, err error) {
	var blocks []string = []string{}
	var block string

	var id string
	for _, id = range ids_after("v=", matches...) {
		block, err = format_video_block(id)
		if err != nil {
			return
		}
		blocks = append(blocks, block)
	}

	comment = strings.Join(blocks, "\n\n")
	return
}

func video_short(matches ...string) (comment string, err error) {
	var blocks []string = []string{}
	var block string

	var id string
	for _, id = range ids_after("/", matches...) {
		block, err = format_video_block(id)
		if err != nil {
			return
		}
		blocks = append(blocks, block)
	}

	comment = strings.Join(blocks, "\n\n")
	return
}
