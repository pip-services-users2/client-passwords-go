package version1

import (
	"time"
)

type UserPasswordV1 struct {
	/* Identification */
	Id       string `json:"id"`
	Password string `json:"password"`

	/* Password management */
	ChangeTime    time.Time `json:"changed_time"`
	Locked        bool      `json:"locked"`
	LockTime      time.Time `json:"lock_time"`
	FailCount     int32     `json:"fail_count"`
	FailTime      time.Time `json:"fail_time"`
	RecCode       string    `json:"rec_code"`
	RecExpireTime time.Time `json:"rec_expire_time"`

	/* Custom fields */
	CustomHdr interface{} `json:"custom_hdr"`
	CustomDat interface{} `json:"custom_dat"`
}

func EmptyUserPasswordV1() *UserPasswordV1 {
	return &UserPasswordV1{}
}

func NewUserPasswordV1(id string, password string) *UserPasswordV1 {
	return &UserPasswordV1{
		Id:        id,
		Password:  password,
		Locked:    false,
		FailCount: 0,
	}
}
