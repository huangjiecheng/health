package enum

// IpTypeEnum
type IpTypeEnum int64

type AccountStatusEnum int8

type IpStatusEnum uint8

type IpRunMode uint8

const (
	IPV4           IpTypeEnum = 400001
	IPV6           IpTypeEnum = 400002
	IPV4V6         IpTypeEnum = 400003
	UnExpectIpType IpTypeEnum = 0

	AccountNotFound AccountStatusEnum = -1 // 没找到账号状态
	AccountNormal   AccountStatusEnum = 1  // 正常
	AccountFault    AccountStatusEnum = 2  // 故障
	AccountOffline  AccountStatusEnum = 3  // 掉线

	IpOk               IpStatusEnum = 1 // 正常
	IpAbnormal         IpStatusEnum = 2 // 异常
	IpDetectFault      IpStatusEnum = 3 // 探测三线IP不通
	IpAccountFault     IpStatusEnum = 4 // 账号登故
	MachineFaultByItil IpStatusEnum = 5 // itil登故
	IpTypeFault        IpStatusEnum = 6 // ip规划类型错误

	UndefinedMode   IpRunMode = 0 // 未定义模式
	IpRunFogMode    IpRunMode = 1 // 跑雾
	IpRunParentMode IpRunMode = 2 // 跑父
)
