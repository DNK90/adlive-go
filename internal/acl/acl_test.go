package acl

import (
	"fmt"
	"github.com/dnk90/adlive/config"
	"github.com/dnk90/adlive/internal/stores"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

type base struct {
	name string
	parent string
}

type permission struct {
	role string
	resource string
}

type ACLTestSuite struct {
	suite.Suite
	acl *ACL
	resources []*base
	resourceMap map[string]int
	roles []*base
	rolesMap map[string]int
	permissions []*permission
	permissionMap map[string]int
	do chan bool
}

func (s *ACLTestSuite) SetupSuite() {
	// remove file test db
	os.Remove(config.Cfg.SQLite)

	config.Cfg.Environment = "T"
	s.acl = &ACL{
		ResourceStore:   stores.NewResourceStore(),
		RoleStore:       stores.NewRoleStore(),
		PermissionStore: stores.NewPermissionStore(),
	}
	s.rolesMap = make(map[string]int)
	s.resourceMap = make(map[string]int)

	s.resources = []*base{
		{
			name: "login",
		},
		{
			name: "action1",
		},
		{
			name:   "action2",
			parent: "action1",
		},
		{
			name: "action3",
		},
	}
	s.roles = []*base{
		{
			name: "user",
		},
		{
			name:   "moderator",
			parent: "user",
		},
		{
			name:   "super moderator",
			parent: "moderator",
		},
		{
			name: "admin",
		},
	}
	s.permissions = []*permission{
		{
			role: "user",
			resource: "action1",
		},
		{
			role: "*",
			resource: "login",
		},
		{
			role: "moderator",
			resource: "action2",
		},
		{
			role: "super moderator",
			resource: "action3",
		},
		{
			role: "admin",
			resource: "*",
		},
	}
}

func (s *ACLTestSuite) TearDownSuite() {
	// remove file test db
	os.Remove(config.Cfg.SQLite)
}

func (s *ACLTestSuite) Test_1_AddRoles() {
	// add resources and roles
	for _, r := range s.roles {
		id, err := s.acl.RoleStore.Add(r.name, r.parent)
		s.NoError(err)
		println(fmt.Sprintf("roleId=%d", id))
		s.rolesMap[r.name] = id
	}
}

func (s *ACLTestSuite) Test_2_AddResources() {
	for _, resource := range s.resources {
		id, err := s.acl.ResourceStore.Add(resource.name, resource.parent)
		s.NoError(err)
		println(fmt.Sprintf("resourceId=%d", id))
		s.resourceMap[resource.name] = id
	}
}

func (s *ACLTestSuite) Test_3_AddPermissions() {
	for _, p := range s.permissions {
		s.True(p.role == ANY_STRING || s.rolesMap[p.role] > 0)
		s.True(p.resource == ANY_STRING || s.resourceMap[p.resource] > 0)

		var roleId, resourceId int
		if p.role == ANY_STRING {
			roleId = -1
		} else {
			roleId = s.rolesMap[p.role]
		}
		if p.resource == ANY_STRING {
			resourceId = -1
		} else {
			resourceId = s.resourceMap[p.resource]
		}
		id, err := s.acl.PermissionStore.Add(roleId, resourceId, ALLOWED)
		s.NoError(err)
		println(fmt.Sprintf("permissionId=%d", id))
	}
}

func (s *ACLTestSuite) Test_4_PermissionTest() {
	type strategy struct {
		description string
		role int
		resource int
		expected bool
	}
	strategies := []*strategy{
		{
			description: "test login permission",
			role: -1,
			resource: s.resourceMap["login"],
			expected: true,
		},
		{
			description: "test action1 permission with role user",
			role: s.rolesMap["user"],
			resource: s.resourceMap["action1"],
			expected: true,
		},
		{
			description: "test action2 permission with role user",
			role: s.rolesMap["user"],
			resource: s.resourceMap["action2"], // action2 is action1's parent
			expected: true,
		},
		{
			description: "test action3 permission with role user",
			role: s.rolesMap["user"],
			resource: s.resourceMap["action3"],
			expected: false,
		},
		{
			description: "test action1 permission with role moderator",
			role: s.rolesMap["moderator"],
			resource: s.resourceMap["action1"],
			expected: true,
		},
		{
			description: "test action2 permission with role moderator",
			role: s.rolesMap["moderator"],
			resource: s.resourceMap["action2"],
			expected: true,
		},
		{
			description: "test action2 permission with role super moderator",
			role: s.rolesMap["super moderator"],
			resource: s.resourceMap["action2"],
			expected: true,
		},
		{
			description: "test action3 permission with role super moderator",
			role: s.rolesMap["super moderator"],
			resource: s.resourceMap["action3"],
			expected: true,
		},
		{
			description: "test action3 permission with role moderator",
			role: s.rolesMap["moderator"],
			resource: s.resourceMap["action3"],
			expected: false,
		},
		{
			description: "test action1 permission with role admin",
			role: s.rolesMap["admin"],
			resource: s.resourceMap["action1"],
			expected: true,
		},
		{
			description: "test action2 permission with role admin",
			role: s.rolesMap["admin"],
			resource: s.resourceMap["action2"],
			expected: true,
		},
		{
			description: "test action3 permission with role admin",
			role: s.rolesMap["admin"],
			resource: s.resourceMap["action3"],
			expected: true,
		},
		{
			description: "test login permission with role admin",
			role: s.rolesMap["admin"],
			resource: s.resourceMap["login"],
			expected: true,
		},
	}
	for _, str := range strategies {
		println(fmt.Sprintf("%s-%d-%d", str.description, str.role, str.resource))
		result := s.acl.Query(str.role, str.resource, ALLOWED)
		s.Equal(str.expected, result)
	}
}

func TestACL(t *testing.T) {
	suite.Run(t, new(ACLTestSuite))
}


