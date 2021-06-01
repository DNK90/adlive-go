package acl

import (
	"errors"
	l "github.com/dnk90/adlive/helpers/log"
	"github.com/dnk90/adlive/internal/models"
	"github.com/dnk90/adlive/internal/stores"
	"github.com/dnk90/adlive/pb"
	"gorm.io/gorm"
)

const (
	ANY = -1
	ANY_STRING = "*"
	ALLOWED = 1
	DENIED = 2
)

type Base interface {
	Add(params ...interface{}) (int, error)
	Exist(params ...interface{}) bool
	Get(output interface{}, params ...interface{}) (bool, error)
}

type Resource interface {
	Base
	GetParent(roleId int) ([]int, error)
}

var (
	ll = l.New()
	InvalidValue = errors.New("invalid value")
	acl *ACL
)

type ACL struct {
	ResourceStore Resource
	RoleStore Resource
	PermissionStore Base
}

func Get() *ACL {
	if acl == nil {
		acl = &ACL{
			ResourceStore:   stores.NewResourceStore(),
			RoleStore:       stores.NewRoleStore(),
			PermissionStore: stores.NewPermissionStore(),
		}
	}
	return acl
}

func (a *ACL) AddRole(role, parent string) (int, error) {
	return a.RoleStore.Add(role, parent)
}

func (a *ACL) AddResource(resource, parent string) (int, error) {
	return a.ResourceStore.Add(resource, parent)
}

func (a *ACL) AddPermission(roleName, resourceName string, allowed int) (int, error) {
	role := new(models.Role)
	resource := new(models.Resource)
	if roleName != ANY_STRING {
		existed, err := a.RoleStore.Get(&role, roleName)
		if err != nil {
			return 0, InvalidValue
		}
		if !existed {
			return 0, gorm.ErrRecordNotFound
		}
	}
	if resourceName != ANY_STRING {
		existed, err := a.ResourceStore.Get(&resource, resourceName)
		if err != nil {
			return 0, InvalidValue
		}
		if !existed {
			return 0, gorm.ErrRecordNotFound
		}
	}
	id, err := a.PermissionStore.Add(roleName, resourceName, allowed)
	if err != nil {
		ll.S.Errorw("error while adding new permission", "err", err)
		return 0, err
	}
	return id, nil
}

func (a *ACL) Query(roleId, resourceId, allowed int) bool {
	roles := make([]int, 0)
	resources := make([]int, 0)

	roles = append(roles, ANY)
	resources = append(resources, ANY)

	if roleId > 0 {
		parents, err := a.RoleStore.GetParent(roleId)
		if err != nil {
			ll.S.Errorw("error while getting role's parent", "roleId", roleId, "err", err)
			return false
		}
		roles = append(roles, parents...)
	}

	if resourceId > 0 {
		parents, err := a.ResourceStore.GetParent(resourceId)
		if err != nil {
			ll.S.Errorw("error while getting resource's parent", "resourceId", resourceId, "err", err)
			return false
		}
		resources = append(resources, parents...)
	}

	for _, role := range roles {
		for _, resource := range resources {
			var permission *models.Permission
			_, err := a.PermissionStore.Get(&permission, role, resource, allowed, int(pb.Status_activated))
			if err != nil {
				ll.S.Errorw("error while getting permission", "err", err)
				continue
			}
			if permission.Id > 0 {
				return true
			}
		}
	}
	return false
}
