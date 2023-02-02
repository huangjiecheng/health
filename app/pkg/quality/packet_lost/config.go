package packet_lost

import "health/share"

type Config struct {
	Switch            bool                                `json:"switch"`
	Period            int                                 `json:"period"`
	SilentRoundNumber int                                 `json:"silent_round_number"`
	CoverMap          map[share.CoverName]*PkgLostRateCfg `json:"cover_map"`
}

type PkgLostRateCfg struct {
	AdjustLineRateList []float64
	PlanLineRateList   []float64
	PkgLostRateList    []float64
}
