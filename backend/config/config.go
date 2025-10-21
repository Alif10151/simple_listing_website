package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var configurations *Config

type DBConfig struct {
	Host      string
	Port      int
	Name      string
	User      string
	Password  string
	EnableSSL bool
}

type Config struct {
	Version      string
	ServiceName  string
	HttpPort     int
	JwtSecretKey string
	DB           *DBConfig
}

func loadConfig() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("failed to load env fileand err:", err)
		os.Exit(1)
	}
	version := os.Getenv("VERSION")
	if version == "" {
		fmt.Println("Version is required")
		os.Exit(1)
	}

	serviceName := os.Getenv("Service_Name")
	if serviceName == "" {
		fmt.Println("service name its required")
		os.Exit(1)
	}

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		fmt.Println("port its required")
		os.Exit(1)
	}

	port, err := strconv.ParseInt(httpPort, 10, 64)

	if err != nil {
		fmt.Println("port must be number")
		os.Exit(1)
	}

	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	if jwtSecretKey == "" {
		fmt.Println("Jwt secret key is needed")
		os.Exit(1)
	}

	dbhost := os.Getenv("DB_HOST")
	if dbhost == "" {
		fmt.Println("DB Host is needed")
		os.Exit(1)
	}

	dbport := os.Getenv("DB_PORT")
	if dbport == "" {
		fmt.Println("DB Port is required")
		os.Exit(1)
	}

	dbprt, err := strconv.ParseInt(dbport, 10, 64)

	if err != nil {
		fmt.Println("port must be number")
		os.Exit(1)
	}

	dbname := os.Getenv("DB_NAME")
	if dbname == "" {
		fmt.Println("DB Name is needed")
		os.Exit(1)
	}

	dbuser := os.Getenv("DB_USER")
	if dbuser == "" {
		fmt.Println("DB User is needed")
		os.Exit(1)
	}

	dbpassword := os.Getenv("DB_PASSWORD")
	if dbpassword == "" {
		fmt.Println("DB Password is needed")
		os.Exit(1)
	}

	enableSSL := os.Getenv("DB_ENABLESSL")

	SSLMODE, err := strconv.ParseBool(enableSSL)
	if err != nil {
		fmt.Println("Inavlid SSLMODE", err)
		os.Exit(1)
	}

	dbConfig := &DBConfig{
		Host:      dbhost,
		Port:      int(dbprt),
		Name:      dbname,
		User:      dbuser,
		Password:  dbpassword,
		EnableSSL: SSLMODE,
	}

	configurations = &Config{
		Version:      version,
		ServiceName:  serviceName,
		HttpPort:     int(port),
		JwtSecretKey: jwtSecretKey,
		DB:           dbConfig,
	}
}

func GetConfig() *Config {
	if configurations == nil {
		loadConfig()
	}
	return configurations
}
