package dto

import "time"

// CompanyDTO representa los datos que se reciben o envían de la entidad Company.
type CompanyDTO struct {
	ID        uint      `json:"id,omitempty"`                                     // ID de la empresa, omitido si no está disponible.
	Name      string    `json:"name" validate:"required,min=3,max=100"`           // Nombre de la empresa.
	Address   string    `json:"address,omitempty" validate:"min=5,max=250"`       // Dirección de la empresa, opcional.
	Phone     string    `json:"phone,omitempty" validate:"required,min=5,max=12"` // Teléfono de la empresa, opcional.
	Email     string    `json:"email" validate:"min=5,max=80"`                    // Correo electrónico, requerido.
	IsActive  bool      `json:"is_active"`                                        // Estado de la empresa.
	CreatedAt time.Time `json:"created_at,omitempty"`                             // Fecha de creación, opcional.
	UpdatedAt time.Time `json:"updated_at,omitempty"`                             // Fecha de actualización, opcional.
	Users     []UserDTO `json:"users,omitempty"`                                  // Relación con usuarios, opcional.
}
