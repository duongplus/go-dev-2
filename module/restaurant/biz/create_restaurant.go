package restaurantbiz

import (
	"context"
	restaurantmodel "go-dev/module/restaurant/model"
)

type CreateRestaurantStore interface {
	InsertRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error
}

type createRestaurantBiz struct {
	store CreateRestaurantStore
}

func NewCreateRestaurantBiz(store CreateRestaurantStore) *createRestaurantBiz {
	return &createRestaurantBiz{store: store}
}

func (res *createRestaurantBiz) CreateRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	if err := res.store.InsertRestaurant(ctx, data); err != nil {
		return err
	}

	return nil
}
