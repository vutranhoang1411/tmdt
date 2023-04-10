package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (server *Server)GetTeacherByPost(ctx *gin.Context){
	form_id:=ctx.PostForm("post_id")
	id,err:=strconv.ParseInt(form_id,10,64)
	if err !=nil{
		ctx.JSON(http.StatusBadRequest,handleError(fmt.Errorf("invalid post id")))
		return
	}
	teacher_list,err:=server.model.GetTeacherByPost(ctx,id)
	if err!=nil{
		ctx.JSON(http.StatusInternalServerError,handleError(err))
		return
	}
	ctx.JSON(http.StatusOK,teacher_list)
	
}