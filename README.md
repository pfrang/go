to compile as ingle file

go build -o test test.go



if multiple files should be included in the bundle

go build -o import import.go export.go
