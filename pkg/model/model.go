package model

type BiliUserDetails struct {
	Mid            int    `json:"mid"`
	Name           string `json:"name"`
	Sex            string `json:"sex"`
	Face           string `json:"face"`
	Sign           string `json:"sign"`
	Rank           int    `json:"rank"`
	Level          int    `json:"level"`
	Jointime       int    `json:"jointime"`
	Moral          int    `json:"moral"`
	Silence        int    `json:"silence"`
	EmailStatus    int    `json:"email_status"`
	TelStatus      int    `json:"tel_status"`
	Identification int    `json:"identification"`
	Vip            struct {
		Type       int   `json:"type"`
		Status     int   `json:"status"`
		DueDate    int64 `json:"due_date"`
		VipPayType int   `json:"vip_pay_type"`
		ThemeType  int   `json:"theme_type"`
		Label      struct {
			Path                  string `json:"path"`
			Text                  string `json:"text"`
			LabelTheme            string `json:"label_theme"`
			TextColor             string `json:"text_color"`
			BgStyle               int    `json:"bg_style"`
			BgColor               string `json:"bg_color"`
			BorderColor           string `json:"border_color"`
			UseImgLabel           bool   `json:"use_img_label"`
			ImgLabelURIHans       string `json:"img_label_uri_hans"`
			ImgLabelURIHant       string `json:"img_label_uri_hant"`
			ImgLabelURIHansStatic string `json:"img_label_uri_hans_static"`
			ImgLabelURIHantStatic string `json:"img_label_uri_hant_static"`
		} `json:"label"`
		AvatarSubscript    int    `json:"avatar_subscript"`
		NicknameColor      string `json:"nickname_color"`
		Role               int    `json:"role"`
		AvatarSubscriptURL string `json:"avatar_subscript_url"`
		TvVipStatus        int    `json:"tv_vip_status"`
		TvVipPayType       int    `json:"tv_vip_pay_type"`
	} `json:"vip"`
	Pendant struct {
		Pid               int    `json:"pid"`
		Name              string `json:"name"`
		Image             string `json:"image"`
		Expire            int    `json:"expire"`
		ImageEnhance      string `json:"image_enhance"`
		ImageEnhanceFrame string `json:"image_enhance_frame"`
	} `json:"pendant"`
	Nameplate struct {
		Nid        int    `json:"nid"`
		Name       string `json:"name"`
		Image      string `json:"image"`
		ImageSmall string `json:"image_small"`
		Level      string `json:"level"`
		Condition  string `json:"condition"`
	} `json:"nameplate"`
	Official struct {
		Role  int    `json:"role"`
		Title string `json:"title"`
		Desc  string `json:"desc"`
		Type  int    `json:"type"`
	} `json:"official"`
	Birthday      int  `json:"birthday"`
	IsTourist     int  `json:"is_tourist"`
	IsFakeAccount int  `json:"is_fake_account"`
	PinPrompting  int  `json:"pin_prompting"`
	IsDeleted     int  `json:"is_deleted"`
	InRegAudit    int  `json:"in_reg_audit"`
	IsRipUser     bool `json:"is_rip_user"`
	Profession    struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		ShowName    string `json:"show_name"`
		IsShow      int    `json:"is_show"`
		CategoryOne string `json:"category_one"`
		Realname    string `json:"realname"`
		Title       string `json:"title"`
		Department  string `json:"department"`
	} `json:"profession"`
	FaceNft        int `json:"face_nft"`
	FaceNftNew     int `json:"face_nft_new"`
	IsSeniorMember int `json:"is_senior_member"`
	Honours        struct {
		Mid    int `json:"mid"`
		Colour struct {
			Dark   string `json:"dark"`
			Normal string `json:"normal"`
		} `json:"colour"`
		Tags any `json:"tags"`
	} `json:"honours"`
	DigitalID   string `json:"digital_id"`
	DigitalType int    `json:"digital_type"`
	LevelExp    struct {
		CurrentLevel int   `json:"current_level"`
		CurrentMin   int   `json:"current_min"`
		CurrentExp   int   `json:"current_exp"`
		NextExp      int   `json:"next_exp"`
		LevelUp      int64 `json:"level_up"`
	} `json:"level_exp"`
	Coins     float64 `json:"coins"`
	Following int     `json:"following"`
	Follower  int     `json:"follower"`
}

type SpaceAccInfo struct {
	Mid         int     `json:"mid"`
	Name        string  `json:"name"`
	Sex         string  `json:"sex"`
	Face        string  `json:"face"`
	FaceNft     int     `json:"face_nft"`
	FaceNftType int     `json:"face_nft_type"`
	Sign        string  `json:"sign"`
	Rank        int     `json:"rank"`
	Level       int     `json:"level"`
	Jointime    int     `json:"jointime"`
	Moral       int     `json:"moral"`
	Silence     int     `json:"silence"`
	Coins       float64 `json:"coins"`
	FansBadge   bool    `json:"fans_badge"`
	FansMedal   struct {
		Show  bool `json:"show"`
		Wear  bool `json:"wear"`
		Medal struct {
			UID              int    `json:"uid"`
			TargetID         int    `json:"target_id"`
			MedalID          int    `json:"medal_id"`
			Level            int    `json:"level"`
			MedalName        string `json:"medal_name"`
			MedalColor       int    `json:"medal_color"`
			Intimacy         int    `json:"intimacy"`
			NextIntimacy     int    `json:"next_intimacy"`
			DayLimit         int    `json:"day_limit"`
			MedalColorStart  int    `json:"medal_color_start"`
			MedalColorEnd    int    `json:"medal_color_end"`
			MedalColorBorder int    `json:"medal_color_border"`
			IsLighted        int    `json:"is_lighted"`
			LightStatus      int    `json:"light_status"`
			WearingStatus    int    `json:"wearing_status"`
			Score            int    `json:"score"`
		} `json:"medal"`
	} `json:"fans_medal"`
	Official struct {
		Role  int    `json:"role"`
		Title string `json:"title"`
		Desc  string `json:"desc"`
		Type  int    `json:"type"`
	} `json:"official"`
	Vip struct {
		Type       int   `json:"type"`
		Status     int   `json:"status"`
		DueDate    int64 `json:"due_date"`
		VipPayType int   `json:"vip_pay_type"`
		ThemeType  int   `json:"theme_type"`
		Label      struct {
			Path                  string `json:"path"`
			Text                  string `json:"text"`
			LabelTheme            string `json:"label_theme"`
			TextColor             string `json:"text_color"`
			BgStyle               int    `json:"bg_style"`
			BgColor               string `json:"bg_color"`
			BorderColor           string `json:"border_color"`
			UseImgLabel           bool   `json:"use_img_label"`
			ImgLabelURIHans       string `json:"img_label_uri_hans"`
			ImgLabelURIHant       string `json:"img_label_uri_hant"`
			ImgLabelURIHansStatic string `json:"img_label_uri_hans_static"`
			ImgLabelURIHantStatic string `json:"img_label_uri_hant_static"`
		} `json:"label"`
		AvatarSubscript    int    `json:"avatar_subscript"`
		NicknameColor      string `json:"nickname_color"`
		Role               int    `json:"role"`
		AvatarSubscriptURL string `json:"avatar_subscript_url"`
		TvVipStatus        int    `json:"tv_vip_status"`
		TvVipPayType       int    `json:"tv_vip_pay_type"`
	} `json:"vip"`
	Pendant struct {
		Pid               int    `json:"pid"`
		Name              string `json:"name"`
		Image             string `json:"image"`
		Expire            int    `json:"expire"`
		ImageEnhance      string `json:"image_enhance"`
		ImageEnhanceFrame string `json:"image_enhance_frame"`
	} `json:"pendant"`
	Nameplate struct {
		Nid        int    `json:"nid"`
		Name       string `json:"name"`
		Image      string `json:"image"`
		ImageSmall string `json:"image_small"`
		Level      string `json:"level"`
		Condition  string `json:"condition"`
	} `json:"nameplate"`
	UserHonourInfo struct {
		Mid    int   `json:"mid"`
		Colour any   `json:"colour"`
		Tags   []any `json:"tags"`
	} `json:"user_honour_info"`
	IsFollowed bool   `json:"is_followed"`
	TopPhoto   string `json:"top_photo"`
	Theme      struct {
	} `json:"theme"`
	SysNotice struct {
	} `json:"sys_notice"`
	LiveRoom struct {
		RoomStatus    int    `json:"roomStatus"`
		LiveStatus    int    `json:"liveStatus"`
		URL           string `json:"url"`
		Title         string `json:"title"`
		Cover         string `json:"cover"`
		Roomid        int    `json:"roomid"`
		RoundStatus   int    `json:"roundStatus"`
		BroadcastType int    `json:"broadcast_type"`
		WatchedShow   struct {
			Switch       bool   `json:"switch"`
			Num          int    `json:"num"`
			TextSmall    string `json:"text_small"`
			TextLarge    string `json:"text_large"`
			Icon         string `json:"icon"`
			IconLocation string `json:"icon_location"`
			IconWeb      string `json:"icon_web"`
		} `json:"watched_show"`
	} `json:"live_room"`
	Birthday string `json:"birthday"`
	School   struct {
		Name string `json:"name"`
	} `json:"school"`
	Profession struct {
		Name       string `json:"name"`
		Department string `json:"department"`
		Title      string `json:"title"`
		IsShow     int    `json:"is_show"`
	} `json:"profession"`
	Tags   any `json:"tags"`
	Series struct {
		UserUpgradeStatus int  `json:"user_upgrade_status"`
		ShowUpgradeWindow bool `json:"show_upgrade_window"`
	} `json:"series"`
	IsSeniorMember int  `json:"is_senior_member"`
	McnInfo        any  `json:"mcn_info"`
	GaiaResType    int  `json:"gaia_res_type"`
	GaiaData       any  `json:"gaia_data"`
	IsRisk         bool `json:"is_risk"`
	Elec           struct {
		ShowInfo struct {
			Show    bool   `json:"show"`
			State   int    `json:"state"`
			Title   string `json:"title"`
			Icon    string `json:"icon"`
			JumpURL string `json:"jump_url"`
		} `json:"show_info"`
	} `json:"elec"`
	Contract struct {
		IsDisplay       bool `json:"is_display"`
		IsFollowDisplay bool `json:"is_follow_display"`
	} `json:"contract"`
}

type CoinLog struct {
	List []struct {
		Time   string  `json:"time"`
		Delta  float64 `json:"delta"`
		Reason string  `json:"reason"`
	} `json:"list"`
	Count int `json:"count"`
}

type LiveRoom struct {
	Info struct {
		UID            int    `json:"uid"`
		Uname          string `json:"uname"`
		Face           string `json:"face"`
		OfficialVerify struct {
			Type int    `json:"type"`
			Desc string `json:"desc"`
		} `json:"official_verify"`
		Gender int `json:"gender"`
	} `json:"info"`
	Exp struct {
		MasterLevel struct {
			Level   int   `json:"level"`
			Color   int   `json:"color"`
			Current []int `json:"current"`
			Next    []int `json:"next"`
		} `json:"master_level"`
	} `json:"exp"`
	FollowerNum  int    `json:"follower_num"`
	RoomID       int    `json:"room_id"`
	MedalName    string `json:"medal_name"`
	GloryCount   int    `json:"glory_count"`
	Pendant      string `json:"pendant"`
	LinkGroupNum int    `json:"link_group_num"`
	RoomNews     struct {
		Content   string `json:"content"`
		Ctime     string `json:"ctime"`
		CtimeText string `json:"ctime_text"`
	} `json:"room_news"`
}

type MedalWall struct {
	List []struct {
		MedalInfo struct {
			TargetID         int    `json:"target_id"`
			Level            int    `json:"level"`
			MedalName        string `json:"medal_name"`
			MedalColorStart  int    `json:"medal_color_start"`
			MedalColorEnd    int    `json:"medal_color_end"`
			MedalColorBorder int    `json:"medal_color_border"`
			GuardLevel       int    `json:"guard_level"`
			WearingStatus    int    `json:"wearing_status"`
			MedalID          int    `json:"medal_id"`
			Intimacy         int    `json:"intimacy"`
			NextIntimacy     int    `json:"next_intimacy"`
			TodayFeed        int    `json:"today_feed"`
			DayLimit         int    `json:"day_limit"`
			GuardIcon        string `json:"guard_icon"`
			HonorIcon        string `json:"honor_icon"`
		} `json:"medal_info"`
		TargetName string `json:"target_name"`
		TargetIcon string `json:"target_icon"`
		Link       string `json:"link"`
		LiveStatus int    `json:"live_status"`
		Official   int    `json:"official"`
	} `json:"list"`
	Count           int    `json:"count"`
	CloseSpaceMedal int    `json:"close_space_medal"`
	OnlyShowWearing int    `json:"only_show_wearing"`
	Name            string `json:"name"`
	Icon            string `json:"icon"`
	UID             int    `json:"uid"`
	Level           int    `json:"level"`
}

type ExpRewardStatus struct {
	Login        bool `json:"login"`
	Watch        bool `json:"watch"`
	Coins        int  `json:"coins"`
	Share        bool `json:"share"`
	Email        bool `json:"email"`
	Tel          bool `json:"tel"`
	SafeQuestion bool `json:"safe_question"`
	IdentifyCard bool `json:"identify_card"`
}

type RegionRank struct {
	Aid         string `json:"aid"`
	Bvid        string `json:"bvid"`
	Typename    string `json:"typename"`
	Title       string `json:"title"`
	Subtitle    string `json:"subtitle"`
	Play        int    `json:"play"`
	Review      int    `json:"review"`
	VideoReview int    `json:"video_review"`
	Favorites   int    `json:"favorites"`
	Mid         int    `json:"mid"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Create      string `json:"create"`
	Pic         string `json:"pic"`
	Coins       int    `json:"coins"`
	Duration    string `json:"duration"`
	Badgepay    bool   `json:"badgepay"`
	Pts         int    `json:"pts"`
	Rights      struct {
		Bp            int `json:"bp"`
		Elec          int `json:"elec"`
		Download      int `json:"download"`
		Movie         int `json:"movie"`
		Pay           int `json:"pay"`
		Hd5           int `json:"hd5"`
		NoReprint     int `json:"no_reprint"`
		Autoplay      int `json:"autoplay"`
		UgcPay        int `json:"ugc_pay"`
		IsCooperation int `json:"is_cooperation"`
		UgcPayPreview int `json:"ugc_pay_preview"`
		NoBackground  int `json:"no_background"`
		ArcPay        int `json:"arc_pay"`
		PayFreeWatch  int `json:"pay_free_watch"`
	} `json:"rights"`
	RedirectURL string `json:"redirect_url"`
}

type VideoDetails struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    struct {
		Bvid      string `json:"bvid"`
		Aid       int    `json:"aid"`
		Videos    int    `json:"videos"`
		Tid       int    `json:"tid"`
		Tname     string `json:"tname"`
		Copyright int    `json:"copyright"`
		Pic       string `json:"pic"`
		Title     string `json:"title"`
		Pubdate   int    `json:"pubdate"`
		Ctime     int    `json:"ctime"`
		Desc      string `json:"desc"`
		DescV2    []struct {
			RawText string `json:"raw_text"`
			Type    int    `json:"type"`
			BizID   int    `json:"biz_id"`
		} `json:"desc_v2"`
		State     int `json:"state"`
		Duration  int `json:"duration"`
		MissionID int `json:"mission_id"`
		Rights    struct {
			Bp            int `json:"bp"`
			Elec          int `json:"elec"`
			Download      int `json:"download"`
			Movie         int `json:"movie"`
			Pay           int `json:"pay"`
			Hd5           int `json:"hd5"`
			NoReprint     int `json:"no_reprint"`
			Autoplay      int `json:"autoplay"`
			UgcPay        int `json:"ugc_pay"`
			IsCooperation int `json:"is_cooperation"`
			UgcPayPreview int `json:"ugc_pay_preview"`
			NoBackground  int `json:"no_background"`
			CleanMode     int `json:"clean_mode"`
			IsSteinGate   int `json:"is_stein_gate"`
			Is360         int `json:"is_360"`
			NoShare       int `json:"no_share"`
			ArcPay        int `json:"arc_pay"`
			FreeWatch     int `json:"free_watch"`
		} `json:"rights"`
		Owner struct {
			Mid  int    `json:"mid"`
			Name string `json:"name"`
			Face string `json:"face"`
		} `json:"owner"`
		Stat struct {
			Aid        int    `json:"aid"`
			View       int    `json:"view"`
			Danmaku    int    `json:"danmaku"`
			Reply      int    `json:"reply"`
			Favorite   int    `json:"favorite"`
			Coin       int    `json:"coin"`
			Share      int    `json:"share"`
			NowRank    int    `json:"now_rank"`
			HisRank    int    `json:"his_rank"`
			Like       int    `json:"like"`
			Dislike    int    `json:"dislike"`
			Evaluation string `json:"evaluation"`
			ArgueMsg   string `json:"argue_msg"`
		} `json:"stat"`
		Dynamic   string `json:"dynamic"`
		Cid       int    `json:"cid"`
		Dimension struct {
			Width  int `json:"width"`
			Height int `json:"height"`
			Rotate int `json:"rotate"`
		} `json:"dimension"`
		Premiere           any  `json:"premiere"`
		TeenageMode        int  `json:"teenage_mode"`
		IsChargeableSeason bool `json:"is_chargeable_season"`
		IsStory            bool `json:"is_story"`
		NoCache            bool `json:"no_cache"`
		Pages              []struct {
			Cid       int    `json:"cid"`
			Page      int    `json:"page"`
			From      string `json:"from"`
			Part      string `json:"part"`
			Duration  int    `json:"duration"`
			Vid       string `json:"vid"`
			Weblink   string `json:"weblink"`
			Dimension struct {
				Width  int `json:"width"`
				Height int `json:"height"`
				Rotate int `json:"rotate"`
			} `json:"dimension"`
		} `json:"pages"`
		Subtitle struct {
			AllowSubmit bool  `json:"allow_submit"`
			List        []any `json:"list"`
		} `json:"subtitle"`
		Staff []struct {
			Mid   int    `json:"mid"`
			Title string `json:"title"`
			Name  string `json:"name"`
			Face  string `json:"face"`
			Vip   struct {
				Type       int   `json:"type"`
				Status     int   `json:"status"`
				DueDate    int64 `json:"due_date"`
				VipPayType int   `json:"vip_pay_type"`
				ThemeType  int   `json:"theme_type"`
				Label      struct {
					Path                  string `json:"path"`
					Text                  string `json:"text"`
					LabelTheme            string `json:"label_theme"`
					TextColor             string `json:"text_color"`
					BgStyle               int    `json:"bg_style"`
					BgColor               string `json:"bg_color"`
					BorderColor           string `json:"border_color"`
					UseImgLabel           bool   `json:"use_img_label"`
					ImgLabelURIHans       string `json:"img_label_uri_hans"`
					ImgLabelURIHant       string `json:"img_label_uri_hant"`
					ImgLabelURIHansStatic string `json:"img_label_uri_hans_static"`
					ImgLabelURIHantStatic string `json:"img_label_uri_hant_static"`
				} `json:"label"`
				AvatarSubscript    int    `json:"avatar_subscript"`
				NicknameColor      string `json:"nickname_color"`
				Role               int    `json:"role"`
				AvatarSubscriptURL string `json:"avatar_subscript_url"`
				TvVipStatus        int    `json:"tv_vip_status"`
				TvVipPayType       int    `json:"tv_vip_pay_type"`
			} `json:"vip"`
			Official struct {
				Role  int    `json:"role"`
				Title string `json:"title"`
				Desc  string `json:"desc"`
				Type  int    `json:"type"`
			} `json:"official"`
			Follower   int `json:"follower"`
			LabelStyle int `json:"label_style"`
		} `json:"staff"`
		IsSeasonDisplay bool `json:"is_season_display"`
		UserGarb        struct {
			URLImageAniCut string `json:"url_image_ani_cut"`
		} `json:"user_garb"`
		HonorReply struct {
			Honor []struct {
				Aid                int    `json:"aid"`
				Type               int    `json:"type"`
				Desc               string `json:"desc"`
				WeeklyRecommendNum int    `json:"weekly_recommend_num"`
			} `json:"honor"`
		} `json:"honor_reply"`
		LikeIcon string `json:"like_icon"`
	} `json:"data"`
}

type CoinInfo struct {
	Money float64 `json:"money"`
}

type DonateCoinInfo struct {
	Multiply int `json:"multiply"`
}

type LiveWallet struct {
	Gold          int    `json:"gold"`
	Silver        int    `json:"silver"`
	Bp            string `json:"bp"`
	Metal         int    `json:"metal"`
	NeedUseNewBp  bool   `json:"need_use_new_bp"`
	IosBp         int    `json:"ios_bp"`
	CommonBp      int    `json:"common_bp"`
	NewBp         string `json:"new_bp"`
	Bp2GoldAmount int    `json:"bp_2_gold_amount"`
}

type LiveCheckIn struct {
	Text        string `json:"text"`
	SpecialText string `json:"specialText"`
	AllDays     int    `json:"allDays"`
	HadSignDays int    `json:"hadSignDays"`
	IsBonusDay  int    `json:"isBonusDay"`
}

type GiftList struct {
	List []struct {
		BagID       int    `json:"bag_id"`
		GiftID      int    `json:"gift_id"`
		GiftName    string `json:"gift_name"`
		GiftNum     int    `json:"gift_num"`
		GiftType    int    `json:"gift_type"`
		ExpireAt    int64  `json:"expire_at"`
		CornerMark  string `json:"corner_mark"`
		CornerColor string `json:"corner_color"`
		CountMap    []struct {
			Num  int    `json:"num"`
			Text string `json:"text"`
		} `json:"count_map"`
		BindRoomid   int    `json:"bind_roomid"`
		BindRoomText string `json:"bind_room_text"`
		Type         int    `json:"type"`
		CardImage    string `json:"card_image"`
		CardGif      string `json:"card_gif"`
		CardID       int    `json:"card_id"`
		CardRecordID int    `json:"card_record_id"`
		IsShowSend   bool   `json:"is_show_send"`
		ExpireText   string `json:"expire_text"`
	} `json:"list"`
	Time int `json:"time"`
}

type ChargeInfo struct {
	BpWallet struct {
		DefaultBp         int    `json:"default_bp"`
		IosBp             int    `json:"ios_bp"`
		CouponBalance     int    `json:"coupon_balance"`
		AvailableBp       int    `json:"available_bp"`
		UnavailableBp     int    `json:"unavailable_bp"`
		TotalBp           int    `json:"total_bp"`
		Mid               int    `json:"mid"`
		UnavailableReason string `json:"unavailable_reason"`
		IsBpRemainsPrior  bool   `json:"is_bp_remains_prior"`
		Tip               string `json:"tip"`
	} `json:"bp_wallet"`
	BatteryList []struct {
		Title       string `json:"title"`
		Groups      int    `json:"groups"`
		IsCustomize int    `json:"is_customize"`
		ElecNum     int    `json:"elec_num"`
		MinElec     int    `json:"min_elec"`
		MaxElec     int    `json:"max_elec"`
		IsChecked   int    `json:"is_checked"`
		Seq         int    `json:"seq"`
		BpNum       string `json:"bp_num"`
		MinBp       string `json:"min_bp"`
		MaxBp       string `json:"max_bp"`
	} `json:"battery_list"`
}

type ChargeResponse struct {
	Mid     int    `json:"mid"`
	UpMid   int    `json:"up_mid"`
	OrderNo string `json:"order_no"`
	BpNum   int    `json:"bp_num"`
	Exp     string `json:"exp"`
	Status  int    `json:"status"`
	Msg     string `json:"msg"`
}
