package pkg

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

func FormatDateTime(d time.Time) string {
	str := "02 January 2006, 15:04:05"
	format := d.Format(str)
	return format
}

func StringToUUID(str string) (uuid.UUID, error) {
	ID, err := uuid.Parse(str)
	if err != nil {
		fmt.Println("Error parsing UUID: ", err)
		return uuid.UUID{}, nil
	}
	return ID, nil
}

func Capitalize(s string) string {
	if len(s) != 0 {
		var res string
		for _, c := range s {
			if len(res) == 0 {
				res += strings.ToUpper(string(c))
				continue
			}
			res += strings.ToLower(string(c))
		}
		return res
	}
	return ""
}
