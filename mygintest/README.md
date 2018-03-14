程序文档和说明

程序文档和说明

> 输出：一个简单的基于GIN框架的接口程序。 实现简易博客系统的以下接口：文章列表、文章编辑、文章删除、文章查看

### Gin框架

Gin是一个golang的微框架，有快速灵活，容错方便等特点。

#### 1 接口程序流程图

![流程图](http://ooi2jt2e8.bkt.clouddn.com/GIN.jpg)

1. `router`：绑定用来处理路由事项，并且，将HTTP的请求绑定到相关的处理函数

2. `cotroller`:处理操作逻辑

3. `mode`: 一般用来操作数据库等

4. `view`:处理视图，一般用模板,`html`,`css`,等处理要返回的数据表现。

   #### 2 接口程序文档

   1. ##### 数据库（定义、连接）

      表的结构（aid、atitle、adetail)

      连接：

      ```

      var SqlDB *sql.DB

      func init(){
      	var err error
      	SqlDB,err = sql.Open("mysql","root:123456@tcp(127.0.0.1:3306)/test?parseTime=true")
      	if err !=nil{
      		log.Fatal(err)
      	}
      	err = SqlDB.Ping()
      	if err !=nil{
      		log.Fatal(err)
      	}
      }

      ```

      ##### 2 数据库的操纵

      1.表结构

      ```
      //文章结构
      type Article struct{
      	Aid        int        `json:"aid" form:"aid"`
      	Atitle     string     `json:"atitle" form:"atitle"`
      	Adetail    string     `json:"adetail" form:"adetail"` 
      }
      ```

      2 文章编辑

      ```
      func (art *Article)EditArticle() bool {

          rs,err := db.SqlDB.Exec("UPDEATE articles SET atitle=?,adetail=? WHERE aid=?",art.Atitle,art.Adetail,art.Aid)

          if err!=nil{

              return false

          }

          id , err := rs.RowsAffected()

          fmt.Print(id)

          if err!=nil{

              log.Print(err)

              return false

          }else{

              return true

          }

      }

      ```

      3 文章查找

      ```
      //文章查找（单个）
      func  GetArticleById(aaid int) (a Article) {
      	var article Article
      	err := db.SqlDB.QueryRow("SELECT aid,atitle,adetail FROM articles WHERE aid=?",aaid).Scan(
      		&article.Aid,&article.Atitle,&article.Adetail)
      	if err!=nil{
      		log.Print(err)
      	}
      	return article
      			
      }
      ```

      4 文章列表

      ```
      func (art *Article) GetAllArticles()(articles []Article){
      	articles = make([]Article,0)
      	rows,err :=db.SqlDB.Query("SELECT aid,atitle,adetail FROM articles")
      	if err != nil{
      		return nil
      	}
      	defer rows.Close()
      	for rows.Next(){
      		var arti Article
      		rows.Scan(&arti.Aid,&arti.Atitle,&arti.Adetail)
      		articles=append(articles,arti)
      	}
      	return articles
      ```


      }
      ```
    
      5 文章删除
    
      ```
      func DeleteArticle(aid int) bool{
      	rs,err :=db.SqlDB.Exec("DELETE FROM articles WHERE aid=?",aid)
    
      	if err !=nil {
      		log.Fatal(err)
      		return false
      	}
      	id,err := rs.RowsAffected()
      	fmt.Print(id)
      	if err!=nil{
      		log.Fatal(err)
      		return false
      	}else{
      		return true
      	}
      }
    
      ```
    
      #### 3 handler处理函数
    
      ```
      //首页
      func IndexFunc(c *gin.Context){
      	c.String(http.StatusOK,"Hello world !")
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
      
      ```
    
      ### main函数（路由和处理函数的绑定）
    
      1 为了简便，将主函数和处理函数与写在一起
    
      ```
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
      ```
    
      结果展示
    
      文章列表：
    
      ![文章列表](http://ooi2jt2e8.bkt.clouddn.com/%E6%96%87%E7%AB%A0%E5%88%97%E8%A1%A8.jpg)
    
      文章查找：
    
      ![](http://ooi2jt2e8.bkt.clouddn.com/%E6%9F%A5%E6%89%BE.jpg)
    
      ​
    
      删除：
    
      ![删除](http://ooi2jt2e8.bkt.clouddn.com/%E5%88%A0%E9%99%A4.jpg)
    
      ​
    
      ​
    
      ​
    
      ​

   ​
