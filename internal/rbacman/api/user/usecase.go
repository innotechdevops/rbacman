package user

type UseCase interface {
	Count(params Params) int64
	FindList(params Params) []User
	FindById(id string) User
	Create(data *CreateUser) (User, error)
	Update(data *UpdateUser) (User, error)
	Delete(id string) error
}

type useCase struct {
	Repo Repository
}

func (u *useCase) Count(params Params) int64 {
	return u.Repo.Count(params)
}

func (u *useCase) FindList(params Params) []User {
	return u.Repo.FindList(params)
}

func (u *useCase) FindById(id string) User {
	return u.Repo.FindById(id)
}

func (u *useCase) Create(data *CreateUser) (User, error) {
	err := u.Repo.Create(data)
	if err != nil {
		return User{}, err
	}
	return u.Repo.FindById(data.Id), nil
}

func (u *useCase) Update(data *UpdateUser) (User, error) {
	err := u.Repo.Update(data)
	if err != nil {
		return User{}, err
	}
	return u.Repo.FindById(data.Id), nil
}

func (u *useCase) Delete(id string) error {
	return u.Repo.Delete(id)
}

func NewUseCase(repo Repository) UseCase {
	return &useCase{
		Repo: repo,
	}
}
