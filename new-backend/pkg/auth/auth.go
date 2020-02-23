package auth

type Auth struct {
}

func (a *Auth) User() {
	// return false
}

func (a *Auth) ID() {
	// return false
}

func (a *Auth) Check() bool {
	return false
}

func (a *Auth) Logout() {
	// return false
}
