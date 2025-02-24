package user

import (
	"beautyProject/internal/pkg/dto"
	"beautyProject/internal/pkg/enum/mouseAction"
	"beautyProject/internal/pkg/model"
	"beautyProject/internal/pkg/repository"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"strconv"
)

type MouseAction struct {
	RecordRepoRaw *repository.MouseActionStatusRecordRaw
}

func (h *MouseAction) Record(action *mouseAction.MouseAction, value string, headers map[string]string) {
	log.Infof("%s is working...", action.Name)
	var err error
	var data map[string]interface{}
	err = json.Unmarshal([]byte(value), &data)
	if err != nil {
		log.Error(err)
	}

	xFloat, ok := data["x"].(float64) // JSON 數字會被解析為 float64
	if !ok {
		fmt.Println("x 轉換失敗")
		return
	}
	x := int64(xFloat)
	yFloat, ok := data["y"].(float64) // JSON 數字會被解析為 float64
	if !ok {
		fmt.Println("y 轉換失敗")
		return
	}
	y := int64(yFloat)
	time, err := strconv.ParseInt(headers["time"], 10, 64)
	if err != nil {
		log.Error(err)
	}
	actionId := uint(action.Number)
	userIDUint, err := strconv.ParseUint(headers["user_id"], 10, 64)
	if err != nil {
		log.Error(err)
	}
	userId := uint(userIDUint)

	record := &model.MouseActionStatusRecordRaw{
		UserId:        userId,
		MouseActionId: actionId,
		X:             x,
		Y:             y,
		Time:          time}
	err = h.RecordRepoRaw.Add(record)
	if err != nil {
		msg := dto.Msg{Success: false, Message: "寫入失敗"}
		log.Infof("%v", msg)
	}
}
