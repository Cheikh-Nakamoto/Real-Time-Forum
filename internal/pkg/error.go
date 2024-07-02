package pkg

import (
	"net/http"
	"strconv"
)

func Error(code int) map[string]string {
	tab := make(map[string]string)
	switch code {
	case http.StatusBadRequest:
		tab["code"] = strconv.Itoa(code)
		tab["msg"] = "Bad request"
	case http.StatusNoContent:
		tab["code"] = strconv.Itoa(code)
		tab["msg"] = "No Content"
	case http.StatusCreated:
		tab["code"] = strconv.Itoa(code)
		tab["msg"] = "Created"
	case http.StatusUnauthorized:
		tab["code"] = strconv.Itoa(code)
		tab["msg"] = "Unauthorized"
	case http.StatusForbidden:
		tab["code"] = strconv.Itoa(code)
		tab["msg"] = "Forbidden"
	case http.StatusNotFound:
		tab["code"] = strconv.Itoa(code)
		tab["msg"] = "Not Found"
	case http.StatusMethodNotAllowed:
		tab["code"] = strconv.Itoa(code)
		tab["msg"] = "Method Not Allowed"
	case http.StatusInternalServerError:
		tab["code"] = strconv.Itoa(code)
		tab["msg"] = "Internal Server Error"
	}
	return tab
}
