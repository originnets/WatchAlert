package repo

import (
	"watchAlert/models"
	"watchAlert/public/globals"
)

type NoticeRepo struct{}

func (nr NoticeRepo) GetData(uuid string) models.AlertNotice {

	var alertNoticeData models.AlertNotice
	globals.DBCli.Model(&models.AlertNotice{}).Where("uuid = ?", uuid).Find(&alertNoticeData)
	return alertNoticeData

}

func (nr NoticeRepo) GetQuota(id string) bool {
	var (
		db     = globals.DBCli.Model(&models.Tenant{})
		data   models.Tenant
		Number int64
	)

	db.Where("id = ?", id)
	db.Find(&data)

	globals.DBCli.Model(&models.AlertNotice{}).Where("tenant_id = ?", id).Count(&Number)

	if Number < data.NoticeNumber {
		return true
	}

	return false
}

func (nr NoticeRepo) Search(req models.NoticeQuery) ([]models.AlertNotice, error) {
	var data []models.AlertNotice
	var db = globals.DBCli.Model(&models.AlertNotice{})
	db.Where("name LIKE ? OR env LIKE ? OR notice_type LIKE ?", "%"+req.Query+"%", "%"+req.Query+"%", "%"+req.Query+"%")
	err := db.Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}
