package config

func loadTestConfig(cfg *Config) {
	cfg.DatabaseHost = "127.0.0.1"
	cfg.DatabaseName = "goyagi"
	cfg.DatabaseUser = "goyagi_admin"
	cfg.Environment = "test"
}
