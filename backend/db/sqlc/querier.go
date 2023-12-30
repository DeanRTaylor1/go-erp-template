// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type Querier interface {
	CreateAddress(ctx context.Context, arg CreateAddressParams) (Address, error)
	CreateGender(ctx context.Context, genderName string) (Gender, error)
	CreateMaritalStatus(ctx context.Context, statusName string) (MaritalStatus, error)
	CreateProfile(ctx context.Context, arg CreateProfileParams) (Profile, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	CreateUserRole(ctx context.Context, roleName string) (UserRole, error)
	CreateUserStatus(ctx context.Context, statusName string) (UserStatus, error)
	DeleteAddress(ctx context.Context, id int32) error
	DeleteGender(ctx context.Context, id int32) error
	DeleteMaritalStatus(ctx context.Context, id int32) error
	DeleteProfile(ctx context.Context, id int32) error
	DeleteUser(ctx context.Context, id int32) error
	DeleteUserRole(ctx context.Context, id int32) error
	DeleteUserStatus(ctx context.Context, id int32) error
	GetAddress(ctx context.Context, id int32) (Address, error)
	GetAddressForUpdate(ctx context.Context, id int32) (Address, error)
	GetAddressesByProfileId(ctx context.Context, profileID pgtype.Int4) ([]Address, error)
	GetAllAddresses(ctx context.Context, arg GetAllAddressesParams) ([]Address, error)
	GetGender(ctx context.Context, id int32) (Gender, error)
	GetGenderByName(ctx context.Context, genderName string) (GetGenderByNameRow, error)
	GetGenderForUpdate(ctx context.Context, id int32) (Gender, error)
	GetGenders(ctx context.Context, arg GetGendersParams) ([]Gender, error)
	GetMaritalStatus(ctx context.Context, id int32) (MaritalStatus, error)
	GetMaritalStatusByName(ctx context.Context, statusName string) (GetMaritalStatusByNameRow, error)
	GetMaritalStatusForUpdate(ctx context.Context, id int32) (MaritalStatus, error)
	GetMaritalStatuses(ctx context.Context, arg GetMaritalStatusesParams) ([]MaritalStatus, error)
	GetProfile(ctx context.Context, id int32) (Profile, error)
	GetProfileForUpdate(ctx context.Context, id int32) (Profile, error)
	GetProfiles(ctx context.Context, arg GetProfilesParams) ([]Profile, error)
	GetUser(ctx context.Context, id int32) (GetUserRow, error)
	GetUserByEmail(ctx context.Context, email string) (GetUserByEmailRow, error)
	GetUserForUpdate(ctx context.Context, id int32) (User, error)
	GetUserRole(ctx context.Context, id int32) (UserRole, error)
	GetUserRoleByName(ctx context.Context, roleName string) (GetUserRoleByNameRow, error)
	GetUserRoleForUpdate(ctx context.Context, id int32) (UserRole, error)
	GetUserRoles(ctx context.Context, arg GetUserRolesParams) ([]UserRole, error)
	GetUserStatus(ctx context.Context, id int32) (UserStatus, error)
	GetUserStatusForUpdate(ctx context.Context, id int32) (UserStatus, error)
	GetUserStatuses(ctx context.Context, arg GetUserStatusesParams) ([]UserStatus, error)
	GetUserUserStatusByName(ctx context.Context, statusName string) (GetUserUserStatusByNameRow, error)
	GetUsers(ctx context.Context, arg GetUsersParams) ([]GetUsersRow, error)
	UpdateAddress(ctx context.Context, arg UpdateAddressParams) error
	UpdateGender(ctx context.Context, arg UpdateGenderParams) error
	UpdateLastLogin(ctx context.Context, id int32) error
	UpdateMaritalStatus(ctx context.Context, arg UpdateMaritalStatusParams) error
	UpdateProfile(ctx context.Context, arg UpdateProfileParams) error
	UpdateUser(ctx context.Context, arg UpdateUserParams) error
	UpdateUserRole(ctx context.Context, arg UpdateUserRoleParams) error
	UpdateUserStatus(ctx context.Context, arg UpdateUserStatusParams) error
}

var _ Querier = (*Queries)(nil)
