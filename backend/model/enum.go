package model

type Difficulty string
type Role string
type Status string

const (
	DifficultyEasy   Difficulty = "easy"
	DifficultyMedium Difficulty = "medium"
	DifficultyHard   Difficulty = "hard"
)

const (
	StatusPassed    Status = "passed"
	StatusNotPassed Status = "not_passed"
)

const (
	RoleAdmin      Role = "admin"
	RoleUser       Role = "user"
	RoleSuperAdmin Role = "super_admin"
	RoleLecturer   Role = "lecturer"
)
