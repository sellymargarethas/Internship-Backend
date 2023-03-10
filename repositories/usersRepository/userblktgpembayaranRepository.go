package usersrepository

import (
	"Internship-Backend/constants"
	"Internship-Backend/models"
	"Internship-Backend/repositories"
	. "Internship-Backend/utils"
	"database/sql"
	"log"
	"strings"
	"time"
)

type blktgRepository struct {
	RepoDB repositories.Repository
}

func NewBlktgsRepository(repoDB repositories.Repository) blktgRepository {
	return blktgRepository{
		RepoDB: repoDB,
	}
}

func (ctx blktgRepository) IsBlkgPembayaranExistsByIndex(blktgPembayaran models.BlktgPembayaran) (models.BlktgPembayaran, bool, error) {
	var args []interface{}
	var result models.BlktgPembayaran

	query := `
	SELECT id, iduser
	FROM userblktgpembayaran
	WHERE`

	if blktgPembayaran.Id != constants.EMPTY_VALUE_INT {
		query += ` id = ? `
		args = append(args, blktgPembayaran.Id)
	}

	if blktgPembayaran.IdUser != constants.EMPTY_VALUE_INT {
		query += ` iduser = ? `
		args = append(args, blktgPembayaran.IdUser)
	}

	query += ` ORDER BY id ASC`

	query = ReplaceSQL(query, "?")

	rows, err := ctx.RepoDB.DB.Query(query, args...)

	if err != nil {
		return result, constants.FALSE_VALUE, err
	}
	defer rows.Close()

	data, err := BlkgPembayaranDto(rows)
	if err != nil {
		return result, constants.FALSE_VALUE, err
	}

	if len(data) == constants.EMPTY_VALUE_INT {
		return result, constants.FALSE_VALUE, err
	}
	return data[0], constants.TRUE_VALUE, err
}

func (ctx blktgRepository) GetBlktgPembayaran(id models.RequestGetBlktgPembayaranByIdUser) (response []models.ResponseBlktgPembayaran, err error) {
	query := `
	SELECT DISTINCT 
		kategoripembayaran.id,
		kategoripembayaran.uraian,
		CASE
			WHEN 
				kategoripembayaran.id
			NOT IN
				(select userblktgpembayaran.idkategoripembayaran FROM userblktgpembayaran WHERE userblktgpembayaran.iduser=?)
			THEN false
			ELSE true
		END AS status
	FROM kategoripembayaran
	ORDER BY kategoripembayaran.id ASC`
	query = ReplaceSQL(query, "?")

	rows, err := ctx.RepoDB.DB.Query(query, id.IdUser)
	if err != nil {
		log.Println("Error querying: GetBlktgPembayaran", err)
	}
	for rows.Next() {
		var list models.ResponseBlktgPembayaran
		rows.Scan(&list.Idkategoripembayaran, &list.KategoriPembayaran, &list.Status)
		response = append(response, list)
	}

	return
}

func (ctx blktgRepository) InsertBlktgPembayaran(request models.RequestUpdateBlktgPembayaran) (status bool, err error) {
	vals := []interface{}{}

	sqlStatement := `INSERT INTO userblktgpembayaran (iduser, created_at, idkategoripembayaran) VALUES`

	for _, row := range request.IdKategoriPembayaran {
		sqlStatement += "(?,?,?),"

		vals = append(vals, request.IdUser, time.Now(), row.IdKategoriPembayaran)
	}
	sqlStatement = strings.TrimRight(sqlStatement, ",")
	sqlStatement = ReplaceSQL(sqlStatement, "?")

	stmt, err := ctx.RepoDB.DB.Prepare(sqlStatement)
	if err != nil {
		log.Println("error preparing sql statement: InsertBlktgPembayaran", err)
		return constants.FALSE_VALUE, err
	}

	_, err = stmt.Exec(vals...)
	if err != nil {
		log.Println("error executing sql statement: InsertBlktgPembayaran", err)
		return constants.FALSE_VALUE, err
	}

	return constants.TRUE_VALUE, nil
}

func (ctx blktgRepository) DeleteBlktgPembayaran(id int, tx *sql.Tx) (result bool, err error) {
	query := `
		DELETE FROM userblktgpembayaran WHERE iduser=? RETURNING id
	`
	query = ReplaceSQL(query, "?")

	err = ctx.RepoDB.DB.QueryRow(query, id).Scan(&id)

	if err != nil {
		log.Println("Error querying DeleteBlktgPembayaran : ", err)
		return constants.FALSE_VALUE, err
	}

	return constants.TRUE_VALUE, err
}

func (ctx blktgRepository) UpdateBlktgPembayaran(request models.RequestUpdateBlktgPembayaran, tx *sql.Tx) (status bool, err error) {
	vals := []interface{}{}
	sqlStatement := `INSERT INTO userblktgpembayaran (iduser, created_at, idkategoripembayaran) VALUES`

	for _, row := range request.IdKategoriPembayaran {
		sqlStatement += "(?,?,?),"

		vals = append(vals, request.IdUser, time.Now(), row.IdKategoriPembayaran)
	}
	sqlStatement = strings.TrimRight(sqlStatement, ",")
	sqlStatement = ReplaceSQL(sqlStatement, "?")

	stmt, err := ctx.RepoDB.DB.Prepare(sqlStatement)
	if err != nil {
		log.Println("error preparing sql statement: UpdateBlktgPembayaran", err)
		return constants.FALSE_VALUE, err
	}

	_, err = stmt.Exec(vals...)
	if err != nil {
		log.Println("error executing sql statement: UpdateBlktgPembayaran", err)
		return constants.FALSE_VALUE, err
	}

	return constants.TRUE_VALUE, nil
}

func BlkgPembayaranDto(rows *sql.Rows) (result []models.BlktgPembayaran, err error) {
	for rows.Next() {
		var blktgPembayaran models.BlktgPembayaran
		err = rows.Scan(&blktgPembayaran.Id, &blktgPembayaran.IdUser)
		if err != nil {
			return
		}
		result = append(result, blktgPembayaran)
	}

	return
}
