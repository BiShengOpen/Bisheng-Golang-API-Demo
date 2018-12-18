package config

const (
	// API KEY
	// Todo: 此处用自己的API Key
	API_KEY = "8gUnVP8IrVWzH62T"
	// Todo: 此处用自己的Secret Key
	SECRET_KEY = "37bdaef95dd4f4094076b83637dad703c844948fbd8d7a4682fb733b5a53d874"

	// API请求地址
	BISHENG_URL = "http://test164.yibeix.com"
	HOST = "test164.yibeix.com"
	GATEWAY = "test1"

)

//系统接口
const (
	TIKER = "/market" //行情

	BALANCE = "/account/balance" //用户持仓

	ORDER_LIST = "/orders/list" //用户订单记录
	ORDER_CREATE = "/orders/create" //创建订单
	ORDER_CANCEL = "/orders/cancel" //撤单
	ORDER_STATUS = "/orders/status" //查询订单详情
	ORDER_UNDEAL = "/orders/undeal/status" //查询所有待成交和部分成交的订单

	RECORD_LIST = "/records/list" //成交回报
	RECORD_MAXTRDNUM = "/records/maxtrdnum" //获取最大的成交编号
	RECORD_TRDNUM = "/records/trdnum" //查询成交记录(根据成交编号查询)
	RECORD_ORDNUM = "/records/ordnum" //查询成交记录(根据订单编号)

)