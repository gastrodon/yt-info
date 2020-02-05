package youtube

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var yt_query string = fmt.Sprintf("part=snippet,statistics&key=%s&id=%%s", os.Getenv("YOUTUBE_API_KEY"))
var yt_video string = fmt.Sprintf("https://www.googleapis.com/youtube/v3/videos?%s", yt_query)

type Video struct {
	Snippet struct {
		Title     string `json:"title"`
		Channel   string `json:"channelTitle"`
		Published string `json:"publishedAt"`
	} `json:"snippet"`
	Statistics struct {
		ViewCountRaw    string `json:"viewCount"`
		LikeCountRaw    string `json:"likeCount"`
		DislikeCountRaw string `json:"dislikeCount"`
		ViewCount       int64
		LikeCount       int64
		DislikeCount    int64
	} `json:"statistics"`
}

type VideoResults struct {
	Items []Video `json:"items"`
}

func get_video_data(ids ...string) (data VideoResults, err error) {
	var response *http.Response
	response, err = http.Get(fmt.Sprintf(yt_video, strings.Join(ids, ",")))
	if err != nil {
		return
	}

	var body []byte
	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &data)
	return
}

func VideoInfo(id string) (info Video, err error) {
	var results VideoResults
	results, err = get_video_data(id)
	if err != nil {
		log.Println(err)
		return
	}

	if len(results.Items) == 0 {
		err = fmt.Errorf("No such video id:%s", id)
		return
	}

	info = results.Items[0]
	info.Statistics.ViewCount, _ = strconv.ParseInt(info.Statistics.ViewCountRaw, 10, 64)
	info.Statistics.LikeCount, _ = strconv.ParseInt(info.Statistics.LikeCountRaw, 10, 64)
	info.Statistics.DislikeCount, _ = strconv.ParseInt(info.Statistics.DislikeCountRaw, 10, 64)
	return
}
