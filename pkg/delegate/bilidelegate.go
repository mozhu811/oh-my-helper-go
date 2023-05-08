package delegate

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"net/url"
	bizerr "oh-my-helper-go/pkg/bizerr"
	"oh-my-helper-go/pkg/model"
	"strconv"
	"strings"
)

var (
	GET_QR_CODE_LOGIN_URL_OLD      = "https://passport.bilibili.com/qrcode/getLoginUrl"
	GET_QR_CODE_LOGIN_URL          = "https://passport.bilibili.com/x/passport-login/web/qrcode/generate"
	GET_QR_CODE_LOGIN_INFO_URL_OLD = "https://passport.bilibili.com/qrcode/getLoginInfo"
	GET_QR_CODE_LOGIN_INFO_URL     = "https://passport.bilibili.com/x/passport-login/web/qrcode/poll"
	GetUserInfoNav                 = "https://api.bilibili.com/x/web-interface/nav"
	GetUserSpaceInfo               = "https://api.bilibili.com/x/space/acc/info"
	GetUserSpaceDetailsInfo        = "https://api.bilibili.com/x/space/myinfo"
	GetCoinChangeLog               = "https://api.bilibili.com/x/member/web/coin/log"
	GetExpRewardStatus             = "https://api.bilibili.com/x/member/web/exp/reward"
	GET_FOLLOWED_UP_POST_VIDEO     = "https://api.vc.bilibili.com/dynamic_svr/v1/dynamic_svr/dynamic_new"
	GetTrendVideo                  = "https://api.bilibili.com/x/web-interface/ranking/region"
	ReportHeartbeat                = "https://api.bilibili.com/x/click-interface/web/heartbeat"
	VIDEO_CLICK                    = "https://api.bilibili.com/x/click-interface/click/web/h5"
	GET_VIDEO_DETAILS              = "https://api.bilibili.com/x/web-interface/view"
	ShareVideo                     = "https://api.bilibili.com/x/web-interface/share/add"
	MangaSign                      = "https://manga.bilibili.com/twirp/activity.v1.Activity/ClockIn"
	GetCoinExpToday                = "https://api.bilibili.com/x/web-interface/coin/today/exp"
	GetCoin                        = "https://account.bilibili.com/site/getCoin"
	CheckDonateCoin                = "https://api.bilibili.com/x/web-interface/archive/coins"
	DonateCoin                     = "https://api.bilibili.com/x/web-interface/coin/add"
	Silver2Coin                    = "https://api.live.bilibili.com/xlive/revenue/v1/wallet/silver2coin"
	BiliLiveWallet                 = "https://api.live.bilibili.com/xlive/revenue/v1/wallet/myWallet?need_bp=1&need_metal=1&platform=pc"
	LIKE_VIDEO                     = "https://api.bilibili.com/x/web-interface/archive/like"
	BiliLiveDoSign                 = "https://api.live.bilibili.com/xlive/web-ucenter/v1/sign/DoSign"
	ListGifts                      = "https://api.live.bilibili.com/xlive/web-room/v1/gift/bag_list"
	GET_LIVE_ROOM_INFO_OLD         = "https://api.live.bilibili.com/room/v1/Room/getRoomInfoOld"
	GetLiveRoomInfo                = "https://api.live.bilibili.com/live_user/v1/Master/info"
	SendGift                       = "https://api.live.bilibili.com/xlive/revenue/v2/gift/sendBag"
	Charge                         = "https://api.bilibili.com/x/ugcpay/web/v2/trade/elec/pay/quick"
	GetChargeInfo                  = "https://api.bilibili.com/x/ugcpay/web/v2/trade/elec/panel"
	CommitChargeComment            = "https://api.bilibili.com/x/ugcpay/trade/elec/message"
	GetVipReward                   = "https://api.bilibili.com/x/vip/privilege/receive"
	GetMangaVipReward              = "https://manga.bilibili.com/twirp/user.v1.User/GetVipReward"
	ReadManga                      = "https://manga.bilibili.com/twirp/bookshelf.v1.Bookshelf/AddHistory?device=pc&platform=web"
	GetMedalWall                   = "https://api.live.bilibili.com/xlive/web-ucenter/user/MedalWall"
	RelationModify                 = "https://api.bilibili.com/x/relation/modify"
	GET_JURY_CASE                  = "https://api.bilibili.com/x/credit/v2/jury/case/next"
	JURY_VOTE                      = "https://api.bilibili.com/x/credit/v2/jury/vote"
	VIP_SIGN                       = "https://api.bilibili.com/pgc/activity/score/task/sign"
	VIP_QUEST_RECEIVE              = "https://api.biliapi.com/pgc/activity/score/task/receive"
	VIP_QUEST_INFO                 = "https://api.biliapi.com/x/vip_point/task/combine"
	VIP_QUEST_VIEW_COMPLETE        = "https://api.bilibili.com/pgc/activity/deliver/task/complete"
	VIP_QUEST_COMPLETE             = "https://api.bilibili.com/pgc/activity/score/task/complete"
	GET_QUESTIONS                  = "https://api.bilibili.com/x/esports/guess/collection/question"
	GET_GUESS_INFO                 = "https://api.bilibili.com/x/esports/guess/collection/statis"
	GUESS_ADD                      = "https://api.bilibili.com/x/esports/guess/add"
)
var userAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36"

type BiliTaskConfig struct {
	Dedeuserid         int
	Sessdata           string
	BiliJct            string
	DonateCoins        float64
	ReserveCoins       int
	AutoCharge         bool
	DonateGift         bool
	DonateGiftTarget   int
	AutoChargeTarget   int
	DevicePlatform     string
	DonateCoinStrategy string
	UserAgent          string
	FollowDeveloper    bool
}

type BiliDelegate struct {
	Config   BiliTaskConfig
	SafeMode bool
}

func NewDelegate(config BiliTaskConfig, safeMode bool) *BiliDelegate {
	return &BiliDelegate{
		Config:   config,
		SafeMode: safeMode,
	}
}

func (bd *BiliDelegate) GetUserDetails() (*model.BiliUserDetails, error) {
	req, err := http.NewRequest("GET", GetUserSpaceDetailsInfo, nil)
	if err != nil {
		return nil, errors.Wrap(err, "构建请求失败")
	}

	req.Header.Set("Referer", "https://www.bilibili.com/")

	res, err := bd.call(req)
	if err != nil {
		return nil, err
	}

	if res.Code != 0 {
		return nil, bizerr.ErrorCookieExpired
	}

	var ud model.BiliUserDetails
	var mapConfig = &mapstructure.DecoderConfig{
		ErrorUnused: true,
		Result:      &ud,
		TagName:     "json",
	}
	decoder, err := mapstructure.NewDecoder(mapConfig)
	if err != nil {
		return nil, err
	}
	err = decoder.Decode(res.Data)
	if err != nil {
		return nil, err
	}
	return &ud, nil
}

func (bd *BiliDelegate) GetSpaceAccInfo(dedeuserid string) (*model.SpaceAccInfo, error) {
	params := url.Values{
		"mid": {dedeuserid},
	}
	req, err := http.NewRequest("GET", GetUserSpaceInfo+"?"+params.Encode(), nil)
	if err != nil {
		return nil, errors.Wrap(err, "构建请求失败")
	}

	req.Header.Set("Referer", "https://www.bilibili.com/")

	res, err := bd.call(req)
	if err != nil {
		return nil, err
	}

	var sai model.SpaceAccInfo
	var mapConfig = &mapstructure.DecoderConfig{
		ErrorUnused: true,
		Result:      &sai,
		TagName:     "json",
	}
	decoder, err := mapstructure.NewDecoder(mapConfig)
	if err != nil {
		return nil, err
	}
	err = decoder.Decode(res.Data)
	if err != nil {
		return nil, err
	}

	return &sai, nil
}

func (bd *BiliDelegate) GetCoinLog() (*model.CoinLog, error) {
	req, err := http.NewRequest("GET", GetCoinChangeLog, nil)
	if err != nil {
		return nil, errors.Wrap(err, "构建请求失败")
	}

	req.Header.Set("Referer", "https://www.bilibili.com/")

	res, err := bd.call(req)
	if err != nil {
		return nil, err
	}

	var cl model.CoinLog
	var mapConfig = &mapstructure.DecoderConfig{
		ErrorUnused: true,
		Result:      &cl,
		TagName:     "json",
	}
	decoder, err := mapstructure.NewDecoder(mapConfig)
	if err != nil {
		return nil, err
	}
	err = decoder.Decode(res.Data)
	if err != nil {
		return nil, err
	}

	return &cl, nil
}

func (bd *BiliDelegate) GetLiveRoomInfo(dedeuserid int) (*model.LiveRoom, error) {
	params := url.Values{
		"uid": {strconv.Itoa(dedeuserid)},
	}

	req, err := http.NewRequest("GET", GetLiveRoomInfo+"?"+params.Encode(), nil)
	if err != nil {
		return nil, errors.Wrap(err, "构建请求失败")
	}

	req.Header.Set("Referer", "https://www.bilibili.com/")

	res, err := bd.call(req)
	if err != nil {
		return nil, err
	}

	var lr model.LiveRoom
	var mapConfig = &mapstructure.DecoderConfig{
		ErrorUnused: true,
		Result:      &lr,
		TagName:     "json",
	}
	decoder, err := mapstructure.NewDecoder(mapConfig)
	if err != nil {
		return nil, err
	}
	err = decoder.Decode(res.Data)
	if err != nil {
		return nil, err
	}

	return &lr, nil
}

func (bd *BiliDelegate) GetMedalWall() (*model.MedalWall, error) {
	params := url.Values{
		"target_id": {strconv.Itoa(bd.Config.Dedeuserid)},
	}
	req, err := http.NewRequest("GET", GetMedalWall+"?"+params.Encode(), nil)
	if err != nil {
		return nil, errors.Wrap(err, "构建请求失败")
	}

	req.Header.Set("Referer", "https://www.bilibili.com/")

	res, err := bd.call(req)
	if err != nil {
		return nil, err
	}

	var mw model.MedalWall
	var mapConfig = &mapstructure.DecoderConfig{
		ErrorUnused: true,
		Result:      &mw,
		TagName:     "json",
	}
	decoder, err := mapstructure.NewDecoder(mapConfig)
	if err != nil {
		return nil, err
	}
	err = decoder.Decode(res.Data)
	if err != nil {
		return nil, err
	}

	return &mw, nil
}

func (bd *BiliDelegate) CheckCookie() (ok bool) {
	req, err := http.NewRequest("GET", GetUserInfoNav, nil)
	if err != nil {
		log.Errorln(errors.Wrap(err, "构建请求失败"))
		return false
	}

	req.Header.Set("Referer", "https://www.bilibili.com/")

	res, err := bd.call(req)
	if err != nil {
		return false
	}

	return res.Code == 0
}

func (bd *BiliDelegate) GetExpRewardStatus() (*model.ExpRewardStatus, error) {
	params := url.Values{
		"csrf": {bd.Config.BiliJct},
	}
	req, err := http.NewRequest("GET", GetExpRewardStatus+"?"+params.Encode(), nil)
	if err != nil {
		return nil, errors.Wrap(err, "构建请求失败")
	}

	req.Header.Set("Referer", "https://account.bilibili.com/")
	req.Header.Set("Origin", "https://account.bilibili.com/")

	res, err := bd.call(req)
	if err != nil {
		return nil, err
	}

	var ers model.ExpRewardStatus
	var mapConfig = &mapstructure.DecoderConfig{
		ErrorUnused: true,
		Result:      &ers,
		TagName:     "json",
	}
	decoder, err := mapstructure.NewDecoder(mapConfig)
	if err != nil {
		return nil, err
	}
	err = decoder.Decode(res.Data)
	if err != nil {
		return nil, err
	}

	return &ers, nil
}

// GetTrendVideo 根据分区ID获取3日热榜视频
func (bd *BiliDelegate) GetTrendVideo(rid int) ([]model.RegionRank, error) {
	params := url.Values{
		"rid": {strconv.Itoa(rid)},
		"day": {"3"},
	}

	req, err := http.NewRequest("GET", GetTrendVideo+"?"+params.Encode(), nil)
	if err != nil {
		return nil, errors.Wrap(err, "构建请求失败")
	}

	req.Header.Set("Referer", "https://www.bilibili.com/")

	res, err := bd.call(req)
	if err != nil {
		return nil, err
	}
	if res.Code != 0 {
		return nil, errors.New(res.Message)
	}
	var regionRank []model.RegionRank
	var mapConfig = &mapstructure.DecoderConfig{
		ErrorUnused: true,
		Result:      &regionRank,
		TagName:     "json",
	}
	decoder, err := mapstructure.NewDecoder(mapConfig)
	if err != nil {
		return nil, err
	}
	err = decoder.Decode(res.Data)
	if err != nil {
		return nil, err
	}

	return regionRank, nil
}

// PlayVideo 观看视频
func (bd *BiliDelegate) PlayVideo(vid string, playedTime int) (ok bool) {
	data := url.Values{
		"mid":              {strconv.Itoa(bd.Config.Dedeuserid)},
		"type":             {"4"},
		"sub_type":         {"1"},
		"play_type":        {"2"},
		"played_time":      {strconv.Itoa(playedTime)},
		"real_played_time": {strconv.Itoa(playedTime)},
		"csrf":             {bd.Config.BiliJct},
	}

	if isNumeric(vid) {
		data.Add("aid", vid)
	} else {
		data.Add("bvid", vid)
	}

	req, err := http.NewRequest("POST", ReportHeartbeat, strings.NewReader(data.Encode()))
	if err != nil {
		log.Errorln(errors.Wrap(err, "构建请求失败"))
		return false
	}

	req.Header.Set("Referer", "https://www.bilibili.com/")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	res, err := bd.call(req)
	if err != nil {
		return false
	}

	return res.Code == 0
}

// ShareVideo 分享视频
func (bd *BiliDelegate) ShareVideo(bvid string) (ok bool) {
	data := url.Values{
		"bvid": {bvid},
		"csrf": {bd.Config.BiliJct},
	}

	req, err := http.NewRequest("POST", ShareVideo, strings.NewReader(data.Encode()))
	if err != nil {
		log.Errorln(errors.Wrap(err, "构建请求失败"))
		return false
	}

	req.Header.Set("Referer", "https://www.bilibili.com/")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	res, err := bd.call(req)
	if err != nil {
		return false
	}

	return res.Code == 0
}

// MangaCheckIn 漫画签到
func (bd *BiliDelegate) MangaCheckIn(platform string) (map[string]interface{}, error) {
	data := url.Values{
		"platform": {platform},
	}

	req, err := http.NewRequest("POST", MangaSign, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return nil, errors.Wrap(err, "构建请求失败")
	}

	req.Header.Set("Origin", "https://manga.bilibili.com")
	req.Header.Set("Referer", "https://www.bilibili.com/")

	//return bd.call(req)
	return nil, nil
}

// GetCoinExpToday 获取今日通过投币获得的经验值
func (bd *BiliDelegate) GetCoinExpToday() (float64, error) {
	req, err := http.NewRequest("GET", GetCoinExpToday, nil)
	if err != nil {
		return 0, errors.Wrap(err, "构建请求失败")
	}

	req.Header.Set("Referer", "https://www.bilibili.com/")

	res, err := bd.call(req)
	if err != nil {
		return 0, err
	}

	return res.Data.(float64), nil
}

// GetVideoDetails 获取视频详细信息
func (bd *BiliDelegate) GetVideoDetails(vid string) (*model.VideoDetails, error) {
	data := url.Values{}
	if isNumeric(vid) {
		data.Add("aid", vid)
	} else {
		data.Add("bvid", vid)
	}
	req, err := http.NewRequest("GET", GET_VIDEO_DETAILS+"?"+data.Encode(), nil)
	if err != nil {
		return nil, errors.Wrap(err, "构建请求失败")
	}

	req.Header.Set("Referer", "https://www.bilibili.com/")

	res, err := bd.call(req)
	if err != nil {
		return nil, err
	}

	var videoDetails model.VideoDetails
	var mapConfig = &mapstructure.DecoderConfig{
		ErrorUnused: true,
		Result:      &videoDetails,
		TagName:     "json",
	}
	decoder, err := mapstructure.NewDecoder(mapConfig)
	if err != nil {
		return nil, err
	}
	err = decoder.Decode(res.Data)
	if err != nil {
		return nil, err
	}

	return &videoDetails, nil
}

// GetCoin 获取硬币账户余额
func (bd *BiliDelegate) GetCoin() (*model.CoinInfo, error) {
	req, err := http.NewRequest("GET", GetCoin, nil)
	if err != nil {
		return nil, errors.Wrap(err, "构建请求失败")
	}

	req.Header.Set("Referer", "https://www.bilibili.com/")

	res, err := bd.call(req)
	if err != nil {
		return nil, err
	}

	var coinInfo model.CoinInfo
	var mapConfig = &mapstructure.DecoderConfig{
		ErrorUnused: true,
		Result:      &coinInfo,
		TagName:     "json",
	}
	decoder, err := mapstructure.NewDecoder(mapConfig)
	if err != nil {
		return nil, err
	}
	err = decoder.Decode(res.Data)
	if err != nil {
		return nil, err
	}

	return &coinInfo, nil
}

// CheckDonateCoin 判断视频是否被投币
func (bd *BiliDelegate) CheckDonateCoin(bvid string) (*model.DonateCoinInfo, error) {
	data := url.Values{
		"bvid": {bvid},
	}

	req, err := http.NewRequest("GET", CheckDonateCoin+"?"+data.Encode(), nil)
	if err != nil {
		return nil, errors.Wrap(err, "构建请求失败")
	}

	req.Header.Set("Referer", "https://www.bilibili.com/")

	res, err := bd.call(req)
	if err != nil {
		return nil, err
	}

	var donateCoinInfo model.DonateCoinInfo
	var mapConfig = &mapstructure.DecoderConfig{
		ErrorUnused: true,
		Result:      &donateCoinInfo,
		TagName:     "json",
	}
	decoder, err := mapstructure.NewDecoder(mapConfig)
	if err != nil {
		return nil, err
	}
	err = decoder.Decode(res.Data)
	if err != nil {
		return nil, err
	}

	return &donateCoinInfo, nil
}

// DonateCoin 投币
func (bd *BiliDelegate) DonateCoin(bvid string, numCoin, isLike int) (ok bool) {
	data := url.Values{
		"bvid":         {bvid},
		"multiply":     {strconv.Itoa(numCoin)},
		"select_like":  {strconv.Itoa(isLike)},
		"cross_domain": {"true"},
		"csrf":         {bd.Config.BiliJct},
	}

	req, err := http.NewRequest("POST", DonateCoin, bytes.NewBufferString(data.Encode()))
	if err != nil {
		log.Errorln(errors.Wrap(err, "构建请求失败"))
		return false
	}

	req.Header.Set("Referer", fmt.Sprintf("https://www.bilibili.com/video/%s", bvid))
	req.Header.Set("Origin", "https://www.bilibili.com")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	res, err := bd.call(req)
	if err != nil {
		log.Errorln(err)
		return false
	}

	return res.Code == 0
}

// GetLiveWallet 获取直播间钱包
func (bd *BiliDelegate) GetLiveWallet() (*model.LiveWallet, error) {
	req, err := http.NewRequest("GET", BiliLiveWallet, nil)
	if err != nil {
		return nil, errors.Wrap(err, "构建请求失败")
	}

	req.Header.Set("Referer", "https://www.bilibili.com/")

	res, err := bd.call(req)
	if err != nil {
		return nil, err
	}

	if res.Code != 0 {
		return nil, fmt.Errorf("failed to get live wallet: %s", res.Message)
	}

	var wallet model.LiveWallet
	var mapConfig = &mapstructure.DecoderConfig{
		ErrorUnused: true,
		Result:      &wallet,
		TagName:     "json",
	}
	decoder, err := mapstructure.NewDecoder(mapConfig)
	if err != nil {
		return nil, err
	}
	err = decoder.Decode(res.Data)
	if err != nil {
		return nil, err
	}

	return &wallet, nil
}

// Silver2Coin 银瓜子兑换硬币
func (bd *BiliDelegate) Silver2Coin() error {
	data := url.Values{
		"csrf_token": {bd.Config.BiliJct},
		"csrf":       {bd.Config.BiliJct},
	}

	req, err := http.NewRequest("POST", Silver2Coin, bytes.NewBufferString(data.Encode()))
	//req, err := http.NewRequest("POST", Silver2Coin+"?"+"csrf_token="+bd.Config.BiliJct+"&csrf="+bd.Config.BiliJct, nil)
	if err != nil {
		return errors.Wrap(err, "构建请求失败")
	}

	req.Header.Set("Referer", "https://link.bilibili.com/p/center/index")
	req.Header.Set("Origin", "https://link.bilibili.com")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	res, err := bd.call(req)
	if err != nil {
		return err
	}
	if res.Code != 0 {
		return errors.New(res.Message)
	}
	return nil
}

// LiveCheckIn 直播签到
func (bd *BiliDelegate) LiveCheckIn() (*model.LiveCheckIn, error) {
	req, err := http.NewRequest("GET", BiliLiveDoSign, nil)
	if err != nil {
		log.Errorln(errors.Wrap(err, "构建请求失败"))
		return nil, err
	}

	req.Header.Set("Referer", "https://live.bilibili.com/")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	res, err := bd.call(req)
	if err != nil {
		return nil, err
	}

	if res.Code != 0 {
		return nil, fmt.Errorf("failed to check in: %s", res.Message)
	}

	var checkIn model.LiveCheckIn
	var mapConfig = &mapstructure.DecoderConfig{
		ErrorUnused: true,
		Result:      &checkIn,
		TagName:     "json",
	}
	decoder, err := mapstructure.NewDecoder(mapConfig)
	if err != nil {
		return nil, err
	}
	err = decoder.Decode(res.Data)
	if err != nil {
		return nil, err
	}
	return &checkIn, nil
}

// ListGifts 获取背包礼物
func (bd *BiliDelegate) ListGifts() (*model.GiftList, error) {
	req, err := http.NewRequest("GET", ListGifts, nil)
	if err != nil {
		return nil, errors.Wrap(err, "构建请求失败")
	}

	req.Header.Set("Referer", "https://live.bilibili.com/")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	res, err := bd.call(req)
	if err != nil {
		return nil, err
	}

	var glist model.GiftList

	var mapConfig = &mapstructure.DecoderConfig{
		ErrorUnused: true,
		Result:      &glist,
		TagName:     "json",
	}
	decoder, err := mapstructure.NewDecoder(mapConfig)
	if err != nil {
		return nil, err
	}
	err = decoder.Decode(res.Data)
	if err != nil {
		return nil, err
	}

	return &glist, nil
}

// DonateGift 赠送礼物
func (bd *BiliDelegate) DonateGift(uid, roomId, bagId, giftId, num int) error {
	params := url.Values{
		"biz_id":        {strconv.Itoa(roomId)},
		"ruid":          {strconv.Itoa(uid)},
		"bag_id":        {strconv.Itoa(bagId)},
		"gift_id":       {strconv.Itoa(giftId)},
		"gift_num":      {strconv.Itoa(num)},
		"uid":           {strconv.Itoa(bd.Config.Dedeuserid)},
		"csrf":          {bd.Config.BiliJct},
		"send_ruid":     {"0"},
		"storm_beat_id": {"0"},
		"price":         {"0"},
		"platform":      {"pc"},
		"biz_code":      {"live"},
	}

	req, err := http.NewRequest("POST", SendGift, bytes.NewBufferString(params.Encode()))
	if err != nil {
		log.Errorln(errors.Wrap(err, "构建请求失败"))
		return err
	}

	req.Header.Set("Referer", "https://www.bilibili.com/")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	res, err := bd.call(req)
	if err != nil {
		return err
	}

	if res.Code != 0 {
		return errors.New(res.Message)
	}
	return nil
}

// GetChargeInfo 获取充电信息
func (bd *BiliDelegate) GetChargeInfo() (*model.ChargeInfo, error) {
	params := url.Values{
		"mid": {strconv.Itoa(bd.Config.Dedeuserid)},
	}

	req, err := http.NewRequest("GET", GetChargeInfo+"?"+params.Encode(), nil)
	if err != nil {
		return nil, errors.Wrap(err, "构建请求失败")
	}

	req.Header.Set("Referer", "https://www.bilibili.com/")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	res, err := bd.call(req)
	if err != nil {
		return nil, err
	}
	if res.Code != 0 {
		return nil, errors.New(res.Message)
	}

	var info model.ChargeInfo
	var mapConfig = &mapstructure.DecoderConfig{
		ErrorUnused: true,
		Result:      &info,
		TagName:     "json",
	}
	decoder, err := mapstructure.NewDecoder(mapConfig)
	if err != nil {
		return nil, err
	}
	err = decoder.Decode(res.Data)
	if err != nil {
		return nil, err
	}

	return &info, nil
}

// DoCharge 充电
func (bd *BiliDelegate) DoCharge(coupon, uid int) (*model.ChargeResponse, error) {
	params := url.Values{
		"bp_num":              {strconv.Itoa(coupon)},
		"is_bp_remains_prior": {"true"},
		"up_mid":              {strconv.Itoa(uid)},
		"otype":               {"up"},
		"oid":                 {strconv.Itoa(bd.Config.Dedeuserid)},
		"csrf":                {bd.Config.BiliJct},
	}

	req, err := http.NewRequest("POST", Charge, bytes.NewBufferString(params.Encode()))
	if err != nil {
		return nil, errors.Wrap(err, "构建请求失败")
	}

	req.Header.Set("Referer", "https://www.bilibili.com/")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	res, err := bd.call(req)
	if err != nil {
		return nil, err
	}
	var cr model.ChargeResponse
	var mapConfig = &mapstructure.DecoderConfig{
		ErrorUnused: true,
		Result:      &cr,
		TagName:     "json",
	}
	decoder, err := mapstructure.NewDecoder(mapConfig)
	if err != nil {
		return nil, err
	}
	err = decoder.Decode(res.Data)
	if err != nil {
		return nil, err
	}
	return &cr, nil
}

// DoChargeComment 充电评论
func (bd *BiliDelegate) DoChargeComment(orderNo string) (ok bool) {
	params := url.Values{
		"order_id": {orderNo},
		"message":  {"oh-my-helper auto charge."},
		"csrf":     {bd.Config.BiliJct},
	}

	req, err := http.NewRequest("POST", CommitChargeComment, bytes.NewBufferString(params.Encode()))
	if err != nil {
		log.Errorln(errors.Wrap(err, "构建请求失败"))
		return false
	}

	req.Header.Set("Referer", "https://www.bilibili.com/")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	res, err := bd.call(req)
	if err != nil {
		log.Errorln(err)
		return false
	}
	return res.Code == 0
}

// GetMangaVipReward 漫画权益
func (bd *BiliDelegate) GetMangaVipReward() (map[string]interface{}, error) {
	params := url.Values{
		"reason_id": {"1"},
	}

	req, err := http.NewRequest("POST", GetMangaVipReward, bytes.NewBufferString(params.Encode()))
	if err != nil {
		return nil, errors.Wrap(err, "构建请求失败")
	}

	req.Header.Set("Origin", "https://manga.bilibili.com")
	req.Header.Set("Referer", "https://manga.bilibili.com/")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	//return bd.call(req)
	return nil, err
}

// GetVipReward 领取大会员权益
func (bd *BiliDelegate) GetVipReward(tid int) int64 {
	params := url.Values{
		"type": {strconv.Itoa(tid)},
		"csrf": {bd.Config.BiliJct},
	}

	req, err := http.NewRequest("POST", GetVipReward, bytes.NewBufferString(params.Encode()))
	if err != nil {
		log.Errorln(errors.Wrap(err, "构建请求失败"))
		return -1
	}
	req.Header.Set("Referer", "https://www.bilibili.com/")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	res, err := bd.call(req)
	if err != nil {
		return -1
	}

	return res.Code
}

// ReadManga 阅读漫画
func (bd *BiliDelegate) ReadManga() (map[string]interface{}, error) {
	params := url.Values{
		"device":   {"pc"},
		"platform": {"web"},
		"comic_id": {"26009"},
		"ep_id":    {"300318"},
	}

	req, err := http.NewRequest("POST", ReadManga, bytes.NewBufferString(params.Encode()))
	if err != nil {
		return nil, errors.Wrap(err, "构建请求失败")
	}
	req.Header.Set("Referer", "https://www.bilibili.com/")
	req.Header.Set("Origin", "https://manga.bilibili.com")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	//return bd.call(req)
	return nil, err
}

// FollowUser 关注用户
func (bd *BiliDelegate) FollowUser(uid string) (ok bool) {
	params := url.Values{
		"fid":    {uid},
		"act":    {"1"},
		"re_src": {"11"},
		"csrf":   {bd.Config.BiliJct},
	}

	req, err := http.NewRequest("POST", RelationModify, bytes.NewBufferString(params.Encode()))
	if err != nil {
		log.Error(err)
		return false
	}
	req.Header.Set("Referer", "https://www.bilibili.com/")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	res, err := bd.call(req)
	if err != nil {
		return false
	}

	return res.Code == 0
}

func (bd *BiliDelegate) call(r *http.Request) (*BiliResponse, error) {
	r.Header.Set("User-Agent", userAgent)
	// 设置Cookie
	r.AddCookie(&http.Cookie{Name: "SESSDATA", Value: bd.Config.Sessdata})
	r.AddCookie(&http.Cookie{Name: "bili_jct", Value: bd.Config.BiliJct})
	r.AddCookie(&http.Cookie{Name: "DedeUserID", Value: strconv.Itoa(bd.Config.Dedeuserid)})
	r.AddCookie(&http.Cookie{Name: "buvid3", Value: "3140CB9C-51C2-DC5B-F1A9-1F432B3D6E7B79263infoc"})
	r.AddCookie(&http.Cookie{Name: "innersign", Value: "0"})

	// 设置http代理
	//proxy, _ := url.Parse("http://localhost:7890")
	//transport := &http.Transport{
	//	Proxy: http.ProxyURL(proxy),
	//}
	client := &http.Client{
		//Transport: transport,
	}
	resp, err := client.Do(r)
	if err != nil {
		return nil, errors.Wrapf(err, "请求API[%v]失败", r.URL)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "读取API[%v]响应数据失败", r.URL)
	}

	var res BiliResponse
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, errors.Wrapf(err, "解析API[%v]响应数据失败: %v", r.URL, string(body))
	}
	return &res, nil
}

// 判断是否为数字
func isNumeric(str string) bool {
	_, err := strconv.Atoi(str)
	return err == nil
}

type BiliResponse struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
