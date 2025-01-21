package utils

import(
    "care-vault/models"
)

var roles = models.Roles

func IsValidRole(role models.Role) bool {
    switch role {
    case roles.SystemAdmin, roles.Manager, roles.User, roles.AppOwner:
        return true
    }
    return false
}
