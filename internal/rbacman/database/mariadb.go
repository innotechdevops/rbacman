package database

import (
	"github.com/innotechdevops/maria-driver/pkg/mariadriver"
	"github.com/spf13/viper"
)

// NewMariaDbDriver for create mariadb drivers
func NewMariaDbDriver() mariadriver.MariaDBDriver {
	return mariadriver.New(mariadriver.Config{
		User:         viper.GetString("mariadb.user"),
		Pass:         viper.GetString("mariadb.pass"),
		Host:         viper.GetString("mariadb.host"),
		DatabaseName: viper.GetString("mariadb.database"),
		Port:         viper.GetString("mariadb.port"),
	})
}
