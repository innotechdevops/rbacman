package user_role

type UseCase interface {
	Count(params Params) int64
	FindList(params Params) []UserRole
	FindById(id int64) UserRole
	Create(data *CreateUserRole) (UserRole, error)
	Update(data *UpdateUserRole) (UserRole, error)
	Delete(id int64) error
}

type useCase struct {
	Repo Repository
}

func (u *useCase) Count(params Params) int64 {
	return u.Repo.Count(params)
}

func (u *useCase) FindList(params Params) []UserRole {
	return u.Repo.FindList(params)
}

func (u *useCase) FindById(id int64) UserRole {
	return u.Repo.FindById(id)
}

func (u *useCase) Create(data *CreateUserRole) (UserRole, error) {
	err := u.Repo.Create(data)
	if err != nil {
		return UserRole{}, err
	}
	return u.Repo.FindById(data.Id), nil
}

func (u *useCase) Update(data *UpdateUserRole) (UserRole, error) {
	err := u.Repo.Update(data)
	if err != nil {
		return UserRole{}, err
	}
	return u.Repo.FindById(data.Id), nil
}

func (u *useCase) Delete(id int64) error {
	return u.Repo.Delete(id)
}

func NewUseCase(repo Repository) UseCase {
	return &useCase{
		Repo: repo,
	}
}
