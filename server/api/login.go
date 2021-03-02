package api

import (
	"Lacs/pkg/app"
	"Lacs/pkg/e"
	mogo "Lacs/pkg/setting"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

func Login(c *gin.Context)  {
	g := app.Gin{c}
	var result user
	userName := c.PostForm("userName")
	password := c.PostForm("password")
	collection := mogo.NewMongoClient("user")
	filter := bson.D{{"_userName", userName},{"_password", password}}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		g.Response(http.StatusInternalServerError,e.ERROR,nil)
		return
	}

	if len(result.UserID) >0 {

	}

	fmt.Println(result)
	g.Response(http.StatusOK,e.SUCCESS,result)
}