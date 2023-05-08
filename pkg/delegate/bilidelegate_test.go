package delegate

import (
	"testing"
)

var bd = &BiliDelegate{
	Config: BiliTaskConfig{
		Dedeuserid:         "287969457",
		Sessdata:           "ed19b2a9%2C1698916149%2C26a7e%2A52",
		BiliJct:            "69f9e060c1fff062f1e688d4604e6caa",
		DonateCoins:        5,
		ReserveCoins:       50,
		AutoCharge:         true,
		DonateGift:         true,
		DonateGiftTarget:   "287969457",
		AutoChargeTarget:   "287969457",
		DevicePlatform:     "ios",
		DonateCoinStrategy: "0",
		UserAgent:          "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36",
		FollowDeveloper:    false,
	},
	SafeMode: false,
}

func TestGetUserDetails(t *testing.T) {
	resp, err := bd.GetUserDetails()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(resp)
}

func TestGetSpaceAccInfo(t *testing.T) {
	resp, err := bd.GetSpaceAccInfo("287969457")
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	t.Log(resp)
}

func TestGetCoinLog(t *testing.T) {
	resp, err := bd.GetCoinLog()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(resp)
}

func TestGetLiveRoomInfo(t *testing.T) {
	resp, err := bd.GetLiveRoomInfo("287969457")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(resp)
}

func TestGetMedalWall(t *testing.T) {
	resp, err := bd.GetMedalWall()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(resp)
}

func TestCheckCookie(t *testing.T) {
	ok := bd.CheckCookie()
	if !ok {
		t.Fatal("Cookie is invalid")
	}

	t.Log("Cookie is valid")
}

func TestGetExpRewardStatus(t *testing.T) {
	resp, err := bd.GetExpRewardStatus()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(resp)
}

func TestGetTrendVideo(t *testing.T) {
	resp, err := bd.GetTrendVideo(129)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(resp)
}

func TestPlayVideo(t *testing.T) {
	ok := bd.PlayVideo("BV1ch4y1s7zX", 90)
	if !ok {
		t.Fatal("Play video failed")
	}

	t.Log("Play video success")
}

func TestShareVideo(t *testing.T) {
	ok := bd.ShareVideo("BV1ch4y1s7zX")
	if !ok {
		t.Fatal("Share video failed")
	}

	t.Log("Share video success")
}

// DEPRECATED
func TestMangaCheckIn(t *testing.T) {
	resp, err := bd.MangaCheckIn(bd.Config.DevicePlatform)
	if err != nil {
		t.Fatal(err)
	}

	code, _ := resp["code"].(float64)
	if code != 0 {
		t.Fatalf("Actual code: %v, expected: 0", code)
	}
	t.Log(resp)
}

func TestGetCoinExpToday(t *testing.T) {
	resp, err := bd.GetCoinExpToday()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(resp)
}

func TestGetVideoDetails(t *testing.T) {
	resp, err := bd.GetVideoDetails("BV1ch4y1s7zX")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(resp)
}

func TestGetCoin(t *testing.T) {
	resp, err := bd.GetCoin()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(resp)
}

func TestDonateCoin(t *testing.T) {
	ok := bd.DonateCoin("BV1ch4y1s7zX", 1, 1)
	if !ok {
		t.Fatal("Donate coin failed")
	}

	t.Log("Donate coin success")
}

func TestCheckDonateCoin(t *testing.T) {
	resp, err := bd.CheckDonateCoin("BV1ch4y1s7zX")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(resp)
}

func TestGetLiveWallet(t *testing.T) {
	resp, err := bd.GetLiveWallet()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(resp)
}

func TestSilver2Coin(t *testing.T) {
	err := bd.Silver2Coin()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Silver2Coin success")
}
