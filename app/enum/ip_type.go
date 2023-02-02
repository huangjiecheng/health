package enum

type IpTypeEnum int64

const (
	IPV4           IpTypeEnum = 400001
	IPV6           IpTypeEnum = 400002
	IPV4V6         IpTypeEnum = 400003
	UnExpectIpType IpTypeEnum = 0
)
