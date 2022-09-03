package fwk

type AppCfgInfo struct {
	X string
	Y int
}

func GetAppConfiguration() AppCfgInfo {
	return AppCfgInfo{
		X: "xxx",
		Y: 42,
	}
}
