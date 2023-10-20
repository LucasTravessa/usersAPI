package model

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int8
	GetID() string

	SetID(string)

	GetJSONValue() (string, error)
	EncryptPassword()
}

func NewUserDomain(email, password, name string, age int8) *userDomain {
	return &userDomain{
		email:    email,
		password: password,
		name:     name,
		age:      age,
	}
}

func NewUserUpdateDomain(name string, age int8) *userDomain {
	return &userDomain{
		name: name,
		age:  age,
	}
}
