module github.com/0xsj/kakao/conduit

go 1.22.5

require (
	github.com/0xsj/kakao/lib/core v0.0.0
	go.uber.org/dig v1.17.1 // indirect
	go.uber.org/fx v1.22.1 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	go.uber.org/zap v1.26.0 // indirect
	golang.org/x/sys v0.0.0-20220412211240-33da011f77ad // indirect
)

replace github.com/0xsj/kakao/lib/core => ../lib/core
