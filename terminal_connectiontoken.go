package stripe

// TerminalConnectionTokenParams is the set of parameters that can be used when creating a terminal connection token.
type TerminalConnectionTokenParams struct {
	Params `form:"*"`

	// This feature has been deprecated and should not be used anymore.
	OperatorAccount *string `form:"operator_account"`
}

// TerminalConnectionToken is the resource representing a Stripe terminal connection token.
type TerminalConnectionToken struct {
	Object string `json:"object"`
	Secret string `json:"secret"`
}
