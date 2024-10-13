package group_permission

type UseCase interface {
	Count(params Params) int64
	FindList(params Params) []GroupPermission
	FindById(id int64) GroupPermission
	Create(data *CreateGroupPermission) (GroupPermission, error)
	Update(data *UpdateGroupPermission) (GroupPermission, error)
	Delete(id int64) error
}

type useCase struct {
	Repo Repository
}

func (u *useCase) Count(params Params) int64 {
	return u.Repo.Count(params)
}

func (u *useCase) FindList(params Params) []GroupPermission {
	return u.Repo.FindList(params)
}

func (u *useCase) FindById(id int64) GroupPermission {
	return u.Repo.FindById(id)
}

func (u *useCase) Create(data *CreateGroupPermission) (GroupPermission, error) {
	err := u.Repo.Create(data)
	if err != nil {
		return GroupPermission{}, err
	}
	return u.Repo.FindById(data.Id), nil
}

func (u *useCase) Update(data *UpdateGroupPermission) (GroupPermission, error) {
	err := u.Repo.Update(data)
	if err != nil {
		return GroupPermission{}, err
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
