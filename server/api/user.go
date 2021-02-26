package api

import (
	"Lacs/pkg/app"
	"Lacs/pkg/e"
	mogo "Lacs/pkg/setting"
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

type user struct {
	UserID string `bson:"_id,omitempty" json:"userId"`
	UserName string `bson:"_userName,omitempty" json:"userName"`
	Password string `bson:"_password,omitempty" json:"password"`
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

func Login(c *gin.Context)  {
	g := app.Gin{c}
	var result user
	userName := c.PostForm("userName")
	password := c.PostForm("password")
	collection := mogo.NewMongoClient("user")
	filter := bson.D{{"_userName", userName},{"_password", password}}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		panic(err)
	}
	g.Response(http.StatusOK,e.SUCCESS,result)
}