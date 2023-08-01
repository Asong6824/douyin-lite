package model

import "time"

type VideoComment struct {
    ID int
    UserID int
    VideoID int
    Content string
    CommentTime time.Time
}