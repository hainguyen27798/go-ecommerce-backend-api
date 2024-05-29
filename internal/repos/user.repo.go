package repos

type UserRepo struct{}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (ur *UserRepo) GetUsers() []string {
	return []string{"hai", "harry"}
}
