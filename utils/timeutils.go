package utils

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func NewTimestampPB(t *time.Time) *timestamppb.Timestamp {
	if t == nil {
		return nil
	}

	return timestamppb.New(*t)
}
