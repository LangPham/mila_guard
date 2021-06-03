package mila_guard

import (
	"bytes"
	"strings"
)

type Guard struct {
	Role   string
	UserID string
}

func (gu Guard) ToString() string {
	var str bytes.Buffer
	str.WriteString("role")
	str.WriteString("=")
	str.WriteString(gu.Role)
	str.WriteString(",")
	str.WriteString("userID")
	str.WriteString("=")
	str.WriteString(gu.UserID)
	return str.String()
}

func ToGuard(str string) (gu Guard) {
	arrStr := strings.Split(str, ",")
	for _, elem := range arrStr {
		arrElem := strings.Split(elem, "=")
		switch arrElem[0] {
		case "role":
			gu.Role = arrElem[1]
		case "userID":
			gu.UserID = arrElem[1]
		}
	}
	return
}