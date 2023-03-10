package models

type RoleDetails struct {
	Menu string `json:"menu"`
	Task string `json:"task"`
}
type RequestRoleDetails struct {
	IdMenu string `json:"idMenu"`
	IdTask string `json:"idTask"`
}

type RequestAddUsersRole struct {
	HirarkiId string               `json:"hirarkiId" validate:"required"`
	Role      string               `json:"role" validate:"required"`
	Privilege []RequestRoleDetails `json:"privilege" validate:"required"`
}

type ResponseUsersRole struct {
	Privilege []RoleDetails `json:"privilege"`
}

type RequestListRole struct {
	HirarkiId string  `json:"hirarkiId"`
	Order     string  `json:"order" validate:"required"` //asc atau desc
	OrderBy   string  `json:"orderBy"`                   //order by per field
	Limit     int     `json:"limit" validate:"required"` //jumlah per page
	Page      int     `json:"page" validate:"required"`  //Halaman ke brp yg mau di load
	Keyword   *string `json:"keyword"`                   //untuk search
}
type RequestSingleRole struct {
	IdRole int `json:"idRole" validate:"required"` //
}
type ResponseListRole struct {
	Id            int    `json:"id"`
	HirarkiId     string `json:"hirarkiId"`
	NamaCorporate string `json:"namaCorporate"`
	Role          string `json:"roles"`
}
type ResponseRoleDetails struct {
	Id            int           `json:"id"`
	HirarkiId     string        `json:"hirarkiId"`
	NamaCorporate string        `json:"namaCorporate"`
	Role          string        `json:"role"`
	Privilege     []RoleDetails `json:"privilege"`
}

type RequestUpdateRole struct {
	Id        int    `json:"id" validate:"required"`
	HirarkiId string `json:"hirarkiId" validate:"required"`
	// NamaCorporate string               `json:"namaCorporate"`
	Role      string               `json:"role" validate:"required"`
	Privilege []RequestRoleDetails `json:"privilege" validate:"required"`
}

type RoleMenu struct {
	Id   int    `json:"id"`
	Menu string `json:"menu"`
}
type RoleTask struct {
	Id   int    `json:"id"`
	Task string `json:"task"`
}
