package pc

import (
	"beautyProject/internal/pkg/dto"
	"beautyProject/internal/pkg/enum"
	"beautyProject/internal/pkg/model"
	"beautyProject/internal/pkg/repository"
	log "github.com/sirupsen/logrus"
	"strconv"
)

type Hardware struct {
}

func (h *Hardware) Record(hardware *enum.Hardware, key string, value string, headers map[string]string, minutes int, repo repository.StatusRecord) {
	//log.Info(minutes)
	//log.Info(key, value)
	//log.Info(headers)
	log.Infof("%s is working...", hardware.Name)
	var err error
	percent, err := strconv.ParseFloat(value, 64)
	if err != nil {
		log.Error(err)
	}
	time, err := strconv.ParseInt(headers["time"], 10, 64)
	if err != nil {
		log.Error(err)
	}
	hardwareId := uint(hardware.Number)
	userIDUint, err := strconv.ParseUint(headers["user_id"], 10, 64)
	if err != nil {
		log.Error(err)
	}
	userId := uint(userIDUint)
	record := &model.StatusRecord{
		UserId:     userId,
		HardwareId: hardwareId,
		Percent:    percent,
		Time:       time,
		Min:        minutes}
	err = repo.Add(record)
	if err != nil {
		msg := dto.Msg{Success: false, Message: "寫入失敗"}
		log.Infof("%v", msg)
		//todo something
	}
}
