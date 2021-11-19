package articles

import (
	"time"
)

type Core struct {
	ID        int
	Title     string
	Status    bool
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	User      User
}

type User struct {
	Name  string
	Media string
}

type CategoryCore struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Bussiness interface {
	CreateData(data Core) (resp Core, err error)
	GetAllData(search string) (resp []Core)
	// continue the function abtraction
}

type Data interface {
	InsertData(data Core) (resp Core, err error)
	SelectData(title string) (resp []Core)
}
