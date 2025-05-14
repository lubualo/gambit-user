git add .
git commit -m "New commit"
git push

set GOOS=linux
set GOARCH=amd64

go build -o bootstrap main.go
del main.zip
zip function.zip bootstrap
