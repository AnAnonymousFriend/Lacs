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

type role struct {
	RoleID string `bson:"_id,omitempty" json:"roleId"`
	RoleName string `bson:"roleName" json:roleName`
}


// @Summary 添加角色
// @Produce  json
// @Success 200 {string} string "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/login [Get]
func AddRole(c *gin.Context) {
	g := app.Gin{c}
	roleNameParm := c.PostForm("roleName")

	var parm = role{
		RoleName: roleNameParm,
	}
	collection := mogo.NewMongoClient("role")

	if collection ==nil {
		println("Collection")
		g.Response(http.StatusInternalServerError,e.ERROR,e.GetMsg(e.ERROR))
	}
	install, err := collection.InsertOne(context.Background(), parm)

	if err !=nil || install ==nil {
		println(err)

		g.Response(http.StatusInternalServerError,e.ERROR,e.GetMsg(e.ERROR))
	}
	g.Response(http.StatusOK,e.SUCCESS,e.GetMsg(e.SUCCESS))
}

func FindOne(c *gin.Context)  {
	g := app.Gin{c}
	var result role
	roleNameParm := c.Query("roleName")
	collection := mogo.NewMongoClient("role")
	filter := bson.D{{"roleName", roleNameParm}}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		panic(err)
	}
	g.Response(http.StatusOK,e.SUCCESS,result)

}

func FindAll(c *gin.Context)  {
	g := app.Gin{c}
	collection := mogo.NewMongoClient("role")
	var roles []*role


	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		panic(err)
	}

	for cur.Next(context.TODO()) {
		// create a value into which the single document can be decoded
		var elem role
		err := cur.Decode(&elem)
		if err != nil {
			println(err)
		}
		roles = append(roles, &elem)
	}




	g.Response(http.StatusOK,e.SUCCESS,roles)
}


func DeleteRoleById(c *gin.Context)  {
	g := app.Gin{c}
	roleId := c.Query("roleId")
	collection := mogo.NewMongoClient("role")
	filter := bson.D{{"roleId", roleId}}
	result,err := collection.DeleteOne(context.Background(),filter)
	if err !=nil {
		g.Response(http.StatusInternalServerError,e.ERROR,nil)
	}
	g.Response(http.StatusOK,e.SUCCESS,result)
}
