build:
   jekyll build
   mkdir -p functions
   GOOS=linux
   GOARCH=amd64
   GOBIN=${PWD}/functions go get ./...
   GOBIN=${PWD}/functions go install ./...