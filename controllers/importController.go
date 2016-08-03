package controllers

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"

	"gopkg.in/mgo.v2/bson"

	"github.com/kcwinner/visualKPIs/data"
	"github.com/kcwinner/visualKPIs/models"
)

const _24K = (1 << 20) * 240

//ImportData HTTP POST - /api/v1/import
func ImportData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := r.ParseMultipartForm(_24K)
	if err != nil {
		sendError(w, err, http.StatusInternalServerError, "Error importing data.")
		return
	}

	var infile multipart.File
	// var outfile *os.File

	for _, fheaders := range r.MultipartForm.File {
		for _, hdr := range fheaders {
			if infile, err = hdr.Open(); err != nil || infile == nil {
				sendError(w, err, http.StatusInternalServerError, "Error importing data.")
				return
			}

			defer infile.Close()
		}
	}

	if infile == nil {
		sendError(w, err, http.StatusInternalServerError, "Error importing data.")
		return
	}

	soldiers, err := parseCSV(infile)
	if err != nil {
		sendError(w, err, http.StatusInternalServerError, "Error importing data.")
		return
	}

	session := data.GetDb()
	defer session.Close()

	coll := session.DB("goDb").C("Soldiers")

	for _, soldier := range soldiers {

		var querySoldiers []models.Soldier
		query := bson.M{"firstname": soldier.FirstName, "lastname": soldier.LastName}
		err := coll.Find(query).All(&querySoldiers)
		if err != nil {
			sendError(w, err, http.StatusInternalServerError, "Error creating soldier.")
			return
		}

		if querySoldiers != nil {
			if len(querySoldiers) == 1 {
				updateSoldier := querySoldiers[0]

				updateSoldier.APFTPass = soldier.APFTPass
				updateSoldier.APFTScore = soldier.APFTScore
				updateSoldier.ETSDate = soldier.ETSDate
				updateSoldier.MOS = soldier.MOS
				updateSoldier.Rank = soldier.Rank
				updateSoldier.Section = soldier.Section

				err = coll.UpdateId(updateSoldier.ID, updateSoldier)
			} else {
				// for _, s := range querySoldiers {
				//
				// }
			}

			// coll.UpdateId(id, update)
			continue
		}

		err = coll.Insert(soldier)
		if err != nil {
			sendError(w, err, http.StatusInternalServerError, "Error creating soldier.")
			return
		}
	}

	respJSON, _ := json.Marshal("Good")

	w.WriteHeader(http.StatusOK)
	w.Write(respJSON)
}

func parseCSV(infile multipart.File) ([]models.Soldier, error) {
	var headers []string
	var soldiers []models.Soldier

	reader := csv.NewReader(infile)

	lineCount := 0

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		if lineCount == 0 {
			for index := range record {
				headers = append(headers, record[index])
			}

			lineCount++
			continue
		}

		soldier := models.Soldier{}

		for index := range record {
			val := record[index]
			switch headers[index] {
			case "FIRST":
				soldier.FirstName = val
			case "LAST":
				soldier.LastName = val
			case "RANK":
				soldier.Rank = val[:3]
			case "PMOS":
				soldier.MOS = val
			case "PARA":
				soldier.Section = val
			case "CivilianEmployment":
				soldier.CivilianEmployment = val
			case "ETS":
				if strings.Compare(val, "") == 0 {
					soldier.ETSDate = val
					break
				}
				formattedTime := fmt.Sprintf("%s-%s-%s", val[:4], val[4:6], val[6:])
				soldier.ETSDate = formattedTime
			case "NCOERDate":
				if strings.Compare(val, "") == 0 {
					soldier.NCOERDate = val
					break
				}

				formattedTime := fmt.Sprintf("%s-%s-%s", val[:4], val[4:6], val[6:8])
				soldier.NCOERDate = formattedTime
			case "APFTSCORE":
				soldier.APFTScore, _ = strconv.Atoi(val)
			case "APFTRSLT":
				soldier.APFTPass = strings.Compare(val, "P") == 0
			case "HTWT":
				soldier.HTWT = val
			}
		}

		soldier.ID = bson.NewObjectId()
		soldiers = append(soldiers, soldier)

		lineCount++
	}

	return soldiers, nil
}
