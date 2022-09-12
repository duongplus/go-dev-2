package restaurantbiz

import (
	"context"
	restaurantmodel "go-dev/module/restaurant/model"
)

type SoftDeleteRestaurantStore interface {
	FindRestaurant(
		ctx context.Context,
		cond map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)

	SoftDeleteRestaurant(ctx context.Context,
		cond map[string]interface{},
	) error
}

type softDeleteRestaurantBiz struct {
	store SoftDeleteRestaurantStore
}

func NewSoftDeleteRestaurantBiz(store SoftDeleteRestaurantStore) *softDeleteRestaurantBiz {
	return &softDeleteRestaurantBiz{store: store}
}

func (biz *softDeleteRestaurantBiz) SoftDeleteRestaurantById(ctx context.Context, id int) error {
	if _, err := biz.store.FindRestaurant(ctx, map[string]interface{}{"id": id, "status": 1}); err != nil {
		return err
	}

	if err := biz.store.SoftDeleteRestaurant(ctx, map[string]interface{}{"id": id}); err != nil {
		return err
	}

	return nil
}
