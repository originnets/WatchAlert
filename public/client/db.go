package client

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"watchAlert/models"
	"watchAlert/public/globals"
)

func InitDB() *gorm.DB {

	// 初始化本地 test.db 数据库文件
	//db, err := gorm.Open(sqlite.Open("data/sql.db"), &gorm.Config{})

	sql := globals.Config.MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4,utf8&parseTime=True&loc=Local&timeout=%s", sql.User, sql.Pass, sql.Host, sql.Port, sql.DBName, sql.Timeout)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("failed to connect database")
	}

	// 检查 Product 结构是否变化，变化则进行迁移
	err = db.AutoMigrate(
		&models.DutySchedule{},
		&models.DutyManagement{},
		&models.AlertNotice{},
		&models.AlertDataSource{},
		&models.AlertRule{},
		&models.AlertCurEvent{},
		&models.AlertHisEvent{},
		&models.AlertSilences{},
		&models.Member{},
		&models.UserRole{},
		&models.UserPermissions{},
		&models.NoticeTemplateExample{},
		&models.RuleGroups{},
		&models.RuleTemplateGroup{},
		&models.RuleTemplate{},
		&models.ServiceResource{},
		&models.Tenant{},
		&models.Dashboard{},
		&models.AuditLog{},
	)
	if err != nil {
		return nil
	}

	return db
}
