package version1

import (
	"time"
)

type UserPasswordInfoV1 struct {
	Id         string    `json:"id"`
	ChangeTime time.Time `json:"changed_time"`
	Locked     bool      `json:"locked"`
	LockTime   time.Time `json:"lock_time"`
}

func EmptyUserPasswordInfoV1() *UserPasswordInfoV1 {
	return &UserPasswordInfoV1{}
}

func NewUserPasswordInfoV1(id string, changeTime time.Time,
	locked bool, lockTime time.Time) *UserPasswordInfoV1 {
	return &UserPasswordInfoV1{
		Id:         id,
		ChangeTime: changeTime,
		Locked:     locked,
		LockTime:   lockTime,
	}
}
