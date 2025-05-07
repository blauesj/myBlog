package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"myBlog/define"
	"myBlog/models"
	"net/http"
	"strconv"
)

func GetCommentList(context *gin.Context) {
	postId, _ := strconv.ParseUint(context.Query("postId"), 10, 64)
	comments, err := models.ListComment(uint(postId))
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		context.JSON(0, gin.H{"error": err})
		return
	}
	context.JSON(http.StatusOK, comments)

}

func AddComment(context *gin.Context) {
	var commentReq CommentRequest

	if err := context.ShouldBindJSON(&commentReq); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, _ := context.Get("user")

	comment := models.Comment{
		Content: commentReq.Content,
		PostID:  commentReq.PostId,
	}

	comment.UserID = user.(*define.UserClaims).Id
	err := models.AddComment(&comment)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Comment added successfully"})
}
