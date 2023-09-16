package http

import (
	"encoding/json"
	"net/http"

	"github.com/myapp/internal/app/domain"
	"github.com/myapp/internal/app/usecase"
)

type TemplateHandler struct {
	TemplateUsecase *usecase.TemplateUsecase
}

func NewTemplateHandler(templateUsecase *usecase.TemplateUsecase) *TemplateHandler {
	// Inicializar y devolver una instancia de TemplateHandler
}

func (h *TemplateHandler) CreateTemplate(w http.ResponseWriter, r *http.Request) {
	// Manejar la solicitud HTTP para crear templates aquí
}
5. internal/app/delivery/http/router.go:

go
Copy code
package http

import (
	"github.com/gorilla/mux"
)

func NewRouter(templateHandler *TemplateHandler) *mux.Router {
	// Configuración del enrutador HTTP y asignación de controladores aquí
}



