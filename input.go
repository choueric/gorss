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
	t      string
}

// add new RSS here
var IDQuerys = []*QueryElement{
	&QueryElement{
		id:     "zhi_japan",
		name:   "知日",
		openid: "oIWsFt3YfRKPuRZmMDZAdlPJgIPU",
		eqs:    "vVszo3Bguw%2BpoUyfUb7gSu7N7CSPLLzqm1DpF5tvTnfaP1JKRtX%2BIxaW3PH%2BFZuKmHrTW",
	},
	&QueryElement{
		id:     "seekingbeta_earl",
		name:   "earletf",
		openid: "oIWsFt2WedIo0nqhDIeHAEazr7Vw",
		eqs:    "43sQoeuglO7EobHMZZaQvuTvFqL8H8xqIXx52SOvhJPxUqP29e1rV1S%2B%2BunSk2bTuXoDE",
	},
	&QueryElement{
		id:     "seniorplayer",
		name:   "大玩家张磊",
		openid: "oIWsFt3sSJFYcdEcqQeePTe55UEM",
		eqs:    "tvsGofAgS%2F0Xoeo2TnXHnuBdq2l1gU5LIn8NyrRsfRDyd4qp6IJThHRVst9kKHCBPW9c3",
	},
	&QueryElement{
		id:     "etfhefenji",
		name:   "ETF和分级圈",
		openid: "oIWsFtz7JJpbnCAaoFXDp-DIQ5LQ",
		eqs:    "HVsGouugIb%2B4oY6QbnUjQuLcOPTseRudnhX4bEmQYtYb%2F879npHpvYjisFtzZ2CLZm8wp",
	},
	&QueryElement{
		id:     "IELTS_Online2015",
		name:   "IELTS在线",
		openid: "oIWsFt_2iP_TI_P6jX1fNgL_SoP4",
		eqs:    "4UsVoongqtkro%2FEq0%2BWFku43OcFCbbmkxWVcb%2F68zo4V31EteCu3NHqIfUtOVu4n4CCsC",
	},
	&QueryElement{
		id:     "ilianyue",
		name:   "连岳",
		openid: "oIWsFt0e_MEZmRrjEbLsh99_H13E",
		eqs:    "NtspokCg2stAoCGXiWRp0uJr8VdDfCXoN%2B2KraHab6e9tXh4EQ%2BLrAH8UfuG2occr3kzH",
	},
	&QueryElement{
		id:     "bitsea",
		name:   "槽边往事",
		openid: "oIWsFtxG-2J2sGx3l5-pknZDv60g",
		eqs:    "busoo9%2Bgtcq6oQsWluIJzuM2p7w6sBt09VL7nn8BrmmYN%2BGKWxdO%2FcAwvEC57oVHNuq7I",
	},
	&QueryElement{
		id:     "shi_che_ju_lan",
		name:   "矢车菊蓝石楠花",
		openid: "oIWsFt0h74dXz_vzojW41tUYj_H8",
		eqs:    "kQs7oVTgHvaaoRqvkfi3SuIiO4VQQtdcRbm5G2Q%2By1TB01FO4xawDBd4cRThMKwHTLEIk",
	},
	&QueryElement{
		id:     "DingXiangYiSheng",
		name:   "丁香医生",
		openid: "oIWsFt-n9zARBylWc-BYhaEITE20",
		eqs:    "B5ssoIyg9sfmocAhz1uFjuDeeSUDKjxGHorL0EGLP76KvorKxMStWdjuV2mfQW60DxAlu",
	},
	&QueryElement{
		id:     "rayli-herstyle",
		name:   "瑞丽伊人风尚",
		openid: "oIWsFt41SMSxev3KfhErrS2ZKMrA",
		eqs:    "%2Bls6oOZg152ZoHIYEhqKkuq27PP9593FkcyKH0TbBVyDyQ5CbJNueIvwAn7e281TpXNUM",
	},
	&QueryElement{
		id:     "SeriousEats",
		name:   "吃早餐是件正经事",
		openid: "oIWsFt1Kb3kZB-99E5xIz_za0CJ0",
		eqs:    "nIsRoMgglLpAo3J7SXrvcuwyBDWJPH6N55GTpYYY2uXw2bqPOLh5AZL96%2FkJi66h11R9L",
	},
	&QueryElement{
		id:     "xuetouzilicai",
		name:   "零基础学投资理财",
		openid: "oIWsFtyZhhvfHzCsmBndbT2SOUh8",
		eqs:    "%2BYsloa4g5cYZoWRAWT%2F1Du7UtQxh1H2YNSEoqSBW5T29leixy88S64Cy7inAiuqm6Snly",
	},
	&QueryElement{
		id:     "jane7ducai",
		name:   "简七理财",
		openid: "oIWsFt0qY9YvyYESHey3MOPfbNy0",
		eqs:    "2WsQoD2gmsBJowgHMJCNwuI3qpyc%2BbRKPZSuHhkLtg6M%2F0a75gcMlccIJv3XzfigvWymn",
	},
}
