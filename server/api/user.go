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
	u "Lacs/pkg/util"
)

type user struct {
	UserID string `bson:"_id,omitempty" json:"userId"`
	UserName string `bson:"_userName,omitempty" json:"userName"`
	Password string `bson:"_password,omitempty" json:"password"`
	RoleId string `bson:"_roleId,omitempty" json:"roleId"`
}

// @Summary 查找角色列表
// @Produce  json
// @Success 200 {string} string "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/findAll [Get]
func FindUserAll(c *gin.Context)  {
	g := app.Gin{c}
	collection := mogo.NewMongoClient("user")
	var users []*user


	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		panic(err)
	}

	for cur.Next(context.TODO()) {
		// create a value into which the single document can be decoded
		var elem user
		err := cur.Decode(&elem)
		if err != nil {
			println(err)
		}
		users = append(users, &elem)
	}
	g.Response(http.StatusOK,e.SUCCESS,users)
}

func AddUser(c *gin.Context)  {
	g := app.Gin{c}

	name := c.PostForm("userName")
	password := c.PostForm("passWord")
	if len(password) > 0 {
		password = u.Md5V(password)
	}
	collection := mogo.NewMongoClient("user")

	var parm = user{
		UserName: name,
		Password: password,
	}

	inert, err := collection.InsertOne(context.Background(), parm)
	if err !=nil {
		println(err)
	}
	fmt.Println(inert)
	g.Response(http.StatusOK,e.SUCCESS,inert)

}
