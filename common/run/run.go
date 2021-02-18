package run

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"project/common/router"
	"project/utils"
	"project/utils/config"

	"go.uber.org/zap"
)

const (
	Version = "0.1.0"
)

var LogoContent = `
 ┏┓		 ┏┓
┏┛┻━━━━━━┛┻┓
┃　　	   ┃ 　
┃　　　━	   ┃
┃　┳┛	┗┳ ┃
┃　	       ┃
┃　　　┻　　 ┃
┃　　　　	   ┃
┗━┓　　	 ┏━┛
  ┃	  	 ┃
  ┃	  	 ┃
  ┃	  	 ┗━━━┓
  ┃		     ┣┓
  ┃	  		 ┏┛
  ┗┓┓┏━━━━┳┓┏┛

`

func Run() {
	r := router.Setup(config.ApplicationConfig)

	// 启动服务（优雅关机）
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", config.ApplicationConfig.Host, config.ApplicationConfig.Port),
		Handler: r,
	}
	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Fatal("listen: ", zap.Error(err))
		}
	}()

	fmt.Println(utils.Red(string(LogoContent)))
	tip()
	fmt.Println(utils.Green("Server run at:"))
	fmt.Printf("-  Local:   http://localhost:%s/ \r\n", config.ApplicationConfig.Port)
	fmt.Printf("-  Network: http://%s:%s/ \r\n", utils.GetLocaHonst(), config.ApplicationConfig.Port)
	fmt.Println(utils.Green("Swagger run at:"))
	fmt.Printf("-  Local:   http://localhost:%s/swagger/index.html \r\n", config.ApplicationConfig.Port)
	fmt.Printf("-  Network: http://%s:%s/swagger/index.html \r\n", utils.GetLocaHonst(), config.ApplicationConfig.Port)
	fmt.Printf(utils.Red("%s Enter Control + C Shutdown Server \r\n"), utils.GetCurrentTimeStr())

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	zap.L().Info(utils.Red("Shutdown Server ..."))
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal(utils.Red("Server Shutdown"), zap.Error(err))
	}
	zap.L().Info(utils.Red("Server exiting"))
}

func tip() {
	usageStr := `欢迎使用 ` + utils.Green(`GoSword `+config.ApplicationConfig.Version) + ` 可以使用 ` + utils.Red(`-help`) + ` 查看命令`
	fmt.Printf("%s \n\n", usageStr)
}
