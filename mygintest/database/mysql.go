package database
import(

	_ "github.com/go-sql-driver/mysql"
	//"github.com/gin-gonic/gin"
	"database/sql"
	"log"
	//"net/http"
	//"fmt"	
)
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


