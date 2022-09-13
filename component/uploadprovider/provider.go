package uploadprovider

import (
	"context"
	"go-dev/common"
)

type UploadProvider interface {
	SaveFileUploaded(ctx context.Context, data []byte, dst string) (*common.Image, error)
	DeleteItem(ctx context.Context, item *string) error
}
