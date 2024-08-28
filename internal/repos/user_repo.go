package repos

type IUserRepo interface {
	CheckUserByEmail(email string) bool
	GetUsers() []string
}

type userRepo struct{}

func (ur userRepo) GetUsers() []string {
	return []string{"hai", "harry"}
}

func (ur userRepo) CheckUserByEmail(email string) bool {
	//TODO implement me
	panic("implement me")
}

func NewUserRepo() IUserRepo {
	return &userRepo{}
}
