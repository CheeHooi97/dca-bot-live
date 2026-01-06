package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	BinanceApiKey    string
	BinanceApiSecret string
	BTC2             string
	BTC1             string
	BTC1x5           string
	ETH2             string
	ETH1             string
	ETH1x5           string
	TelegramChatId   string
	BTC1_1h          string
	BTC1_4h          string
	ETH1_1h          string
	ETH1_4h          string
	ADA1_1h          string
	BNB1_1h          string
	SOL1_1h          string

	SystemAesKey string
	GatewayApi   string
	Env          string
	DBHost       string
	DBPort       string
	DBUser       string
	DBPassword   string
	DBName       string
)

// LoadConfig
func LoadConfig() {
	_ = godotenv.Load()

	BinanceApiKey = GetEnv("BINANCE_API_KEY")
	BinanceApiSecret = GetEnv("BINANCE_API_SECRET")
	BTC2 = GetEnv("BTC2")
	BTC1 = GetEnv("BTC1")
	BTC1x5 = GetEnv("BTC1x5")
	ETH2 = GetEnv("ETH2")
	ETH1 = GetEnv("ETH1")
	ETH1x5 = GetEnv("ETH1x5")
	TelegramChatId = GetEnv("TELEGRAM_CHAT_ID")
	BTC1_1h = GetEnv("BTC1_1h")
	BTC1_4h = GetEnv("BTC1_4h")
	ETH1_1h = GetEnv("ETH1_1h")
	ETH1_4h = GetEnv("ETH1_4h")
	ADA1_1h = GetEnv("ADA1_1h")
	BNB1_1h = GetEnv("BNB1_1h")
	SOL1_1h = GetEnv("SOL1_1h")
	SystemAesKey = GetEnv("SYSTEM_AES_KEY")
	GatewayApi = GetEnv("GATEWAY_API")

	Env = GetEnv("ENV")

	DBHost = GetEnv("MYSQL_HOST")
	DBPort = GetEnv("MYSQL_PORT")
	DBUser = GetEnv("MYSQL_USER")
	DBPassword = GetEnv("MYSQL_PASSWORD")
	DBName = GetEnv("MYSQL_DATABASE")
}

func GetEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("%s environment variable not set", key)
	}
	return value
}

var IsDevelopment = func() bool {
	return Env == "dev"
}

var IsLocal = func() bool {
	return IsDevelopment()
}
