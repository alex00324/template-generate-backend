package usecase

import (
	"github.com/myapp/internal/app/domain"
)

type TemplateUsecase struct {
	// Dependencias o interfaces de repositorio aquí
}

func NewTemplateUsecase() *TemplateUsecase {
	// Inicializar y devolver una instancia de TemplateUsecase
}

func (u *TemplateUsecase) CreateTemplate(data *domain.TemplateData) error {
	// Lógica para crear templates aquí
}
