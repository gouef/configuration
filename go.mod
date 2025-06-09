module github.com/gouef/configuration

go 1.24.2

require (
	github.com/stretchr/testify v1.10.0
	gopkg.in/yaml.v3 v3.0.1
)

replace (
	github.com/gouef/configuration/cache => ./cache
	github.com/gouef/configuration/cache/file => ./cache/file
	github.com/gouef/configuration/cache/memory => ./cache/memory
	github.com/gouef/configuration/cache/redis => ./cache/redis
	github.com/gouef/configuration/helper => ./helper
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rogpeppe/go-internal v1.13.1 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
)
