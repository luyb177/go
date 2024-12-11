package main 
// import (
// 	"github.com/dgrijalva/jwt-go"
// 	"time"
// 	"github.com/gin-gonic/gin"
// 	"fmt"
// )
// type User struct {
//     ID       uint   `json:"id"`
//     Username string `json:"username"`
//     Role     string `json:"role"` // "admin" 或 "user"
// }

// type Claims struct {
//     User User `json:"user"`
//     jwt.StandardClaims
// }
// func GenerateToken(user User) (string, error) {
//     claims := Claims{
//         User: user,
//         StandardClaims: jwt.StandardClaims{
//             ExpiresAt: time.Now().Add(time.Hour * 72).Unix(), // 过期时间
//         },
//     }

//     token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//     return token.SignedString([]byte("your_secret_key")) // 替换为你的密钥
// }
// func AuthMiddleware() gin.HandlerFunc {
//     return func(c *gin.Context) {
//         tokenString := c.GetHeader("Authorization")
//         if tokenString == "" {
//             c.JSON(401, gin.H{"message": "unauthorized"})
//             c.Abort()
//             return
//         }

//         token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
//             // 检查签名方法是否正确
//             if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//                 return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
//             }
//             return []byte("your_secret_key"), nil // 替换为你的密钥
//         })

//         if err != nil || !token.Valid {
//             c.JSON(401, gin.H{"message": "unauthorized"})
//             c.Abort()
//             return
//         }

//         claims, ok := token.Claims.(jwt.MapClaims)
//         if !ok || !token.Valid {
//             c.JSON(401, gin.H{"message": "unauthorized"})
//             c.Abort()
//             return
//         }

//         user := claims["user"].(map[string]interface{})
//         c.Set("user", User{ID: uint(user["id"].(float64)), Username: user["username"].(string), Role: user["role"].(string)})
//         c.Next()
//     }
// }

// func AdminMiddleware() gin.HandlerFunc {
//     return func(c *gin.Context) {
//         user := c.MustGet("user").(User)
//         if user.Role != "admin" {
//             c.JSON(403, gin.H{"message": "forbidden"})
//             c.Abort()
//             return
//         }
//         c.Next()
//     }
// }
// func main() {
//     r := gin.Default()

//     // 用户登录示例，生成 Token
//     r.POST("/login", func(c *gin.Context) {
//         var user User
//         if err := c.ShouldBindJSON(&user); err != nil {
//             c.JSON(400, gin.H{"message": "invalid input"})
//             return
//         }

//         // 假设你已经验证了用户的身份
//         token, err := GenerateToken(user)
//         if err != nil {
//             c.JSON(500, gin.H{"message": "could not generate token"})
//             return
//         }

//         c.JSON(200, gin.H{"token": token})
//     })

//     // 普通用户路由
//     r.GET("/user", AuthMiddleware(), func(c *gin.Context) {
//         c.JSON(200, gin.H{"message": "Hello User!"})
//     })

//     // 管理员路由
//     admin := r.Group("/admin")
//     admin.Use(AuthMiddleware(), AdminMiddleware())
//     {
//         admin.GET("/dashboard", func(c *gin.Context) {
//             c.JSON(200, gin.H{"message": "Welcome to admin dashboard!"})
//         })
//     }

//     r.Run() // 启动服务
// }