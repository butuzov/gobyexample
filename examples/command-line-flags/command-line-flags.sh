# Щоб проексперементувати з прапорцями командного рядку,
# краще всього спершу скомпілювати нашу програму і вже
# потім запускати збудований файл напряму.
$ go build command-line-flags.go

# Спробуйте спершу скомпільовану програму - надавши
# значення усім прапорцям.
$ ./command-line-flags -word=opt -numb=7 -fork -svar=flag
word: opt
numb: 7
fork: true
svar: flag
tail: []

# Зауважте, що якщо ви пропустите прапорці, вони
# автоматично приймуть свої налаштування за замовчуванням.
$ ./command-line-flags -word=opt
word: opt
numb: 42
fork: false
svar: bar
tail: []

# Заключні або позиційні аргументи можуть бути вказані після
# будь яких прапорців.
$ ./command-line-flags -word=opt a1 a2 a3
word: opt
...
tail: [a1 a2 a3]

# Зауваже що пакет `flag` потребує наявності усіх прапорців
# до позиційних аргументів (інакше прапорці будуть розпізнані
# як позиційні елементи).
$ ./command-line-flags -word=opt a1 a2 a3 -numb=7
word: opt
numb: 42
fork: false
svar: bar
tail: [a1 a2 a3 -numb=7]

# Скористайтесь прапорцями `-h` або `--help` щоб отримати автоматично
# згенеровану довідку по роботі з команднним рядком програми
# (якщо було використано пакет `flags`).
$ ./command-line-flags -h
Usage of ./command-line-flags:
  -fork=false: a bool
  -numb=42: an int
  -svar="bar": a string var
  -word="foo": a string

# Якщо ви вкажете прапорець, який не було вказано в нашій програмі,
# у відповідь, ви отримаєте повідомлення помилки та
# (повторно) тест довідки.
$ ./command-line-flags -wat
flag provided but not defined: -wat
Usage of ./command-line-flags:
...

# Далі, ми розглядемо змінні середовища, інший
# загальноприйнятий метод передачі параметрів програмам.
