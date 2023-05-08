package runnner

import (
	"github.com/sirupsen/logrus"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/zero-contrib/logx/logrusx"
	"io"
	"oh-my-helper-go/pkg/delegate"
	"os"
	"testing"
)

var executor *BiliExecutor

func init() {
	writer := logrusx.NewLogrusWriter(func(logger *logrus.Logger) {
		logger.SetFormatter(&logrus.TextFormatter{})
		stdout := os.Stdout
		fwriter, err := os.OpenFile("runner.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			panic(err)
		}
		logger.SetOutput(io.MultiWriter(stdout, fwriter))
	})

	logx.SetWriter(writer)

	d := delegate.NewDelegate(delegate.BiliTaskConfig{
		Dedeuserid:         287969457,
		Sessdata:           "f1eafd50%2C1699117920%2Cf8d8d%2A52",
		BiliJct:            "d5d73da8a344d382ccd7bb4070bb31f3",
		DonateCoins:        5,
		ReserveCoins:       50,
		AutoCharge:         true,
		DonateGift:         true,
		DonateGiftTarget:   387636363,
		AutoChargeTarget:   287969457,
		DevicePlatform:     "ios",
		DonateCoinStrategy: "0",
		UserAgent:          "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36",
		FollowDeveloper:    false,
	}, false)
	u, err := d.GetUserDetails()
	if err != nil {
		logx.Errorw("get user details failed", logx.Field("uid", d.Config.Dedeuserid))
		panic(err)
	}
	executor = NewExecutor(d, u)
}

func TestGetCoinChangeLog(t *testing.T) {
	err := executor.GetCoinChangeLog()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("ok")
}

func TestWatchVideo(t *testing.T) {
	err := executor.WatchVideo()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("ok")
}

func TestGetVipPrivilege(t *testing.T) {
	err := executor.GetVipPrivilege()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("ok")
}

func TestDonateCoin(t *testing.T) {
	err := executor.DonateCoin()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("ok")
}

func TestSilver2Coin(t *testing.T) {
	err := executor.Silver2Coin()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("ok")
}

func TestLiveCheckIn(t *testing.T) {
	err := executor.LiveCheckIn()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("ok")
}

func TestDonateGift(t *testing.T) {
	err := executor.DonateGift()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("ok")
}

func TestBiliExecutor_AutoCharge(t *testing.T) {
	err := executor.AutoCharge()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("ok")
}
