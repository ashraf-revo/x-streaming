package engine

import "time"

//Parameters for engine
type Parameters struct {
	EnableWaitStream bool
	RingSize         int
}

//ExtendInfo for extension
type ExtendInfo struct {
	Version   *string
	StartTime time.Time //启动时间
	Params    *Parameters
}

//Settings for engine
var (
	Config = &Parameters{true, 10}

	// ConfigRaw 配置信息的原始数据
	ConfigRaw []byte
	// Version 引擎版本号
	Version string
	// EngineInfo 引擎信息
	EngineInfo = &ExtendInfo{&Version, time.Now(), Config}
)
