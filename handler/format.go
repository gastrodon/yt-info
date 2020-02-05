package handler

import (
	"github.com/gastrodon/yt-info/youtube"

	"fmt"
	"strings"
	"time"
)

const BAR_SIZE int = 30

func bar_at(index int) (bar string) {
	bar = fmt.Sprintf(
		"%s|%s",
		strings.Repeat("-", index),
		strings.Repeat("-", BAR_SIZE-index),
	)
	return
}

func time_from_string(time_string string) (parsed time.Time) {
	parsed, _ = time.Parse(time.RFC3339, time_string)
	return
}

func format_video_block(id string) (pretty string, err error) {
	var video youtube.Video
	video, err = youtube.VideoInfo(id)
	if err != nil {
		return
	}

	var when time.Time = time_from_string(video.Snippet.Published)

	var rating int
	if video.Statistics.LikeCount*video.Statistics.DislikeCount > 0 {
		rating = int((100 * video.Statistics.LikeCount) / (video.Statistics.LikeCount + video.Statistics.DislikeCount))
	} else {
		rating = 100
	}
	var bar_scale float32 = 100.0 / float32(BAR_SIZE)
	var bar string = bar_at(BAR_SIZE - int(float32(rating)/bar_scale))
	pretty = fmt.Sprintf(
		"%s (%d years ago)\n\n%03d%% liked `%s`",
		video.Snippet.Title, time.Now().Year()-when.Year(),
		rating, bar,
	)
	return
}
