package middlewares

import (
    "github.com/gin-gonic/gin"
    "github.com/appleboy/gin-jwt"
    "time"
    "github.com/app8izer/go-gin-ng6-starter/backend/utils"
)

var authMiddleware *jwt.GinJWTMiddleware

// only example code
type User struct {
    UserName  string
    FirstName string
    LastName  string
    Role string
}

func init() {
    // TODO: check where to use util getenv
    authMiddleware = &jwt.GinJWTMiddleware{
        Realm:      utils.GetEnv("REALM", "b-rpc test"),
        Key:        []byte(utils.GetEnv("KEY", "secret key")),
        Timeout:    time.Hour,
        MaxRefresh: time.Hour,
        // login handler
        Authenticator: authenticate,
        // access permission check
        Authorizator: authorize,
        // unauthorized handler
        Unauthorized: unauthorized,
        PayloadFunc: payload,
        // TokenLookup is a string in the form of "<source>:<name>" that is used
        // to extract token from the request.
        // Optional. Default value "header:Authorization".
        // Possible values:
        // - "header:<name>"
        // - "query:<name>"
        // - "cookie:<name>"
        TokenLookup: "header: Authorization, query: token, cookie: jwt",
        // TokenLookup: "query:token",
        // TokenLookup: "cookie:token",

        // TokenHeadName is a string in the header. Default value is "Bearer"
        TokenHeadName: "Bearer",

        // TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
        TimeFunc: time.Now,
    }

}

func GetAuthMiddleware() *jwt.GinJWTMiddleware {
    return authMiddleware
}

func authenticate(userId string, password string, c *gin.Context) (interface{}, bool) {

    // find user in db -> verify pw -> let handler create token (what to return)

    if (userId == "admin" && password == "admin") || (userId == "test" && password == "test") {
        return &User{
        UserName:  userId,
        LastName:  "Bo-Yi",
        FirstName: "Wu",
        Role: "chief",
        }, true
    }
        return nil, false
}

func authorize(user interface{}, c *gin.Context) bool {

    if v, ok := user.(string); ok && v == "admin" {
        return true
    }

    return false
}

func unauthorized(c *gin.Context, code int, message string) {
    c.JSON(code, gin.H{
        "code":    code,
        "message": message,
    })
}

func payload(data interface{}) jwt.MapClaims {
    // in this method, you'd want to fetch some user info
    // based on their email address (which is provided once
    // they've successfully logged in).  the information
    // you set here will be available the lifetime of the
    // user's session
    //val := data.(*User)
    return map[string]interface{}{
        "id":   "1349",
        "role": "admin",
    }
}
