package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

// var flag int = 0
// var num = 0

type User struct {
	Name     string `json:"name"`
	Passward string `json:"password"`
	genter   string `json:"genter,omitempty"`
	Phone    string `json:"phone,omitempty"`
	QQ       string `json:"qq,omitempty"`
	Email    string `json:"email,omitempty"`
}

var Use []User

func index(context *gin.Context) { //首页
	context.HTML(200, "index.html", nil)
}
func login(context *gin.Context) { //登录页面
	context.HTML(200, "login.html", nil)
}
func register(context *gin.Context) { //注册页面
	context.HTML(200, "register.html", nil)
}
func auth1(context *gin.Context) { //控制注册

	name := context.PostForm("username")
	password := context.PostForm("password")
	//连接数据库
	dsn := "root:Lu03150079@tcp(172.30.247.67:3307)/userinfo"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	defer db.Close() //关闭数据库
	//检查数据库连接
	if err := db.Ping(); err != nil {
		log.Fatalln(err.Error())
		return
	}
	_, err = db.Exec("insert user (username,password) values(?,?)", name, password)
	//unique约束
	if err != nil {
		if driverErr, ok := err.(*mysql.MySQLError); ok {
			if driverErr.Number == 1062 {
				context.HTML(200, "register2.html", nil)
			} else {
				log.Fatalln(err.Error())

			}
		} else {
			log.Fatalln(err.Error())
		}
	} else {
		context.HTML(200, "register1.html", nil)
	}

}
func auth2(context *gin.Context) { //控制登录

	name := context.PostForm("username")
	password := context.PostForm("password")
	var db_name, db_pas string
	//连接数据库
	dsn := "root:Lu03150079@tcp(172.30.247.67:3307)/userinfo"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln(err.Error())

	}
	defer db.Close() //关闭数据库
	//检查数据库连接
	if err := db.Ping(); err != nil {
		log.Fatalln(err.Error())

	}
	err = db.QueryRow("select username, password from user where username = ? and password = ?", name, password).Scan(&db_name, &db_pas)
	if err != nil {
		if err == sql.ErrNoRows {
			//处理未找到记录的情况
			context.HTML(200, "login2.html", nil)
		} else {
			log.Fatalln(err.Error())

		}
	} else {
		//登录成功，设置cookie
		context.SetCookie("user_name", name, 3600, "/", "", false, true) //使用localhost 或 127.0.0.1
		//填写"localhost" 或 ""
		//返回登录成功的页面
		context.HTML(200, "login1.html", gin.H{
			"username": name,
		})
	}
}
func alter(context *gin.Context) {

	switch context.Request.Method {
	case http.MethodGet:
		context.HTML(200, "alter.html", nil)
	case http.MethodPost:

		alter1(context)
	}

}
func alter1(context *gin.Context) {
	name := context.PostForm("username")
	password := context.PostForm("password")
	password_new := context.PostForm("password_new")
	var db_name, db_pas string

	//连接数据库
	dsn := "root:Lu03150079@tcp(172.30.247.67:3307)/userinfo"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln(err.Error())

	}
	defer db.Close() //关闭数据库
	//检查数据库连接
	if err := db.Ping(); err != nil {
		log.Fatalln(err.Error())

	}
	err = db.QueryRow("select username, password from user where username = ? and password = ?", name, password).Scan(&db_name, &db_pas)
	if err != nil {
		if err == sql.ErrNoRows {
			//处理未找到记录的情况
			context.HTML(200, "alter1.html", nil)
		} else {
			log.Fatalln(err.Error())

		}
	} else {
		_, err := db.Exec("update user set password = ? where username = ?", password_new, name)
		if err != nil {
			log.Fatalln(err.Error())
		} else {
			context.HTML(200, "alter2.html", nil)
		}
	}
}
func inquire(context *gin.Context) {
	switch context.Request.Method {
	case http.MethodGet:
		context.HTML(200, "inquire.html", nil)
	case http.MethodPost:

		inquire1(context)
	}
}
func inquire1(context *gin.Context) {

	name := context.PostForm("username")
	password := context.PostForm("password")

	var use User
	//连接数据库
	dsn := "root:Lu03150079@tcp(172.30.247.67:3307)/userinfo"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln(err.Error())

	}
	defer db.Close() //关闭数据库
	//检查数据库连接
	if err := db.Ping(); err != nil {
		log.Fatalln(err.Error())

	}
	err = db.QueryRow("select username,password ,gender from user where username = ? and password = ?", name, password).Scan(&use.Name, &use.Passward, &use.genter)
	if err != nil {
		if err == sql.ErrNoRows {
			//处理未找到记录的情况
			context.HTML(200, "inquire1.html", nil)
		} else {
			log.Fatalln(err.Error())

		}
	} else {
		context.HTML(200, "inquire2.html", nil)

		context.JSON(http.StatusOK, gin.H{
			"masage": "inquire success",
			"use":    use,
		})

	}

}

func mimi(context *gin.Context) {

	user_name, err := context.Cookie("user_name")
	if err != nil {
		context.HTML(http.StatusUnauthorized, "error.html", gin.H{
			"error": "Unauthorized: Please log in first!",
		})
		return
	}
	// 根据 cookie 获取用户信息
	var db_name string

	//连接数据库
	dsn := "root:Lu03150079@tcp(172.30.247.67:3307)/userinfo"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln(err.Error())

	}
	defer db.Close() //关闭数据库
	//检查数据库连接
	if err := db.Ping(); err != nil {
		log.Fatalln(err.Error())

	}
	err = db.QueryRow("select username from user where username = ? ", user_name).Scan(&db_name)
	if err != nil {
		if err == sql.ErrNoRows {
			// 如果找不到对应的用户
			context.HTML(http.StatusUnauthorized, "error.html", gin.H{
				"error": "User not found:Invaild session!",
			})
		}
	}

	context.HTML(http.StatusOK, "mimi.html", nil) // 渲染页面
}
func logOut(context *gin.Context) {

	//设置一个同名的cookie并将有效时间改为-1
	context.SetCookie("user_name", "", -1, "/", "", false, true)
	context.HTML(200, "login.html", nil)
}
func main() {
	//获取路由对象
	router := gin.Default()
	//加载响应的HTML文件
	router.LoadHTMLGlob("./templates/*") // * 是加载文件夹里的所有文件
	router.GET("/", index)
	router.GET("login", login)
	router.GET("register", register)
	router.POST("auth1", auth1)
	router.POST("auth2", auth2)
	router.GET("inquire", inquire)
	router.POST("inquire", inquire)
	router.GET("alter", alter)
	router.POST("alter", alter)
	router.GET("logout", logOut)
	router.GET("mimi", mimi)

	//连接数据库
	dsn := "root:Lu03150079@tcp(127.0.0.1:3306)/user"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	defer db.Close() //关闭数据库
	//检查数据库连接
	if err := db.Ping(); err != nil {
		log.Fatalln(err.Error())
		return
	}

	router.Run(":8090")
}
