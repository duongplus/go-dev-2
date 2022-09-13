package storagerestaurant

import (
	"context"
	"go-dev/common"
	restaurantmodel "go-dev/module/restaurant/model"
)

func (store *sqlStore) ListRestaurant(ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	moreKey ...string,
) ([]restaurantmodel.Restaurant, error) {
	var result []restaurantmodel.Restaurant

	db := store.db

	db = db.
		Table(restaurantmodel.Restaurant{}.TableName()).
		Where("status = 1").
		Count(&paging.Total)

	if v := filter.OwnerId; v > 0 {
		db = db.Where("owner_id = ?", v)
	}

	if v := paging.FakeCursor; v != "" {
		if uid, err := common.FromBase58(v); err == nil {
			db = db.Where("id < ?", uid.GetLocalID())
		}
	} else {
		offset := (paging.Page - 1) * paging.Limit
		db = db.Offset(offset)
	}

	if err := db.
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if len(result) > 0 {
		result[len(result)-1].Mask(true)
		paging.NextCursor = result[len(result)-1].FakeId.String()
	}

	return result, nil
}
