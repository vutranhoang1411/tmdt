package api

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	db "github.com/vutranhoang1411/tmdt/db/sqlc"
)
type Server struct{
	model db.Model
	router *gin.Engine
}

func NewServer(conn *sql.DB)*Server{
	server:=&Server{
		model: db.NewSQLModel(conn),
	}

	server.router=gin.Default()
	
	///add route
	server.router.POST("/api/post",server.CreatePost)
	server.router.GET("/api/post",server.GetPost)
	server.router.GET("/api/post/customer",server.GetPostByCustomer)

	server.router.POST("/api/post/register",server.RegisterPost)
	server.router.POST("/api/post/approved",server.ApproveRegister)
	server.router.POST("/api/post/teacher",server.GetTeacherByPost)
	
	return server
}
func handleError(err error)gin.H{
	return gin.H{"error":err.Error()}
}
func (server *Server)Start(addr string)error{
	return server.router.Run(addr)
}