package main

import "github.com/gin-gonic/gin"

func registerRouter(r *gin.Engine) {
	r.Static("/frontend", "./frontend")

	r.GET("/", doIndex)

	v1 := r.Group("/dcmp/v1")
	v1.GET("/key/list", doKeyList)
	v1.POST("/key/new", doKeyNew)
	v1.POST("/key/delete", doKeyDelete)
	v1.POST("/key/save", doKeySave)
}
