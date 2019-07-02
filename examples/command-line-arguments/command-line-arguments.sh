# Для експерементів з аргементами командного рядка,
# спершу скомпілюємо файл за допомогою `go build`.
$ go build command-line-arguments.go

# Запустимо нашу програму передавши їй якісь аргументи.
$ ./command-line-arguments a b c d
[./command-line-arguments a b c d]
[a b c d]
c

# У наступному прикладі ми познайомимось з прапорцями - більш
# просунутим концептом передачі параметрів у командному рядку.
