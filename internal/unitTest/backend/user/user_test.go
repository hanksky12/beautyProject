package record

import (
	"beautyProject/internal/pkg/mocks"
	"beautyProject/internal/pkg/model"
	"beautyProject/internal/pkg/util/str"
	"beautyProject/internal/pkg/web/request"
	"beautyProject/internal/services/backend/user/authentication"
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
	"testing"
)

type Suite struct {
	suite.Suite
	MockUserRepo *mocks.IRepoUser
	MockJwt      *mocks.ISetCookie
	UserReq      request.UserReq
	ModelUser    *model.User
}

func (suite *Suite) SetupTest() {
	suite.MockUserRepo = new(mocks.IRepoUser)
	suite.MockJwt = new(mocks.ISetCookie)
	suite.UserReq = request.UserReq{
		UserName:     "test",
		UserPassword: "test",
	}
	suite.ModelUser = &model.User{
		Name:               suite.UserReq.UserName,
		Password:           str.HashPassword(suite.UserReq.UserPassword),
		AuthorizationLevel: "user",
	}
}

func TestStart(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (suite *Suite) TestUser_Login_Success() {
	log.Info("TestUser_Login_Success")
	suite.MockUserRepo.On("FindByName", suite.UserReq.UserName).
		Return(suite.ModelUser, true)
	suite.MockUserRepo.On("Add", suite.ModelUser).
		Return(nil)
	suite.MockJwt.On("SetCookie", uint(0)).Return(nil)

	userServ := authentication.User{}
	msgDto := userServ.Login(suite.MockJwt, suite.MockUserRepo, suite.UserReq)

	suite.Equal(true, msgDto.Success)
	suite.Equal("登入成功", msgDto.Message)
}

func (suite *Suite) TestUser_Login_Not_Register() {
	log.Info("TestUser_Login_Not_Register")
	suite.MockUserRepo.On("FindByName", suite.UserReq.UserName).
		Return(suite.ModelUser, false)
	//suite.MockUserRepo.On("Add", suite.ModelUser).
	//	Return(nil)
	//suite.MockJwt.On("SetCookie", uint(0)).Return(nil)

	userServ := authentication.User{}
	msgDto := userServ.Login(suite.MockJwt, suite.MockUserRepo, suite.UserReq)

	suite.Equal(false, msgDto.Success)
	suite.Equal("用戶未註冊", msgDto.Message)
}

func (suite *Suite) TestUser_Login_Wrong_Password() {
	log.Info("TestUser_Login_Wrong_Password")
	modelUserWrong := &model.User{
		Name:               suite.UserReq.UserName,
		Password:           str.HashPassword("test1"),
		AuthorizationLevel: "user",
	}
	suite.MockUserRepo.On("FindByName", suite.UserReq.UserName).
		Return(modelUserWrong, true)
	//suite.MockUserRepo.On("Add", suite.ModelUser).
	//	Return(nil)
	//suite.MockJwt.On("SetCookie", uint(0)).Return(nil)

	userServ := authentication.User{}
	msgDto := userServ.Login(suite.MockJwt, suite.MockUserRepo, suite.UserReq)

	suite.Equal(false, msgDto.Success)
	suite.Equal("密碼錯誤", msgDto.Message)
}

func (suite *Suite) TestUser_Login_SetCookie_Failed() {
	log.Info("TestUser_Login_SetCookie_Failed")
	suite.MockUserRepo.On("FindByName", suite.UserReq.UserName).
		Return(suite.ModelUser, true)
	suite.MockUserRepo.On("Add", suite.ModelUser).
		Return(nil)
	suite.MockJwt.On("SetCookie", uint(0)).Return(errors.New("cookie error"))

	userServ := authentication.User{}
	msgDto := userServ.Login(suite.MockJwt, suite.MockUserRepo, suite.UserReq)

	suite.Equal(false, msgDto.Success)
	suite.Equal("設定cookie失敗", msgDto.Message)
}

func (suite *Suite) TestUser_Register_Success() {
	log.Info("TestUser_Register_Success")
	suite.MockUserRepo.On("FindByName", suite.UserReq.UserName).
		Return(nil, false)
	suite.MockUserRepo.On("Add", suite.ModelUser).
		Return(nil)
	suite.MockJwt.On("SetCookie", uint(0)).Return(nil)

	userServ := authentication.User{}
	msgDto := userServ.Register(suite.MockJwt, suite.MockUserRepo, suite.UserReq)

	suite.Equal(true, msgDto.Success)
	suite.Equal("註冊成功", msgDto.Message)
}

func (suite *Suite) TestUser_Register_Repeat() {
	log.Info("TestUser_Register_Repeat")
	suite.MockUserRepo.On("FindByName", suite.UserReq.UserName).
		Return(nil, true)
	//suite.MockUserRepo.On("Add", suite.ModelUser).
	//	Return(nil)
	//suite.MockJwt.On("SetCookie", uint(0)).Return(nil)
	userServ := authentication.User{}
	msgDto := userServ.Register(suite.MockJwt, suite.MockUserRepo, suite.UserReq)

	suite.Equal(false, msgDto.Success)
	suite.Equal("用戶名被註冊", msgDto.Message)
}

func (suite *Suite) TestUser_Register_Sql_Insert_Failed() {
	log.Info("TestUser_Register_Sql_Insert_Failed")
	suite.MockUserRepo.On("FindByName", suite.UserReq.UserName).
		Return(nil, false)
	suite.MockUserRepo.On("Add", suite.ModelUser).
		Return(errors.New("add error"))
	//suite.MockJwt.On("SetCookie", uint(0)).Return(nil)
	userServ := authentication.User{}
	msgDto := userServ.Register(suite.MockJwt, suite.MockUserRepo, suite.UserReq)

	suite.Equal(false, msgDto.Success)
	suite.Equal("註冊失敗", msgDto.Message)
}

func (suite *Suite) TestUser_Register_SetCookie_Failed() {
	log.Info("TestUser_Register_SetCookie_Failed")
	suite.MockUserRepo.On("FindByName", suite.UserReq.UserName).
		Return(nil, false)
	suite.MockUserRepo.On("Add", suite.ModelUser).
		Return(nil)
	suite.MockJwt.On("SetCookie", uint(0)).Return(errors.New("cookie error")).Once()

	userServ := authentication.User{}
	msgDto := userServ.Register(suite.MockJwt, suite.MockUserRepo, suite.UserReq)

	suite.Equal(false, msgDto.Success)
	suite.Equal("設定cookie失敗", msgDto.Message)
}
