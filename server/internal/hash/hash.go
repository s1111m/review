package hash

import (
	"crypto/sha512"
	"fmt"
	log "github.com/sirupsen/logrus"
)

func GetHash(data string) string {
	bytes := sha512.Sum512([]byte (data))
	value := fmt.Sprintf("%x",[]byte(bytes[:]))
	log.WithFields(log.Fields{
		"original": data,
		"hash":     value,
	}).Info("Sucessfuly hashed")
	return value

}

 func GetHashesFromStrings(data []string) []string {
 	var hashes []string
 	for _, str := range data {
 		hashes = append(hashes, GetHash(str))
 	}
 	return hashes
 }