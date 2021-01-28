package handler

import (
	"context"
	"errors"
	"router/internal/config"
	"router/models"
	"router/pkg/hashservice"
	"time"

	"google.golang.org/grpc"
)

func convertStringToProtoStringArray(strings models.ArrayOfStrings) *hashservice.ProtoArrayOfStrings {
	strs := &hashservice.ProtoArrayOfStrings{}
	for _, str := range strings {
		strs.StrToConvert = append(strs.StrToConvert, &hashservice.ProtoStringToHash{Str: str})
	}
	return strs

}

func convertProtoHashArrayToArrayOfHashes(hashes *hashservice.ProtoArrayOfHashes) models.ArrayOfHash {
	arrayOfHashes := models.ArrayOfHash{}
	for _, str := range hashes.Hashes {
		arrayOfHashes = append(arrayOfHashes, &models.Hash{Hash: str.GetHash()})
	}
	return arrayOfHashes
}
func sendToHashservice(strings models.ArrayOfStrings) (*hashservice.ProtoArrayOfHashes, error) {

	cwt, _ := context.WithTimeout(context.Background(), time.Second*60)

	conn, err := grpc.DialContext(cwt, config.Cfg.GRPC_SERVICE_ADDR+":"+config.Cfg.GRPC_SERVICE_PORT, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	uc := hashservice.NewHashServiceClient(conn)

	us, err := uc.CreateHash(cwt, convertStringToProtoStringArray(strings))
	if err != nil {
		return nil, errors.New("Can't make hashes")
	}
	return us, nil
}
