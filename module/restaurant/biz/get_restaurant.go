package restaurantbiz

import (
	"context"
	"errors"
	restaurantmodel "go-dev/module/restaurant/model"
)

type GetRestaurantStore interface {
	FindRestaurant(ctx context.Context, cond map[string]interface{}, moreKey ...string) (*restaurantmodel.Restaurant, error)
}

type getRestaurantBiz struct {
	store GetRestaurantStore
}

func NewGetRestaurantBiz(store GetRestaurantStore) *getRestaurantBiz {
	return &getRestaurantBiz{store: store}
}

func (res *getRestaurantBiz) GetRestaurant(ctx context.Context, id int) (*restaurantmodel.Restaurant, error) {
	result, err := res.store.FindRestaurant(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}

	if result.Status == 0 {
		return nil, errors.New("record not found")
	}

	return result, nil

}
