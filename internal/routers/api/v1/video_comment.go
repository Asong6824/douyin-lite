package v1

import "github.com/gin-gonic/gin"

type VideoComment struct {}

func NewVideoComment() VideoComment {
    return VideoComment{}
}

func (vc VideoComment) CommentAction(c *gin.Context) {}
func (vc VideoComment) CommentList(c *gin.Context) {}