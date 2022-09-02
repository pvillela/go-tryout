package fwk

type AppCfg struct {
	X string
	Y int
}

func getAppConfiguration() AppCfg {
	return AppCfg{
		X: "xxx",
		Y: 42,
	}
}
