package base

import (
	"beautyProject/internal/pkg/dto"
	"fmt"
	log "github.com/sirupsen/logrus"
	"time"
)

type Recorder struct {
	hardware WorkAndName
	userId   string
}

func NewRecorder(hardware WorkAndName, userId string) *Recorder {
	return &Recorder{
		hardware: hardware,
		userId:   userId,
	}
}

func (r *Recorder) Initialize() {
	if personalStatusRecord == nil {
		personalStatusRecord = make(map[string]hardware)
	}
	if _, exists := personalStatusRecord[r.userId]; !exists {
		personalStatusRecord[r.userId] = hardware{
			Name: make(map[string]state),
		}
	}
	if _, exists := personalStatusRecord[r.userId].Name[r.hardware.Name()]; !exists {
		personalStatusRecord[r.userId].Name[r.hardware.Name()] = state{
			IsWorking: false,
			StopChan:  make(chan bool), // 初始化通道
		}
	}
}

func (r *Recorder) Start() dto.Msg {
	if r.GetIsWorking() {
		msg := dto.Msg{Success: false, Message: "已在記錄中..."}
		log.Infof("%v", msg)
		return msg
	}

	r.SetStopChan(make(chan bool)) // 創建一個新的 channel 用於停止信號
	r.SetsWorking(true)
	panicChan := make(chan string, 1)
	go r.Run(panicChan) // 啟動工作線程

	var msg dto.Msg
	select {
	case result := <-panicChan:
		msg = dto.Msg{Success: false, Message: result}
	case <-time.After(3 * time.Second):
		msg = dto.Msg{Success: true, Message: "開始記錄"}
	}
	log.Infof("%v", msg)
	return msg

}

func (r *Recorder) Stop() dto.Msg {
	if !r.GetIsWorking() {
		msg := dto.Msg{Success: false, Message: "沒有在記錄..."}
		log.Infof("%v", msg)
		return msg
	}
	r.GetStopChan() <- true // 發送停止信號
	close(r.GetStopChan())  // 關閉 channel
	r.SetsWorking(false)

	msg := dto.Msg{Success: true, Message: "停止記錄"}
	log.Infof("%v", msg)
	return msg
}

func (r *Recorder) Run(panicChan chan string) {
	log.Info("Worker started")
	defer func() {
		if r := recover(); r != nil {
			panicChan <- fmt.Sprintf("Recovered from panic: %v\n", r)
		}
	}()

	for {
		select {
		case <-r.GetStopChan():
			log.Info("Worker stopped")
			return
		default:
			r.hardware.Work(r.userId)
		}
	}
}

func (r *Recorder) GetIsWorking() bool {
	return personalStatusRecord[r.userId].Name[r.hardware.Name()].IsWorking
}

func (r *Recorder) GetStopChan() chan bool {
	return personalStatusRecord[r.userId].Name[r.hardware.Name()].StopChan
}

func (r *Recorder) SetsWorking(bool2 bool) {
	state := personalStatusRecord[r.userId].Name[r.hardware.Name()] // 取出複製值
	state.IsWorking = bool2                                         // 修改值
	personalStatusRecord[r.userId].Name[r.hardware.Name()] = state  // 更新回 map
}

func (r *Recorder) SetStopChan(chan2 chan bool) {
	state := personalStatusRecord[r.userId].Name[r.hardware.Name()]
	state.StopChan = chan2
	personalStatusRecord[r.userId].Name[r.hardware.Name()] = state
}
