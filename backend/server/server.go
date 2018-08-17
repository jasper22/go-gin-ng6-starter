package server

import "github.com/app8izer/go-gin-ng6-starter/backend/utils"

func Init() {
    // setup router
    r := SetupRouter()

    // get port from env or use default 8080
    port := utils.GetEnv("PORT", "8080")

    // TODO: add ENV here (prod, dev, test) and setup router, server, db based on config

    // run http server
    r.Run(":" + port)

}
