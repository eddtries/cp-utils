pre_publish:
	go mod tidy
	go test ./...

# publish:
# 	git commit -m "eddtries/cputils: changes for $(VERSION)"
# 	git tag $(VERSION)
# 	git push origin $(VERSION)

post_publish:
	go list -m github.com/eddtries/cp-utils