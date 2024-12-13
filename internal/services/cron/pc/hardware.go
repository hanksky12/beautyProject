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

type Hardware struct {
	AverageRow        int // 平均多少筆資料做一次分析
	RecordRepo        *repository.StatusRecord
	RecordAverageRepo *repository.StatusRecordAverage
	UserRepo          *repository.User
}

func (h *Hardware) Analyze() {
	defer timeUtils.TrackExecutionTime("Analyze", time.Now())
	users := h.UserRepo.FindAll()
	var wg = &sync.WaitGroup{}
	for _, user := range users {
		h.handleSingleUser(user, wg)
	}
	wg.Wait()
}

func (h *Hardware) handleSingleUser(user model.User, wg *sync.WaitGroup) {
	log.Infof("處理單一 User資料: %s", user.Name)
	for _, hardware := range []*enum.Hardware{enum.Cpu, enum.Disk, enum.Memory} {
		wg.Add(1)
		go h.handleSingleHardware(hardware, user, wg)
	}
}

func (h *Hardware) handleSingleHardware(hardware *enum.Hardware, user model.User, wg *sync.WaitGroup) {
	defer wg.Done()
	log.Infof("處理User資料: %s 硬體資料: %s", user.Name, hardware.Name)
	records, err := h.RecordRepo.FindByUserAndHardware(user.ID, hardware.Number, 500)
	if err != nil {
		log.Error("查詢失敗", err)
		return
	}
	if len(records) == 0 {
		log.Infof("%s 無資料", hardware.Name)
		return
	}
	lengthRows := len(records)
	log.Infof("資料總筆數: %v ", lengthRows)
	var wgAverage = &sync.WaitGroup{}
	for index := 0; index < lengthRows; index += h.AverageRow {
		endIndex := index + h.AverageRow
		strData := fmt.Sprintf("處理第%v到%v筆資料", index, endIndex)
		log.Info("開始", strData)
		if endIndex > lengthRows {
			wgAverage.Add(1)
			go h.computeAverage(hardware, user, records[index:], wgAverage, strData)
			break
		}
		wgAverage.Add(1)
		go h.computeAverage(hardware, user, records[index:endIndex], wgAverage, strData)
	}
	wgAverage.Wait()
}

func (h *Hardware) computeAverage(hardware *enum.Hardware, user model.User, partRecords []model.MiniStatusRecord, wgAverage *sync.WaitGroup, strData string) {
	defer wgAverage.Done()
	record := &model.StatusRecordAverage{
		UserId:     user.ID,
		HardwareId: uint(hardware.Number),
		Percent:    h.getAveragePercent(partRecords),
		Time:       partRecords[0].Time,
	}
	ids := h.getIds(partRecords)
	err := h.updateTables(ids, record)
	if err != nil {
		log.Error(strData, "=>更新失敗=>", err)
	}
	log.Info(strData, "=>更新完成")
}

func (h *Hardware) getAveragePercent(partRecords []model.MiniStatusRecord) float64 {
	var total float64
	for _, record := range partRecords {
		total += record.Percent
	}
	return total / float64(len(partRecords))
}

func (h *Hardware) getIds(partRecords []model.MiniStatusRecord) []int {
	ids := make([]int, len(partRecords))
	for i, record := range partRecords {
		ids[i] = int(record.ID)
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
