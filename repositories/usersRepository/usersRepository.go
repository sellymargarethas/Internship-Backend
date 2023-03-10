package usersrepository

import (
	"Internship-Backend/constants"
	"Internship-Backend/models"
	"Internship-Backend/repositories"
	. "Internship-Backend/utils"
	"database/sql"
	"fmt"
	"log"
	"time"
)

type usersRepository struct {
	RepoDB repositories.Repository
}

func NewUsersRepository(repoDB repositories.Repository) usersRepository {
	return usersRepository{
		RepoDB: repoDB,
	}
}

func (ctx usersRepository) InsertUser(user models.RequestAddUser) (id int, err error) {
	query := `INSERT INTO users (nama, username, email, password, jenis, idcorporate, created_at, enccredential, idrole)
	VALUES (?,?,?,?,?,?,?,?,?) returning id`
	query = ReplaceSQL(query, "?")

	err = ctx.RepoDB.DB.QueryRow(query, user.Nama, user.Username, user.Email, user.Password, user.Jenis, user.IdCorporate, time.Now(), user.EncCredentials, user.IdRole).Scan(&id)
	if err != nil {
		log.Println("Error querying InsertUser", err)
	}

	return
}

func (ctx usersRepository) GetAllUser(request models.RequestList) (users []models.ResponseUser, err error) {
	var args []interface{}
	query := `SELECT users.id, users.nama, users.username, jenisuser.uraian, corporate.uraian, users.email, users.enccredential, users.idrole, roles.rolename, users.idcorporate
	FROM users
	LEFT JOIN roles ON users.idrole=roles.id
	INNER JOIN corporate on corporate.id=users.idcorporate
	INNER JOIN jenisuser on users.jenis=jenisuser.id
	WHERE users.deleted_at IS NULL 
	AND corporate.hirarki_id LIKE ?||'%' 
	`
	args = append(args, request.HirarkiId)
	if request.Keyword != "" {
		query += ` AND (CAST( users.id AS TEXT) ILIKE '%' || ? || '%'
		OR users.nama ILIKE '%' || ? || '%'
		OR users.username ILIKE '%' || ? || '%'
		OR jenisuser.uraian ILIKE '%' || ? || '%'
		OR corporate.uraian ILIKE '%' || ? || '%'
		OR users.email ILIKE '%' || ? || '%'
		OR users.enccredential ILIKE '%' || ? || '%')
		`
		args = append(args, request.Keyword, request.Keyword, request.Keyword, request.Keyword, request.Keyword, request.Keyword, request.Keyword)
	}

	orderby := fmt.Sprintf(request.OrderBy)
	order := fmt.Sprintf(request.Order)

	query += ` ORDER BY ` + orderby + ` ` + order

	query = ReplaceSQL(query, "?")

	rows, err := ctx.RepoDB.DB.Query(query, args...)
	if err != nil {
		log.Println("Error querying GetUser", err)
	}
	for rows.Next() {
		var user models.ResponseUser
		rows.Scan(&user.Id, &user.Nama, &user.Username, &user.JenisUser, &user.NamaCorporate, &user.Email, &user.EncCredentials, &user.IdRole, &user.Role, &user.IdCorporate)
		users = append(users, user)
	}

	return
}

func (ctx usersRepository) GetUserByIndex(user models.Users) (response []models.ResponseUser, err error) {
	var args []interface{}
	var users models.ResponseUser
	query := `SELECT users.id, users.nama, users.username, users.jenis, users.idcorporate, users.email, users.enccredential, users.idrole, roles.rolename
	FROM users
	LEFT JOIN roles on users.idrole=roles.id
	WHERE users.deleted_at IS NULL
	`
	if user.Id != constants.EMPTY_VALUE_INT {
		query += `AND users.id=?`
		args = append(args, user.Id)
	}
	if user.Username != constants.EMPTY_VALUE {
		query += `AND users.username=?`
		args = append(args, user.Username)
	}
	if user.Email != constants.EMPTY_VALUE {
		query += `AND users.email=?`
		args = append(args, user.Email)
	}
	if user.IdCorporate != constants.EMPTY_VALUE_INT {
		query += `AND users.idcorporate=?`
		args = append(args, user.IdCorporate)
	}

	query = ReplaceSQL(query, "?")

	rows, err := ctx.RepoDB.DB.Query(query, args...)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("Error querying GetOneUser", err)
		}
	}
	for rows.Next() {

		err = rows.Scan(&users.Id, &users.Nama, &users.Username, &users.IdJenis, &users.IdCorporate, &users.Email, &users.EncCredentials, &users.IdRole, &users.Role)

		if err != nil {
			log.Println("Error Scanning Users: GetSingleUser", err)
		}
		response = append(response, users)
	}
	return
}

func (ctx usersRepository) UpdateUser(user models.RequestUpdateUser) (id int, err error) {
	query := `UPDATE users 
	SET 
	nama=?, username=?,  
	jenis=?, idcorporate=?, email=?, 
	enccredential=?, updated_at=?, idrole=?
	WHERE id=? AND deleted_at IS NULL 
	returning id`
	query = ReplaceSQL(query, "?")

	err = ctx.RepoDB.DB.QueryRow(query, user.Nama, user.Username, user.Jenis, user.IdCorporate, user.Email, user.EncCredentials, time.Now(), user.IdRole, user.Id).Scan(&id)
	if err != nil {
		log.Println("Error querying  UpdateUser", err)
	}
	return

}
func (ctx usersRepository) UpdateUserPassword(request models.RequestUpdatePassword) (id int, err error) {
	query := `UPDATE users SET password=? WHERE id=? returning id`
	query = ReplaceSQL(query, "?")

	err = ctx.RepoDB.DB.QueryRow(query, request.Password, request.Id).Scan(&id)
	if err != nil {
		log.Println("Error querying  UpdateUserPassword", err)
	}
	return
}

func (ctx usersRepository) UpdateEnc(id int, enc string) (id2 int, err error) {
	query := `UPDATE users SET enccredential=? WHERE id=? returning id`
	query = ReplaceSQL(query, "?")

	err = ctx.RepoDB.DB.QueryRow(query, enc, id).Scan(&id2)
	if err != nil {
		log.Println("Error querying  UpdateUserEnc", err)
	}
	return
}

func (ctx usersRepository) DeleteUser(user models.RequestDeleteUser) (id int, err error) {
	query := `UPDATE users SET deleted_at=$1 WHERE id=$2 returning id`

	err = ctx.RepoDB.DB.QueryRow(query, time.Now(), user.Id).Scan(&id)
	if err != nil {
		log.Println("Error querying DeleteUser", err)
	}
	return
}

func (ctx usersRepository) GetJenisUser() (jenis []models.ResponseJenisUser, err error) {
	query := `SELECT id, uraian FROM jenisuser`
	rows, err := ctx.RepoDB.DB.Query(query)
	if err != nil {
		log.Println("Error querying GetJenisUser", err)
	}
	for rows.Next() {
		var jenisuser models.ResponseJenisUser
		rows.Scan(&jenisuser.Id, &jenisuser.JenisUser)
		jenis = append(jenis, jenisuser)
	}

	return

}
