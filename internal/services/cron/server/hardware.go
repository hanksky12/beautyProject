package server

import (
	"beautyProject/internal/pkg/entity"
	"beautyProject/internal/pkg/enum/hardware"
	"beautyProject/internal/pkg/model"
	"beautyProject/internal/pkg/repository"
	"beautyProject/internal/pkg/util/db/sql"
	timeUtils "beautyProject/internal/pkg/util/time"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"sync"
	"time"
)

var enumHardware = []*hardware.Hardware{hardware.Cpu, hardware.Disk, hardware.Memory}

type Hardware struct {
	AverageRow        int // 平均多少筆資料做一次分析
	RecordRepo        *repository.HardwareStatusRecord
	RecordAverageRepo *repository.HardwareStatusRecordAverage
	RecordQueryRepo   *repository.HardwareStatusRecordQuery
}

func (h *Hardware) Analyze() {
	defer timeUtils.TrackExecutionTime("Analyze", time.Now())
	records, err := h.RecordQueryRepo.FindByNoProcessedWithUser(5000000)
	if err != nil {
		log.Error("查詢失敗", err)
		return
	}
	if len(records) == 0 {
		log.Info("無資料")
		return
	}
	userMap := h.splitStatusRecordByUser(records)
	var wg = &sync.WaitGroup{}
	for userID, statusRecords := range userMap {
		//user 一定有資料
		wg.Add(1)
		go h.handleSingleUser(userID, statusRecords, wg)
	}
	wg.Wait()
}

func (h *Hardware) splitStatusRecordByUser(records []entity.HardwareStatusRecordWithUser) map[uint][]entity.HardwareStatusRecordWithUser {
	userMap := make(map[uint][]entity.HardwareStatusRecordWithUser)
	for _, record := range records {
		userMap[record.UserId] = append(userMap[record.UserId], record)
	}
	return userMap
}

func (h *Hardware) handleSingleUser(userID uint, records []entity.HardwareStatusRecordWithUser, wg *sync.WaitGroup) {
	defer wg.Done()
	logger := log.WithFields(log.Fields{"userID": userID})
	logger.Infof("處理單一User資料: %s", userID)
	var userWg = &sync.WaitGroup{}
	for _, hardware := range enumHardware {
		hardwareRecords := h.getHardwareRecord(records, hardware)
		if len(hardwareRecords) == 0 {
			logger.Infof("無%s資料", hardware.Name)
			continue
		}
		userWg.Add(1)
		go h.handleSingleHardware(hardware, userID, userWg, hardwareRecords, logger)
	}
	userWg.Wait()
}

func (h *Hardware) getHardwareRecord(records []entity.HardwareStatusRecordWithUser, hardware *hardware.Hardware) []entity.HardwareStatusRecordWithUser {
	partRecords := make([]entity.HardwareStatusRecordWithUser, 0)
	for _, record := range records {
		if record.HardwareId == uint(hardware.Number) {
			partRecords = append(partRecords, record)
		}
	}
	return partRecords
}

func (h *Hardware) handleSingleHardware(hardware *hardware.Hardware, userID uint, userWg *sync.WaitGroup, records []entity.HardwareStatusRecordWithUser, logger *log.Entry) {
	defer userWg.Done()
	logger = logger.WithFields(log.Fields{"hardwareName": hardware.Name})
	lengthRows := len(records)
	logger.Infof("筆數: %v ", lengthRows)
	var averageWg = &sync.WaitGroup{}
	for index := 0; index < lengthRows; index += h.AverageRow {
		endIndex := index + h.AverageRow
		strData := fmt.Sprintf("處理第%v到%v筆資料", index, endIndex)
		logger.Info("開始", strData)
		if endIndex > lengthRows {
			averageWg.Add(1)
			go h.computeAverage(hardware, userID, records[index:], averageWg, strData, logger)
			break
		}
		averageWg.Add(1)
		go h.computeAverage(hardware, userID, records[index:endIndex], averageWg, strData, logger)
	}
	averageWg.Wait()
}

func (h *Hardware) computeAverage(hardware *hardware.Hardware, userID uint, partRecords []entity.HardwareStatusRecordWithUser, averageWg *sync.WaitGroup, strData string, logger *log.Entry) {
	defer averageWg.Done()
	record := &model.HardwareStatusRecordAverage{
		UserId:     userID,
		HardwareId: uint(hardware.Number),
		Percent:    h.getAveragePercent(partRecords),
		Time:       partRecords[0].Time,
	}
	ids := h.getIds(partRecords)
	err := h.updateTables(ids, record)
	if err != nil {
		logger.Error(strData, "=>更新失敗=>", err)
	}
	logger.Info(strData, "=>更新完成")
}

func (h *Hardware) getAveragePercent(partRecords []entity.HardwareStatusRecordWithUser) float64 {
	var total float64
	for _, record := range partRecords {
		total += record.Percent
	}
	return total / float64(len(partRecords))
}

func (h *Hardware) getIds(partRecords []entity.HardwareStatusRecordWithUser) []int {
	ids := make([]int, len(partRecords))
	for i, record := range partRecords {
		ids[i] = int(record.UserId)
	}
	return ids
}

func (h *Hardware) updateTables(ids []int, averageRecord *model.HardwareStatusRecordAverage) error {
	return sql.RunTransaction(func(tx *gorm.DB) error {
		// 更新 HardwareStatusRecord
		if err := h.RecordRepo.UpdateProcessed(tx, ids); err != nil {
			return err
		}
		// 添加 HardwareStatusRecordAverage
		if err := h.RecordAverageRepo.Add(tx, averageRecord); err != nil {
			return err
		}
		return nil
	})
}
