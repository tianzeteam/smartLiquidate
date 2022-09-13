package config

type Config struct {
	App       *App       `yaml:"app"`
	Account   *Account   `yaml:"account"`
	Contract  *Contract  `yaml:"contract"`
	Liquidate *Liquidate `yaml:"liquidate"`
	BaseToken *BaseToken `yaml:"baseToken"`
	ChainId   *ChainId   `yaml:"chainId"`
}

type App struct {
	DatabaseUrl string `yaml:"dburl"`
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
	Username    string `yaml:"username"`
	Password    string `yaml:"password"`
	Runing      bool   `yaml:"Runing"`
}

type Account struct {
	RpcHttpUrl    string `yaml:"RpcHttpUrl"`
	AccountPriKey string `yaml:"AccountPriKey"`
	AccountAddr   string `yaml:"AccountAddr"`
	GasPrice      string `yaml:"GasPrice"`
	GasLimit      int    `yaml:"GasLimit"`
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

type BaseToken struct {
	Aave string `yaml:"Aave"`
	Usdt string `yaml:"Usdt"`
	Weth string `yaml:"Weth"`
	Dai  string `yaml:"Dai"`
}

type ChainId struct {
	Network int `yaml:"Network"`
}
