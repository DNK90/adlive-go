//AUTO-GENERATED: DO NOT EDIT

package config

import (
	"bytes"
	l "github.com/dnk90/adlive/helpers/log"
	"github.com/dnk90/adlive/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"strings"

	"github.com/spf13/viper"
)

var (
	Cfg *Config
	DB *gorm.DB
)

func init() {
	Load()
	DB = LoadDB()
}

var (
	ll = l.New()
)

type Tracing struct {
	Enabled bool
}

type Base struct {
	HTTPAddress int `yaml:"http_address" mapstructure:"http_address"`
	GRPCAddress int `mapstructure:"grpc_address"`
	Environment string

	Tracing Tracing
}

func Load() {
	if Cfg == nil {
		Cfg = &Config{}
		viper.SetConfigType("yaml")
		err := viper.ReadConfig(bytes.NewBuffer(defaultConfig))
		if err != nil {
			ll.Fatal("Failed to read viper config", l.Error(err))
		}
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
		viper.AutomaticEnv()
		err = viper.Unmarshal(&Cfg)
		if err != nil {
			ll.Fatal("Failed to unmarshal config", l.Error(err))
		}
		ll.Info("Config loaded", l.Object("config", Cfg))
	}
}

func LoadDB() *gorm.DB {
	if DB == nil {
		// config db
		if Cfg.Environment == "D" || Cfg.Environment == "P" {
			DB = mustConnectMySQL()
		} else {
			db, err := gorm.Open(sqlite.Open(Cfg.SQLite), &gorm.Config{
				CreateBatchSize: 10,
			})
			if err != nil {
				panic(err)
			}
			DB = db
		}
		DB = DB.Debug()
		if Cfg.Environment != "P" {
			// auto migrate
			m := []interface{}{
				new(models.User),
				new(models.Role),
				new(models.Resource),
				new(models.Permission),
				new(models.Area),
				new(models.Location),
				new(models.Campaign),
				new(models.CampaignLocation),
				new(models.Video),
				new(models.CampaignVideo),
				new(models.Screen),
				new(models.CampaignScreen),
				new(models.Log),
			}
			if err := DB.AutoMigrate(m...); err != nil {
				panic(err)
			}
		}
	}
	return DB
}

func DropDB() {
	m := []interface{}{
		new(models.CampaignScreen),
		new(models.CampaignVideo),
		new(models.CampaignLocation),
		new(models.Screen),
		new(models.Area),
		new(models.Location),
		new(models.Video),
		new(models.Campaign),
		new(models.User),
		new(models.Role),
		new(models.Resource),
		new(models.Permission),
		new(models.Log),
	}
	for _, table := range m {
		if err := DB.Migrator().DropTable(table); err != nil {
			ll.S.Errorw("error while dropping table", "err", err)
			return
		}
	}

}
