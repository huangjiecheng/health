package share

type CoverName string
type ViewName string
type CacheGroupName string
type IpCard string
type CardIp string
type SvgName string
type ComponentName string
type StrategyName string
type RegionName string
type TimestampMs int64 //毫秒级时间戳

type CoverRegionCacheInfo struct {
	CoverName         CoverName      `json:"cover_name"`
	RegionName        RegionName     `json:"region_name"`
	Name              CacheGroupName `json:"name"`
	PassReqCount      int64          `json:"pass_req_count"`      //放行请求数
	InterceptReqCount int64          `json:"intercept_req_count"` //拦截请求数
	AllReq            int64          `json:"all_req"`             // 放行+未放行
	ExpectWeight      int64          `json:"expect_weight"`
	TotalExpectWeight int64          `json:"total_expect_weight"`
	ExpectWeightRate  int64          `json:"expect_weight_rate"`
	Qps               int64          `json:"qps"`
}

const (
	DefaultRegion RegionName = "_default_region_4.0"
)
