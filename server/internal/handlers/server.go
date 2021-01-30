//package internal/handlers/server.go
package handlers

import (
	"context"
	"server/internal/hash"
	"server/pkg/hashservice"
)

// grpc server
type Server struct {
	hashservice.UnimplementedHashServiceServer
	HashesResponse *hashservice.ProtoArrayOfHashes
}

// create hash from strings
func (s *Server) CreateHash(ctx context.Context, convertToHashes *hashservice.ProtoArrayOfStrings) (*hashservice.ProtoArrayOfHashes, error) {
	s.HashesResponse = hash.GetHashesFromProtoArrayOfStrings(convertToHashes)
	return s.HashesResponse, nil
}
