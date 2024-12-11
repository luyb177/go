package main

import (
	"database/sql"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	swaggerfiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"log"
	"net/http"
	"time"

	// 	"github.com/gin-gonic/gin"
	_ "library/docs"
)

// @title 图书管理系统
// @version 1.0
// description 这是一个可以区分管理员和普通访客以及基本的增删改查图书的图书管理系统
// @host 127.0.0.1:8080
// @basepath /user
type Book struct {
	Id     int    `json : "id"`     //数据库里面的主键
	Name   string `json : "name"`   //书的姓名
	Author string `json : "author"` //书的作者
	Pid    string `json : "pid"`    //书的索引号
}
type User struct {
	Id       int    `json : "id"`
	Username string `json : "username"`
	Password string `json : "password"`
	Role     string `json : "role"` //"admin" 或 "user"
}
type Claims struct {
	User User `json:"user"`
	jwt.StandardClaims
}

// 处理跨域请求,支持options访问
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}

// 生成token
func GenerateToken(user User) (string, error) {
	claims := Claims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(), // 过期时间
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("1234567890")) // 替换为你的密钥
}

// 创建一个中间件来验证 JWT 并检查用户角色：
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(401, gin.H{"message": "unauthorized"})
			fmt.Println(1)
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte("1234567890"), nil // 替换为你的密钥
		})

		if err != nil {
			c.JSON(401, gin.H{"message": "Invalid token", "error": err.Error()})
			c.Abort()
			return
		} //检查jwt是否错误

		if !token.Valid {
			c.JSON(401, gin.H{"message": "Invalid or expired token"})
			c.Abort()
			return
		} //检查token是否有效

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.JSON(401, gin.H{"message": "unauthorized"})
			fmt.Println(3)

			c.Abort()
			return
		}
		//补充认证环节
		user, ok := claims["user"].(map[string]interface{})
		if !ok {
			c.JSON(401, gin.H{
				"message": "无法获取用户消息",
			})
			c.Abort() //终止
			return
		}

		id, ok := user["Id"].(float64) //json解码后整形为float64
		if !ok {
			c.JSON(401, gin.H{
				"msg": "用户id不存在或格式错误",
			})
			c.Abort()
			return
		}
		username, ok := user["Username"].(string)
		if !ok {
			c.JSON(401, gin.H{
				"msg": "用户名不存在或格式错误",
			})
			c.Abort()
			return
		}
		role, ok := user["Role"].(string)
		if !ok {
			c.JSON(401, gin.H{
				"msg": "role不存在或格式错误",
			})
			c.Abort()
			return
		}
		c.Set("user", User{
			Id:       int(id), //整形强转
			Username: username,
			Role:     role,
		})
		c.Next()
	}
}

// 创建一个专门用于验证管理员权限的中间件：
func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet("user").(User)
		if user.Role != "admin" {
			c.JSON(403, gin.H{"message": "forbidden"})
			c.Abort()
			return
		}
		c.Next()
	}
}

// Index 首页
// @summary 首页
// @description 用户可以与前端交互选择进入管理员页面或者用户页面
// @tags page
// @accept application/json
// @produce 	text/html
// @success 200 {string} string "首页"
// @router / [get]
func Index(c *gin.Context) {

}

// GetLogin 用户登录 (get请求页面)
// @summary 用户登录
// @description 用户获取登录页面
// @tags page
// @accept application/json
// @produce 	text/html
// @success 200 {string} string "登录页面"
// @router /user/login [get]

func GetLogin(c *gin.Context) {

}

// PostLogin  用户登录 (post登录)
// @summary 用户登录
// @description 用户通过用户名和密码进行登录
// @tags user
// @accept multipart/form-data
// @produce application/json
// @param username body string true "用户名"
// @param password body string true "密码"
// @success 200 {object}  map[string]interface{} "登陆成功返回json消息包含token"
// @failure 400 {object} map[string]interface{} "登录失败返回json消息"
// @router /user/login [post]
func PostLogin(c *gin.Context) {
	name := c.PostForm("username")
	password := c.PostForm("password")
	var user User
	//连接数据库
	dsn := "root:Lu03150079@tcp(172.30.247.67:3307)/user"
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

	if err := db.Ping(); err != nil {
		log.Fatalln(err.Error())

	}
	err = db.QueryRow("select name, password,role from userin where name = ? and password = ?", name, password).Scan(&user.Username, &user.Password, &user.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			//处理未找到记录的情况
			c.JSON(400, gin.H{
				"msg": "用户未注册或者用户用户名或者密码错误",
			})
		} else {
			log.Fatalln(err.Error())
		}
	} else {
		token, err := GenerateToken(user)
		if err != nil {
			log.Fatalln(err.Error())
			return
		}

		c.JSON(200, gin.H{
			"msg":   "登录成功",
			"token": token,
		})
	}
}

// GetRegister 用户注册 (get请求页面)
// @summary 用户注册
// @description 用户获取注册页面
// @tags page
// @accept application/json
// @produce 	text/html
// @success 200 {string} string "获取注册页面"
// @router /user/register [get]
func GetRegister(c *gin.Context) {
}

// PostRegister 用户注册 (post注册)
// @summary 用户注册
// @description 用户通过用户名和密码进行注册
// @tags user
// @accept multipart/form-data
// @produce application/json
// @param username body string true "用户名"
// @param password body string true "密码"
// @success 200 {object} map[string]interface{} "注册成功返回json消息"
// @failure 400 {object} map[string]interface{} "注册失败返回json消息"
// @router /user/register [post]
func PostRegister(c *gin.Context) {
	name := c.PostForm("username")
	password := c.PostForm("password")
	//连接数据库
	dsn := "root:Lu03150079@tcp(172.30.247.67:3307)/user"
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
	_, err = db.Exec("insert userin (name,password,Role) values(?,?,?)", name, password, "user")
	//unique约束
	if err != nil {
		if driverErr, ok := err.(*mysql.MySQLError); ok {
			if driverErr.Number == 1062 { //重复
				c.JSON(400, gin.H{
					"msg": err.Error(),
				})
			} else {
				log.Fatalln(err.Error())
			}
		} else {
			log.Fatalln(err.Error())
		}
	} else {
		c.JSON(200, gin.H{
			"msg": "注册成功",
		})
	}
}

// GetInsert 用户查询页面
// @summary 用户查询书籍页面
// @description 用户通过书名查询书本
// @tags page
// @accept application/json
// @produce 	 text/html
// @success 200 {string} string "获取查询页面成功"
// @router /user/insert [get]
func GetInsert(c *gin.Context) {

}

// PostInsert 用户查询书籍
//
//		@summary 用户查询书籍
//		@description 用户通过书名查询书籍信息
//		@tags user
//		@accept multipart/form-data
//		@produce application/json
//	    @param Authorization headers string true "token"
//		@param bookName body string true "书名"
//		@success 200 {object} map[string]interface{} "查询成功返回书本信息"
//		@failure 400 {object} map[string]interface{} "查询失败"
//		@router /user/insert [post]
func PostInsert(c *gin.Context) {
	bookName := c.PostForm("bookName")
	//连接数据库
	var book Book
	dsn := "root:Lu03150079@tcp(172.30.247.67:3307)/book"
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
	//查询
	err = db.QueryRow("select * from books where name = ?", bookName).Scan(&book.Id, &book.Name, &book.Author, &book.Pid)
	if err != nil {
		if err == sql.ErrNoRows { //处理未找到记录的情况
			c.JSON(400, gin.H{
				"msg": err.Error(),
			})
		} else {
			log.Fatalln(err.Error())
		}
	} else {
		c.JSON(200, gin.H{
			"msg":  "查询成功",
			"book": book,
		})
	}
}

// GetAdminLogin 管理员登录 (get请求页面)
// @summary 管理员登录
// @description 管理员获取登录页面
// @tags page
// @accept application/json
// @produce 	text/html
// @success 200 {string} string "登录页面"
// @router /admin/login [get]
func GetAdminLogin(c *gin.Context) {

}

// PostAdminLogin  管理员登录 (post登录)
// @summary  管理员登录
// @description 管理员通过用户名和密码进行登录
// @tags admin
// @accept multipart/form-data
// @produce 	application/json
// @param username body string true "用户名"
// @param password body string true "密码"
// @success 200 {object} map[string]interface{} "登陆成功返回json消息"
// @failure 400 {object} map[string]interface{} "登录失败返回json消息"
// @router /admin/login [post]
func PostAdminLogin(c *gin.Context) {
	name := c.PostForm("username")
	password := c.PostForm("password")
	var user User
	//连接数据库
	dsn := "root:Lu03150079@tcp(172.30.247.67:3307)/user"
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

	if err := db.Ping(); err != nil {
		log.Fatalln(err.Error())

	}
	err = db.QueryRow("select id, name, password,role from userin where name = ? and password = ?", name, password).Scan(&user.Id, &user.Username, &user.Password, &user.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			//处理未找到记录的情况
			c.JSON(400, gin.H{
				"msg": "用户未注册或者用户用户名或者密码错误",
			})
		} else {
			log.Fatalln(err.Error())
		}
	} else {
		token, err := GenerateToken(user)
		if err != nil {
			log.Fatalln(err.Error())
			return
		}

		c.JSON(200, gin.H{
			"msg":   "登录成功",
			"token": token,
		})
	}
}

// GetAdminRegister 管理员注册 (get请求页面)
// @summary 管理员注册
// @description 管理员获取注册页面
// @tags page
// @accept application/json
// @produce 	text/html
// @success 200 {string} string "获取注册页面"
// @router /admin/register [get]
func GetAdminRegister(c *gin.Context) {

}

// PostAdminRegister 管理员注册 (post注册)
// @summary 管理员注册
// @description 管理员通过用户名和密码进行注册
// @tags admin
// @accept multipart/form-data
// @produce application/json
// @param username body string true "用户名"
// @param password body string true "密码"
// @success 200 {object} map[string]interface{} "注册成功返回Json消息"
// @failure 400 {object} map[string]interface{} "注册失败返回json消息"
// @router /admin/register [post]
func PostAdminRegister(c *gin.Context) {
	name := c.PostForm("username")
	password := c.PostForm("password")
	//连接数据库
	dsn := "root:Lu03150079@tcp(172.30.247.67:3307)/user"
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
	_, err = db.Exec("insert userin (name,password,Role) values(?,?,?)", name, password, "admin")
	//unique约束
	if err != nil {
		if driverErr, ok := err.(*mysql.MySQLError); ok {
			if driverErr.Number == 1062 { //重复
				c.JSON(400, gin.H{
					"msg": err.Error(),
				})
			} else {
				log.Fatalln(err.Error())
			}
		} else {
			log.Fatalln(err.Error())
		}
	} else {
		c.JSON(200, gin.H{
			"msg": "注册成功",
		})
	}
}

// GetAdminInsert 管理员查询 (get请求页面)
// @summary 管理员查询书
// @description 管理员通过书名查询书本
// @tags page
// @accept application/json
// @produce 	 text/html
// @success 200 {string} string "获取查询页面成功"
// @router /admin/insert [get]
func GetAdminInsert(c *gin.Context) {

}

// PostAdminInsert 查询 (post查询)
// @summary 管理员查询
// @description 管理员通过书名进行查询
// @tags admin
// @accept multipart/form-data
// @produce 	application/json
// @param Authorization headers string true "token"
// @param booKName body string true "书名"
// @success 200 {object} map[string]interface{} "查询成功返回json消息"
// @failure 400 {object} map[string]interface{} "查询失败返回json消息"
// @router /admin/insert [post]
func PostAdminInsert(c *gin.Context) {
	bookName := c.PostForm("bookName")
	//连接数据库
	var book Book
	dsn := "root:Lu03150079@tcp(172.30.247.67:3307)/book"
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
	//查询
	err = db.QueryRow("select * from books where name = ?", bookName).Scan(&book.Id, &book.Name, &book.Author, &book.Pid)
	if err != nil {
		if err == sql.ErrNoRows { //处理未找到记录的情况
			c.JSON(400, gin.H{
				"msg": err.Error(),
			})
		} else {
			log.Println(err.Error())
		}
	} else {
		c.JSON(200, gin.H{
			"msg":  "查询成功",
			"book": book,
		})
	}
}

// GetUpdate 管理员更新 (get请求页面)
// @summary 管理员更新书籍
// @description 管理员通过书名更新书本
// @tags page
// @accept application/json
// @produce 	 text/html
// @success 200 {string} string "获取更新页面成功"
// @router /admin/update [get]
func GetUpdate(c *gin.Context) {

}

// PostUpdate 更新
// @summary 管理员更新书籍信息
// @description 管理员通过书名进行更新
// @tags admin
// @accept  multipart/form-data
// @produce application/json
// @param Authorization headers string true "token"
// @param BbookName body string true "旧书名"
// @param AbookName body string true "新书名"
// @param Author body string true "作者"
// @param Pid body string true "索引号"
// @success 200 {object} map[string]interface{} "更新成功返回书本的信息"
// @failure 400 {object} map[string]interface{} "更新失败返回失败的消息"
// @router /admin/update [post]
func PostUpdate(c *gin.Context) {
	var book Book
	BbookName := c.PostForm("BbookName")
	book.Name = c.PostForm("AbookName")
	book.Author = c.PostForm("Author")
	book.Pid = c.PostForm("Pid")
	var db_name string

	//连接数据库

	dsn := "root:Lu03150079@tcp(172.30.247.67:3307)/book"
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
	err = db.QueryRow("select name from books where name = ? ", BbookName).Scan(&db_name)
	if err != nil {
		if err == sql.ErrNoRows { //未找到数据
			c.JSON(400, gin.H{
				"msg": err.Error(),
			})
		} else {
			log.Fatalln(err.Error())
		}
	} else {
		_, err = db.Exec("update books set  name = ? ,author = ? ,pid = ? where name = ?", book.Name, book.Author, book.Pid, BbookName)
		if err != nil {
			log.Fatalln(err.Error())
			return
		}
		c.JSON(200, gin.H{
			"msg":  "更新成功",
			"book": book,
		})
	}

}

// GetDelete 管理员删除 (get请求页面)
// @summary 管理员删除书籍
// @description 管理员通过书名删除书本
// @tags page
// @accept application/json
// @produce 	 text/html
// @success 200 {string} string "获取删除页面成功"
// @router /admin/delete [get]
func GetDelete(c *gin.Context) {

}

// PostDelete 删除
// @summary 管理员删除书籍
// @description 管理员通过书名进行删除
// @tags admin
// @accept  multipart/form-data
// @produce application/json
// @param Authorization headers string true "Bearer {token}" // 这里说明 token 的格式
// @param bookName body string true "书名"
// @success 200 {object} map[string]interface{} "删除成功返回json消息"
// @failure 400 {object} map[string]interface{} "删除失败返回json消息"
// @router /admin/delete [post]
func PostDelete(c *gin.Context) {
	bookName := c.PostForm("bookName")

	//连接数据库
	dsn := "root:Lu03150079@tcp(172.30.247.67:3307)/book"
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
	//删除书籍
	result, err := db.Exec("delete from books where name = ?", bookName)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": err.Error(),
		})
	} else {
		//检查是否有行删除
		affect, err := result.RowsAffected()
		if err != nil {
			c.JSON(400, gin.H{
				"msg": err.Error(),
			})
		}
		if affect == 0 {
			c.JSON(400, gin.H{
				"msg": "没有这本书",
			})
		} else {
			c.JSON(200, gin.H{
				"msg": "删除成功",
			})
		}
	}
}

// GetAdd 管理员添加 (get请求页面)
// @summary 管理员添加书籍
// @description 管理员通过书名、索引号、作者添加书本
// @tags page
// @accept application/json
// @produce 	 text/html
// @success 200 {string} string "获取添加页面成功"
// @router /admin/add [get]
func GetAdd(c *gin.Context) {

}

// PostAdd 添加
// @summary 管理员添加书籍
// @description 管理员通过书名、索引号、作者进行添加
// @tags admin
// @accept  multipart/form-data
// @produce application/json
// @param Authorization headers string true "token"
// @param bookName body string true "书名"
// @param author body string true "作者"
// @param Pid body string true "索引号"
// @success 200 {object} map[string]interface{}"添加成功返回json消息"
// @failure 400 {object} map[string]interface{} "添加失败返回失败消息"
// @router /admin/add [post]
func PostAdd(c *gin.Context) {
	var book Book
	book.Name = c.PostForm("bookName")
	book.Author = c.PostForm("author")
	book.Pid = c.PostForm("Pid")
	//连接数据库
	dsn := "root:Lu03150079@tcp(172.30.247.67:3307)/book"
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
	//添加书籍
	_, err = db.Exec("insert books (name, author, pid) values(?, ?, ?)", book.Name, book.Author, book.Pid)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"msg":  "添加成功",
			"book": book,
		})
	}

}

func main() {
	r := gin.Default()
	// 跨域设置
	r.Use(Cors())

	r.GET("/", Index)

	v1 := r.Group("/user")
	{
		//登录和注册不需要认证
		v1.GET("login", GetLogin)
		v1.POST("login", PostLogin)
		v1.GET("register", GetRegister)
		v1.POST("register", PostRegister)
		v1.Use(AuthMiddleware()) //查询需要认证
		{
			v1.GET("insert", GetInsert)
			v1.POST("insert", PostInsert) //查
		}

	}
	v2 := r.Group("admin")
	{
		//登录注册不需要认证
		v2.GET("login", GetAdminLogin)
		v2.POST("login", PostAdminLogin)
		v2.GET("register", GetAdminRegister)
		v2.POST("register", PostAdminRegister)
		v2.Use(AuthMiddleware(), AdminMiddleware()) //用户认证和管理员权限
		{
			v2.GET("insert", GetAdminInsert)
			v2.POST("insert", PostAdminInsert) //查
			v2.GET("update", GetUpdate)
			v2.POST("update", PostUpdate) //改
			v2.GET("delete", GetDelete)
			v2.POST("delete", PostDelete) //删
			v2.GET("add", GetAdd)
			v2.POST("add", PostAdd) //增
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run()
}
