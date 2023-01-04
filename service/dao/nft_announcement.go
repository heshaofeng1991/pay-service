package dao

import (
	"context"
	"encoding/json"
	inter "pay-service/internal"
	"pay-service/model"
	"pay-service/service/cache"
	tp "pay-service/types"
	"strconv"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func exist(ctx context.Context) error {
	flag, err := cache.RedisClient(ctx).Exists(ctx, inter.Announcement).Result()
	zap.S().Infof("GetIssuerInfoByID flag : %v, err %v", flag, err)
	if err != nil {
		return err
	}

	if flag == 0 {
		num, err := cache.Expire(ctx, cache.RedisClient(ctx), inter.Announcement)
		zap.S().Infof("GetIssuerInfoByID num : %v, err %v", num, err)
		if err != nil {
			return err
		}
	}

	return nil
}

func GetAnnouncementInfoByID(ctx context.Context, id int64, data *model.NftAnnouncement) (err error) {
	idStr := strconv.FormatInt(id, 10)

	if err := exist(ctx); err != nil {
		return err
	}

	existed, err := cache.HExists(ctx, cache.RedisClient(ctx), inter.Announcement, idStr)
	if err != nil {
		return err
	}

	timeNow := time.Now()

	if !existed {
		if err := GetDB(ctx).Where("(start_time is NULL OR start_time <= ?) AND (end_time is NULL OR end_time >= ?)",
			timeNow, timeNow).First(data, id).Error; err != nil {
			return err
		}

		by, _ := json.Marshal(data)
		if err := cache.HSet(ctx, cache.RedisClient(ctx), inter.Announcement, idStr, string(by)); err != nil {
			return err
		}

		return nil
	}

	var total int64
	if err := GetDB(ctx).Where("(start_time is NULL OR start_time <= ?) AND (end_time is NULL OR end_time >= ?)",
		timeNow, timeNow).First(data, id).Count(&total).Error; err != nil {
		return err
	}

	if total == 0 {
		if _, err := cache.HDel(ctx, cache.RedisClient(ctx), inter.Announcement, idStr); err != nil {
			return err
		}

		return gorm.ErrRecordNotFound
	}

	result, err := cache.HGet(ctx, cache.RedisClient(ctx), inter.Announcement, idStr)
	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(result), data)
}

func GetAnnouncementList(ctx context.Context, req *tp.GetAnnouncementReq) ([]*model.NftAnnouncement, int64, error) {
	var (
		list   []*model.NftAnnouncement
		total  int64
		err    error
		result map[string]string
		detail *model.NftAnnouncement
	)

	list = make([]*model.NftAnnouncement, 0)
	result = make(map[string]string, 0)

	if err := exist(ctx); err != nil {
		return list, total, err
	}

	result, total, err = cache.HGetAll(ctx, cache.RedisClient(ctx), inter.Announcement)
	if err != nil {
		return list, total, err
	}

	for _, val := range result {
		detail = &model.NftAnnouncement{}

		_ = json.Unmarshal([]byte(val), &detail)

		list = append(list, detail)
	}

	where := getAnnouncementListCondition(req)

	// 无缓存
	if total == 0 {
		return getAnnouncementDataAndSetCache(ctx, false, where)
	}

	var count int64
	timeNow := time.Now()
	if err := GetDB(ctx).Model(model.NftAnnouncement{}).Where("(start_time is NULL OR start_time <= ?) AND (end_time is NULL OR end_time >= ?)",
		timeNow, timeNow).Where(where).Count(&count).Error; err != nil {
		return list, total, err
	}

	// 数据库与redis数据不一致
	if total != count {
		return getAnnouncementDataAndSetCache(ctx, true, where)
	}

	return list, total, err
}

func getAnnouncementDataAndSetCache(ctx context.Context, isSame bool, where map[string]interface{}) ([]*model.NftAnnouncement, int64, error) {
	var (
		list  []*model.NftAnnouncement
		total int64
		err   error
	)

	timeNow := time.Now()
	query := GetDB(ctx).Model(model.NftAnnouncement{})
	if err = query.Where("(start_time is NULL OR start_time <= ?) AND (end_time is NULL OR end_time >= ?)",
		timeNow, timeNow).Order("created_at desc").Find(&list, where).Count(&total).Error; err != nil {
		return list, total, err
	}

	if total == 0 {
		return nil, 0, gorm.ErrRecordNotFound
	}

	if isSame {
		// 删除缓存
		if _, err = cache.Del(ctx, cache.RedisClient(ctx), inter.Announcement); err != nil {
			return list, total, err
		}
	}

	// 设置缓存
	for _, val := range list {
		by, _ := json.Marshal(val)
		if err = cache.HSet(ctx, cache.RedisClient(ctx), inter.Announcement, val.Id, string(by)); err != nil {
			_, err = cache.Del(ctx, cache.RedisClient(ctx), inter.Announcement)
			if err != nil {
				return list, total, err
			}

			break
		}
	}

	return list, total, err
}

func getAnnouncementListCondition(req *tp.GetAnnouncementReq) map[string]interface{} {
	//根据实标查询条件写
	where := map[string]interface{}{}

	if req.Status != nil {
		where["status"] = *req.Status
	}

	if req.Title != "" {
		where["title"] = req.Title
	}

	return where
}
