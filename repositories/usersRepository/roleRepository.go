package usersrepository

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

type roleRepository struct {
	RepoDB repositories.Repository
}

func NewRoleRepository(repoDB repositories.Repository) roleRepository {
	return roleRepository{
		RepoDB: repoDB,
	}
}

func (ctx roleRepository) GetAllRoleMenu() (response []models.RoleMenu, err error) {
	query := `SELECT id, uraian FROM rolemenu where deleted_at is null`
	rows, err := ctx.RepoDB.DB.Query(query)
	for rows.Next() {
		var data models.RoleMenu
		rows.Scan(&data.Id, &data.Menu)
		response = append(response, data)
	}
	return
}
func (ctx roleRepository) GetAllRoleTask() (response []models.RoleTask, err error) {
	query := `SELECT id, uraian FROM roletask where deleted_at is null`
	rows, err := ctx.RepoDB.DB.Query(query)
	for rows.Next() {
		var data models.RoleTask
		rows.Scan(&data.Id, &data.Task)
		response = append(response, data)
	}
	// fmt.Println(response)
	return
}

func (ctx roleRepository) AddUserRole(request models.RequestAddUsersRole) (err error) {
	var args []interface{}

	query := `INSERT INTO roles (idcorporate, rolename, menubar, task, created_at) VALUES((SELECT id FROM corporate WHERE hirarki_id =?),?,ARRAY[`
	args = append(args, request.HirarkiId, request.Role)
	for _, row := range request.Privilege {
		query += "(SELECT id from rolemenu where uraian=?)::int,"

		args = append(args, row.IdMenu)
	}
	query = strings.TrimRight(query, ",")
	query += `]::int[], ARRAY[`
	for _, row := range request.Privilege {
		query += "(SELECT id from roletask where uraian=?)::int,"

		args = append(args, row.IdTask)
	}
	query = strings.TrimRight(query, ",")
	query += `]::int[],?)`
	args = append(args, time.Now())
	query = ReplaceSQL(query, "?")
	// fmt.Println(query)
	stmt, _ := ctx.RepoDB.DB.Prepare(query)

	_, err = stmt.Exec(args...)
	return
}

func (ctx roleRepository) GetListRole(request models.RequestListRole) (response []models.ResponseListRole, err error) {
	var args []interface{}
	query := `SELECT roles.id, corporate.hirarki_id, roles.rolename, corporate.uraian from roles 
	INNER JOIN corporate on roles.idcorporate=corporate.id
	WHERE corporate.hirarki_id LIKE ?||'%' AND roles.deleted_at IS NULL`
	args = append(args, request.HirarkiId)

	if request.Keyword != nil {
		query += ` 
		AND (CAST(roles.id AS TEXT) ILIKE '%'||? || '%' OR
		CAST(roles.idcorporate AS TEXT) ILIKE '%' || ? || '%' OR
		roles.rolename ILIKE '%' || ? || '%' OR
		corporate.uraian ILIKE '%' || ? || '%')`
		args = append(args, request.Keyword, request.Keyword, request.Keyword, request.Keyword)
	}
	orderby := fmt.Sprintf(request.OrderBy)
	order := fmt.Sprintf(request.Order)

	query += ` ORDER BY ` + orderby + ` ` + order

	query = ReplaceSQL(query, "?")
	// fmt.Println(query)
	rows, err := ctx.RepoDB.DB.Query(query, args...)
	if err != nil {
		log.Println("Error querying GetRole", err)
	}
	for rows.Next() {
		var data models.ResponseListRole
		rows.Scan(&data.Id, &data.HirarkiId, &data.Role, &data.NamaCorporate)
		response = append(response, data)
	}
	return
}

func (ctx roleRepository) GetSingleRoleDetails(request models.RequestSingleRole) (response []models.RoleDetails, err error) {
	var args []interface{}
	query := `SELECT rolemenu.uraian as menu, roletask.uraian	as task
	from 
	(SELECT unnest(menubar)as menu, unnest(task)as task from roles WHERE id=? AND deleted_at IS NULL) as roles
	INNER JOIN rolemenu on roles.menu = rolemenu.id
	INNER JOIN roletask on roles.task = roletask.id`
	args = append(args, request.IdRole)
	query = ReplaceSQL(query, "?")
	rows, err := ctx.RepoDB.DB.Query(query, args...)
	if err != nil {
		log.Println("Error querying GetUserRoleDetails", err)
	}
	for rows.Next() {
		var data models.RoleDetails
		rows.Scan(&data.Menu, &data.Task)
		response = append(response, data)
	}
	return
}
func (ctx roleRepository) GetSingleRole(request models.RequestSingleRole) (response models.ResponseRoleDetails, err error) {
	var args []interface{}
	query := `SELECT roles.id, corporate.hirarki_id, corporate.uraian, roles.rolename FROM roles 
	INNER JOIN corporate ON roles.idcorporate=corporate.id
	WHERE roles.id=? AND roles.deleted_at is null`
	args = append(args, request.IdRole)
	query = ReplaceSQL(query, "?")
	err = ctx.RepoDB.DB.QueryRow(query, args...).Scan(&response.Id, &response.HirarkiId, &response.NamaCorporate, &response.Role)
	if err != nil {
		log.Println("Error querying GetUserRole", err)
	}
	return
}

// func(ctx usersRepository) UpdateRole(request )
func (ctx roleRepository) UpdateRole(request models.RequestUpdateRole) (status bool, err error) {
	var args []interface{}
	query := `UPDATE roles SET idcorporate=(SELECT id FROM corporate WHERE hirarki_id =?), rolename=?, menubar=ARRAY[`
	args = append(args, request.HirarkiId, request.Role)
	for _, row := range request.Privilege {
		query += "(SELECT id from rolemenu where uraian=?)::int,"

		args = append(args, row.IdMenu)
	}
	query = strings.TrimRight(query, ",")
	query += `]::int[], task=ARRAY[`
	for _, row := range request.Privilege {
		query += "(SELECT id from roletask where uraian=?)::int,"

		args = append(args, row.IdTask)
	}
	query = strings.TrimRight(query, ",")
	query += `]::int[], updated_at=? WHERE id=? AND deleted_at is null`
	args = append(args, time.Now(), request.Id)

	query = ReplaceSQL(query, "?")
	// fmt.Println(query)
	_, err = ctx.RepoDB.DB.Query(query, args...)
	if err != nil {
		log.Println("Error querying GetUserRole", err)
		return constants.FALSE_VALUE, err
	}
	return constants.TRUE_VALUE, err
}
func (ctx roleRepository) DeleteRole(request models.RequestSingleRole) (status bool, err error) {
	query := `UPDATE roles SET deleted_at = $1 WHERE id=$2`
	_, err = ctx.RepoDB.DB.Query(query, time.Now(), request.IdRole)
	if err != nil {
		log.Println("Failed to delete roles", err)
		return constants.FALSE_VALUE, err
	}
	return constants.TRUE_VALUE, err
}

func (ctx roleRepository) GetRoleIdByIdUser(id int) (roles int, err error) {
	query := `SELECT users.idrole from users
	INNER JOIN roles ON users.idrole=roles.id
	WHERE users.id=$1 and  roles.deleted_at is NULL`
	err = ctx.RepoDB.DB.QueryRow(query, id).Scan(&roles)
	return
}
