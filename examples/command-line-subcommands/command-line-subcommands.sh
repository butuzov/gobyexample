$ go build command-line-subcommands.go

# Для початку викликати команду foo.
$ ./command-line-subcommands foo -enable -name=joe a1 a2
subcommand 'foo'
  enable: true
  name: joe
  tail: [a1 a2]

# Тепер викликати команду bar.
$ ./command-line-subcommands bar -level 8 a1
subcommand 'bar'
  level: 8
  tail: [a1]

# Переконатись, що bar не приймає параметри визначенні в foo.
$ ./command-line-subcommands bar -enable a1
flag provided but not defined: -enable
Usage of bar:
  -level int
    	level

# Далі будуть розглянуті змінні середовища,
# ще один поширений спосіб для заданя необхідних параметрів.
