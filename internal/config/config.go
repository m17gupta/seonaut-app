package config

import (
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/spf13/viper"
)

// CrawlerConfig stores the configuration for the crawler.
type CrawlerConfig struct {
	Agent string `mapstructure:"agent"`
}

// HTTPServerConfig stores the configuration for the HTTP server.
type HTTPServerConfig struct {
	Server     string `mapstructure:"host"`
	Port       int    `mapstructure:"port"`
	URL        string `mapstructure:"url"`
	FrontendURL string `mapstructure:"frontend_url"`
}

// DBConfig stores the configuration for the database store.
type DBConfig struct {
	Server string `mapstructure:"server"`
	Port   int    `mapstructure:"port"`
	User   string `mapstructure:"user"`
	Pass   string `mapstructure:"password"`
	Name   string `mapstructure:"database"`
}

// Config stores the configuration for the application.
type Config struct {
	Crawler    *CrawlerConfig    `mapstructure:"crawler"`
	HTTPServer *HTTPServerConfig `mapstructure:"server"`
	DB         *DBConfig         `mapstructure:"database"`
}

// NewConfig loads the configuration from the specified file and path.
func NewConfig(configFile string) (*Config, error) {
	viper.AddConfigPath(filepath.Dir(configFile))
	viper.SetConfigName(filepath.Base(configFile))
	viper.SetConfigType("toml")

	// Allow environment variables to override config file values
	viper.AutomaticEnv()
	viper.SetEnvPrefix("SEONAUT")

	// Set defaults for production
	viper.SetDefault("server.host", "0.0.0.0")
	viper.SetDefault("server.port", 10000)
	viper.SetDefault("crawler.agent", "Mozilla/5.0 (compatible; SEOnautBot/1.0; +https://seonaut.org/bot)")

	if err := viper.ReadInConfig(); err != nil {
		// If config file is not found, continue with environment variables and defaults
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, err
		}
	}

	var config Config

	err := viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	// Parse DATABASE_URL if provided (common in cloud deployments like Render)
	if databaseURL := os.Getenv("DATABASE_URL"); databaseURL != "" {
		dbConfig, err := parseDatabaseURL(databaseURL)
		if err == nil {
			config.DB = dbConfig
		}
	}

	// Override port with PORT environment variable if provided (Render requirement)
	if port := os.Getenv("PORT"); port != "" {
		if p, err := strconv.Atoi(port); err == nil {
			config.HTTPServer.Port = p
		}
	}

	return &config, nil
}

// parseDatabaseURL parses a database URL into a DBConfig struct
func parseDatabaseURL(databaseURL string) (*DBConfig, error) {
	u, err := url.Parse(databaseURL)
	if err != nil {
		return nil, err
	}

	password, _ := u.User.Password()
	port := 3306
	if u.Port() != "" {
		if p, err := strconv.Atoi(u.Port()); err == nil {
			port = p
		}
	}

	return &DBConfig{
		Server: u.Hostname(),
		Port:   port,
		User:   u.User.Username(),
		Pass:   password,
		Name:   strings.TrimPrefix(u.Path, "/"),
	}, nil
}
