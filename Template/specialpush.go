package Template

type Specialpush struct {
	Status Status `json:"status"`
	Data   struct {
		NovelSpecialSubject string `json:"novelSpecialSubject"`
		HomePush            []struct {
			ImgURL    string `json:"imgUrl"`
			Link      string `json:"link"`
			Type      string `json:"type"`
			Desc      string `json:"desc"`
			BeginDate string `json:"beginDate"`
			EndDate   string `json:"endDate"`
			AddDate   string `json:"addDate"`
		} `json:"homePush"`
		VipPush []struct {
			ImgURL    string `json:"imgUrl"`
			Link      string `json:"link"`
			Type      string `json:"type"`
			Desc      string `json:"desc"`
			BeginDate string `json:"beginDate"`
			EndDate   string `json:"endDate"`
			AddDate   string `json:"addDate"`
		} `json:"vipPush"`
		HomeSecondPush []struct {
			ImgURL    string      `json:"imgUrl"`
			Link      string      `json:"link"`
			Type      string      `json:"type"`
			BeginDate string      `json:"beginDate"`
			EndDate   string      `json:"endDate"`
			Desc      string      `json:"desc"`
			Area      interface{} `json:"area"`
			Expand    struct {
			} `json:"expand"`
			EntityID int    `json:"entityId"`
			AddDate  string `json:"addDate"`
		} `json:"homeSecondPush"`
		PocketLoopPush  []interface{} `json:"pocketLoopPush"`
		PocketCoverPush []struct {
			ImgURL     string `json:"imgUrl"`
			Link       string `json:"link"`
			Type       string `json:"type"`
			Desc       string `json:"desc"`
			EndDate    string `json:"endDate"`
			EntityName string `json:"entityName"`
		} `json:"pocketCoverPush"`
		ChatNovelPush []struct {
			ImgURL    string      `json:"imgUrl"`
			Link      string      `json:"link"`
			Type      string      `json:"type"`
			BeginDate string      `json:"beginDate"`
			EndDate   string      `json:"endDate"`
			Desc      string      `json:"desc"`
			Area      interface{} `json:"area"`
			Expand    struct {
			} `json:"expand"`
			EntityID int    `json:"entityId"`
			AddDate  string `json:"addDate"`
		} `json:"chatNovelPush"`
		WelfarePush []struct {
			ImgURL    string `json:"imgUrl"`
			Link      string `json:"link"`
			Type      string `json:"type"`
			Desc      string `json:"desc"`
			BeginDate string `json:"beginDate"`
			EndDate   string `json:"endDate"`
			AddDate   string `json:"addDate"`
		} `json:"welfarePush"`
		AnnouncementPush []interface{} `json:"announcementPush"`
		PopupPush        []struct {
			ImgURL    string      `json:"imgUrl"`
			Link      string      `json:"link"`
			Type      string      `json:"type"`
			BeginDate string      `json:"beginDate"`
			EndDate   string      `json:"endDate"`
			Desc      string      `json:"desc"`
			Area      interface{} `json:"area"`
			Expand    struct {
			} `json:"expand"`
			EntityID int    `json:"entityId"`
			AddDate  string `json:"addDate"`
		} `json:"popupPush"`
		VersionPush        []interface{} `json:"versionPush"`
		MinipChatNovelPush []struct {
			ImgURL    string      `json:"imgUrl"`
			Link      string      `json:"link"`
			Type      string      `json:"type"`
			BeginDate string      `json:"beginDate"`
			EndDate   string      `json:"endDate"`
			Desc      string      `json:"desc"`
			Area      interface{} `json:"area"`
			Expand    struct {
			} `json:"expand"`
			EntityID int    `json:"entityId"`
			AddDate  string `json:"addDate"`
		} `json:"minipChatNovelPush"`
		FanNovelHomePush []struct {
			ImgURL    string      `json:"imgUrl"`
			Link      string      `json:"link"`
			Type      string      `json:"type"`
			BeginDate string      `json:"beginDate"`
			EndDate   string      `json:"endDate"`
			Desc      string      `json:"desc"`
			Area      interface{} `json:"area"`
			Expand    struct {
			} `json:"expand"`
			EntityID int    `json:"entityId"`
			AddDate  string `json:"addDate"`
		} `json:"fanNovelHomePush"`
		HomeFloatPush       []interface{} `json:"homeFloatPush"`
		FirstChargePush     []interface{} `json:"firstChargePush"`
		ChatnovelbannerPush []struct {
			ImgURL    string      `json:"imgUrl"`
			Link      string      `json:"link"`
			Type      string      `json:"type"`
			BeginDate string      `json:"beginDate"`
			EndDate   string      `json:"endDate"`
			Desc      string      `json:"desc"`
			Area      interface{} `json:"area"`
			Expand    struct {
			} `json:"expand"`
			EntityID int    `json:"entityId"`
			AddDate  string `json:"addDate"`
		} `json:"chatnovelbannerPush"`
		ComicCustomizePush struct {
			Title   string `json:"title"`
			Content string `json:"content"`
		} `json:"comicCustomizePush"`
		CompositionbannerPush []struct {
			ImgURL    string `json:"imgUrl"`
			Link      string `json:"link"`
			Type      string `json:"type"`
			Desc      string `json:"desc"`
			BeginDate string `json:"beginDate"`
			EndDate   string `json:"endDate"`
			Title     string `json:"title"`
			Status    int    `json:"status"`
		} `json:"compositionbannerPush"`
		RoleBanner                 []interface{} `json:"roleBanner"`
		NovelChatNovelSecondBanner []struct {
			ImgURL    string      `json:"imgUrl"`
			Link      string      `json:"link"`
			Type      string      `json:"type"`
			BeginDate string      `json:"beginDate"`
			EndDate   string      `json:"endDate"`
			Desc      string      `json:"desc"`
			Area      interface{} `json:"area"`
			Expand    struct {
			} `json:"expand"`
			EntityID int    `json:"entityId"`
			AddDate  string `json:"addDate"`
		} `json:"novelChatNovelSecondBanner"`
		BottomButton       []interface{} `json:"bottomButton"`
		HomeBroadcastPush  []interface{} `json:"homeBroadcastPush"`
		HomeBottomPush     []interface{} `json:"homeBottomPush"`
		BookMarkPush       []interface{} `json:"bookMarkPush"`
		HomeTopRreshenPush []struct {
			ImgURL    string      `json:"imgUrl"`
			Link      string      `json:"link"`
			Type      string      `json:"type"`
			BeginDate string      `json:"beginDate"`
			EndDate   string      `json:"endDate"`
			Desc      string      `json:"desc"`
			Area      interface{} `json:"area"`
			Expand    struct {
			} `json:"expand"`
			EntityID int    `json:"entityId"`
			AddDate  string `json:"addDate"`
		} `json:"homeTopRreshenPush"`
		HomeBottomTabPush []interface{} `json:"homeBottomTabPush"`
		MyBanner          []struct {
			ImgURL    string      `json:"imgUrl"`
			Link      string      `json:"link"`
			Type      string      `json:"type"`
			BeginDate string      `json:"beginDate"`
			EndDate   string      `json:"endDate"`
			Desc      string      `json:"desc"`
			Area      interface{} `json:"area"`
			Expand    struct {
			} `json:"expand"`
			EntityID int    `json:"entityId"`
			AddDate  string `json:"addDate"`
		} `json:"myBanner"`
		InteractNovel []struct {
			ImgURL  string `json:"imgUrl"`
			Link    string `json:"link"`
			Type    string `json:"type"`
			Desc    string `json:"desc"`
			EndDate string `json:"endDate"`
		} `json:"interactNovel"`
		RecomposeIP     []interface{} `json:"recomposeIp"`
		Mutualnovel     []interface{} `json:"mutualnovel"`
		MutualnovelMore []struct {
			ImgURL    string      `json:"imgUrl"`
			Link      string      `json:"link"`
			Type      string      `json:"type"`
			BeginDate string      `json:"beginDate"`
			EndDate   string      `json:"endDate"`
			Desc      string      `json:"desc"`
			Area      interface{} `json:"area"`
			Expand    struct {
			} `json:"expand"`
			EntityID int    `json:"entityId"`
			AddDate  string `json:"addDate"`
			Name     string `json:"name"`
		} `json:"mutualnovelMore"`
		ChatNovelHotbanners []struct {
			ImgURL    string      `json:"imgUrl"`
			Link      string      `json:"link"`
			Type      string      `json:"type"`
			BeginDate string      `json:"beginDate"`
			EndDate   string      `json:"endDate"`
			Desc      string      `json:"desc"`
			Area      interface{} `json:"area"`
			Expand    struct {
			} `json:"expand"`
			EntityID int    `json:"entityId"`
			AddDate  string `json:"addDate"`
			Entity   struct {
				AuthorID       int     `json:"authorId"`
				LastUpdateTime string  `json:"lastUpdateTime"`
				MarkCount      int     `json:"markCount"`
				NovelCover     string  `json:"novelCover"`
				BgBanner       string  `json:"bgBanner"`
				NovelID        int     `json:"novelId"`
				NovelName      string  `json:"novelName"`
				Point          float64 `json:"point"`
				IsFinish       bool    `json:"isFinish"`
				AuthorName     string  `json:"authorName"`
				CharCount      int     `json:"charCount"`
				ViewTimes      int     `json:"viewTimes"`
				TypeID         int     `json:"typeId"`
				AllowDown      bool    `json:"allowDown"`
				AddTime        string  `json:"addTime"`
				IsSensitive    bool    `json:"isSensitive"`
				SignStatus     string  `json:"signStatus"`
				CategoryID     int     `json:"categoryId"`
				Expand         struct {
				} `json:"expand"`
			} `json:"entity,omitempty"`
		} `json:"chatNovelHotbanners"`
		ChatNovelHotEditor []struct {
			ImgURL    string      `json:"imgUrl"`
			Link      string      `json:"link"`
			Type      string      `json:"type"`
			BeginDate string      `json:"beginDate"`
			EndDate   string      `json:"endDate"`
			Desc      string      `json:"desc"`
			Area      interface{} `json:"area"`
			Expand    struct {
			} `json:"expand"`
			EntityID int    `json:"entityId"`
			AddDate  string `json:"addDate"`
			Entity   struct {
				AuthorID       int     `json:"authorId"`
				LastUpdateTime string  `json:"lastUpdateTime"`
				MarkCount      int     `json:"markCount"`
				NovelCover     string  `json:"novelCover"`
				BgBanner       string  `json:"bgBanner"`
				NovelID        int     `json:"novelId"`
				NovelName      string  `json:"novelName"`
				Point          float64 `json:"point"`
				IsFinish       bool    `json:"isFinish"`
				AuthorName     string  `json:"authorName"`
				CharCount      int     `json:"charCount"`
				ViewTimes      int     `json:"viewTimes"`
				TypeID         int     `json:"typeId"`
				AllowDown      bool    `json:"allowDown"`
				AddTime        string  `json:"addTime"`
				IsSensitive    bool    `json:"isSensitive"`
				SignStatus     string  `json:"signStatus"`
				CategoryID     int     `json:"categoryId"`
				Expand         struct {
				} `json:"expand"`
			} `json:"entity"`
		} `json:"chatNovelHotEditor"`
		ChatNovelHotDaily []struct {
			ImgURL    string      `json:"imgUrl"`
			Link      string      `json:"link"`
			Type      string      `json:"type"`
			BeginDate string      `json:"beginDate"`
			EndDate   string      `json:"endDate"`
			Desc      string      `json:"desc"`
			Area      interface{} `json:"area"`
			Expand    struct {
			} `json:"expand"`
			EntityID int    `json:"entityId"`
			AddDate  string `json:"addDate"`
			Entity   struct {
				AuthorID       int     `json:"authorId"`
				LastUpdateTime string  `json:"lastUpdateTime"`
				MarkCount      int     `json:"markCount"`
				NovelCover     string  `json:"novelCover"`
				BgBanner       string  `json:"bgBanner"`
				NovelID        int     `json:"novelId"`
				NovelName      string  `json:"novelName"`
				Point          float64 `json:"point"`
				IsFinish       bool    `json:"isFinish"`
				AuthorName     string  `json:"authorName"`
				CharCount      int     `json:"charCount"`
				ViewTimes      int     `json:"viewTimes"`
				TypeID         int     `json:"typeId"`
				AllowDown      bool    `json:"allowDown"`
				AddTime        string  `json:"addTime"`
				IsSensitive    bool    `json:"isSensitive"`
				SignStatus     string  `json:"signStatus"`
				CategoryID     int     `json:"categoryId"`
				Expand         struct {
				} `json:"expand"`
			} `json:"entity"`
		} `json:"chatNovelHotDaily"`
		ChatNovelHotType    []interface{} `json:"chatNovelHotType"`
		ChatNovelHotLooking []struct {
			ImgURL    string      `json:"imgUrl"`
			Link      string      `json:"link"`
			Type      string      `json:"type"`
			BeginDate string      `json:"beginDate"`
			EndDate   string      `json:"endDate"`
			Desc      string      `json:"desc"`
			Area      interface{} `json:"area"`
			Expand    struct {
			} `json:"expand"`
			EntityID int    `json:"entityId"`
			AddDate  string `json:"addDate"`
			Entity   struct {
				AuthorID       int     `json:"authorId"`
				LastUpdateTime string  `json:"lastUpdateTime"`
				MarkCount      int     `json:"markCount"`
				NovelCover     string  `json:"novelCover"`
				BgBanner       string  `json:"bgBanner"`
				NovelID        int     `json:"novelId"`
				NovelName      string  `json:"novelName"`
				Point          float64 `json:"point"`
				IsFinish       bool    `json:"isFinish"`
				AuthorName     string  `json:"authorName"`
				CharCount      int     `json:"charCount"`
				ViewTimes      int     `json:"viewTimes"`
				TypeID         int     `json:"typeId"`
				AllowDown      bool    `json:"allowDown"`
				AddTime        string  `json:"addTime"`
				IsSensitive    bool    `json:"isSensitive"`
				SignStatus     string  `json:"signStatus"`
				CategoryID     int     `json:"categoryId"`
				Expand         struct {
				} `json:"expand"`
			} `json:"entity"`
		} `json:"chatNovelHotLooking"`
		ChatNovelNewHot []struct {
			ImgURL    string      `json:"imgUrl"`
			Link      string      `json:"link"`
			Type      string      `json:"type"`
			BeginDate string      `json:"beginDate"`
			EndDate   string      `json:"endDate"`
			Desc      string      `json:"desc"`
			Area      interface{} `json:"area"`
			Expand    struct {
			} `json:"expand"`
			EntityID int    `json:"entityId"`
			AddDate  string `json:"addDate"`
			Entity   struct {
				AuthorID       int     `json:"authorId"`
				LastUpdateTime string  `json:"lastUpdateTime"`
				MarkCount      int     `json:"markCount"`
				NovelCover     string  `json:"novelCover"`
				BgBanner       string  `json:"bgBanner"`
				NovelID        int     `json:"novelId"`
				NovelName      string  `json:"novelName"`
				Point          float64 `json:"point"`
				IsFinish       bool    `json:"isFinish"`
				AuthorName     string  `json:"authorName"`
				CharCount      int     `json:"charCount"`
				ViewTimes      int     `json:"viewTimes"`
				TypeID         int     `json:"typeId"`
				AllowDown      bool    `json:"allowDown"`
				AddTime        string  `json:"addTime"`
				IsSensitive    bool    `json:"isSensitive"`
				SignStatus     string  `json:"signStatus"`
				CategoryID     int     `json:"categoryId"`
				Expand         struct {
				} `json:"expand"`
			} `json:"entity"`
		} `json:"chatNovelNewHot"`
		ChatNovelNewPotential []struct {
			ImgURL    string      `json:"imgUrl"`
			Link      string      `json:"link"`
			Type      string      `json:"type"`
			BeginDate string      `json:"beginDate"`
			EndDate   string      `json:"endDate"`
			Desc      string      `json:"desc"`
			Area      interface{} `json:"area"`
			Expand    struct {
			} `json:"expand"`
			EntityID int    `json:"entityId"`
			AddDate  string `json:"addDate"`
			Entity   struct {
				AuthorID       int     `json:"authorId"`
				LastUpdateTime string  `json:"lastUpdateTime"`
				MarkCount      int     `json:"markCount"`
				NovelCover     string  `json:"novelCover"`
				BgBanner       string  `json:"bgBanner"`
				NovelID        int     `json:"novelId"`
				NovelName      string  `json:"novelName"`
				Point          float64 `json:"point"`
				IsFinish       bool    `json:"isFinish"`
				AuthorName     string  `json:"authorName"`
				CharCount      int     `json:"charCount"`
				ViewTimes      int     `json:"viewTimes"`
				TypeID         int     `json:"typeId"`
				AllowDown      bool    `json:"allowDown"`
				AddTime        string  `json:"addTime"`
				IsSensitive    bool    `json:"isSensitive"`
				SignStatus     string  `json:"signStatus"`
				CategoryID     int     `json:"categoryId"`
				Expand         struct {
				} `json:"expand"`
			} `json:"entity"`
		} `json:"chatNovelNewPotential"`
		ChatNovelNewWin []struct {
			ImgURL    string      `json:"imgUrl"`
			Link      string      `json:"link"`
			Type      string      `json:"type"`
			BeginDate string      `json:"beginDate"`
			EndDate   string      `json:"endDate"`
			Desc      string      `json:"desc"`
			Area      interface{} `json:"area"`
			Expand    struct {
			} `json:"expand"`
			EntityID int    `json:"entityId"`
			AddDate  string `json:"addDate"`
			Entity   struct {
				AuthorID       int     `json:"authorId"`
				LastUpdateTime string  `json:"lastUpdateTime"`
				MarkCount      int     `json:"markCount"`
				NovelCover     string  `json:"novelCover"`
				BgBanner       string  `json:"bgBanner"`
				NovelID        int     `json:"novelId"`
				NovelName      string  `json:"novelName"`
				Point          float64 `json:"point"`
				IsFinish       bool    `json:"isFinish"`
				AuthorName     string  `json:"authorName"`
				CharCount      int     `json:"charCount"`
				ViewTimes      int     `json:"viewTimes"`
				TypeID         int     `json:"typeId"`
				AllowDown      bool    `json:"allowDown"`
				AddTime        string  `json:"addTime"`
				IsSensitive    bool    `json:"isSensitive"`
				SignStatus     string  `json:"signStatus"`
				CategoryID     int     `json:"categoryId"`
				Expand         struct {
				} `json:"expand"`
			} `json:"entity"`
		} `json:"chatNovelNewWin"`
		ChatNovelFinishHighPoint []struct {
			ImgURL    string      `json:"imgUrl"`
			Link      string      `json:"link"`
			Type      string      `json:"type"`
			BeginDate string      `json:"beginDate"`
			EndDate   string      `json:"endDate"`
			Desc      string      `json:"desc"`
			Area      interface{} `json:"area"`
			Expand    struct {
			} `json:"expand"`
			EntityID int    `json:"entityId"`
			AddDate  string `json:"addDate"`
			Entity   struct {
				AuthorID       int     `json:"authorId"`
				LastUpdateTime string  `json:"lastUpdateTime"`
				MarkCount      int     `json:"markCount"`
				NovelCover     string  `json:"novelCover"`
				BgBanner       string  `json:"bgBanner"`
				NovelID        int     `json:"novelId"`
				NovelName      string  `json:"novelName"`
				Point          float64 `json:"point"`
				IsFinish       bool    `json:"isFinish"`
				AuthorName     string  `json:"authorName"`
				CharCount      int     `json:"charCount"`
				ViewTimes      int     `json:"viewTimes"`
				TypeID         int     `json:"typeId"`
				AllowDown      bool    `json:"allowDown"`
				AddTime        string  `json:"addTime"`
				IsSensitive    bool    `json:"isSensitive"`
				SignStatus     string  `json:"signStatus"`
				CategoryID     int     `json:"categoryId"`
				Expand         struct {
				} `json:"expand"`
			} `json:"entity"`
		} `json:"chatNovelFinishHighPoint"`
		ChatNovelFinishGood []struct {
			ImgURL    string      `json:"imgUrl"`
			Link      string      `json:"link"`
			Type      string      `json:"type"`
			BeginDate string      `json:"beginDate"`
			EndDate   string      `json:"endDate"`
			Desc      string      `json:"desc"`
			Area      interface{} `json:"area"`
			Expand    struct {
			} `json:"expand"`
			EntityID int    `json:"entityId"`
			AddDate  string `json:"addDate"`
			Entity   struct {
				AuthorID       int     `json:"authorId"`
				LastUpdateTime string  `json:"lastUpdateTime"`
				MarkCount      int     `json:"markCount"`
				NovelCover     string  `json:"novelCover"`
				BgBanner       string  `json:"bgBanner"`
				NovelID        int     `json:"novelId"`
				NovelName      string  `json:"novelName"`
				Point          float64 `json:"point"`
				IsFinish       bool    `json:"isFinish"`
				AuthorName     string  `json:"authorName"`
				CharCount      int     `json:"charCount"`
				ViewTimes      int     `json:"viewTimes"`
				TypeID         int     `json:"typeId"`
				AllowDown      bool    `json:"allowDown"`
				AddTime        string  `json:"addTime"`
				IsSensitive    bool    `json:"isSensitive"`
				SignStatus     string  `json:"signStatus"`
				CategoryID     int     `json:"categoryId"`
				Expand         struct {
				} `json:"expand"`
			} `json:"entity"`
		} `json:"chatNovelFinishGood"`
		MerchPush []struct {
			ImgURL    string `json:"imgUrl"`
			Link      string `json:"link"`
			Type      string `json:"type"`
			BeginDate string `json:"beginDate"`
			EndDate   string `json:"endDate"`
			Desc      string `json:"desc"`
			Area      string `json:"area"`
			Expand    struct {
				Price string `json:"price"`
			} `json:"expand"`
			EntityID int    `json:"entityId"`
			AddDate  string `json:"addDate"`
		} `json:"merchPush"`
		EntityDetailPush []struct {
			ImgURL    string `json:"imgUrl"`
			Link      string `json:"link"`
			Type      string `json:"type"`
			BeginDate string `json:"beginDate"`
			EndDate   string `json:"endDate"`
			Desc      string `json:"desc"`
			Area      string `json:"area"`
			Expand    struct {
			} `json:"expand"`
			EntityID int    `json:"entityId"`
			AddDate  string `json:"addDate"`
		} `json:"entityDetailPush"`
		NewbiepackagePush  []interface{} `json:"newbiepackagePush"`
		EntityLastPagePush []struct {
			ImgURL    string `json:"imgUrl"`
			Link      string `json:"link"`
			Type      string `json:"type"`
			BeginDate string `json:"beginDate"`
			EndDate   string `json:"endDate"`
			Desc      string `json:"desc"`
			Area      string `json:"area"`
			Expand    struct {
			} `json:"expand"`
			EntityID int    `json:"entityId"`
			AddDate  string `json:"addDate"`
		} `json:"entityLastPagePush"`
		BigBrainPush []struct {
			ImgURL    string      `json:"imgUrl"`
			Link      string      `json:"link"`
			Type      string      `json:"type"`
			BeginDate string      `json:"beginDate"`
			EndDate   string      `json:"endDate"`
			Desc      string      `json:"desc"`
			Area      interface{} `json:"area"`
			Expand    struct {
			} `json:"expand"`
			EntityID int    `json:"entityId"`
			AddDate  string `json:"addDate"`
		} `json:"bigBrainPush"`
		SignPush        []interface{} `json:"signPush"`
		HomeHolidayPush []interface{} `json:"homeHolidayPush"`
		NewSignInPush   []struct {
			ImgURL    string      `json:"imgUrl"`
			Link      string      `json:"link"`
			Type      string      `json:"type"`
			BeginDate string      `json:"beginDate"`
			EndDate   string      `json:"endDate"`
			Desc      string      `json:"desc"`
			Area      interface{} `json:"area"`
			Expand    struct {
			} `json:"expand"`
			EntityID int    `json:"entityId"`
			AddDate  string `json:"addDate"`
			Entity   struct {
				AuthorID       int     `json:"authorId"`
				LastUpdateTime string  `json:"lastUpdateTime"`
				MarkCount      int     `json:"markCount"`
				NovelCover     string  `json:"novelCover"`
				BgBanner       string  `json:"bgBanner"`
				NovelID        int     `json:"novelId"`
				NovelName      string  `json:"novelName"`
				Point          float64 `json:"point"`
				IsFinish       bool    `json:"isFinish"`
				AuthorName     string  `json:"authorName"`
				CharCount      int     `json:"charCount"`
				ViewTimes      int     `json:"viewTimes"`
				TypeID         int     `json:"typeId"`
				AllowDown      bool    `json:"allowDown"`
				AddTime        string  `json:"addTime"`
				IsSensitive    bool    `json:"isSensitive"`
				SignStatus     string  `json:"signStatus"`
				CategoryID     int     `json:"categoryId"`
				Expand         struct {
				} `json:"expand"`
			} `json:"entity"`
		} `json:"newSignInPush"`
		Contest2Push []struct {
			ImgURL    string      `json:"imgUrl"`
			Link      string      `json:"link"`
			Type      string      `json:"type"`
			BeginDate string      `json:"beginDate"`
			EndDate   string      `json:"endDate"`
			Desc      string      `json:"desc"`
			Area      interface{} `json:"area"`
			Expand    struct {
			} `json:"expand"`
			EntityID int    `json:"entityId"`
			AddDate  string `json:"addDate"`
			Entitys  []struct {
				AuthorID       int     `json:"authorId"`
				LastUpdateTime string  `json:"lastUpdateTime"`
				MarkCount      int     `json:"markCount"`
				NovelCover     string  `json:"novelCover"`
				BgBanner       string  `json:"bgBanner"`
				NovelID        int     `json:"novelId"`
				NovelName      string  `json:"novelName"`
				Point          float64 `json:"point"`
				IsFinish       bool    `json:"isFinish"`
				AuthorName     string  `json:"authorName"`
				CharCount      int     `json:"charCount"`
				ViewTimes      int     `json:"viewTimes"`
				TypeID         int     `json:"typeId"`
				AllowDown      bool    `json:"allowDown"`
				SignStatus     string  `json:"signStatus"`
				CategoryID     int     `json:"categoryId"`
				IsSensitive    bool    `json:"isSensitive"`
				Expand         struct {
				} `json:"expand"`
			} `json:"entitys"`
		} `json:"contest2Push"`
	} `json:"data"`
}
