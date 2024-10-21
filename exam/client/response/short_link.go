package response

type CreateResponse struct {
	Code    int       `json:"code"`
	Message string    `json:"msg"`
	Data    ShortLink `json:"data"`
}

type UpdateResponse struct {
	Code    int       `json:"code"`
	Message string    `json:"msg"`
	Data    ShortLink `json:"data"`
}

type ShareResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

type ListResp struct {
	Code    int      `json:"code"`
	Message string   `json:"msg"`
	Data    ListData `json:"data"`
}

type SearchResp struct {
	Code    int        `json:"code"`
	Message string     `json:"msg"`
	Data    SearchData `json:"data"`
}

type DeleteResp struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

type SummarizeResp struct {
	Code    int           `json:"code"`
	Message string        `json:"msg"`
	Data    SummarizeData `json:"data"`
}

type ShortLink struct {
	ID          uint64 `json:"id"`           // 主键id
	LinkID      string `json:"link_id"`      // 短链id, 核心业务id, 由短链生成器生成
	GroupID     string `json:"group_id"`     // 组
	Title       string `json:"title"`        // 短链标题
	OriginalURL string `json:"original_url"` // 原始url地址
	Sign        string `json:"sign"`         // 长链的md5码，方便查找
	ExpiredAt   int64  `json:"expired_at"`   // 过期时间，时间戳格式
	CreatedAt   int64  `json:"created_at"`   // 创建时间，时间戳格式
	UpdatedAt   int64  `json:"updated_at"`   // 更新时间，时间戳格式
	Del         uint8  `json:"del"`          // 0是默认，1是删除
	UserID      string `json:"user_id"`      // 用户ID
}

type ListData struct {
	Total int64        `json:"total"`
	List  []*ShortLink `json:"list"`
}

type SearchData struct {
	Total int64        `json:"total"`
	List  []*ShortLink `json:"list"`
}

type SummarizeData struct {
	Pv int64 `json:"pv"`
	Uv int64 `json:"uv"`
}
