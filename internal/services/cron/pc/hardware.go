package pc

import (
	"beautyProject/internal/pkg/enum"
	"beautyProject/internal/pkg/model"
	"beautyProject/internal/pkg/repository"
	timeUtils "beautyProject/internal/pkg/util/time"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"sync"
	"time"
)

var enumHardware = []*enum.Hardware{enum.Cpu, enum.Disk, enum.Memory}

type Hardware struct {
	AverageRow        int // 平均多少筆資料做一次分析
	RecordRepo        *repository.StatusRecord
	RecordAverageRepo *repository.StatusRecordAverage
	RecordQueryRepo   *repository.StatusRecordQuery
	UserRepo          *repository.User
}

func (h *Hardware) Analyze() {
	defer timeUtils.TrackExecutionTime("Analyze", time.Now())
	records, err := h.RecordQueryRepo.FindByNoProcessedWithUser(5000)
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

func (h *Hardware) splitStatusRecordByUser(records []repository.StatusRecordWithUser) map[uint][]repository.StatusRecordWithUser {
	userMap := make(map[uint][]repository.StatusRecordWithUser)
	for _, record := range records {
		userMap[record.UserId] = append(userMap[record.UserId], record)
	}
	return userMap
}

func (h *Hardware) handleSingleUser(userID uint, records []repository.StatusRecordWithUser, wg *sync.WaitGroup) {
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

func (h *Hardware) getHardwareRecord(records []repository.StatusRecordWithUser, hardware *enum.Hardware) []repository.StatusRecordWithUser {
	partRecords := make([]repository.StatusRecordWithUser, 0)
	for _, record := range records {
		if record.HardwareId == uint(hardware.Number) {
			partRecords = append(partRecords, record)
		}
	}
	return partRecords
}

func (h *Hardware) handleSingleHardware(hardware *enum.Hardware, userID uint, userWg *sync.WaitGroup, records []repository.StatusRecordWithUser, logger *log.Entry) {
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

func (h *Hardware) computeAverage(hardware *enum.Hardware, userID uint, partRecords []repository.StatusRecordWithUser, averageWg *sync.WaitGroup, strData string, logger *log.Entry) {
	defer averageWg.Done()
	record := &model.StatusRecordAverage{
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

func (h *Hardware) getAveragePercent(partRecords []repository.StatusRecordWithUser) float64 {
	var total float64
	for _, record := range partRecords {
		total += record.Percent
	}
	return total / float64(len(partRecords))
}

func (h *Hardware) getIds(partRecords []repository.StatusRecordWithUser) []int {
	ids := make([]int, len(partRecords))
	for i, record := range partRecords {
		ids[i] = int(record.UserId)
	}
	return ids
}

func (h *Hardware) updateTables(ids []int, averageRecord *model.StatusRecordAverage) error {
	return repository.RunTransaction(func(tx *gorm.DB) error {
		// 更新 StatusRecord
		if err := h.RecordRepo.UpdateProcessed(tx, ids); err != nil {
			return err
		}
		// 添加 StatusRecordAverage
		if err := h.RecordAverageRepo.Add(tx, averageRecord); err != nil {
			return err
		}
		return nil
	})
}
