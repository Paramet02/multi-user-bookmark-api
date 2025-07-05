package security

type PasswordHasher interface{
	Hash(password string) (string , error)
	Compare(password , hash string) error
}

type PasswordPolicy interface{
	Validate(password string) error
}