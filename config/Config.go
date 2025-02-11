package config

import (
	appconfig "anierp/config/app_config"
	dbconfig "anierp/config/db_config"
)

func InitConfig(){
	appconfig.InitAppConfig()
	dbconfig.InitDBConfig()
}