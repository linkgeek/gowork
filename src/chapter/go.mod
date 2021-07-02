module chapter

go 1.15

require (
	github.com/go-redis/redis v6.15.9+incompatible // indirect
	github.com/go-redis/redis/v8 v8.9.0 // indirect
	github.com/myuser/calculator v0.0.0
	rsc.io/quote v1.5.2
)

replace github.com/myuser/calculator => ../calculator
