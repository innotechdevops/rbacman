package user_group

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewDataSource,
	NewRepository,
	NewUseCase,
	NewHandler,
	NewRouter,
	NewValidate,
)
