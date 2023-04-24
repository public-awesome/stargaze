package types

func SetCodeAuthorizationProposalFixture(mutators ...func(*SetCodeAuthorizationProposal)) *SetCodeAuthorizationProposal {
	p := &SetCodeAuthorizationProposal{
		Title:       "Foo",
		Description: "Bar",
		CodeAuthorization: &CodeAuthorization{
			CodeId:  1,
			Methods: []string{"mint", "unlist"},
		},
	}
	for _, m := range mutators {
		m(p)
	}
	return p
}

func RemoveCodeAuthorizationProposalFixture(mutators ...func(*RemoveCodeAuthorizationProposal)) *RemoveCodeAuthorizationProposal {
	p := &RemoveCodeAuthorizationProposal{
		Title:       "Foo",
		Description: "Bar",
		CodeId:      1,
	}
	for _, m := range mutators {
		m(p)
	}
	return p
}

func SetContractAuthorizationProposalFixture(mutators ...func(*SetContractAuthorizationProposal)) *SetContractAuthorizationProposal {
	const anyAddress = "cosmos1qyqszqgpqyqszqgpqyqszqgpqyqszqgpjnp7du"
	p := &SetContractAuthorizationProposal{
		Title:       "Foo",
		Description: "Bar",
		ContractAuthorization: &ContractAuthorization{
			ContractAddress: anyAddress,
			Methods:         []string{"mint", "unlist"},
		},
	}
	for _, m := range mutators {
		m(p)
	}
	return p
}

func RemoveContractAuthorizationProposalFixture(mutators ...func(*RemoveContractAuthorizationProposal)) *RemoveContractAuthorizationProposal {
	const anyAddress = "cosmos1qyqszqgpqyqszqgpqyqszqgpqyqszqgpjnp7du"
	p := &RemoveContractAuthorizationProposal{
		Title:           "Foo",
		Description:     "Bar",
		ContractAddress: anyAddress,
	}
	for _, m := range mutators {
		m(p)
	}
	return p
}
