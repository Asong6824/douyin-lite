
package v1

import "github.com/gin-gonic/gin"

type VideoLike struct {}

func NewVideoLike() VideoLike {
    return VideoLike{}
}

func (vl VideoLike) FavoriteAction(c *gin.Context) {}
func (vl VideoLike) FavoriteList(c *gin.Context) {}
