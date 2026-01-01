package domain

import "time"

type Loan struct {
	ID        int64
	BookID    int64
	UserID    int64
	Status    string
	CreatedAt time.Time
}
