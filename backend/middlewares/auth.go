package middlewares

import (
    "github.com/gin-gonic/gin"
    "github.com/appleboy/gin-jwt"
    "time"
    "github.com/app8izer/go-gin-ng6-starter/backend/services"
    "github.com/app8izer/go-gin-ng6-starter/backend/models"
    "github.com/app8izer/go-gin-ng6-starter/backend/config"
    "net/http"
    "errors"
)

type JwtMiddleware struct {
    *jwt.GinJWTMiddleware
}

var authMiddleware *JwtMiddleware
var userService = services.GetUserServiceInstance()

// auth error types

var (
    ErrNotAdmin        = errors.New("Has not admin role.")
    ErrNotContentAdmin = errors.New("Has not content admin role.")
)

func init() {
    cfg := config.GetConfig()
    authMiddleware = &JwtMiddleware{
        &jwt.GinJWTMiddleware{
            Realm:         cfg.GetString("JWT_REALM"),
            Key:           []byte(cfg.GetString("JWT_KEY")),
            Timeout:       time.Hour,
            MaxRefresh:    time.Hour,
            Authenticator: authenticate,
            Authorizator:  authorize,
            Unauthorized:  unauthorized,
            PayloadFunc:   payload,
            TokenLookup:   "header: Authorization, query: token, cookie: jwt",
            TokenHeadName: "Bearer",
            TimeFunc:      time.Now,
        },
    }

}

func GetAuthMiddleware() *JwtMiddleware {
    return authMiddleware
}

func (mw *JwtMiddleware) AdminMiddleware(c *gin.Context) {
    claims := jwt.ExtractClaims(c)

    if !hasRole(int32(claims["role"].(float64)), models.User_ROLE_ADMIN) {
        c.Abort()
        mw.Unauthorized(c, http.StatusUnauthorized, mw.HTTPStatusMessageFunc(ErrNotAdmin, c))
        return
    }

    c.Next()
}

func (mw *JwtMiddleware) ContentAdminMiddleware(c *gin.Context) {
    claims := jwt.ExtractClaims(c)

    if !hasRole(int32(claims["role"].(float64)), models.User_ROLE_CONTENT_ADMIN) {
        c.Abort()
        mw.Unauthorized(c, http.StatusUnauthorized, mw.HTTPStatusMessageFunc(ErrNotContentAdmin, c))
        return
    }
    c.Next()
}

// login
func authenticate(username string, password string, c *gin.Context) (interface{}, bool) {

    user := userService.GetByUsername(username)

    if user != nil && services.CheckPasswordHash(password, user.Password) {
        return user, true
    }
    return nil, false

}

func authorize(user interface{}, c *gin.Context) bool {

    claims := jwt.ExtractClaims(c)

    return hasRole(int32(claims["role"].(float64)), models.User_ROLE_USER)

}

func unauthorized(c *gin.Context, code int, message string) {
    c.JSON(code, gin.H{
        "code":    code,
        "message": message,
    })
}

func payload(data interface{}) jwt.MapClaims {

    user := data.(*models.User)

    return map[string]interface{}{
        "role": user.Role,
    }
}

func hasRole(role int32, desiredRole models.User_Role) bool {
    return role >= int32(desiredRole)
}
