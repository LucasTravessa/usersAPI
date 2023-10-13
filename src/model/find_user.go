package model

import "example.com/m/src/configuration/rest_err"

func (*UserDomain) FindUser(string) *rest_err.RestErr {
	return nil
}
