package api

import (
	// "database/sql"
	"database/sql"
	"fmt"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
	db "github.com/vutranhoang1411/tmdt/db/sqlc"
)
func (server *Server)GetPost(ctx *gin.Context){
	posts,err:=server.model.GetPost(ctx)
	if err!=nil{
		ctx.JSON(http.StatusInternalServerError,handleError(err))
		return
	}
	ctx.JSON(http.StatusOK,posts)
}
func (server *Server)GetPostByCustomer(ctx *gin.Context){
	form_id:=ctx.PostForm("customer_id")
	id,err:=strconv.ParseInt(form_id,10,64)
	if err !=nil{
		ctx.JSON(http.StatusBadRequest,handleError(fmt.Errorf("invalid customer id")))
		return
	}

	post,err:=server.model.GetPostByOwner(ctx,id)
	if err!=nil{
		ctx.JSON(http.StatusInternalServerError,handleError(err))
		return
	}
	ctx.JSON(http.StatusOK,post)
}
type CreatePostRequest struct {
	Grade     int32          `json:"grade" form:"grade"`
	Subject   []string       `json:"subject" form:"subject"`
	Address   string         `json:"address" form:"address"`
	OwnerID   int64          `json:"owner_id" form:"owner_id"`
	ExtraInfo string `json:"extra_info" form:"extra_info"`
}
func (server *Server)CreatePost(ctx *gin.Context){
	var req CreatePostRequest
	err:=ctx.ShouldBind(&req)
	if err!=nil{
		ctx.JSON(http.StatusBadRequest,handleError(err))
		return
	}
	var extra_info sql.NullString
	if len(req.ExtraInfo)>0{
		extra_info.String=req.ExtraInfo
		extra_info.Valid=true
	}else{
		extra_info.Valid=false
	}
	
	post,err:=server.model.CreatePost(ctx,db.CreatePostParams{
		Grade: req.Grade,
		Subject: req.Subject,
		Address: req.Address,
		OwnerID: req.OwnerID,
		ExtraInfo: extra_info,
	})
	if err!=nil{
		ctx.JSON(http.StatusBadRequest,handleError(err))
		return
	}
	ctx.JSON(http.StatusOK,post)

}