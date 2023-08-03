package main

//
//import (
//	"bt.baishancloud.com/resdev/y/util/file"
//	"context"
//	"fmt"
//	"github.com/go-playground/validator/v10"
//	"health/enum"
//	"health/model"
//	"health/oms"
//	"health/share"
//	"sort"
//	"time"
//)
//
//// User contains user information
//type User struct {
//	FirstName      string     `validate:"required"`
//	LastName       string     `validate:"required"`
//	Age            uint8      `validate:"gte=0,lte=130"`
//	Email          string     `validate:"required,email"`
//	FavouriteColor string     `validate:"iscolor"`                // alias for 'hexcolor|rgb|rgba|hsl|hsla'
//	Addresses      []*Address `validate:"required,dive,required"` // a person can have a home and cottage...
//}
//
//// Address houses a users address information
//type Address struct {
//	Street string `validate:"required"`
//	City   string `validate:"required"`
//	Planet string `validate:"required"`
//	Phone  string `validate:"required"`
//}
//
//// use a single instance of Validate, it caches struct info
//var validate *validator.Validate
//
//type ipManager struct {
//	FogIpCardMap       map[share.IpCard]int64 `json:"fog_ip_card_map"`
//	smallNodeIpCardMap map[share.IpCard]int64 `json:"small_node_ip_card_map"`
//	ChangeIpCard       share.IpCard           `json:"change_ip_card"`
//}
//
//type (
//	SMasterData struct {
//		CoverList []*CoverUnit     `json:"cover_list"`
//		SNodeInfo *SNodeConfigInfo `json:"s_node_info"`
//	}
//
//	CoverUnit struct {
//		Name       share.CoverName `json:"name"`
//		Version    string          `json:"version"`
//		ViewList   []*ViewUnit     `json:"view_list"`
//		DomainList []string        `json:"domain_list"`
//	}
//
//	ViewUnit struct {
//		Name      share.ViewName `json:"name"`
//		Weight    float64        `json:"weight"`
//		CacheList []*CacheUnit   `json:"cache_list"`
//	}
//
//	CacheUnit struct {
//		CacheGroupName share.CacheGroupName `json:"cache_group_name"`
//		HunterUrlV4    string               `json:"hunter_url_v4"`
//		HunterUrlV6    string               `json:"hunter_url_v6"`
//		MasterHost     CacheHostUnit        `json:"master_host"`
//	}
//
//	CacheHostUnit struct {
//		Hostname string `json:"hostname"`
//		IPV4     string `json:"ip_v4"`
//		IPV6     string `json:"ip_v6"`
//	}
//
//	SNodeConfigInfo struct {
//		SvgName        share.SvgName   `json:"svg_name"`
//		Isp            string          `json:"isp"`
//		Province       string          `json:"province"`
//		Region         string          `json:"region"`
//		Stype          string          `json:"s_type"`
//		BuildLine      int64           `json:"build_line"`
//		PlanLine       int64           `json:"plan_line"`
//		IpType         enum.IpTypeEnum `json:"ip_type"`
//		MachineType    string          `json:"machine_type"`
//		AccountList    []AccountInfo   `json:"account_list"`
//		IsFaultByItil  bool            `json:"is_fault_by_itil"`
//		MergeAliasName string          `json:"merge_alias_name"`
//	}
//
//	AccountInfo struct {
//		CardName      share.IpCard `json:"card_name"`
//		AccountStatus int          `json:"account_status"`
//	}
//)
//
//var (
//	waterConfig *SMasterData
//)
//
//const (
//	dumpSMasterPath  = "/Users/huangjiecheng/test_hjc.json"
//	dumpSMasterPath2 = "/Users/huangjiecheng/test_hjc222.json"
//)
//
//func init() {
//	if waterConfig != nil {
//		return
//	}
//	var result = &SMasterData{}
//	err := file.LoadFile(dumpSMasterPath, result)
//	if err != nil {
//		fmt.Println(err)
//	}
//	waterConfig = result
//	err = file.DumpFile(dumpSMasterPath2, waterConfig)
//	if err != nil {
//		fmt.Println(err)
//	}
//}
//
//type omsAcl struct {
//}
//
//type OmsAcl interface {
//	GetSNodeNameList() (sNodeNameList []string)
//}
//
//type Factory interface {
//	OmsAcl() OmsAcl
//}
//type Manager struct {
//	Factory1
//}
//
//func NewManager(m Factory1) *Manager {
//	return &Manager{Factory1: m}
//}
//
//type Factory1 interface {
//	Acl() Factory
//	Ctx() context.Context
//}
//
//type factory struct {
//	ctx    context.Context
//	omsAcl OmsAcl
//}
//
//func (f *factory) OmsAcl() OmsAcl {
//	if f.omsAcl == nil {
//		f.omsAcl = oms.NewOmsAcl()
//	}
//	return f.omsAcl
//}
//
//type ReportV3 struct {
//	Hostname    string       `json:"hostname"`
//	TTL         int64        `json:"ttl"`
//	Version     int64        `json:"version"`
//	SNodeIpList []*IPAbility `json:"s_node_ip_list"`
//}
//
//type IPAbility struct {
//	IpCard       share.IpCard `json:"ip_card"`
//	CardIp       share.CardIp `json:"card_ip"`
//	ExpectWeight int64        `json:"expect_weight"` // 期望的权重
//	TotalWeight  int64        `json:"total_weight"`  // 总权重
//	AdjustLine   int64        `json:"adjust_line"`   // 调度线
//	Plan         int64        `json:"plan"`          // 规划线
//}
//
//func distributionExpSelf(selfReport *ReportV3) {
//	sort.Slice(selfReport.SNodeIpList, func(i, j int) bool {
//		return selfReport.SNodeIpList[i].CardIp < selfReport.SNodeIpList[j].CardIp
//	})
//	coverToIpAbsMap := assignIpToCover(selfReport.SNodeIpList)
//	if ipAbMap, ok := coverToIpAbsMap["cover"]; ok {
//		execDistribute(ipAbMap)
//	}
//}
//func execDistribute(abMap map[share.CardIp]*IPAbility) {
//	var (
//		allExp       = int64(0)
//		coverAvgSize = 2 * 1e6
//	)
//	for _, ab := range abMap {
//		var (
//			maxExp = int64(float64(ab.AdjustLine) / float64(coverAvgSize))
//		)
//		if ab.ExpectWeight > maxExp {
//			ab.ExpectWeight = 666
//		}
//		allExp += ab.ExpectWeight
//	}
//}
//
//type Mode string
//
//func (m Mode) GetMode() Mode {
//
//	if _, ok := aaaMap["aaa"]; ok {
//		return m
//	}
//
//	return "aa"
//}
//
//var aaaMap = make(map[share.CardIp]Mode, 0)
//var (
//	cardIpToLastDecisionInfoMap = make(map[share.CardIp]*model.DecisionInfo)
//)
//
//func main() {
//	cardIpToLastDecisionInfoMap["aaa"] = &model.DecisionInfo{
//		PlanLine: 1111,
//	}
//
//	item, _ := cardIpToLastDecisionInfoMap["aaa"]
//	item.PlanLine = 2222
//	fmt.Println("=================== %v", cardIpToLastDecisionInfoMap)
//	now := time.Now()
//	fmt.Printf("时间： %d, 准确: %s\n", now.Unix(), now.String())
//	num := 3 * 3 * 3 * 17 * 201 * 10
//	namePlanList := make([]string, 0, num)
//	for i := 0; i < num; i++ {
//		//if i%5 == 0 {
//		//	namePlan = ""
//		//}
//		//namePlan += string("sNode.Name") + "::" + strconv.Itoa(45000000)
//		namePlanList = append(namePlanList, fmt.Sprintf("%s::%d", "sNode.Name", 45000000))
//
//		//fmt.Println(namePlan)
//	}
//
//	end := time.Now()
//	fmt.Printf("开始时间： %d, 准确: %s\n", now.Unix(), now.String())
//	fmt.Printf("结束时间： %d, 准确: %s\n", end.Unix(), end.String())
//
//	fmt.Println(aaaMap["132"].GetMode())
//	r := &ReportV3{
//		Hostname:    "127.0.0.1",
//		TTL:         3600,
//		Version:     1,
//		SNodeIpList: make([]*IPAbility, 0),
//	}
//	r.SNodeIpList = append(r.SNodeIpList, &IPAbility{
//		IpCard:       "127.0.0.1",
//		CardIp:       "127.0.0.1",
//		ExpectWeight: 11,
//		TotalWeight:  11,
//		AdjustLine:   11,
//		Plan:         11,
//	})
//	r.SNodeIpList = append(r.SNodeIpList, &IPAbility{
//		IpCard:       "127.0.0.22",
//		CardIp:       "127.0.0.22",
//		ExpectWeight: 22,
//		TotalWeight:  22,
//		AdjustLine:   22,
//		Plan:         22,
//	})
//
//	distributionExpSelf(r)
//
//	var (
//		lastCardIpToCoverMap = make(map[share.CardIp]share.CoverName) // 上一次线路分配的cover
//		lastCoverIpNum       = make(map[share.CoverName]int64, 0)     // 上一次分配给cover对应的线路个数
//	)
//	lastCardIpToCoverMap["1111"] = "c1"
//	lastCardIpToCoverMap["222"] = "c1"
//	lastCardIpToCoverMap["333"] = "c1"
//	lastCardIpToCoverMap["44"] = "c2"
//	lastCardIpToCoverMap["55"] = "c2"
//	for _, coverName := range lastCardIpToCoverMap {
//		lastCoverIpNum[coverName]++
//	}
//
//	sortedCacheResult := execDemo("")
//	sortedCacheResult2 := execDemo2("")
//	result := make([]*share.CoverRegionCacheInfo, 0, 5)
//	result = append(result, sortedCacheResult...)
//	result = append(result, sortedCacheResult2...)
//
//	for _, i := range result {
//		fmt.Println(fmt.Sprintf("name: %s  allReq: %d", i.CoverName, i.AllReq))
//	}
//
//	validate = validator.New()
//
//	validateStruct()
//	validateVariable()
//}
//
//func assignIpToCover(list []*IPAbility) map[share.CoverName]map[share.CardIp]*IPAbility {
//	coverToIpAbsMap := make(map[share.CoverName]map[share.CardIp]*IPAbility, 0) // cover对应分配的ip列表
//	ipAbMap := make(map[share.CardIp]*IPAbility)                                // 分配的线路
//
//	for _, ipAb := range list {
//		ipAbMap[ipAb.CardIp] = ipAb
//	}
//	coverToIpAbsMap["cover"] = ipAbMap
//	return coverToIpAbsMap
//}
//
//func execDemo(name string) []*share.CoverRegionCacheInfo {
//	sortedCacheResult := make([]*share.CoverRegionCacheInfo, 0, 5)
//	sortedCacheResult = append(sortedCacheResult, &share.CoverRegionCacheInfo{
//		CoverName: "c1",
//		AllReq:    2,
//	})
//	sortedCacheResult = append(sortedCacheResult, &share.CoverRegionCacheInfo{
//		CoverName: "c1",
//		AllReq:    6,
//	})
//	sortedCacheResult = append(sortedCacheResult, &share.CoverRegionCacheInfo{
//		CoverName: "c1",
//		AllReq:    8,
//	})
//	sortedCacheResult = append(sortedCacheResult, &share.CoverRegionCacheInfo{
//		CoverName: "c1",
//		AllReq:    4,
//	})
//	// 按放行+未放行之和从大到小排序
//	sort.Slice(sortedCacheResult, func(i, j int) bool {
//		return sortedCacheResult[i].AllReq > sortedCacheResult[j].AllReq
//	})
//	return sortedCacheResult
//}
//func execDemo2(name string) []*share.CoverRegionCacheInfo {
//	sortedCacheResult := make([]*share.CoverRegionCacheInfo, 0, 5)
//	sortedCacheResult = append(sortedCacheResult, &share.CoverRegionCacheInfo{
//		CoverName: "c2",
//		AllReq:    11,
//	})
//	sortedCacheResult = append(sortedCacheResult, &share.CoverRegionCacheInfo{
//		CoverName: "c2",
//		AllReq:    13,
//	})
//	sortedCacheResult = append(sortedCacheResult, &share.CoverRegionCacheInfo{
//		CoverName: "c2",
//		AllReq:    5,
//	})
//	sortedCacheResult = append(sortedCacheResult, &share.CoverRegionCacheInfo{
//		CoverName: "c2",
//		AllReq:    7,
//	})
//	// 按放行+未放行之和从大到小排序
//	sort.Slice(sortedCacheResult, func(i, j int) bool {
//		return sortedCacheResult[i].AllReq > sortedCacheResult[j].AllReq
//	})
//	return sortedCacheResult
//}
//
//func validateStruct() {
//
//	address := &Address{
//		Street: "Eavesdown Docks",
//		Planet: "Persphone",
//		Phone:  "none",
//	}
//
//	user := &User{
//		FirstName:      "Badger",
//		LastName:       "Smith",
//		Age:            135,
//		Email:          "Badger.Smith@gmail.com",
//		FavouriteColor: "#000-",
//		Addresses:      []*Address{address},
//	}
//
//	// returns nil or ValidationErrors ( []FieldError )
//	err := validate.Struct(user)
//	if err != nil {
//
//		// this check is only needed when your code could produce
//		// an invalid value for validation such as interface with nil
//		// value most including myself do not usually have code like this.
//		if _, ok := err.(*validator.InvalidValidationError); ok {
//			fmt.Println(err)
//			return
//		}
//
//		for _, err := range err.(validator.ValidationErrors) {
//
//			fmt.Println(err.Namespace())
//			fmt.Println(err.Field())
//			fmt.Println(err.StructNamespace())
//			fmt.Println(err.StructField())
//			fmt.Println(err.Tag())
//			fmt.Println(err.ActualTag())
//			fmt.Println(err.Kind())
//			fmt.Println(err.Type())
//			fmt.Println(err.Value())
//			fmt.Println(err.Param())
//			fmt.Println()
//		}
//
//		// from here you can create your own error messages in whatever language you wish
//		return
//	}
//
//	// save user to database
//}
//
//func validateVariable() {
//
//	myEmail := "joeybloggs.gmail.com"
//
//	errs := validate.Var(myEmail, "required,email")
//
//	if errs != nil {
//		fmt.Println(errs) // output: Key: "" Error:Field validation for "" failed on the "email" tag
//		return
//	}
//
//	// email ok, move on
//}
