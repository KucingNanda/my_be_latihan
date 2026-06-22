package config

var allowedOrigins = []string{
	"http://localhost:5173",
	"https://mybelatihan-production.up.railway.app",
	"https://my-fe-latihan.vercel.app",
}

func GetAllowedOrigins() []string {
	return allowedOrigins
}
