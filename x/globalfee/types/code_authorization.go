package types

func (ca CodeAuthorization) Validate() error {
	return validateMethods(ca.Methods)
}
