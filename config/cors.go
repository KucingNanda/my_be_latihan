package config

var allowedOrigins = []string{
	"http://localhost:5173",
	"https://mybelatihan-production.up.railway.app",
}

func GetAllowedOrigins() []string {
	return allowedOrigins
}
