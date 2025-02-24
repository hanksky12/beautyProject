package record

import (
	"beautyProject/internal/pkg/mocks"
	"beautyProject/internal/services/backend/server/hardware/record/base"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
	"testing"
)

type Suite struct {
	suite.Suite
	userId       string
	mockHardware *mocks.IWorkAndName
}

func (suite *Suite) SetupSuite() {
	suite.userId = "1"
	suite.mockHardware = new(mocks.IWorkAndName)
	suite.mockHardware.On("Work", "1")
	suite.mockHardware.On("Name").Return("mock")

}

func TestStart(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (suite *Suite) TestRecorder_Success_Start() {
	log.Info("TestRecorder_Success_Start")
	recorder := base.NewRecorder(suite.mockHardware, suite.userId)
	recorder.Initialize()

	msgDto := recorder.Start()
	suite.Equal(true, msgDto.Success)
	suite.Equal("開始記錄", msgDto.Message)
	recorder.Stop()
}

func (suite *Suite) TestRecorder_Already_Start() {
	log.Info("TestRecorder_Already_Start")
	recorder := base.NewRecorder(suite.mockHardware, suite.userId)
	recorder.Initialize()

	_ = recorder.Start()
	msgDto := recorder.Start()

	suite.Equal(false, msgDto.Success)
	suite.Equal("已在記錄中...", msgDto.Message)
	recorder.Stop()
}

func (suite *Suite) TestRecorder_Not_Start() {
	log.Info("TestRecorder_Not_Start")
	recorder := base.NewRecorder(suite.mockHardware, suite.userId)
	recorder.Initialize()

	msgDto := recorder.Stop()
	suite.Equal(false, msgDto.Success)
	suite.Equal("沒有在記錄...", msgDto.Message)
}

func (suite *Suite) TestRecorder_Success_Stop() {
	log.Info("TestRecorder_Success_Stop")
	recorder := base.NewRecorder(suite.mockHardware, suite.userId)
	recorder.Initialize()

	_ = recorder.Start()
	msgDto := recorder.Stop()

	suite.Equal(true, msgDto.Success)
	suite.Equal("停止記錄", msgDto.Message)
}
