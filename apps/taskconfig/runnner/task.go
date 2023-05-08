package runnner

import (
	"context"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"oh-my-helper-go/pkg/delegate"
)

type Task interface {
	Run(ctx context.Context)
	Name() string
}

type VideoTask struct {
	delegate *delegate.BiliDelegate
	// 随机数生成
	rand *rand.Rand
	// 视频排行
	trend []string
	// 分区
	regions []int64
}

func InitVideoTask(*delegate.BiliDelegate) {

}

func (v *VideoTask) initTrend() {

}

type DonateCoinTask struct {
	delegate *delegate.BiliDelegate // 请求接口代理
	vid      []string               // 待投币的视频id
}

func (d *DonateCoinTask) Run(ctx context.Context) error {
	biliUser, ok := ctx.Value("biliUser").(map[string]interface{})
	if !ok {
		return errors.New("执行任务出错，未获取到用户信息")
	}
	if level, ok := biliUser["level"].(int64); ok {
		if !(level < 6) {
			log.Infof("用户%v等级大于6级，不执行投币任务", biliUser["dedeuserid"])
		}
	}
	fmt.Printf("配置投币数量为: %v \n", d.delegate.Config.DonateCoins)
	return nil
}

func (d *DonateCoinTask) Name() string {
	return "投币任务"
}
