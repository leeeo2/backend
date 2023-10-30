package controller

import (
	"fmt"
	"github.com/leeeo2/backend/pkg/common/jwt"
	"github.com/leeeo2/backend/pkg/common/logger"
	"github.com/leeeo2/backend/pkg/common/response"
	"github.com/leeeo2/backend/pkg/dao"
	"github.com/leeeo2/backend/pkg/domain"
	"github.com/leeeo2/backend/pkg/model"
	"github.com/leeeo2/backend/pkg/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(ctx *gin.Context) {
	// obtain arguments
	var input domain.RegisterInput
	if err := ctx.Bind(&input); err != nil {
		response.Resp(ctx, http.StatusUnprocessableEntity, 422, nil, "arguments error")
		return
	}

	// validate
	if len(input.Telephone) != 11 {
		response.Resp(ctx, http.StatusUnprocessableEntity, 422, nil, "telephone should be 11 bytes")
		return
	}
	if len(input.Password) < 6 {
		response.Resp(ctx, http.StatusUnprocessableEntity, 422, nil, "password should not less than 6 bytes")
		return
	}
	if len(input.Name) == 0 {
		input.Name = util.RandStr(10)
	}

	// check user by telephone
	logger.Info(ctx, "recv register request", "input", input)
	if dao.IsUserExist(input.Telephone) {
		response.Resp(ctx, http.StatusUnprocessableEntity, 422, nil, "user is exist")
		return
	}

	// create user
	id, err := dao.CreateUser(&model.User{
		Name:      input.Name,
		Telephone: input.Telephone,
		Password:  input.Password,
	})
	if err != nil {
		fmt.Println("create user failed,err:", err)
		response.Resp(ctx, http.StatusUnprocessableEntity, 422, nil, "create user failed")
		return
	}

	// register success
	output := domain.RegisterOutput{UserId: id}
	response.Success(ctx, output, "register success")
}

func Login(ctx *gin.Context) {
	// obtain arguments
	var input domain.LoginInput
	if err := ctx.Bind(&input); err != nil {
		response.Resp(ctx, http.StatusUnprocessableEntity, 422, nil, "arguments error")
		return
	}

	// validate
	if len(input.Telephone) != 11 {
		response.Resp(ctx, http.StatusUnprocessableEntity, 422, nil, "telephone should be 11 bytes")
		return
	}
	if len(input.Password) < 6 {
		response.Resp(ctx, http.StatusUnprocessableEntity, 422, nil, "password should not less than 6 bytes")
		return
	}
	// check telephone
	user, err := dao.DescribeUserByTelephone(input.Telephone)
	if err != nil {
		response.Resp(ctx, http.StatusInternalServerError, 500, nil, "query user failed")
		return
	} else if user.ID == 0 {
		response.Resp(ctx, http.StatusUnprocessableEntity, 422, nil, "user is not exist")
		return
	}

	// check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		response.Resp(ctx, http.StatusUnprocessableEntity, 422, nil, "password is wrong")
		return
	}

	// grant token
	token, err := jwt.ReleaseToken(user)
	if err != nil {
		response.Resp(ctx, http.StatusInternalServerError, 500, nil, "internal server error")
		return
	}

	// return result
	response.Success(ctx, domain.LoginOutput{Token: token}, "login success")
}

func UserInfo(ctx *gin.Context) {
	user, ok := ctx.Get("user")
	if !ok || user == nil {
		response.Resp(ctx, http.StatusUnprocessableEntity, 422, nil, "user not exist")
	}
	u := user.(*model.User)

	response.Success(ctx, domain.GetUserInfoOutput{
		UserId:    u.ID,
		Name:      u.Name,
		Telephone: u.Telephone,
	}, "success")
}
