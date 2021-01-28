package db

import (
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

func DbWrite(hashes models.ArrayOfHash) models.ArrayOfHash {
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
		fmt.Println("alarm")
	}
	db.AutoMigrate(&models.Hash{})
	// db.Create(&models.Hash{Hash: "a7ffc6f8bf1ed76651c14756a061d662f580ff4de43b49fa82d80a4b80f8434a"})
	// db.Create(&models.Hash{Hash: "!fc6f8bf1ed76651c14756a061d662f580ff4de43b49fa82d80a4b80f8434lla"})
	// db.Create(&models.Hash{Hash: "!!ffc6f8bf1ed76651c14756a061d662f580ff4de43b49fa82d80a4b80f8434a"})
	// db.Create(&models.Hash{Hash: "!!!fc6f8bf1ed76651c14756a061d662f580ff4de43b49fa82d80a4b80f8434a"})
	//db.First(&user)
	//fmt.Printf("%v", &user)
}
