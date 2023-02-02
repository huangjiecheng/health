package download_speed

import "health/share"

type (
	Config struct {
		Switch               bool                             `json:"switch"`
		SyncInterval         int                              `json:"sync_interval"`
		SpeedLimitMinRate    float64                          `json:"speed_limit_min_rate"`
		SpeedLimitFreezeTime int64                            `json:"speed_limit_freeze_time"`
		CoverMap             map[share.CoverName][]*CoverInfo `json:"cover_map"`
	}

	CoverInfo struct {
		Speed int64   `json:"speed"`
		Rate  float64 `json:"rate"`
	}
)
