// YOU CAN EDIT YOUR CUSTOM CONFIG HERE

package config

import (
	"bytes"
	l "github.com/dnk90/adlive/helpers/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/url"
	"sort"
	"time"
)

// Config holds all settings
var defaultConfig = []byte(`
environment: D
grpc_address: 10000
http_address: 9000

mysql:
  address: localhost:3306
  database: adlive
  username: root
  password: 12345
  protocol: tcp
  parse_time: true
  batches: 100

jwt:
  algorithm: 
  secret: 3924e972-ba55-4730-9962-5deb0c22e504
  refresh_secret: 61b3c0a6-be87-11eb-8529-0242ac130003
  expire: 60
  refresh_expire: 10080

sqlite: data.db
`)

type Config struct {
	Base `mapstructure:",squash"`
	// Custom here
	MySQL                    *MySQL `yaml:"mysql" mapstructure:"mysql"`
	DB                       *gorm.DB
	JWT      JWT       `yaml:"jwt" mapstructure:"jwt"`
	SQLite   string    `yaml:"sqlite" mapstructure:"sqlite"`
}

type JWT struct {
	Secret string `yaml:"secret" mapstructure:"secret"`
	RefreshSecret string `yaml:"refresh_secret" mapstructure:"refresh_secret"`
	Expire int    `yaml:"expire" mapstructure:"expire"`
	RefreshExpire int `yaml:"refresh_expire" mapstructure:"refresh_expire"`
	Name string `yaml:"name" mapstructure:"name"`
}

// MySQL is settings of a MySQL server. It contains almost same fields as mysql.Config,
// but with some different field names and tags.
type MySQL struct {
	Username  string            `yaml:"username" mapstructure:"username"`
	Password  string            `yaml:"password" mapstructure:"password"`
	Protocol  string            `yaml:"protocol" mapstructure:"protocol"`
	Address   string            `yaml:"address" mapstructure:"address"`
	Database  string            `yaml:"database" mapstructure:"database"`
	Params    map[string]string `yaml:"params" mapstructure:"params"`
	Collation string            `yaml:"collation" mapstructure:"collation"`
	Loc       *time.Location    `yaml:"location" mapstructure:"loc"`
	TLSConfig string            `yaml:"tls_config" mapstructure:"tls_config"`

	Timeout      time.Duration `yaml:"timeout" mapstructure:"timeout"`
	ReadTimeout  time.Duration `yaml:"read_timeout" mapstructure:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout" mapstructure:"write_timeout"`

	ConnMaxLifetime int `yaml:"conn_max_life_time" mapstructure:"conn_max_life_time"`
	MaxOpenConns    int `yaml:"max_open_conns" mapstructure:"max_open_conns"`
	MaxIdleConns    int `yaml:"max_idle_conns" mapstructure:"max_idle_conns"`

	AllowAllFiles           bool   `yaml:"allow_all_files" mapstructure:"allow_all_files"`
	AllowCleartextPasswords bool   `yaml:"allow_cleartext_passwords" mapstructure:"allow_cleartext_passwords"`
	AllowOldPasswords       bool   `yaml:"allow_old_passwords" mapstructure:"allow_old_passwords"`
	ClientFoundRows         bool   `yaml:"client_found_rows" mapstructure:"client_found_rows"`
	ColumnsWithAlias        bool   `yaml:"columns_with_alias" mapstructure:"columns_with_alias"`
	InterpolateParams       bool   `yaml:"interpolate_params" mapstructure:"interpolate_params"`
	MultiStatements         bool   `yaml:"multi_statements" mapstructure:"multi_statements"`
	ParseTime               bool   `yaml:"parse_time" mapstructure:"parse_time"`
	GoogleAuthFile          string `yaml:"google_auth_file" mapstructure:"google_auth_file"`

	Batches int `yaml:"batches" mapstructure:"batches"`
}

const defaultCollation = "utf8mb4_general_ci"

func (cfg *MySQL) FormatDSN() string {
	var buf bytes.Buffer

	// [username[:password]@]
	if len(cfg.Username) > 0 {
		buf.WriteString(cfg.Username)
		if len(cfg.Password) > 0 {
			buf.WriteByte(':')
			buf.WriteString(cfg.Password)
		}
		buf.WriteByte('@')
	}

	// [protocol[(address)]]
	if len(cfg.Protocol) > 0 {
		buf.WriteString(cfg.Protocol)
		if len(cfg.Address) > 0 {
			buf.WriteByte('(')
			buf.WriteString(cfg.Address)
			buf.WriteByte(')')
		}
	}

	// /dbname
	buf.WriteByte('/')
	buf.WriteString(cfg.Database)

	// [?param1=value1&...&paramN=valueN]
	hasParam := false

	if cfg.AllowAllFiles {
		hasParam = true
		buf.WriteString("?allowAllFiles=true")
	}

	if col := cfg.Collation; col != defaultCollation && len(col) > 0 {
		writeDSNParam(&buf, &hasParam, "collation", col)
	}

	if cfg.ColumnsWithAlias {
		writeDSNParam(&buf, &hasParam, "columnsWithAlias", "true")
	}

	if cfg.InterpolateParams {
		writeDSNParam(&buf, &hasParam, "interpolateParams", "true")
	}

	if cfg.Loc != time.UTC && cfg.Loc != nil {
		writeDSNParam(&buf, &hasParam, "loc", url.QueryEscape(cfg.Loc.String()))
	}

	if cfg.MultiStatements {
		writeDSNParam(&buf, &hasParam, "multiStatements", "true")
	}

	if cfg.ParseTime {
		writeDSNParam(&buf, &hasParam, "parseTime", "true")
	}

	if cfg.ReadTimeout > 0 {
		writeDSNParam(&buf, &hasParam, "readTimeout", cfg.ReadTimeout.String())
	}

	if cfg.Timeout > 0 {
		writeDSNParam(&buf, &hasParam, "timeout", cfg.Timeout.String())
	}

	if len(cfg.TLSConfig) > 0 {
		writeDSNParam(&buf, &hasParam, "tls", url.QueryEscape(cfg.TLSConfig))
	}

	if cfg.WriteTimeout > 0 {
		writeDSNParam(&buf, &hasParam, "writeTimeout", cfg.WriteTimeout.String())
	}

	// other params
	if cfg.Params != nil {
		var params []string
		for param := range cfg.Params {
			params = append(params, param)
		}
		sort.Strings(params)
		for _, param := range params {
			writeDSNParam(&buf, &hasParam, param, url.QueryEscape(cfg.Params[param]))
		}
	}

	return buf.String()
}

func writeDSNParam(buf *bytes.Buffer, hasParam *bool, name, value string) {
	buf.Grow(1 + len(name) + 1 + len(value))
	if !*hasParam {
		*hasParam = true
		buf.WriteByte('?')
	} else {
		buf.WriteByte('&')
	}
	buf.WriteString(name)
	buf.WriteByte('=')
	buf.WriteString(value)
}

// MySQLDSN returns the MySQL DSN from config.
func (cfg *Config) MySQLDSN() string {
	return cfg.MySQL.FormatDSN()
}

func mustConnectMySQL() *gorm.DB {
	if Cfg.DB != nil {
		return Cfg.DB
	}
	db, err := gorm.Open(mysql.Open(Cfg.MySQLDSN()), &gorm.Config{
		CreateBatchSize: Cfg.MySQL.Batches,
	})
	if err != nil {
		panic(err)
	}

	if Cfg.Environment == "D" {
		db = db.Debug()
	}

	sqlDb, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDb.SetConnMaxLifetime(time.Duration(Cfg.MySQL.ConnMaxLifetime) * time.Hour)
	sqlDb.SetMaxIdleConns(Cfg.MySQL.MaxIdleConns)
	sqlDb.SetMaxOpenConns(Cfg.MySQL.MaxOpenConns)

	err = db.Raw("SELECT 1").Error
	if err != nil {
		ll.Fatal("error querying SELECT 1", l.Error(err))
	}

	return db
}