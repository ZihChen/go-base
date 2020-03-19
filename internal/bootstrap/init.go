package bootstrap

import (
	"os"
	"os/signal"
	"syscall"
)

var sig chan os.Signal
var serverClose chan struct{}

// SetupGracefulSignal 設定優雅關閉的信號
func SetupGracefulSignal() {
	sig = make(chan os.Signal, 1)
	serverClose = make(chan struct{})
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	go func() {
		<-sig
		close(serverClose)
	}()
}

// GracefulDown 優雅結束程式
func GracefulDown() <-chan struct{} {
	return serverClose
}

// WaitOnceSignal 等待一次的訊號
func WaitOnceSignal() (sig chan os.Signal) {
	sig = make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	return
}
