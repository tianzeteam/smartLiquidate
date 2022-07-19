package config

type Config struct {
	App       *App       `yaml:"app"`
	Account   *Account   `yaml:"account"`
	Contract  *Contract  `yaml:"contract"`
	Liquidate *Liquidate `yaml:"liquidate"`
}

type App struct {
	DatabaseUrl string `yaml:"dburl"`
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
	Username    string `yaml:"username"`
	Password    string `yaml:"password"`
}

type Account struct {
	RpcHttpUrl    string `yaml:"RpcHttpUrl"`
	AccountPriKey string `yaml:"AccountPriKey"`
	AccountAddr   string `yaml:"AccountAddr"`
}

type Contract struct {
	UniswapV2Router       string `yaml:"UniswapV2Router"`
	LiquidateLoanContract string `yaml:"LiquidateLoanContract"`
	LendpoolContract      string `yaml:"LendpoolContract"`
	UniswapV2Factory      string `yaml:"UniswapV2Factory"`
}

type Liquidate struct {
	MAX_HEALTH_THRESHOLD     float64 `yaml:"MAX_HEALTH_THRESHOLD"`
	MAX_COLLATERAL_THRESHOLD float64 `yaml:"MAX_COLLATERAL_THRESHOLD"`
	MAX_DEBT_THRESHOLD       float64 `yaml:"MAX_DEBT_THRESHOLD"`
}


