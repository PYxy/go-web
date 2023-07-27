package reject_request

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/atomic"
	"time"
)

type RejectRequest struct {
	status        *atomic.Bool   //允许请求进入
	count         *atomic.Uint32 //当前正在处理的请求数
	checkDuration time.Duration
}

func NewRejecter(duration time.Duration) *RejectRequest {
	return &RejectRequest{
		status:        atomic.NewBool(true),
		count:         atomic.NewUint32(0),
		checkDuration: duration,
	}
}

// Allow 是否运行请求进来
func (r *RejectRequest) Allow() bool {
	return r.status.Load()
}

func (r *RejectRequest) Store(flag bool) {
	r.status.Store(flag)
}

func (r *RejectRequest) Wait(ctx context.Context) {

	if r.count.Load() == 0 {
		return
	}
	ticker := time.NewTicker(r.checkDuration)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if r.count.Load() == 0 {
				return
			}
		case <-ctx.Done():
			return
		}
	}

}

func (r *RejectRequest) Build() gin.HandlerFunc {
	return func(context *gin.Context) {
		if !r.Allow() {
			context.JSON(500, map[string]string{
				"code":   "0",
				"errmsg": "服务关闭中,拒绝请求",
			})
			context.Abort()

		}
		//增加一个
		r.count.Add(1)
		context.Next()
		r.count.Sub(1)
	}
}
