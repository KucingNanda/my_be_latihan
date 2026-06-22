package model

type Response struct {
	Message string      `json:"message" example:"detail pesan"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty" example:"detail error"`
}

type Response200 struct {
	Message string `json:"message" example:"berhasil melakukan operasi"`
}

type Response201 struct {
	Message string `json:"message" example:"berhasil menambahkan data"`
}

type Response401 struct {
	Message string `json:"message" example:"unauthorized"`
}

type Response403 struct {
	Message string `json:"message" example:"user tidak memiliki akses untuk fitur ini"`
}
