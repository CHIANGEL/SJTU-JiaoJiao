package main

import (
	"context"
	db "jiaojiao/database"
	buyinfo "jiaojiao/srv/buyinfo/proto"
	"jiaojiao/srv/content/mock"
	content "jiaojiao/srv/content/proto"
	"jiaojiao/utils"
	"time"

	"github.com/micro/go-micro/client"

	"github.com/jinzhu/gorm"
)

type srv struct{}

/**
 * @api {rpc} /rpc BuyInfo.Query
 * @apiVersion 1.0.0
 * @apiGroup Service
 * @apiName BuyInfo.Query
 * @apiDescription Query buy info
 *
 * @apiParam {int32} buyInfoID buyInfo id.
 * @apiSuccess {int32} buyInfoID buyInfoID
 * @apiSuccess {int32} status 1 for buying <br> 2 for reserved <br> 3 for done <br> 4 for expired
 * @apiSuccess {int64} releaseTime buyInfo release time
 * @apiSuccess {int64} validTime buyInfo validate time
 * @apiSuccess {string} goodName good name
 * @apiSuccess {double} price good price
 * @apiSuccess {string} description good description
 * @apiSuccess {string} contentID multimedia data
 * @apiSuccess {int32} userID userID
 * @apiUse DBServerDown
 */
func (a *srv) Query(ctx context.Context, req *buyinfo.BuyInfoQueryRequest, rsp *buyinfo.BuyInfoMsg) error {
	if req.BuyInfoID == 0 {
		return nil
	}
	info := db.BuyInfo{
		ID: req.BuyInfoID,
	}
	err := db.Ormer.First(&info).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil
	} else if utils.LogContinue(err, utils.Warning) {
		return err
	}
	good := db.Good{
		ID: info.GoodID,
	}
	err = db.Ormer.First(&good).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil
	} else if utils.LogContinue(err, utils.Warning) {
		return err
	}

	rsp.BuyInfoID = info.ID
	rsp.Status = info.Status
	rsp.ReleaseTime = info.ReleaseTime.Unix()
	rsp.ValidTime = info.ValidTime.Unix()
	rsp.GoodName = good.GoodName
	rsp.Price = good.Price
	rsp.Description = good.Description
	rsp.ContentID = good.ContentID
	rsp.UserID = info.UserID
	return nil
}

/**
 * @api {rpc} /rpc BuyInfo.Create
 * @apiVersion 1.0.0
 * @apiGroup Service
 * @apiName BuyInfo.Create
 * @apiDescription Create buy info
 *
 * @apiParam {int32} userID user id
 * @apiParam {int64} validTime valid timestamp
 * @apiParam {string} goodName good name
 * @apiParam {string} [description] description for good
 * @apiParam {double} [price] good price
 * @apiParam {string} [contentID] content id of good
 * @apiParam {string} [contentToken] content token
 * @apiSuccess {int32} status -1 for invalid param <br> 1 for success <br> 2 for invalid token
 * @apiSuccess {int32} buyInfoID created buyInfoID
 * @apiUse DBServerDown
 */
func (a *srv) Create(ctx context.Context, req *buyinfo.BuyInfoCreateRequest, rsp *buyinfo.BuyInfoCreateResponse) error {
	if req.ValidTime == 0 || req.GoodName == "" || req.UserID == 0 {
		rsp.Status = buyinfo.BuyInfoCreateResponse_INVALID_PARAM
		return nil
	}

	good := db.Good{
		GoodName:    req.GoodName,
		Price:       req.Price,
		Description: req.Description,
	}
	info := db.BuyInfo{
		ReleaseTime: time.Now(),
		ValidTime:   time.Unix(req.ValidTime, 0),
		UserID:      req.UserID,
	}

	insert := func() (int32, error) {
		tx := db.Ormer.Begin()
		if utils.LogContinue(tx.Error, utils.Warning) {
			return 0, tx.Error
		}
		err := tx.Create(&good).Error
		if utils.LogContinue(err, utils.Warning) {
			tx.Rollback()
			return 0, err
		}
		info.GoodID = good.ID
		err = tx.Create(&info).Error
		if utils.LogContinue(err, utils.Warning) {
			tx.Rollback()
			return 0, err
		}

		err = tx.Commit().Error
		if utils.LogContinue(err, utils.Warning) {
			tx.Rollback()
			return 0, err
		}
		return info.ID, nil
	}

	if req.ContentID == "" && req.ContentToken == "" {
		id, err := insert()
		if err != nil || id == 0 {
			return nil
		}
		rsp.Status = buyinfo.BuyInfoCreateResponse_SUCCESS
		rsp.BuyInfoID = id
	} else if req.ContentID != "" && req.ContentToken != "" {
		srv := utils.CallMicroService("content", func(name string, c client.Client) interface{} {
			return content.NewContentService(name, c)
		}, func() interface{} {
			return mock.NewContentService()
		}).(content.ContentService)
		microRsp, err := srv.Check(context.TODO(), &content.ContentCheckRequest{
			ContentID:    req.ContentID,
			ContentToken: req.ContentToken,
		})
		if err != nil || microRsp.Status != content.ContentCheckResponse_VALID {
			rsp.Status = buyinfo.BuyInfoCreateResponse_INVALID_TOKEN
			return nil
		}

		good.ContentID = req.ContentID
		id, err := insert()
		if err != nil || id == 0 {
			return nil
		}
		rsp.Status = buyinfo.BuyInfoCreateResponse_SUCCESS
		rsp.BuyInfoID = id
	} else {
		rsp.Status = buyinfo.BuyInfoCreateResponse_INVALID_PARAM
	}
	return nil
}

/**
 * @api {rpc} /rpc BuyInfo.Find
 * @apiVersion 1.0.0
 * @apiGroup Service
 * @apiName BuyInfo.Find
 * @apiDescription Find BuyInfo.
 *
 * @apiParam {int32} [userID] userID
 * @apiParam {int32} [status] status 1 for waiting <br> 2 for reserved <br> 3 for done <br> 4 for expired
 * @apiParam {string} [goodName] good name(fuzzy)
 * @apiParam {double} lowPrice=0 low bound of price, included
 * @apiParam {double} highPrice=inf high bound of price, included
 * @apiParam {uint32} limit=100 row limit
 * @apiParam {uint32} offset=0 row offset
 * @apiSuccess {list} buyInfo see [BuyInfo Service](#api-Service-BuyInfo_Query)
 * @apiUse DBServerDown
 */
func (a *srv) Find(ctx context.Context, req *buyinfo.BuyInfoFindRequest, rsp *buyinfo.BuyInfoFindResponse) error {
	type result struct {
		BuyInfoID   int32
		Status      int32
		ReleaseTime time.Time
		ValidTime   time.Time
		GoodName    string
		Price       float64
		Description string
		ContentID   string
		UserID      int32
	}

	if req.Limit == 0 {
		req.Limit = 100
	}
	if req.LowPrice < 0 {
		req.LowPrice = 0
	}
	if req.HighPrice < 0 {
		req.HighPrice = 0
	}

	var res []*result
	tb := db.Ormer.Table("buy_infos, goods").Select("buy_infos.id as buy_info_id, status, release_time, " +
		"valid_time, good_name, price, description, content_id, user_id").Where("buy_infos.good_id = goods.id")
	if req.UserID != 0 {
		tb = tb.Where("user_id = ?", req.UserID)
	}
	if req.Status != 0 {
		tb = tb.Where("status = ?", req.Status)
	}
	if req.GoodName != "" {
		tb = tb.Where("good_name LIKE ?", "%"+req.GoodName+"%")
	}
	if req.LowPrice != 0 {
		tb = tb.Where("price >= ?", req.LowPrice)
	}
	if req.HighPrice != 0 {
		tb = tb.Where("price <= ?", req.HighPrice)
	}

	err := tb.Limit(req.Limit).Offset(req.Offset).Find(&res).Error
	if utils.LogContinue(err, utils.Warning) {
		return err
	}
	for _, v := range res {
		rsp.BuyInfo = append(rsp.BuyInfo, &buyinfo.BuyInfoMsg{
			BuyInfoID:   v.BuyInfoID,
			Status:      v.Status,
			ReleaseTime: v.ReleaseTime.Unix(),
			ValidTime:   v.ValidTime.Unix(),
			GoodName:    v.GoodName,
			Price:       v.Price,
			Description: v.Description,
			ContentID:   v.ContentID,
			UserID:      v.UserID,
		})
	}
	return nil
}

func main() {
	db.InitORM("buyinfodb", new(db.BuyInfo), new(db.Good))
	defer db.CloseORM()
	service := utils.InitMicroService("buyInfo")
	utils.LogPanic(buyinfo.RegisterBuyInfoHandler(service.Server(), new(srv)))
	utils.RunMicroService(service)
}
