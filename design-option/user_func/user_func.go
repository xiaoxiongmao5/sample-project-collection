package userfunc

type User struct {
	Name string
	Age  int32
}

type Option func(*User)

func NewUser(opts ...Option) *User {
	user := new(User)
	for _, opt := range opts {
		opt(user)
	}
	return user
}

func WithName(name string) Option {
	return func(u *User) {
		u.Name = name
	}
}

func WithAge(age int32) Option {
	return func(u *User) {
		u.Age = age
	}
}

func Do() {
	NewUser(
		WithName("zs"),
		WithAge(17),
	)
}
