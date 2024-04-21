package controller

import (
	"encoding/json"
	"io"
	"net/http"
)

var MemberList []Member

type Member struct {
	Id    int64
	Name  string
	Birth string
}

func CreateMember(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var member Member
	err = json.Unmarshal(b, &member)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	dbService.open()
}
