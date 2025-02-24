package server

import (
	"beautyProject/internal/pkg/dto"
	"beautyProject/internal/pkg/enum/hardware"
	"beautyProject/internal/pkg/model"
	"beautyProject/internal/pkg/repository"
	"fmt"
	log "github.com/sirupsen/logrus"
	"strconv"
)

type Hardware struct {
	RecordRepo    *repository.HardwareStatusRecord
	RecordRepoRaw *repository.HardwareStatusRecordRaw
}

func (h *Hardware) Record(hardware *hardware.Hardware, value string, headers map[string]string, isRaw bool) {
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
	if isRaw {
		record := &model.HardwareStatusRecordRaw{
			UserId:     userId,
			HardwareId: hardwareId,
			Percent:    percent,
			Time:       time}
		err = h.RecordRepoRaw.Add(record)
	} else {
		percent, err = strconv.ParseFloat(fmt.Sprintf("%.1f", percent), 64)
		if err != nil {
			log.Error(err)
		}
		record := &model.HardwareStatusRecord{
			UserId:     userId,
			HardwareId: hardwareId,
			Percent:    percent,
			Time:       time,
			Processed:  false,
		}
		err = h.RecordRepo.Add(record)
	}
	if err != nil {
		msg := dto.Msg{Success: false, Message: "寫入失敗"}
		log.Infof("%v", msg)
	}
}
