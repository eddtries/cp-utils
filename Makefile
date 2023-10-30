pre_publish:
	go mod tidy
	go test ./...

post_publish:
	go list -m eddtries/cp-utils