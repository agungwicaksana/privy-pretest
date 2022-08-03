package container

import (
	"github.com/agungwicaksana/privy-pretest/internal/domain"
	"github.com/agungwicaksana/privy-pretest/internal/infrastructure/mysql"
	"github.com/agungwicaksana/privy-pretest/internal/interface/usecase/cake"
	"github.com/agungwicaksana/privy-pretest/internal/migration"
	"github.com/agungwicaksana/privy-pretest/pkg/config"
	_ "github.com/go-sql-driver/mysql"
)

type Container struct {
	AppConfig   *config.AppConfig
	CakeService cake.CakeService
}

func Setup() *Container {
	config.Load()
	env := config.AssignEnv()
	appConfig := &config.AppConfig{
		AppPort:          config.GetString("APP_PORT", env),
		MySqlUri:         config.GetString("MYSQL_URI", env),
		MySqlTimeout:     config.GetInt("MYSQL_TIMEOUT", env),
		MySqlMaxPool:     config.GetInt("MYSQL_MAX_POOL", env),
		MySqlMinPoolSize: config.GetInt("MYSQL_MIN_POOL", env),
		MySqlDebugMode:   config.GetBool("MYSQL_DEBUG_MODE", env),
	}
	config.ValidateEnv(appConfig)

	db := mysql.NewMySQL(*appConfig)
	migration.NewMigration(db).RegisterSchemas(
		migration.CakeSchema,
	).Migrate()

	cakeRepo := domain.NewCakeRepo(db)
	cakeService := cake.NewService().SetCakeRepo(cakeRepo).Validate()

	container := &Container{
		AppConfig:   appConfig,
		CakeService: cakeService,
	}
	return container
}
