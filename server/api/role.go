package api

import (
	"Lacs/pkg/app"
	"Lacs/pkg/e"
	mogo "Lacs/pkg/setting"
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	 "go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"fmt"
)

type role struct {
	RoleID string `bson:"_id,omitempty" json:"roleId"`
	RoleName string `bson:"roleName" json:roleName`
}


// @Summary 添加角色
// @Produce  json
// @Success 200 {string} string "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/AddRole [Get]
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

// @Summary 查找角色
// @Produce  json
// @Success 200 {string} string "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/findOne [Get]
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

// @Summary 查找角色列表
// @Produce  json
// @Success 200 {string} string "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/findAll [Get]
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

// @Summary 根据角色ID 删除
// @Produce  json
// @Success 200 {string} string "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/deleteRoleById [Get]
func DeleteRoleById(c *gin.Context)  {
	g := app.Gin{c}
	roleId := c.PostForm("roleId")
	obj_id, err := primitive.ObjectIDFromHex(roleId)
	if err != nil {
		fmt.Println(err)
		g.Response(http.StatusInternalServerError,e.ERROR,nil)
		return
	}
	collection := mogo.NewMongoClient("role")
	filter := bson.D{{"_id", obj_id}}
	result,err := collection.DeleteMany(context.Background(),filter)
	if err !=nil {
		g.Response(http.StatusInternalServerError,e.ERROR,nil)
		return
	}
	g.Response(http.StatusOK,e.SUCCESS,result.DeletedCount)
}

// @Summary 删除所有
// @Produce  json
// @Success 200 {string} string "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/deleteRoleById [Get]
func DeleteRoleAll(c *gin.Context)  {
	g := app.Gin{c}
	collection := mogo.NewMongoClient("role")
	result,err := collection.DeleteMany(context.Background(),nil)
	if err !=nil {
		g.Response(http.StatusInternalServerError,e.ERROR,nil)
		return
	}
	g.Response(http.StatusOK,e.SUCCESS,result.DeletedCount)
}
