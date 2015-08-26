package gorss

// RSS query info for a gzh. gzh is the "public account"
type QueryElement struct {
	id     string
	name   string
	openid string
	eqs    string
	cb     string // fixed
	ekv    string // fixed
	page   string // fixed
	t      string // fixed
}

// add new RSS here
var IDQuerys = []*QueryElement{
	&QueryElement{
		id:     "zhi_japan",
		name:   "知日",
		openid: "oIWsFt3YfRKPuRZmMDZAdlPJgIPU",
		eqs:    "vVszo3Bguw%2BpoUyfUb7gSu7N7CSPLLzqm1DpF5tvTnfaP1JKRtX%2BIxaW3PH%2BFZuKmHrTW",
		cb:     "sogou.weixin.gzhcb",
		ekv:    "3",
		page:   "1",
		t:      "1440596043703",
	},
}
