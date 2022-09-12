package storagerestaurant

import (
	"context"
	"go-dev/common"
	restaurantmodel "go-dev/module/restaurant/model"
)

func (store *sqlStore) SoftDeleteRestaurant(ctx context.Context,
	cond map[string]interface{},
) error {
	if err := store.db.
		Table(restaurantmodel.Restaurant{}.TableName()).
		Where(cond).
		Updates(map[string]interface{}{"status": 0}).
		Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
