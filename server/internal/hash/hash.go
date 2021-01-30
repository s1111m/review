package hash

import (
	"crypto/sha512"
	"fmt"
	"server/internal/config"
	"server/pkg/hashservice"
	"sync"

	"github.com/sirupsen/logrus"
)

func GetHash(data string) string {
	bytes := sha512.Sum512([]byte(data))
	value := fmt.Sprintf("%x", []byte(bytes[:]))
	config.Logger.WithFields(logrus.Fields{
		"original": data,
		"hash":     value,
	}).Trace("Sucessfuly hashed")
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
	config.Logger.WithFields(logrus.Fields{
		"request-id": arr.GetRequestId(),
	}).Trace("Start Hashing")
	hashesResponse := hashservice.ProtoArrayOfHashes{}
	//сюда будем наваливать результат работы тредов и добавлять в массив выше
	c := make(chan string)
	//
	// стоп-семафор, вычисляем число процов и тормозим если тредов больше
	semaphoreChan := make(chan struct{}, config.Cfg.MAX_THREADS-3)
	defer close(semaphoreChan)
	// горутина будет получать хэши в порядке запуска горутин и записывать в слайс.
	go func() {
		for {
			readyHash, ok := <-c
			if ok == false {
				break // exit break loop
			} else {
				hashesResponse.Hashes = append(hashesResponse.Hashes, &hashservice.ProtoHash{Hash: readyHash}) // собрали слайс и вернули
			}
		}
		defer close(c)
	}()
	// чтобы дождаться завершения всех тредов
	//погнали по всему массиву
	var wg sync.WaitGroup
	for _, str := range arr.StrToConvert {
		semaphoreChan <- struct{}{} // заминусовали семафор
		go func(str string) {
			wg.Add(1)
			c <- GetHash(str) //посчитали хэш, отправили в канал
			<-semaphoreChan   // почистили семафор
			wg.Done()
		}(str.Str)
	}
	//Блокируем выполнение функции до завершения всех тредов
	wg.Wait() // подождали
	// перекинули requestId и отправили обратным сообщением
	hashesResponse.RequestId = arr.RequestId
	config.Logger.WithFields(logrus.Fields{
		"request-id": hashesResponse.RequestId,
		//"response":   hashesResponse.Hashes,
	}).Trace("Sending response back")

	return &hashesResponse
}
