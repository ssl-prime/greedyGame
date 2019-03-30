package util

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"greedyGame/bid_auction/model"
	"io/ioutil"
	"math/rand"
	"strconv"

	"github.com/asaskevich/govalidator"

	"github.com/gin-gonic/gin"
)

//temporary acquiring of slot
type kt map[string]bool

var slotAcquireMap = make(kt)

//ValidateAdRequest ...
func ValidateAdRequest(c *gin.Context) (string, error) {
	var (
		err           error
		reqPayload    []byte
		adRequestBody model.AdRequest
		adID          string
		isExists      bool
	)
	if reqPayload, err = ioutil.ReadAll(c.Request.Body); err == nil {
		if err = json.Unmarshal(reqPayload, &adRequestBody); err == nil {
			if _, err = govalidator.ValidateStruct(adRequestBody); err == nil {
				if isExists, err = isSlotExists(adRequestBody.AdPlacementID); err == nil && isExists {
					adID = adRequestBody.AdPlacementID
				} else {
					err = errors.New("slot existance error ")
				}
			}
		}
	}
	return adID, err
}

//AdRequest ...
func AdRequest(adSlotID string) (interface{}, error) {
	var (
		//isInProcess bool
		resp   interface{}
		crResp model.ObjectResp
		err    error
	)
	crResp.AdObjectID = RandStringRunes()
	crResp.Price = rand.Intn(100) + rand.Intn(1000)
	if db, err := ConnectDB(); err == nil {
		adObjectID := crResp.AdObjectID
		price := strconv.Itoa(crResp.Price)
		insertBid := `insert into bid (adplacement_id, ad_object_id, price) 
		values( "` + adSlotID + `" , "` + adObjectID + `" , ` + price + ` )`
		if _, err = db.Query(insertBid); err == nil {
			resp = crResp
		} else {
			err = errors.New("bid insertion err" + err.Error())
		}
	}
	return resp, err
}

func isSlotExists(adSlotID string) (bool, error) {
	db, err := ConnectDB()
	var ID int
	if err == nil {
		isExistQry := `select id from ad_slot where ad_slot_id = "` + adSlotID + `"`
		//fmt.Println(isExistQry)
		if results, err := db.Query(isExistQry); err == nil {
			for results.Next() {
				results.Scan(&ID)
			}
			if ID > 0 {
				return true, err
			}
		} else {
			return false, errors.New("err" + err.Error())
		}
	}
	return false, err
}

//ValidateAcceptance ...
// func ValidateAcceptance(c *gin.Context) error {
// 	var (
// 		err            error
// 		reqPayload     []byte
// 		adAcptanceBody model.SlotAcceptance
// 	)
// 	if reqPayload, err = ioutil.ReadAll(c.Request.Body); err == nil {
// 		if err = json.Unmarshal(reqPayload, &adAcptanceBody); err == nil {
// 			if _, err = govalidator.ValidateStruct(adAcptanceBody); err == nil {

// 			}
// 		}
// 	}
// 	return nil
// }

//Acceptance ...
func Acceptance() (interface{}, error) {
	return "accept", nil
}

//ValidateDeleteSlot ...
func ValidateDeleteSlot(c *gin.Context) error {
	return nil
}

//DeleteSlot ...
func DeleteSlot() (interface{}, error) {
	return "accept", nil
}

//ValidateAddSlot ...
func ValidateAddSlot(c *gin.Context) (string, error) {
	var (
		err               error
		reqPayload        []byte
		adSlotRequestBody model.AddAdSlot
		adSlotName        string
	)
	if reqPayload, err = ioutil.ReadAll(c.Request.Body); err == nil {
		if err = json.Unmarshal(reqPayload, &adSlotRequestBody); err == nil {
			if _, err = govalidator.ValidateStruct(adSlotRequestBody); err == nil {
				adSlotName = adSlotRequestBody.Name
			}
		}
	}
	return adSlotName, err
}

//AddSlot ...
func AddSlot(adSlotName string) (interface{}, error) {
	adSlotID := RandStringRunes()
	var (
		resp string
		err  error
		db   *sql.DB
	)
	insertAdSlot := `insert into ad_slot(ad_slot_id,name) values( "` + adSlotID + `" , "` + adSlotName + `" )`
	if db, err = ConnectDB(); err == nil {
		fmt.Println(insertAdSlot)
		if _, err = db.Query(insertAdSlot); err == nil {
			resp = "inserted succesfully"
		} else {
			resp = err.Error()
		}
	}
	return resp, err
}

//GetSlot ... this will give all available slot
func GetSlot() (interface{}, error) {
	return "available slot", nil
}

//GetBidByAdPlacementID ...
func GetBidByAdPlacementID(adPlacementID string) (int, error) {
	db, err := ConnectDB()
	var price int
	if err == nil {
		qry := `SELECT price from bid where adplacement_id = "` + adPlacementID + `"
		 order by price desc limit 1`
		if results, err := db.Query(qry); err == nil {
			for results.Next() {
				err = results.Scan(&price)
			}
		} else {
			fmt.Println(err.Error())
		}
	} else {
		fmt.Println(err.Error())
	}
	return price, err
}
