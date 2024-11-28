package log

import (
	"fmt"
	"github.com/duke-git/lancet/v2/xerror"
	"go.uber.org/zap"
	"testing"
	"time"
)

func TestLog(t *testing.T) {
	logger := New(
		WithDebug(true),
		WithConsoleChannel(),
		WithServerJiang3Channel("sctp3245ta-0i6e7emiwxck4g1blm0tbk03"),
	)
	now := time.Now()
	err := getErr()
	logger.Error("1", zap.Error(err))
	fmt.Println(time.Since(now))
	time.Sleep(time.Second * 2)
}

func getErr() error {
	return xerror.New("报错")
}
