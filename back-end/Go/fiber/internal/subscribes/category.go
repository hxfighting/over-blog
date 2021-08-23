package subscribes

import (
	"context"
	"sync"
	"time"

	"github.com/ohdata/blog/internal/enum"
	"github.com/ohdata/blog/internal/pkg/cache"
	"github.com/ohdata/blog/tools/log"
)

func CategoryCreateSubscribe() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	f := func() (context.Context, *sync.Mutex) {
		return ctx, nil
	}
	if err := cache.Remove(f, enum.CacheCategory.String()); err != nil {
		log.Log.Err(err).Send()
	}
}
