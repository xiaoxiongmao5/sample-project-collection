package userinterface

type User struct {
	Name string
	Age  int32
}

type Option interface {
	Apply(*User)
}

type WithName struct {
	Name string
}

func (w *WithName) Apply(u *User) {
	u.Name = w.Name
}

type WithAge struct {
	Age int32
}

func (w *WithAge) Apply(u *User) {
	u.Age = w.Age
}

func NewUser(opts ...Option) *User {
	user := new(User)
	for _, opt := range opts {
		opt.Apply(user)
	}
	return user
}

func Do() {
	NewUser(
		&WithName{Name: "l4"},
		&WithAge{Age: 20},
	)
}
