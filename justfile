build:
    go build -o xOrgV2Adapter main.go

release version:
    git tag {{version}}
    git push origin {{version}}