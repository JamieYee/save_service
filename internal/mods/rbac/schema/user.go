package schema

import (
	"time"

	"github.com/JamieYee/save_service/internal/config"
	"github.com/JamieYee/save_service/pkg/crypto/hash"
	"github.com/JamieYee/save_service/pkg/errors"
	"github.com/JamieYee/save_service/pkg/util"
)

const (
	UserStatusActivated = "activated"
	UserStatusFreezed   = "freezed"
)

// User management for RBAC
type User struct {
	ID        string    `json:"id" gorm:"size:20;primarykey;"` // Unique ID
	Username  string    `json:"username" gorm:"size:64;index"` // Username for login
	Password  string    `json:"-" gorm:"size:64;"`             // Password for login (encrypted)
	Nickname  string    `json:"nickname" gorm:"size:64;index"` // NickName of user
	Avatar    string    `json:"avatar" gorm:"size:256;"`       // Avatar of user
	Remark    string    `json:"remark" gorm:"size:1024;"`      // Remark of user
	Status    string    `json:"status" gorm:"size:20;index"`   // Status of user (activated, freezed)
	CreatedAt time.Time `json:"created_at" gorm:"index;"`      // Create time
	UpdatedAt time.Time `json:"updated_at" gorm:"index;"`      // Update time
	Roles     UserRoles `json:"roles" gorm:"-"`                // Roles of user
}

func (a *User) TableName() string {
	return config.C.FormatTableName("user")
}

// Defining the query parameters for the `User` struct.
type UserQueryParam struct {
	util.PaginationParam
	LikeUsername string `form:"username"`                                    // Username for login
	LikeNickname string `form:"nickname"`                                    // Name of user
	Status       string `form:"status" binding:"oneof=activated freezed ''"` // Status of user (activated, freezed)
}

// Defining the query options for the `User` struct.
type UserQueryOptions struct {
	util.QueryOptions
}

// Defining the query result for the `User` struct.
type UserQueryResult struct {
	Data       Users
	PageResult *util.PaginationResult
}

// Defining the slice of `User` struct.
type Users []*User

func (a Users) ToIDs() []string {
	var ids []string
	for _, item := range a {
		ids = append(ids, item.ID)
	}
	return ids
}

// Defining the data structure for creating a `User` struct.
type UserForm struct {
	Username string    `json:"username" binding:"required,max=64"`                // Username for login
	Password string    `json:"password" binding:"max=64"`                         // Password for login (md5 hash)
	Nickname string    `json:"nickname" binding:"max=64"`                         // NickName of user
	Avatar   string    `json:"avatar" binding:"max=256"`                          // Avatar of user
	Remark   string    `json:"remark" binding:"max=1024"`                         // Remark of user
	Status   string    `json:"status" binding:"required,oneof=activated freezed"` // Status of user (activated, freezed)
	Roles    UserRoles `json:"roles" binding:"required"`                          // Roles of user
}

// A validation function for the `UserForm` struct.
//func (a *UserForm) Validate() error {
//	if a.NickName != "" && validator.New().Var(a.NickName, "nickname") != nil {
//		return errors.BadRequest("", "Invalid email address")
//	}
//	return nil
//}

// Convert `UserForm` to `User` object.
func (a *UserForm) FillTo(user *User) error {
	user.Username = a.Username
	user.Nickname = a.Nickname
	user.Avatar = a.Avatar
	user.Remark = a.Remark
	user.Status = a.Status

	if avatar := a.Avatar; avatar != "" {
		user.Avatar = avatar
	}

	if pass := a.Password; pass != "" {
		hashPass, err := hash.GeneratePassword(pass)
		if err != nil {
			return errors.BadRequest("", "Failed to generate hash password: %s", err.Error())
		}
		user.Password = hashPass
	}

	return nil
}
