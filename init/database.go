package init

import (
	"fintechpractices/global"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/v2pro/plz/countlog/output/lumberjack"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func getGormLogWriter() logger.Writer {
	var writer io.Writer

	dbCfg := global.AppCfg.DbCfg
	logCfg := global.AppCfg.LogCfg
	if !dbCfg.LogToFile {
		writer = os.Stdout
	} else {
		writer = &lumberjack.Logger{
			Filename:   strings.Join([]string{logCfg.Dir, dbCfg.FileName}, "/"),
			MaxSize:    logCfg.MaxSize,
			MaxBackups: logCfg.MaxBackups,
			MaxAge:     logCfg.MaxAge,
		}
	}

	return log.New(writer, "\r\n", log.LstdFlags)
}

func getGormLogger() logger.Interface {
	var logMode logger.LogLevel

	dbCfg := global.AppCfg.DbCfg
	switch dbCfg.LogMode {
	case "silent":
		logMode = logger.Silent
	case "error":
		logMode = logger.Error
	case "warn":
		logMode = logger.Warn
	case "info":
		logMode = logger.Info
	default:
		logMode = logger.Info
	}

	return logger.New(getGormLogWriter(), logger.Config{
		SlowThreshold:             time.Duration(dbCfg.SlowThreshold) * time.Millisecond,
		LogLevel:                  logMode,
		IgnoreRecordNotFoundError: false,
		Colorful:                  !dbCfg.LogToFile,
	})
}

func InitMySQLGorm() *gorm.DB {
	dbCfg := global.AppCfg.DbCfg

	if dbCfg.Database == "" {
		return nil
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbCfg.User, dbCfg.Password, dbCfg.Host, dbCfg.Port, dbCfg.Database,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   getGormLogger(),
	})
	if err != nil {
		global.Log.Sugar().Errorf("mysql connect failed: %s", err.Error())
		return nil
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(dbCfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(dbCfg.MaxOpenConns)

	return db
}
