package user_group

type UseCase interface {
	Count(params Params) int64
	FindList(params Params) []UserGroup
	FindById(id int64) UserGroup
	Create(data *CreateUserGroup) (UserGroup, error)
	Update(data *UpdateUserGroup) (UserGroup, error)
	Delete(id int64) error
}

type useCase struct {
	Repo Repository
}

func (u *useCase) Count(params Params) int64 {
	return u.Repo.Count(params)
}

func (u *useCase) FindList(params Params) []UserGroup {
	return u.Repo.FindList(params)
}

func (u *useCase) FindById(id int64) UserGroup {
	return u.Repo.FindById(id)
}

func (u *useCase) Create(data *CreateUserGroup) (UserGroup, error) {
	err := u.Repo.Create(data)
	if err != nil {
		return UserGroup{}, err
	}
	return u.Repo.FindById(data.Id), nil
}

func (u *useCase) Update(data *UpdateUserGroup) (UserGroup, error) {
	err := u.Repo.Update(data)
	if err != nil {
		return UserGroup{}, err
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
