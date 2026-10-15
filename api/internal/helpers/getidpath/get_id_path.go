package getidpath

import (
	"net/http"
	"strconv"
)

func GetIdPath(r *http.Request) (id int, err error) {
	id, err = strconv.Atoi(r.PathValue("id"))

	if err != nil {
		return id, err
	}

	return id, err
}
