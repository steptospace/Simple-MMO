package pg

import (
	"github.com/rs/zerolog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"math/rand"
	"time"
)

type Worker struct {
	conn   *gorm.DB
	logger zerolog.Logger
}

func (w *Worker) Close() error {
	db, _ := w.conn.DB()
	return db.Close()
}

type UserData struct {
	Login     string        `json:"login" gorm:"primaryKey"`
	Uid       int           `json:"uid" gorm:"primaryKey"`
	Password  string        `json:"password"`
	OnLine    bool          `json:"on_line"`
	LoginTime time.Duration `json:"login_time"`
}

func Init(pgConfig string, logger zerolog.Logger) (*Worker, error) {
	db, err := gorm.Open(postgres.Open(pgConfig), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(UserData{}); err != nil {
		return nil, err
	}

	return &Worker{
		conn:   db,
		logger: logger,
	}, nil
}

func (w *Worker) GetAccountData(login string, uid int) (acc UserData, err error) {
	w.conn.Where("login = ?", login).Where("uid = ?", uid).Find(acc)
	return acc, nil
}

func (w *Worker) InsertAccount(login string, password string) error {
	data := UserData{
		Login:    login,
		Password: password,
		Uid:      rand.Int(),
		OnLine:   true,
	}
	w.conn.Clauses(clause.OnConflict{DoNothing: true}).Create(&data)
	return nil
}
