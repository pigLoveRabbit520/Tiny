package config

type Config struct {
	Port int `json:"port"`
	DB   struct {
		Host     string `json:"host"`
		Port     uint   `json:"port"`
		User     string `json:"usr"`
		Name     string `json:"name"`
		Password string `json:"password"`
		Prefix   string `json:"prefix"`
	} `json:"db"`
}
