package handlers
import (
	"log"
	"strconv"
	"net/http"
	"github.com/gin-gonic/gin"
	. "myginTest/models"
	"fmt"

)

//首页
func IndexFunc(c *gin.Context){
	c.String(http.StatusOK,"Hello world !")
}

//增加文章
func AddArticleFunc(c *gin.Context){
	//处理请求参数
	//客户端格式：a_id 格式
	//转换后格式： articleId 
	//插入数据库格式：aid 
	//articleId := c.PostForm("a_id")
	//articleTitle:=c.PostForm("a_title")
	//articleDetail:=c.PostForm("a_detail")
	//
	//ar := Article{aid:articleId,atitle:articleTitle,adetail:articleDetail}
	//ra := ar.AddArticle()
	//c.JSON(http.StatusOK,gin.H{
	//	"add success":ra,
	//})

}

//编辑文章
func EditArticleFunc(c *gin.Context){
	num := c.Param("a_id")
	fmt.Printf(num)
	
	aid,err := strconv.Atoi(num)
	if err!=nil {
		log.Fatalln(err)
	}
	//提取转换
	atitle :=c.PostForm("a_title")
	adetail:=c.PostForm("a_detail")
	//查找要修改的文章
	a := GetArticleById(aid)
	a.Atitle=atitle
	a.Adetail=adetail
	//赋值
	ra:=a.EditArticle()
	if ra==false{
		log.Fatal(err)
	}
	//返回结果
	c.JSON(http.StatusOK,gin.H{
		"editsuccess": ra,
	})
 
}


//删除文章
func DeleteArticleFunc(c *gin.Context){
	num := c.Param("a_id")
	fmt.Printf(num)
	aid,err := strconv.Atoi(num)
	if err!=nil {
		log.Fatalln(err)
	}
	//查找
	ra:=DeleteArticle(aid)
	fmt.Println(ra)
	if ra==false{
		log.Fatalln("deleat failed")
	}
	c.JSON(http.StatusOK,gin.H{
		"detele success":ra,
	})

}


//查找文章
func FindArticleFunc(c *gin.Context){
	
	num :=c.Param("a_id")
	fmt.Printf("num:%s",num)

	aid,err := strconv.Atoi(num)
	if err!=nil {
		log.Fatalln(err)
	}
	//find
	rep := GetArticleById(aid)
	c.JSON(http.StatusOK,gin.H{
		"rep": rep,
	})

}


//文章列表
func GetAllArticleFunc(c *gin.Context){
	var ar Article
	ra := ar.GetAllArticles()
	
	c.JSON(http.StatusOK,gin.H{
		"ra":ra,
	})

}
