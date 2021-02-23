package api

import (
	"Lacs/pkg/app"
	"Lacs/pkg/e"

	mogo "Lacs/pkg/setting"
	"context"
	"github.com/gin-gonic/gin"
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
func AddRole(ctx *gin.Context) {
	g := app.Gin{ctx}
	roleNameParm := ctx.PostForm("roleName")

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

func FindOne(ctx *gin.Context)  {
	g := app.Gin{ctx}
	collection := mogo.NewMongoClient("role")
	cn :=mogo.Mgo{
		collection,
	}

	one := cn.FindOne("admin","admin")
	g.Response(http.StatusOK,e.SUCCESS,one)
}
