package str

import "github.com/google/uuid"

var uid string

func SetUid() {
	uid = uuid.New().String()
}

func GetNewUid() string {
	return uid
}
