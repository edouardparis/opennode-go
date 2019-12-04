package opennode

type Env string

const (
	Development = Env("https://dev-api.opennode.co")
	Production  = Env("https://api.opennode.co")
)
