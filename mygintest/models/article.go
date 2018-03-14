package models

import(
	"log"
	"fmt"
	db "mygintest/database"
	
)

//文章结构
type Article struct{
	Aid        int        `json:"aid" form:"aid"`
	Atitle     string     `json:"atitle" form:"atitle"`
	Adetail    string     `json:"adetail" form:"adetail"` 
}
//文章操作

//新增 文章
func (art *Article)AddArticle() bool{
	rs,err := db.SqlDB.Exec("INSERT INTO articles(aid,atitle,adetail) VALUES(?,?,?)",art.Aid,art.Atitle,art.Adetail)
	if err != nil{
		return false
	}
	id,err := rs.LastInsertId()
	fmt.Print(id)
	if err != nil{
		return false
	}else{
		return true
	}
}
//文章编辑
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
//文章删除
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
//文章查找（多个）
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


}