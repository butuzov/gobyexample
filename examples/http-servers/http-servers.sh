# Запустити сервер в фоновому режимі
$ go run http-servers.go &

# Зробити запит до маршруту "hello"
$ curl localhost:8090/hello
hello

# Завершимо наш сервер
$ kill -9 $!
