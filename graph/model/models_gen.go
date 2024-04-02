// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type CreateUserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Medication struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Stock       int    `json:"stock"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

type Mutation struct {
}

type Order struct {
	ID        string       `json:"id"`
	UserID    string       `json:"userId"`
	User      *User        `json:"user"`
	Items     []*OrderItem `json:"items"`
	Total     float64      `json:"total"`
	Status    string       `json:"status"`
	CreatedAt string       `json:"createdAt"`
	UpdatedAt string       `json:"updatedAt"`
}

type OrderItem struct {
	ID           string      `json:"id"`
	MedicationID string      `json:"medicationId"`
	Medication   *Medication `json:"medication"`
	Quantity     int         `json:"quantity"`
	Price        float64     `json:"price"`
}

type OrderItemInput struct {
	MedicationID string `json:"medicationId"`
	Quantity     int    `json:"quantity"`
}

type Query struct {
}

type UpdateUserInput struct {
	Name    *string `json:"name,omitempty"`
	Email   *string `json:"email,omitempty"`
	Phone   *string `json:"phone,omitempty"`
	Address *string `json:"address,omitempty"`
}

type User struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Email     string  `json:"email"`
	Password  string  `json:"password"`
	Phone     *string `json:"phone,omitempty"`
	Address   *string `json:"address,omitempty"`
	CreatedAt string  `json:"createdAt"`
	UpdatedAt string  `json:"updatedAt"`
}
