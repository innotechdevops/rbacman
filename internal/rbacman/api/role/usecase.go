package role

type UseCase interface {
	Count(params Params) int64
	FindList(params Params) []Role
	FindById(id int64) Role
	Create(data *CreateRole) (Role, error)
	Update(data *UpdateRole) (Role, error)
	Delete(id int64) error
}

type useCase struct {
	Repo Repository
}

func (u *useCase) Count(params Params) int64 {
	return u.Repo.Count(params)
}

func (u *useCase) FindList(params Params) []Role {
	return u.Repo.FindList(params)
}

func (u *useCase) FindById(id int64) Role {
	return u.Repo.FindById(id)
}

func (u *useCase) Create(data *CreateRole) (Role, error) {
	err := u.Repo.Create(data)
	if err != nil {
		return Role{}, err
	}
	return u.Repo.FindById(data.Id), nil
}

func (u *useCase) Update(data *UpdateRole) (Role, error) {
	err := u.Repo.Update(data)
	if err != nil {
		return Role{}, err
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
