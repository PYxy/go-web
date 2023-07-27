package customer_app

import (
	"context"
	"fmt"
	"github.com/PYxy/go-web/internal/customer-app/config"
	"github.com/PYxy/go-web/internal/customer-app/pkg"
	"github.com/PYxy/go-web/internal/customer-app/store"
	"github.com/PYxy/go-web/internal/customer-app/store/mysql"
	linuxSingal "github.com/PYxy/go-web/pkg"
	"github.com/PYxy/go-web/pkg/middle/reject_request"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"
)

// App 定义服务接口
type App interface {
	StartAndServer()
	shutdown()
	RUN()
}

type AppOption func(app *HttpApp)

func NewApp(name, addr string, opts ...AppOption) App {
	//, shutdownTimeout, waitTime, cbTimeout, svcStopTimeOut time.Duration
	return (&HttpApp{
		Name:            name,
		Addr:            addr,
		shutdownTimeout: time.Second * 30,
		waitTime:        time.Second * 10,
		cbTimeout:       time.Second * 3,
		svcStopTimeOut:  time.Second * 3,
		reject:          reject_request.NewRejecter(time.Millisecond * 10),
	}).WithOption(opts...)

}

type HttpApp struct {
	svc  http.Server
	Name string

	Addr string
	//===下面都是service级别的参数===
	// 优雅退出整个超时时间，默认30秒
	shutdownTimeout time.Duration
	// 优雅退出时候等待处理已有请求时间，默认10秒钟
	waitTime time.Duration
	// 自定义回调超时时间，默认三秒钟
	cbTimeout time.Duration
	//服务关闭最长时间 Stop
	svcStopTimeOut time.Duration
	reject         *reject_request.RejectRequest
	//回调函数 按需
	cbs []func(ctx context.Context)
	//定时任务
	cronSlice  []func(ctx context.Context)
	ctx        context.Context
	cancelFunc context.CancelFunc
}

func (h *HttpApp) WithOption(opts ...AppOption) *HttpApp {
	for _, opt := range opts {
		opt(h)
	}
	return h
}

// WithCron 添加定时任务
func WithCron(cronSlice []func(ctx context.Context)) AppOption {
	return func(app *HttpApp) {
		app.cronSlice = append(app.cronSlice, cronSlice...)
	}
}

func WithshutdownTimeou(shutdownTimeou time.Duration) AppOption {
	return func(app *HttpApp) {
		app.shutdownTimeout = shutdownTimeou
	}
}

func WithwaitTime(waitTime time.Duration) AppOption {
	return func(app *HttpApp) {
		app.waitTime = waitTime
	}
}

// WithcbTimeout 自定义回调超时时间，默
func WithcbTimeout(cbTimeout time.Duration) AppOption {
	return func(app *HttpApp) {
		app.cbTimeout = cbTimeout
	}
}

// WithsvcStopTimeOut 服务关闭最长时间 StopTimeOut
func WithsvcStopTimeOut(svcStopTimeOut time.Duration) AppOption {
	return func(app *HttpApp) {
		app.svcStopTimeOut = svcStopTimeOut
	}
}

// WithTLS 设置TLS
func WithTLS() AppOption {
	return func(app *HttpApp) {

	}
}

// CronRun 启动定时任务
func (h *HttpApp) CronRun() {
	if len(h.cronSlice) != 0 {
		h.ctx, h.cancelFunc = context.WithCancel(context.Background())
		//TODO 启动定时任务
		for _, cronFunc := range h.cronSlice {
			go cronFunc(h.ctx)
		}
	}
}

func (h *HttpApp) StartAndServer() {
	//TODO implement me
	r := gin.Default()
	//TODO 1.注册禁止请求中间件  下面需要关注请求数 貌似只能这么写,请指教
	h.reject = reject_request.NewRejecter(time.Millisecond * 10)

	r.Use(h.reject.Build()) //TODO 全局中间件

	initRouter(r)
	//这里可以写个option 按需添加
	h.svc = http.Server{
		Addr:              h.Addr,
		Handler:           r,
		TLSConfig:         nil,
		ReadTimeout:       0,
		ReadHeaderTimeout: 0,
		WriteTimeout:      0,
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}
	// 从这里开始优雅退出监听系统信号，强制退出以及超时强制退出。
	signalch := make(chan os.Signal, 2)
	signal.Notify(signalch, linuxSingal.Signals...)
	go func() {
		if err := h.svc.ListenAndServe(); err != nil {
			if err == http.ErrServerClosed {
				log.Printf("服务器[%s]已关闭", h.Name)
			} else {
				log.Printf("服务器[%s]异常退出,异常信息:%v", h.Name, err)
				close(signalch)
			}
		}
	}()

	<-signalch

	h.shutdown()
}

func (h *HttpApp) RUN() {
	//TODO 0.获取必要参数
	configPath := pkg.ParseCommand()
	//项目使用的变量初始化 例如 文件读取, mysql redis  etcd 的连接信息
	//TODO 1.变量初始化 例如 文件读取, mysql redis  etcd 的连接信息,   注意连接的close
	configOption, err := config.ParseAppInI(configPath)
	if err != nil {
		fmt.Println(err)
		panic("读取配置文件失败")
	}
	fmt.Println("configOption", configOption.MysqlOption)
	//TODO 1.1  初始化mysql 在这里可以通过修改 Log 字段来修改日志输出位置
	//option.WithLog(nil)(configOption.MysqlOption)
	mysqlFactory, err := mysql.GetMySQLFactoryOr(configOption.MysqlOption)
	if err != nil {
		panic("连接mysql 异常")
	}
	defer func() {
		fmt.Println("mysql 连接关闭")
		_ = mysqlFactory.Close()
	}()
	store.SetClient("mysql", mysqlFactory)

	//TODO 1.x  初始化redis

	//TODO 1.x  初始化etcd

	//TODO 2.启动定时任务
	h.CronRun()
	//TODO 3.启动grpc 服务 或者 http 服务
	h.StartAndServer()
}

// shutdown  服务终止
func (h *HttpApp) shutdown() {
	//TODO implement me
	//panic("implement me")
	//TODO 1.拒绝请求
	h.reject.Store(false)
	//TODO 2.停止定时任务
	if h.cancelFunc != nil {
		h.cancelFunc()
	}

	//TODO 3.要通知下游 或者注册中心去摘除节点

	//TODO 4.在这里等待一段时间实时统计正在处理的请求数量直到为 0 或 超时 则进行下一步
	//time.Sleep(app.waitTime)
	ctx, cancel := context.WithTimeout(context.Background(), h.waitTime)
	defer cancel()

	h.reject.Wait(ctx)

	//TODO 5.正常关闭服务
	err := h.stop()
	fmt.Printf("服务[%s],关闭是否出现异常:[%v] \n", h.Name, err)
	//TODO 6 做一些收尾工作
	var wg sync.WaitGroup
	for _, cbFun := range h.cbs {
		wg.Add(1)
		go func(fun func(ctx context.Context)) {
			defer wg.Done()
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*h.cbTimeout)
			defer cancel()
			fun(ctx)
		}(cbFun)
	}
	wg.Wait()
}

func (h *HttpApp) stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*h.svcStopTimeOut)
	defer cancel()
	return h.svc.Shutdown(ctx)

}
