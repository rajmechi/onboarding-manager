package controller

import (
	"github.com/rajmechi/onboarding-manager/pkg/controller/user"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, user.Add)
}
