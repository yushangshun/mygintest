package main
	/*
	一个简单的基于GIN框架的接口程序。
	实现简易博客系统的以下接口：
 	文章列表、文章编辑、文章删除、文章查看
	*/
import (
	db "mygintest/database"
	"github.com/gin-gonic/gin"
	. "mygintest/controllers"
	
)

func main(){
	defer db.SqlDB.Close()
	
	router:=gin.Default()

	router.GET("/",IndexFunc)
	//文章查看
	router.GET("/article/:a_id",FindArticleFunc)
	//文章列表
	router.GET("/articles/",GetAllArticleFunc)
	//文章编辑
	router.POST("/articles/edit",EditArticleFunc)
	//文章删除
	router.GET("/articless/:a_id",DeleteArticleFunc)
	
	router.Run(":8000")

	
}