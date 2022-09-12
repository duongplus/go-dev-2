package storagerestaurant

import (
	"context"
	"go-dev/common"
	restaurantmodel "go-dev/module/restaurant/model"
)

func (store *sqlStore) InsertRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	data.PrepareForInsert()
	if err := store.db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
