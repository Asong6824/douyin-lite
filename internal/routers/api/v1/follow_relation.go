package v1

import "github.com/gin-gonic/gin"

type FollowRelation struct {}

func NewFollowRelation() FollowRelation {
    return FollowRelation{}
}

func (fr FollowRelation) RelationAction(c *gin.Context) {}
func (fr FollowRelation) RelationList(c *gin.Context) {}