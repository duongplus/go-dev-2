package restaurantbiz

import (
	"context"
	"go-dev/common"
	restaurantmodel "go-dev/module/restaurant/model"
)

type ListRestaurantStore interface {
	ListRestaurant(ctx context.Context,
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKey ...string,
	) ([]restaurantmodel.Restaurant, error)
}
type listRestaurantBiz struct {
	store ListRestaurantStore
}

func NewListRestaurantBiz(store ListRestaurantStore) *listRestaurantBiz {
	return &listRestaurantBiz{store: store}
}

func (biz *listRestaurantBiz) GetListRestaurant(ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
) ([]restaurantmodel.Restaurant, error) {
	result, err := biz.store.ListRestaurant(ctx, filter, paging)

	if err != nil {
		return nil, err
	}

	return result, nil
}
