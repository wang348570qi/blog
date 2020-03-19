package api

import (
	"github.com/EDDYCJY/go-gin-example/service/auth_service"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	//"log"
	"net/http"
	"blog/pkg/app"
	"blog/pkg/e"
	"blog/pkg/util"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

// @Summary Get Auth
// @Produce json
// @Param username query string true "username"
// @Param password query string true "password"
// @Sucess 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /auth [get]
func GetAuth(c *gin.Context) {
	appG := app.Gin{C: c}
	username := c.Query("username")
	password := c.Query(("password"))
	valid := validation.Validation{}
	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)
	/* data:=make(map[string]interface{})
	code:=e.INVALID_PARAMS */
	if !ok {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
		/* isExist:=models.CheckAuth(username,password)
		if isExist{
			token,err:=util.GenerateToken(username,password)
			if err!=nil{
				code=e.ERROR_AUTH_TOKEN
			}else{
				data["token"]=token
				code=e.SUCCESS
			}
		}else{
			code=e.ERROR_AUTH
		} */
	}
	authServer := auth_service.Auth{Username: username, Password: password}
	isExist, err := authServer.Check()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}
	if !isExist {
		appG.Response(http.StatusUnauthorized, e.ERROR_AUTH, nil)
		return
	}
	token, err := util.GenerateToken(username, password)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"token": token,
	})

}
