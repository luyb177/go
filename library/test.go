package main 
// import (
// 	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
// 	swaggerfiles "github.com/swaggo/files" // swagger embed files
// 	"github.com/gin-gonic/gin"
// 	_ "library/docs"
// )
// // @title 简单的测试
// // @version 1.0
// // @description 测试
// // @host 127.0.0.1:8080
// // @BasePath /
// func main(){
// 	r := gin.Default()
// 	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
// 	r.GET("/",hello)
// 	r.Run()
// }

// //@Summary 一个简单的hello
// //@schemes
// //@Decription 简单的测试，会输出hello
// //@Tags test
// //@accept json
// //@produce json
// //success 200 {string} hello
// //@router / [get] 
// func hello (c *gin.Context){
// 	c.JSON(200,gin.H{
// 		"msg" : "hello",
// 	})
// }