package coresettlementRepository

import (
	"Internship-Backend/constants"
	"Internship-Backend/models"
	"Internship-Backend/repositories"
	. "Internship-Backend/utils"
	"fmt"
	"log"
	"strings"
	"time"
)

type coresettlementRepository struct {
	RepoDB repositories.Repository
}

func NewCoreSettlementRepository(repoDB repositories.Repository) coresettlementRepository {
	return coresettlementRepository{
		RepoDB: repoDB,
	}
}

func (ctx coresettlementRepository) GetListCoreSettlementKeys(request models.RequestList) (coresettlementkeys []models.ResponseGetCoreSettlement, err error) {
	var args []interface{}
	query := `SELECT coresettlementkeys.id, coresettlementkeys.value, coresettlementkeys.name, coresettlementkeys.workerusername
	FROM coresettlementkeys`

	if request.Keyword != "" {
		query += ` WHERE (CAST(coresettlementkeys.id AS TEXT)) ILIKE '%' || ? || '%' 
		OR coresettlementkeys.value ILIKE '%' || ? || '%' 
		OR coresettlementkeys.name ILIKE '%' || ? || '%'
		OR coresettlementkeys.workerusername ILIKE '%' || ? || '%' `
		args = append(args, request.Keyword, request.Keyword, request.Keyword, request.Keyword)
	}

	orderby := fmt.Sprintf(request.OrderBy)
	order := fmt.Sprintf(request.Order)

	query += ` ORDER BY ` + orderby + ` ` + order

	query = ReplaceSQL(query, "?")

	rows, err := ctx.RepoDB.DB.Query(query, args...)

	if err != nil {
		log.Println("Error getting coresettlement keys: GetCoreSettlementKeys", err)
	}
	for rows.Next() {
		var list models.ResponseGetCoreSettlement

		rows.Scan(&list.Id, &list.Value, &list.Name, &list.WorkerUsername)

		coresettlementkeys = append(coresettlementkeys, list)
	}
	return
}

func (ctx coresettlementRepository) GetSingleCoreSettlementKeys(id int) (list models.ResponseGetCoreSettlement, err error) {
	query := `SELECT id, value, name, workerusername
	FROM coresettlementkeys WHERE id=$1`

	err = ctx.RepoDB.DB.QueryRow(query, id).Scan(&list.Id, &list.Value, &list.Name, &list.WorkerUsername)

	if err != nil {
		log.Println("Error querying coresettlementkeys: GetSingleCoreSettlementKeys", err)
	}

	return
}

func (ctx coresettlementRepository) InsertCoreSettlementKey(coresettlement models.RequestAddCoreSettlement) (id int, err error) {

	query := `INSERT INTO coresettlementkeys (value, name, workerusername, created_at) 
	VALUES (?, ?, ?, ?) returning id`

	query = ReplaceSQL(query, "?")

	err = ctx.RepoDB.DB.QueryRow(query, coresettlement.Value, coresettlement.Nama, coresettlement.WorkerUsername, time.Now()).Scan(&id)
	if err != nil {
		log.Println("Error InsertCoreSettlementKey: repository")
	}
	return
}

func (ctx coresettlementRepository) UpdateCoreSettlementKey(coresettlement models.RequestUpdateCoreSettlement) (id int, err error) {

	query := `UPDATE coresettlementkeys SET value=?, name=?, workerusername=?, updated_at=? where id=? returning id`
	query = ReplaceSQL(query, "?")

	err = ctx.RepoDB.DB.QueryRow(query, coresettlement.Value, coresettlement.Nama, coresettlement.WorkerUsername, time.Now(), coresettlement.Id).Scan(&id)

	if err != nil {
		log.Println("Error UpdateCoreSettlementKey: repository")
	}

	return
}

func (ctx coresettlementRepository) DeleteCoreSettlementKey(idcoresettlement models.RequestDeleteCoreSettlement) (bool, error) {
	query := `DELETE FROM coresettlementkeys WHERE id=$1`
	rows, err := ctx.RepoDB.DB.Query(query, idcoresettlement.Id)

	if err != nil {
		log.Println("Error DeleteCoreSettlementKey: repository", rows, err)
		return constants.FALSE_VALUE, err
	}
	return constants.TRUE_VALUE, nil
}

func (ctx coresettlementRepository) IsCoresettlementkeysExistByIndex(csk models.CoreSettlementKeys) (bool, error) {
	var count int
	var args []interface{}
	query := `SELECT COUNT(*) FROM coresettlementkeys WHERE `
	if csk.Id != constants.EMPTY_VALUE_INT {
		query += `id=?`
		args = append(args, csk.Id)
	}

	query = ReplaceSQL(query, "?")
	err := ctx.RepoDB.DB.QueryRow(query, args...).Scan(&count)
	if err != nil {
		log.Println("Error IsCoresettlementkeysExist: repository", err)

	}
	if count != constants.EMPTY_VALUE_INT {
		return constants.TRUE_VALUE, nil
	}
	return constants.FALSE_VALUE, nil
}

func (ctx coresettlementRepository) GetListCSKId() (result []models.SettlementDestination, err error) {
	query := `SELECT name FROM coresettlementkeys`
	rows, err := ctx.RepoDB.DB.Query(query)
	for rows.Next() {
		var data models.SettlementDestination
		rows.Scan(&data.IdCoreSettlement)
		result = append(result, data)
	}
	return
}

func (ctx coresettlementRepository) GetStrAggCSKName(request models.RequestSummarySettled) (result string, err error) {
	var args []interface{}
	query := `SELECT STRING_AGG(name, ',') from coresettlementkeys where name=ANY(array[`
	for _, row := range request.SettlementDestination {
		query += "?,"
		args = append(args, row.IdCoreSettlement)
	}
	query = strings.TrimRight(query, ",")
	query += `]::varchar[])`
	query = ReplaceSQL(query, "?")
	err = ctx.RepoDB.DB.QueryRow(query, args...).Scan(&result)
	return
}
