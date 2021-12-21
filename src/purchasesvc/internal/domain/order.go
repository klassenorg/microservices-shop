package domain

type Order struct {
	ID         int    `json:"order_id,omitempty" db:"id"`
	UserID     string `json:"user_id" db:"user_id"`
	FullName   string `json:"full_name" db:"full_name"`
	TotalPrice string `json:"total_price" db:"total_price"`
	Address    string `json:"address" db:"address"`
	CardNumber string `json:"card_number" db:"card_number"`
	CVC        string `json:"cvc" db:"cvc"`
	CardExp    string `json:"exp" db:"exp"`
}
