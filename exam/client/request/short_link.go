package request

type CreateShortLinkReq struct {
	Title       string `json:"title"`      // 短链标题
	OriginalURL string `json:"url"`        // 短链原始 URL
	ExpiresAt   int64  `json:"expired_at"` // 过期时间(Unix 时间戳,秒)
}

type DeleteShortLinkReq struct {
	LinkId string `json:"link_id"` // link id
}

type UpdateShortLinkReq struct {
	LinkId      string `json:"link_id"`
	Title       string `json:"title"`
	OriginalUrl string `json:"original_url"`
	ExpiredAt   int64  `json:"expired_at"`
}

type ShareShortLinkReq struct {
	LinkId  string   `json:"link_id" `
	UserIds []string `json:"user_ids"`
}

type ListShortLinksReq struct {
	Page     int64 `json:"page" form:"page" binding:"required"`
	PageSize int64 `json:"page_size" form:"page_size" binding:"required"`
}

type ListShortLinksByLinkIDReq struct {
	Keyword  string `json:"keyword" form:"keyword"`
	Page     int64  `json:"page" form:"page" `
	PageSize int64  `json:"page_size" form:"page_size" `
}

type SummarizeReq struct {
	LinkId string `json:"link_id" form:"link_id" binding:"required"`
}
