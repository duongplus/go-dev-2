package uploadbusiness

import (
	"context"
	"go-dev/common"
)

type ListImageStorage interface {
	ListImages(
		context context.Context,
		ids []int,
		moreKeys ...string,
	) ([]common.Image, error)
}

type listImagesBiz struct {
	store ListImageStorage
}

func NewListImageBiz(store ListImageStorage) *listImagesBiz {
	return &listImagesBiz{store: store}
}

func (biz *listImagesBiz) List(ctx context.Context, ids []int) (common.Images, error) {
	result, err := biz.store.ListImages(ctx, ids)

	if err != nil {
		return nil, err
	}

	return result, nil
}
