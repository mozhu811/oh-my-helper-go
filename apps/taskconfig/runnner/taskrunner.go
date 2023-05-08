package runnner

import (
	"errors"
	_ "github.com/sirupsen/logrus"
	"github.com/zeromicro/go-zero/core/logx"
	"math/rand"
	"oh-my-helper-go/pkg/delegate"
	"oh-my-helper-go/pkg/model"
	"strconv"
	"time"
)

type BiliExecutor struct {
	d    *delegate.BiliDelegate
	user *model.BiliUserDetails
}

func NewExecutor(delegate *delegate.BiliDelegate, user *model.BiliUserDetails) *BiliExecutor {
	return &BiliExecutor{
		d:    delegate,
		user: user,
	}
}

func (b *BiliExecutor) GetCoinChangeLog() error {
	coinLog, err := b.d.GetCoinLog()
	if err != nil {
		return err
	}

	count := coinLog.Count
	logx.Infof("最近一周共产生%v条变更日志", count)

	var in float64
	var out float64

	for _, l := range coinLog.List {
		if l.Delta > 0 {
			in += l.Delta
		} else {
			out += l.Delta
		}
	}

	logx.Infof("最近一周共产生%v个硬币", in)
	logx.Infof("最近一周共消耗%v个硬币", out)
	return nil
}

func (b *BiliExecutor) WatchVideo() error {
	ers, err := b.d.GetExpRewardStatus()
	if err != nil {
		return err
	}
	videos, err := b.d.GetTrendVideo(129)
	if err != nil {
		return err
	}
	video := videos[rand.Intn(len(videos))]
	if !ers.Watch {
		// play video
		ptime := rand.Intn(60) + 10
		ok := b.d.PlayVideo(video.Bvid, ptime)
		if ok {
			logx.Infow("今日观看视频任务已完成 ✔️", logx.Field("dedeuserid", b.user.Mid), logx.Field("bvid", video.Bvid), logx.Field("ptime", ptime), logx.Field("title", video.Title))
		} else {
			logx.Error("今日观看视频任务失败 ❌", logx.Field("dedeuserid", b.user.Mid))
		}
	} else {
		logx.Infow("今日观看视频任务已完成 ✔️", logx.Field("dedeuserid", b.user.Mid))
	}

	if !ers.Share {
		// share video
		ok := b.d.ShareVideo(video.Bvid)
		if ok {
			logx.Info("今日分享视频任务已完成 ✔️", logx.Field("dedeuserid", b.user.Mid), logx.Field("bvid", video.Bvid), logx.Field("title", video.Title))
		} else {
			logx.Error("今日分享视频任务失败 ❌", logx.Field("dedeuserid", b.user.Mid))
		}
	} else {
		logx.Info("今日分享视频任务已完成 ✔️", logx.Field("dedeuserid", b.user.Mid))
	}
	return nil
}

func (b *BiliExecutor) GetVipPrivilege() error {
	isBigVip := b.user.Vip.Type > 0 && b.user.Vip.Status == 1
	if !isBigVip {
		logx.Infow("该账号非大会员，取消执行领取大会员权益 ❌", logx.Field("dedeuserid", b.user.Mid))
		return errors.New("user " + strconv.Itoa(b.user.Mid) + " is not big vip")
	}
	code := b.d.GetVipReward(1)
	if code < 0 {
		logx.Errorw("领取大会员权益失败 ❌", logx.Field("dedeuserid", b.user.Mid))
	} else if code == 0 {
		logx.Infow("领取大会员权益成功 ✔️", logx.Field("dedeuserid", b.user.Mid))
	} else if code == 69801 {
		logx.Infow("本月已领取大会员权益 ✔️", logx.Field("dedeuserid", b.user.Mid))
	} else {
		logx.Errorw("未知错误 ❌", logx.Field("dedeuserid", b.user.Mid))
	}
	return nil
}

func (b *BiliExecutor) DonateCoin() error {
	logx.Infow("配置投币数为: "+
		strconv.FormatFloat(b.d.Config.DonateCoins, 'f', 2, 64),
		logx.Field("dedeuserid", b.user.Mid))

	if b.user.Level >= 6 {
		logx.Infow("当前等级为"+strconv.Itoa(b.user.Level)+"级，取消执行投币任务 ❌", logx.Field("dedeuserid", b.user.Mid))
		return nil
	}

	// 获取当日经验
	cet, err := b.d.GetCoinExpToday()
	if err != nil {
		return err
	}

	// 计算需要投币数
	needCoins := b.d.Config.DonateCoins - cet/10

	if needCoins <= 0 {
		logx.Infow("今日投币任务已完成 ✔️", logx.Field("dedeuserid", b.user.Mid))
		return nil
	}

	balance := b.user.Coins
	if balance < b.d.Config.DonateCoins {
		logx.Infow("账户余额不足，取消执行投币任务 ❌", logx.Field("dedeuserid", b.user.Mid))
		return nil
	}

	// 执行投币
	for b.calExpDiff() > 0 && needCoins > 0 {
		// 获取视频列表
		videos, err := b.d.GetTrendVideo(129)
		if err != nil {
			return err
		}
		// 随机选择一个视频
		video := videos[rand.Intn(len(videos))]
		ok := b.d.DonateCoin(video.Bvid, 1, 1)
		if ok {
			logx.Infow("为视频["+video.Title+"]投币成功 ✔️", logx.Field("dedeuserid", b.user.Mid))
			needCoins -= 1
		} else {
			logx.Errorw("为视频["+video.Title+"]投币失败 ❌", logx.Field("dedeuserid", b.user.Mid))
			break
		}

	}
	return nil
}

func (b *BiliExecutor) calExpDiff() float64 {
	cet, err := b.d.GetCoinExpToday()
	if err != nil {
		return 0
	}
	return b.d.Config.DonateCoins - cet/10
}

func (b *BiliExecutor) Silver2Coin() error {
	wallet, err := b.d.GetLiveWallet()
	if err != nil {
		logx.Errorw(err.Error(), logx.Field("dedeuserid", b.user.Mid))
		return err
	}
	nSilver := wallet.Silver
	if nSilver < 700 {
		logx.Errorw("银瓜子余额不足 ❌", logx.Field("dedeuserid", b.user.Mid))
		return nil
	}

	err = b.d.Silver2Coin()
	if err != nil {
		logx.Errorw("兑换银瓜子失败 ❌: "+err.Error(), logx.Field("dedeuserid", b.user.Mid))
		return err
	}
	logx.Infow("兑换银瓜子成功 ✔️️", logx.Field("dedeuserid", b.user.Mid))
	return nil
}

func (b *BiliExecutor) LiveCheckIn() error {
	liveCheckIn, err := b.d.LiveCheckIn()
	if err != nil {
		logx.Errorw("直播签到失败 ❌: "+err.Error(), logx.Field("dedeuserid", b.user.Mid))
		return err
	}
	logx.Infow("直播签到成功 ✔️，本次获得"+liveCheckIn.Text+"，"+liveCheckIn.SpecialText+" ✔️️", logx.Field("dedeuserid", b.user.Mid))
	return nil
}

func (b *BiliExecutor) DonateGift() error {
	config := b.d.Config
	enable := config.DonateGift
	if !enable {
		return nil
	}

	target := config.DonateGiftTarget

	var roomId int
	roomInfo, err := b.d.GetLiveRoomInfo(target)
	if err != nil {
		roomId = 11526309
		target = 287969457
	} else {
		roomId = roomInfo.RoomID
	}

	gifts, err := b.d.ListGifts()
	if err != nil {
		return err
	}
	if len(gifts.List) < 1 {
		logx.Infow("礼物列表为空，取消执行任务 ❌", logx.Field("dedeuserid", b.user.Mid))
		return nil
	}
	// 获取当前时间戳
	now := time.Now().Unix()
	for _, g := range gifts.List {
		if g.ExpireAt != 0 &&
			g.ExpireAt-now < 60*60*24*3 {
			err = b.d.DonateGift(target, roomId, g.BagID, g.GiftID, g.GiftNum)
			if err == nil {
				logx.Infow("赠送礼物成功 ✔️", logx.Field("dedeuserid", b.user.Mid))
			} else {
				logx.Errorw("赠送礼物失败 ❌:"+err.Error(), logx.Field("dedeuserid", b.user.Mid))
			}
		}
	}
	return nil
}

func (b *BiliExecutor) AutoCharge() error {
	enable := b.d.Config.AutoCharge
	if !enable {
		logx.Infow("自动充电未开启，取消执行任务 ❌", logx.Field("dedeuserid", b.user.Mid))
		return nil
	}

	user := b.user
	if user.Vip.Status == 1 && user.Vip.Type == 1 {
		logx.Infow("该账号非大会员，取消执行任务 ❌", logx.Field("dedeuserid", b.user.Mid))
		return nil
	}

	chargeInfo, err := b.d.GetChargeInfo()
	if err != nil {
		logx.Errorw("获取充电信息失败 ❌: "+err.Error(), logx.Field("dedeuserid", b.user.Mid))
		return err
	}
	coupon := chargeInfo.BpWallet.CouponBalance
	//if coupon < 2 {
	//	logx.Infow("B币券余额不足，取消执行任务 ❌", logx.Field("dedeuserid", b.user.Mid))
	//	return nil
	//}

	target := b.d.Config.AutoChargeTarget
	if target == 0 {
		target = b.d.Config.Dedeuserid
	}

	order, err := b.d.DoCharge(coupon, target)
	if err != nil {
		logx.Errorw("充电失败 ❌: "+err.Error(), logx.Field("dedeuserid", b.user.Mid))
		return err
	}
	if order.Status == 4 {
		logx.Infow("充电成功，本次消费"+strconv.Itoa(coupon)+"个B币券 ✔️️", logx.Field("dedeuserid", b.user.Mid))
		go b.d.DoChargeComment(order.OrderNo)
	}
	return nil
}

func (b *BiliExecutor) BigVip() error {
	return nil
}
