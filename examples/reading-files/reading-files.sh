# Спочатку створемо файл, з якм ми будемо працювати надалі.
$ echo "hello" > /tmp/dat
$ echo "go" >>   /tmp/dat

# Запустимо нашу програму.
$ go run reading-files.go
hello
go
5 байт: hello
2 байт @ 6: go
2 байт @ 6: go
5 байт: hello
