package models

import (
	"log"
	"time"

	"github.com/Softinnov/bearded-basket/database"
)

type PDV struct {
	Id     int        `json:"pv_id"`
	Nom    string     `json:"pv_nom"`
	Expire *time.Time `json:"pv_abo_expire"`
}

func GetPDV(id int) (*PDV, error) {
	var pdv PDV

	err := database.DB.
		QueryRow("SELECT pv_id, pv_nom, pv_abo_expire FROM pdv WHERE pv_id=?", id).
		Scan(&pdv.Id, &pdv.Nom, &pdv.Expire)
	if err != nil {
		log.Println("GetPDV:", err)
		return nil, err
	}
	return &pdv, nil
}

func GetPDVs() ([]*PDV, error) {
	var pdvs []*PDV

	rows, err := database.DB.
		Query("SELECT pv_id, pv_nom, pv_abo_expire FROM pdv")
	if err != nil {
		log.Println("GetPDVs:", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var pdv PDV
		err := rows.Scan(&pdv.Id, &pdv.Nom, &pdv.Expire)
		if err != nil {
			log.Println("GetPDVs:", err)
			return nil, err
		}
		pdvs = append(pdvs, &pdv)
	}
	return pdvs, nil
}
