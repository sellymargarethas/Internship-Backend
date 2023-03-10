package corporaterepository

import (
	"Internship-Backend/constants"
	"Internship-Backend/models"
	"Internship-Backend/repositories"
	. "Internship-Backend/utils"
	"log"
	"time"
)

type corporateCategoryRepository struct {
	RepoDB repositories.Repository
}

func NewCorporateCategoryRepository(repoDB repositories.Repository) corporateCategoryRepository {
	return corporateCategoryRepository{
		RepoDB: repoDB,
	}
}

func (ctx corporateCategoryRepository) IsCorporateCategoryExistbyIndex(category models.CorporateCategory) (status bool, err error) {
	var result []models.CorporateCategory
	query := `SELECT id, kode, uraian FROM corporatecategory WHERE deleted_at IS NULL`
	var args []interface{}
	if category.Kode != constants.EMPTY_VALUE {
		query += ` AND kode=?`
		args = append(args, category.Kode)
	}
	if category.IdCategory != constants.EMPTY_VALUE_INT {
		query += ` AND id=?`
		args = append(args, category.IdCategory)
	}
	query = ReplaceSQL(query, "?")
	rows, err := ctx.RepoDB.DB.Query(query, args...)
	if err != nil {
		return constants.FALSE_VALUE, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&category.IdCategory, &category.Kode, &category.Uraian)
		if err != nil {
			return constants.FALSE_VALUE, err
		}
		result = append(result, category)
	}
	if len(result) == constants.EMPTY_VALUE_INT {
		return constants.FALSE_VALUE, err
	}
	return constants.TRUE_VALUE, err
}

func (ctx corporateCategoryRepository) GetCorporateCategoryList() (list []models.CorporateCategory, err error) {
	query := `SELECT id, kode, uraian FROM corporatecategory WHERE deleted_at IS NULL`
	rows, err := ctx.RepoDB.DB.Query(query)

	for rows.Next() {
		var category models.CorporateCategory
		rows.Scan(&category.IdCategory, &category.Kode, &category.Uraian)
		list = append(list, category)
	}
	if err != nil {
		log.Println("ERROR GetCorporateCategory Repository", err)
	}
	return
}

func (ctx corporateCategoryRepository) InsertCorporateCategory(data models.RequestAddCorporateCategory) (id int, err error) {
	query := `INSERT INTO corporatecategory (kode, uraian, created_at) VALUES (?,?,?) returning id`
	query = ReplaceSQL(query, "?")

	err = ctx.RepoDB.DB.QueryRow(query, data.Kode, data.Uraian, time.Now()).Scan(&id)
	if err != nil {
		log.Println("ERROR InsertCorporateCategory Repository", err)
	}

	return
}

func (ctx corporateCategoryRepository) EditCorporateCategory(data models.RequestUpdateCorporateCategory) (id int, err error) {
	query := `UPDATE corporatecategory SET kode=?, uraian=?, updated_at=? WHERE id=? AND deleted_at IS NULL
	returning id`
	query = ReplaceSQL(query, "?")

	err = ctx.RepoDB.DB.QueryRow(query, data.Kode, data.Uraian, time.Now(), data.IdCategory).Scan(&id)
	if err != nil {
		log.Println("ERROR EditCorporateCategory Repository", err)
	}

	return
}

func (ctx corporateCategoryRepository) DeleteCorporateCategory(data models.DeleteCorporateCategory) (id int, err error) {
	query := `UPDATE corporatecategory SET deleted_at=? WHERE id=?
	returning id`
	query = ReplaceSQL(query, "?")

	err = ctx.RepoDB.DB.QueryRow(query, time.Now(), data.IdCategory).Scan(&id)
	if err != nil {
		log.Println("ERROR DeleteCorporateCategory Repository", err)
	}

	return
}
