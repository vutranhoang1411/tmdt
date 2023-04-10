package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/vutranhoang1411/tmdt/db/sqlc"
)

type RegisterPostParam struct{
	PostID int64 `json:"post_id" form:"post_id" binding:"required"`
	TeacherID int64 `json:"teacher_id" form:"teacher_id" binding:"required"`
}
func (server *Server)RegisterPost(ctx *gin.Context){
	var req RegisterPostParam
	err:=ctx.ShouldBind(&req)
	if err!=nil{
		ctx.JSON(http.StatusBadRequest,handleError(err))
		return
	}
	register,err:=server.model.CreateRegister(ctx,db.CreateRegisterParams{
		TeacherID: req.PostID,
		PostID: req.TeacherID,
	})
	if err!=nil{
		ctx.JSON(http.StatusInternalServerError,handleError(err))
		return
	}
	ctx.JSON(http.StatusOK,register)
}
func (server *Server)ApproveRegister(ctx *gin.Context){
	var req RegisterPostParam
	err:=ctx.ShouldBind(&req)
	if err!=nil{
		ctx.JSON(http.StatusBadRequest,handleError(err))
		return
	}
	err=server.model.ApproveRegister(ctx,db.ApproveRegisterParams{
		PostID: req.PostID,
		TeacherID: req.TeacherID,
	})
	if err!=nil{
		ctx.JSON(http.StatusBadRequest,handleError(err))
		return
	}
	ctx.JSON(http.StatusOK,map[string]string{})
}