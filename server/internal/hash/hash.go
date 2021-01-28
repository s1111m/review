package hash

import (
	"crypto/sha512"
	"fmt"
	"server/pkg/hashservice"
	"strconv"
	"sync"
)

func GetHash(data string) string {
	bytes := sha512.Sum512([]byte(data))
	value := fmt.Sprintf("%x", []byte(bytes[:]))
	// log.WithFields(log.Fields{
	// 	"original": data,
	// 	"hash":     value,
	// }).Info("Sucessfuly hashed")
	return value

}

func GetHashesFromStrings(data []string) []string {
	var hashes []string
	for _, str := range data {
		hashes = append(hashes, GetHash(str))
	}
	return hashes
}

func GetHashesFromProtoArrayOfStrings(arr *hashservice.ProtoArrayOfStrings) *hashservice.ProtoArrayOfHashes {
	//выделим память под наш ответ
	hashesResponse := hashservice.ProtoArrayOfHashes{}
	//сюда будем наваливать результат работы тредов и добавлять в массив выше
	c := make(chan string)
	// стоп-семафор, вычисляем число процов и тормозим если тредов больше
	semaphoreChan := make(chan struct{}, 16)
	defer close(semaphoreChan)
	defer close(c)
	// чтобы дождаться завершения всех тредов
	var wg sync.WaitGroup
	//погнали по всему массиву
	for _, str := range arr.StrToConvert {
		semaphoreChan <- struct{}{} // заминусовали семафор
		go func(str string) {
			defer wg.Done()
			wg.Add(1)
			fmt.Println(GetHash(str))
			c <- GetHash(str) //посчитали хэш, отправили в канал
			<-semaphoreChan   // почистили семафор
			return
		}(str.Str)
	}
	wg.Wait() // подождали

	for i := 0; i < len(arr.StrToConvert); i++ { //ожидаем столько сообщений, сколько длина массива
		readyHash := <-c
		fmt.Println(strconv.Itoa(i) + readyHash + "\r\n")
		hashesResponse.Hashes = append(hashesResponse.Hashes, &hashservice.ProtoHash{Hash: readyHash}) // собрали слайс и вернули
	}

	return &hashesResponse
}
