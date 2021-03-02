package api

import (
	"Lacs/pkg/app"
	"Lacs/pkg/e"
	mogo "Lacs/pkg/setting"
	"Lacs/pkg/util"
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

func Login(c *gin.Context)  {
	g := app.Gin{c}
	var result user
	data := make(map[string]interface{})

	userName := c.PostForm("userName")
	passWord := c.PostForm("password")
	collection := mogo.NewMongoClient("user")
	filter := bson.D{{"_userName", userName},{"_password", passWord}}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		g.Response(http.StatusInternalServerError,e.ERROR,nil)
		return
	}

	if len(result.UserID) >0 {
		token, err := util.GenerateToken(userName, passWord)
		if err != nil {
			g.Response(http.StatusInternalServerError,e.ERROR,nil)
			return
		}
		data["token"] = token
	}
	data["info"] = result
	g.Response(http.StatusOK,e.SUCCESS,data)
}