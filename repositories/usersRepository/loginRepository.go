package usersrepository

import (
	"Internship-Backend/models"
	"Internship-Backend/repositories"
	. "Internship-Backend/utils"
	"log"
)

type loginRepository struct {
	RepoDB repositories.Repository
}

func NewloginRepository(repoDB repositories.Repository) loginRepository {
	return loginRepository{
		RepoDB: repoDB,
	}
}

func (ctx loginRepository) CheckLogin(data models.Login) (hash string, jenis int, err error) {
	query := `SELECT password, jenis FROM users WHERE username=?`

	query = ReplaceSQL(query, "?")
	err = ctx.RepoDB.DB.QueryRow(query, data.Username).Scan(&hash, &jenis)

	if err != nil {
		log.Println("Error querying: CheckLogin: ", err)
	}
	return hash, jenis, err

}

func (ctx loginRepository) LoginReturn(data models.Login) (user models.LoginResponse, err error) {
	query := `SELECT users.id, users.nama, corporate.hirarki_id FROM users
	INNER JOIN corporate on users.idcorporate=corporate.id
	WHERE users.username=$1`
	err = ctx.RepoDB.DB.QueryRow(query, data.Username).Scan(&user.Id, &user.Nama, &user.HirarkiId)
	if err != nil {
		log.Println("Error querying: LoginReturn: ", err)
	}
	return
}
