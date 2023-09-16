package domain

type TemplateData struct {
	NombreArchivo      string            `json:"nombreArchivo"`
	NumSecciones       int               `json:"numSecciones"`
	ImagenesPorSeccion [][]string        `json:"imagenesPorSeccion"`
	URLFondo           string            `json:"urlFondo"`
	Color              string            `json:"color"`
	RedesSociales      map[string]string `json:"redesSociales"`
}
