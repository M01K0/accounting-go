package main

import (
	"fmt"

	dao "github.com/alexyslozada/accounting-go/dao/executedao"
	"log"
	"encoding/json"
)

func main() {
	p, err := dao.ProfileDAO.GetByID(1)
	checkErr(err)

	p.ObjectsProfile, err = dao.ObjectProfileDAO.GetByProfileID(p.ID)
	checkErr(err)

	myJson, err := json.Marshal(p)
	checkErr(err)

	fmt.Println(string(myJson))
}

func checkErr(e error) {
	if e != nil {
		log.Println(e)
	}
}
