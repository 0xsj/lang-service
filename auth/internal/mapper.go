package internal

import (
	"fmt"

	pb "github.com/0xsj/kakao/auth/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Mapper defines the methods for mapping between Protobuf messages, DTOs, and internal entities.
type Mapper interface {
	ProtoToEntity(protoMessage interface{}) (interface{}, error)
	EntityToProto(entity interface{}) (interface{}, error)
	
}

// MapperImpl implements the Mapper interface.
type MapperImpl struct{}

// NewMapper creates a new Mapper instance.
func NewMapper() Mapper {
	return &MapperImpl{}
}

// ProtoToEntity converts Protobuf messages to internal entities.
func (m *MapperImpl) ProtoToEntity(protoMessage interface{}) (interface{}, error) {
	switch msg := protoMessage.(type) {
	case *pb.OTP:
		return &OtpEntity{
			UserID:   msg.GetUserId(),
			Token:    msg.GetToken(),
			IssuedAt: msg.GetIssuedAt().AsTime(),
		}, nil
	case *pb.Session:
		return &SessionTokenEntity{
			UserID:   msg.GetUserId(),
			Token:    msg.GetToken(),
			IssuedAt: msg.GetIssuedAt().AsTime(),
		}, nil
	case *pb.Challenge:
		return &ChallengeEntity{
			UserID:   msg.GetUserId(),
			Challenge: msg.GetChallenge(),
			IssuedAt: msg.GetIssuedAt().AsTime(),
		}, nil
	default:
		return nil, fmt.Errorf("unknown proto message type: %T", protoMessage)
	}
}

// EntityToProto converts internal entities to Protobuf messages.
func (m *MapperImpl) EntityToProto(entity interface{}) (interface{}, error) {
	switch e := entity.(type) {
	case *OtpEntity:
		return &pb.OTP{
			UserId:   e.UserID,
			Token:    e.Token,
			IssuedAt: timestamppb.New(e.IssuedAt),
		}, nil
	case *SessionTokenEntity:
		return &pb.Session{
			UserId:   e.UserID,
			Token:    e.Token,
			IssuedAt: timestamppb.New(e.IssuedAt),
		}, nil
	case *ChallengeEntity:
		return &pb.Challenge{
			UserId:   e.UserID,
			Challenge: e.Challenge,
			IssuedAt: timestamppb.New(e.IssuedAt),
		}, nil
	default:
		return nil, fmt.Errorf("unknown entity type: %T", entity)
	}
}
