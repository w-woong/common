package dto_test

import (
	"time"

	"github.com/w-woong/common/dto"
)

var (
	bd      = time.Date(2002, 1, 2, 0, 0, 0, 0, time.Local)
	userDto = dto.User{
		ID:      "22bcbf79-ca5f-42dc-8ca0-29441209a36a",
		LoginID: "wonk",
		CredentialPassword: dto.CredentialPassword{
			ID:     "333cbf79-ca5f-42dc-8ca0-29441209a36a",
			UserID: "22bcbf79-ca5f-42dc-8ca0-29441209a36a",
			Value:  "asdfasdfasdf",
		},
		Personal: dto.Personal{
			ID:          "433cbf79-ca5f-42dc-8ca0-29441209a36a",
			UserID:      "22bcbf79-ca5f-42dc-8ca0-29441209a36a",
			FirstName:   "wonk",
			LastName:    "sun",
			BirthYear:   2002,
			BirthMonth:  1,
			BirthDay:    2,
			BirthDate:   &bd,
			Gender:      "M",
			Nationality: "KOR",
		},
	}
)
