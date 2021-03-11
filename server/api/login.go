package api

import (
	"Lacs/pkg/app"
	"Lacs/pkg/e"
	mogo "Lacs/pkg/setting"
	"Lacs/pkg/util"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"

	ep "Lacs/pkg/encryption"
	"Lacs/pkg/setting"
)

func Login(c *gin.Context)  {
	g := app.Gin{c}
	var result user
	data := make(map[string]interface{})

	userName := c.PostForm("userName")
	passWord := c.PostForm("passWord")

	collection := mogo.NewMongoClient("user")

	// AES 加密
	encryptCode := ep.AesEncrypt(passWord, setting.AppSetting.AesSecret)

	filter := bson.D{{"_userName", userName},{"_password", passWord}}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		fmt.Println(err)
		g.Response(http.StatusInternalServerError,e.ERROR,"查询出错")
		return
	}

	if len(result.UserID) >0 {
		token, err := util.GenerateToken(userName, passWord)
		if err != nil {
			g.Response(http.StatusInternalServerError,e.ERROR,nil)
			return
		}
		data["token"] = token
		data["info"] = result
		data["key"] = encryptCode
	}
	g.Response(http.StatusOK,e.SUCCESS,data)
}

