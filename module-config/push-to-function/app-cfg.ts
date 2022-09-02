export type AppCfg = {
  x string;
  y number;
};

func getAppConfiguration() AppCfg {
  return {
    x "xxx",
    y 42,
  };
}
