package models

type Role string

var Roles = struct {
    SystemAdmin Role
    Manager Role
    User Role
    AppOwner Role
} {
    SystemAdmin: "SYSTEM_ADMIN",
    Manager: "MANAGER",
    User: "USER",
    AppOwner: "APP_OWNER",
}
