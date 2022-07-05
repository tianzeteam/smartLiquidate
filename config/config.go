package config

type Config struct {
	App      *App      `yaml:"app"`
	Account  *Account  `yaml:"account"`
	Contract *Contract `yaml:"contract"`
}

type App struct {
	DatabaseUrl     string `yaml:"dburl"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Account struct {
	BscMainNetHttpUrl string `yaml:"BscMainNetHttpUrl"`
	AccountPriKey     string `yaml:"AccountPriKey"`
	AccountAddr       string `yaml:"AccountAddr"`
}

type Contract struct {
	UniswapV2Router       string `yaml:"UniswapV2Router"`
	LiquidateLoanContract string `yaml:"LiquidateLoanContract"`
	LendpoolContract      string `yaml:"LendpoolContract"`
}
