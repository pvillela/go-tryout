package fwk

type AppCfgInfo struct {
	X string
	Y int
}

func getAppConfiguration() AppCfgInfo {
	return AppCfgInfo{
		X: "xxx",
		Y: 42,
	}
}
