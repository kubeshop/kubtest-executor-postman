module github.com/kubeshop/kubtest-executor-postman

go 1.16

// replace github.com/kubeshop/kubtest-operator v0.1.1 => ../kubtest-operator
// replace github.com/kubeshop/kubtest v0.5.22 => ../kubtest

require (
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/kubeshop/kubtest v0.5.26-beta001
	github.com/stretchr/testify v1.7.0
)
