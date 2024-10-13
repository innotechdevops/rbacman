package permission

type UseCase interface {
	Count(params Params) int64
	FindList(params Params) []Permission
	FindById(id int64) Permission
	Create(data *CreatePermission) (Permission, error)
	Update(data *UpdatePermission) (Permission, error)
	Delete(id int64) error
}

type useCase struct {
	Repo Repository
}

func (u *useCase) Count(params Params) int64 {
	return u.Repo.Count(params)
}

func (u *useCase) FindList(params Params) []Permission {
	return u.Repo.FindList(params)
}

func (u *useCase) FindById(id int64) Permission {
	return u.Repo.FindById(id)
}

func (u *useCase) Create(data *CreatePermission) (Permission, error) {
	err := u.Repo.Create(data)
	if err != nil {
		return Permission{}, err
	}
	return u.Repo.FindById(data.Id), nil
}

func (u *useCase) Update(data *UpdatePermission) (Permission, error) {
	err := u.Repo.Update(data)
	if err != nil {
		return Permission{}, err
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
