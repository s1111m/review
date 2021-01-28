//package internal/handlers/server.go
package handlers

import (
	"context"
	"server/internal/hash"
	"server/pkg/hashservice"
)

type Server struct {
	hashservice.UnimplementedHashServiceServer
	HashesResponse *hashservice.ProtoArrayOfHashes
}

func (s *Server) CreateHash(ctx context.Context, convertToHashes *hashservice.ProtoArrayOfStrings) (*hashservice.ProtoArrayOfHashes, error) {

	s.HashesResponse = hash.GetHashesFromProtoArrayOfStrings(convertToHashes)

	return s.HashesResponse, nil
}
