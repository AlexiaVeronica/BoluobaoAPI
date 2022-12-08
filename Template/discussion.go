package Template

type Discussion struct {
	Status Status `json:"status"`
	Data   []struct {
		PostDate string `json:"postDate"`
		Post     struct {
			Post struct {
				PostID      int    `json:"postId"`
				AccountID   int    `json:"accountId"`
				RepostNum   int    `json:"repostNum"`
				ReplyNum    int    `json:"replyNum"`
				FavNum      int    `json:"favNum"`
				NickName    string `json:"nickName"`
				Avatar      string `json:"avatar"`
				Content     string `json:"content"`
				PostDate    string `json:"postDate"`
				IsCanDelete bool   `json:"isCanDelete"`
				Images      []struct {
					SourceImageURL    string `json:"sourceImageUrl"`
					ThumbnailImageURL string `json:"thumbnailImageUrl"`
					SourceWidth       int    `json:"sourceWidth"`
					SourceHeight      int    `json:"sourceHeight"`
					ThumbnailWidth    int    `json:"thumbnailWidth"`
					ThumbnailHeight   int    `json:"thumbnailHeight"`
					ImageID           int    `json:"imageId"`
				} `json:"images"`
				From struct {
					Name interface{} `json:"name"`
					Type interface{} `json:"type"`
					Link interface{} `json:"link"`
				} `json:"from"`
				IsDelete bool        `json:"isDelete"`
				Expand   interface{} `json:"expand"`
			} `json:"post"`
			Source interface{} `json:"source"`
		} `json:"post"`
		Comment     interface{} `json:"comment"`
		PostComment interface{} `json:"postComment"`
		Expand      struct {
			HasFollowed bool `json:"hasFollowed"`
		} `json:"expand"`
		Entity interface{} `json:"entity"`
	} `json:"data"`
}
