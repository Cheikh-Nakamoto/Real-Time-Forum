package pkg

import (
	"fmt"
	"os"
)

func FileToString(filename string) string {
	file, err := os.ReadFile("./src/main/resources/private/database/tables/" + filename + ".sql")
	if err != nil {
		fmt.Println("Could not read file:", err)
		return ""
	}
	return string(file)
}
