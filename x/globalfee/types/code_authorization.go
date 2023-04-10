package types

// Validate validates the values of code authorizations
func (ca CodeAuthorization) Validate() error {
	return validateMethods(ca.Methods)
}
