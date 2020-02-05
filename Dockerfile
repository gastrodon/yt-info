FROM golang:1.13

go get -u ./...

ENV REDDIT_USER ""
ENV REDDIT_PASS ""
ENV REDDIT_ID ""
ENV REDDIT_SECRET ""
ENV USER_AGENT "github.com/gastrodon/yt-info"
ENV YOUTUBE_API_KEY ""

go run . 
