module helloworld

go 1.15

require (
	github.com/myuser/calculator v0.0.0
	rsc.io/quote v1.5.2 // indirect
)

replace github.com/myuser/calculator => ../calculator
