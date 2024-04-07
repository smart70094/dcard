package ad

type GetAdReqVo struct {
	Age      int    `form:"age"`
	Gender   string `form:"gender"`
	Country  string `form:"country"`
	Platform string `form:"platform"`
	Offset   int    `form:"offset"`
	Limit    int    `form:"limit"`
}
