package db

import (
	"context"
	"errors"
	"fmt"
	"router/internal/config"
	"router/models"
	"strconv"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// для поиска по одному id
func DbGet(id uint) *models.Hash {
	user := models.Hash{}
	db.First(&user, id)
	return &user
}

func DbWrite(hashes models.ArrayOfHash, ctx context.Context) models.ArrayOfHash {

	fmt.Println(ctx.Value("request-id"))
	var response models.ArrayOfHash
	for _, hash := range hashes {
		db.Create(&hash)
		response = append(response, hash)
	}
	return response
}

func DbFind(ids []string) (models.ArrayOfHash, error) {
	var users models.ArrayOfHash
	var int_ids []int
	// конвертим массив строк-id в int
	for _, id := range ids {
		//конвертируем id из строки в число, и если ошибка, вываливаемся из запроса. Заодно при конвертации убираются вопросы sql-injection
		id_int, err := strconv.Atoi(id)
		// если сконвертилось - добавляем индекс в массив. Если была ошибка просто пропускаем этот элемент
		if err == nil {
			int_ids = append(int_ids, id_int)
		}
	}
	result := db.Find(&users, ids)
	if result.RowsAffected == 0 {
		return users, errors.New("No records found")
	}
	return users, nil
}

var db *gorm.DB

func init() {
	err := *new(error)
	db, err = gorm.Open(sqlite.Open(config.Cfg.DB_PATH), &gorm.Config{})
	if err != nil {
		config.Logger.WithError(err)
	}
	db.AutoMigrate(&models.Hash{})
}
