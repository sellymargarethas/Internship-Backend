package deviceRepository

import (
	"Internship-Backend/constants"
	"Internship-Backend/models"
	"Internship-Backend/repositories"
	. "Internship-Backend/utils"
	"fmt"
	"log"
	"time"
)

type deviceRepository struct {
	RepoDB repositories.Repository
}

func NewDeviceRepository(repoDB repositories.Repository) deviceRepository {
	return deviceRepository{
		RepoDB: repoDB,
	}
}

func (ctx deviceRepository) GetListDevice(request models.RequestList) (devices []models.ResponseDevice, err error) {
	var args []interface{}

	query := `SELECT 
	device.id, device.device_id, device.idcorporate,
	corporate.uraian, device.merchantkey, device.dkimid,
	device.dkitid, device.jenisdevice, device.mid, coresettlementkeys.name,
	device.tid, device.tokenfcm, 
	device.acquiringmid, device.acquiringtid 
   FROM device
   INNER JOIN corporate on corporate.id=device.idcorporate
   INNER JOIN coresettlementkeys on device.merchantkey=coresettlementkeys.value
   WHERE device.deleted_at IS NULL
   AND corporate.deleted_at IS NULL
   AND corporate.hirarki_id LIKE ?||'%'
   `
	args = append(args, request.HirarkiId)

	if request.Keyword != "" {
		query += ` AND (
	CAST(device.id AS TEXT) ILIKE '%' || ? || '%' OR
	device.device_id ILIKE '%' || ? || '%' OR
	corporate.uraian ILIKE '%' || ? || '%' OR
	device.merchantkey ILIKE '%' || ? || '%' OR
	device.dkimid ILIKE '%' || ? || '%' OR
	device.dkitid ILIKE '%' || ? || '%' OR
	device.jenisdevice ILIKE '%' || ? || '%' OR
	device.mid ILIKE '%' || ? || '%' OR
	coresettlementkeys.name ILIKE '%' || ? || '%')`
		args = append(args, request.Keyword, request.Keyword, request.Keyword, request.Keyword, request.Keyword, request.Keyword, request.Keyword, request.Keyword, request.Keyword)
	}

	orderby := fmt.Sprintf(request.OrderBy)
	order := fmt.Sprintf(request.Order)

	query += ` ORDER BY ` + orderby + ` ` + order

	query = ReplaceSQL(query, "?")
	// fmt.Println(query)
	rows, err := ctx.RepoDB.DB.Query(query, args...)
	for rows.Next() {
		var data models.ResponseDevice
		rows.Scan(&data.Id, &data.DeviceId, &data.IdCorporate, &data.NamaCorporate, &data.MerchantKey, &data.DkiMid, &data.DkiTid, &data.JenisDevice, &data.Mid, &data.CSKName, &data.Tid, &data.Tokenfcm, &data.AquiringMid, &data.AquiringTid)
		devices = append(devices, data)
	}
	if err != nil {
		log.Println("Error querying GetAllDevice: ", err)
	}

	return

}

func (ctx deviceRepository) InsertDevice(device models.RequestAddDevice) (id int, err error) {
	query := `INSERT INTO device 
	(device_id, idcorporate, merchantkey,
	 dkimid, dkitid, jenisdevice, 
	 created_at, tid, mid,
	 acquiringmid, acquiringtid, tokenfcm)
	VALUES 
	(?, ?, (SELECT value from coresettlementkeys WHERE id=?),
 	?, ?, ?,
  	?, ?, ?, 
	?, ?, ?) returning id`
	query = ReplaceSQL(query, "?")

	err = ctx.RepoDB.DB.QueryRow(query, device.DeviceId, device.IdCorporate, device.IdCorporateMerchantkey, device.DkiMid, device.DkiTid, device.JenisDevice, time.Now(), device.Tid, device.Mid, device.AquiringMid, device.AquiringTid, device.Tokenfcm).Scan(&id)
	if err != nil {
		log.Println("Error querying InsertDevice: ", err)
	}
	return

}

func (ctx deviceRepository) EditDevice(device models.RequestUpdateDevice) (id int, err error) {
	query := `UPDATE device SET device_id=?, idcorporate=?, updated_at=?, 
	merchantkey=(select value from coresettlementkeys where id=?), dkimid=?, dkitid=?, jenisdevice=?, mid=?, tid=?,
	acquiringmid=?, acquiringtid=?, tokenfcm=? 
	WHERE id=? returning id`
	query = ReplaceSQL(query, "?")

	err = ctx.RepoDB.DB.QueryRow(query, device.DeviceId, device.IdCorporate, time.Now(), device.IdCorporateMerchantkey, device.DkiMid, device.DkiTid, device.JenisDevice, device.Mid, device.Tid, device.AquiringMid, device.AquiringTid, device.Tokenfcm, device.Id).Scan(&id)
	if err != nil {
		log.Println("Error querying EditDevice: ", err)
	}
	return
}

func (ctx deviceRepository) DeleteDevice(device models.RequestDeleteDevice) (id int, err error) {
	query := `UPDATE device SET deleted_at=$1 WHERE id=$2 returning id`

	err = ctx.RepoDB.DB.QueryRow(query, time.Now(), device.Id).Scan(&id)
	if err != nil {
		log.Println("Error querying DeleteDevice: ", err)
	}
	return
}

func (ctx deviceRepository) GetSingleDevice(device models.Device) (data models.ResponseSingleDevice, err error) {
	var args []interface{}
	query := `SELECT device.id, device.device_id, device.idcorporate,
	corporate.uraian, (SELECT coresettlementkeys.id FROM coresettlementkeys WHERE coresettlementkeys.value=device.merchantkey), device.dkimid,
	 device.dkitid, device.jenisdevice, device.mid,
	 corporate.uraian, device.tid, device.tokenfcm, 
	 device.acquiringmid, device.acquiringtid
   FROM device
   INNER JOIN corporate on corporate.id=device.idcorporate
   WHERE device.deleted_at IS NULL 
   AND corporate.deleted_at IS NULL
	`
	if device.Id != constants.EMPTY_VALUE_INT {
		query += `AND device.id=?`
		args = append(args, device.Id)
	}
	if device.DeviceId != constants.EMPTY_VALUE {
		query += `AND device.device_id=?`
		args = append(args, device.DeviceId)
	}
	query = ReplaceSQL(query, "?")

	err = ctx.RepoDB.DB.QueryRow(query, args...).Scan(&data.Id, &data.DeviceId, &data.IdCorporate, &data.NamaCorporate, &data.MerchantKey, &data.DkiMid, &data.DkiTid, &data.JenisDevice, &data.Mid, &data.CSKName, &data.Tid, &data.Tokenfcm, &data.AquiringMid, &data.AquiringTid)
	if err != nil {
		log.Println("Error querying GetOneDevice: ", err)
	}
	return
}
