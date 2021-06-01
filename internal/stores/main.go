package stores

import (
	l "github.com/dnk90/adlive/helpers/log"
	"github.com/dnk90/adlive/internal/models"
	"gorm.io/gorm"
)

var ll = l.New()

type MainStore struct {
	*gorm.DB
}

// NewMainStore creates a new instance of MainStore, contains whole common functions
// for a service

func NewMainStore() *MainStore {
	return &MainStore{}
}

type IRoleStore interface {
	AddRole(roleName, parentName string) (int, error)
	Exist(roleName string) bool
	GetRole(roleName string) (*models.Role, bool, error)
	GetParent(roleName string) ([]int, error)
}

type IResourceStore interface {
	AddResource(roleName, parentName string) (int, error)
	Exist(roleName string) bool
	GetResource(roleName string) (*models.Resource, bool, error)
	GetParent(roleName string) ([]int, error)
}

type IPermissionStore interface {
	AddPermission(role, resource int, allowed int) (int, error)
	Exist(role, resource int) bool
	GetPermission(roleName, resourceName int, allowed, status int) (*models.Permission, bool, error)
}

type IStore interface {
	Save(data interface{}) error
	Update(data interface{}) error
	Query(output interface{}, params ...interface{}) error
	DB(params ...interface{}) interface{}
}
