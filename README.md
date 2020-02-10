 # Go за прикладом [![Build Status](https://travis-ci.org/butuzov/gobyexample.svg?branch=ukrainian)](https://travis-ci.org/butuzov/gobyexample)


Наповнення та інструментарій для [Go за Прикладом](https://butuzov.github.io/gobyexample/), сайту що навчає Go за допомогою анотованих прикладів.

### Загальне

Сайт "Go за прикладом" збудовано шляхом обробки коду та коментарів отриманих з першоджерельних файлів (що знаходяться в директорії `examples`) та форматуванню їх за допомогою шаблонів (з директорій `templates`) у статичні файли (що лежатимуть у директорії `public`). Інструменти що забезпечують весь процес створення сайт знаходяться у директорії `tools`, поряд з деякими залежностями (що лежать у `vendor`).


Створена директорія `public` може буде завантажена на будь-який shared/CDN/cloud хостинг.

### Побудова сторінок сайту

Щоб побудувати сайт - вам знадобляться Go та Python. Запустіть:

```bash
> go get github.com/russross/blackfriday
> tools/build
> open public/index.html
```

Щоб будувати безперервно скористайтесь (хоч, насправді, це не так зручно):

```bash
> tools/build-loop
```

### Ліцензія

Ця робота є авторським правом Mark McGranaghan та ліцензована за
[Creative Commons Attribution 3.0 Unported License](http://creativecommons.org/licenses/by/3.0/).

Go's Гофер є авторським правом [Renée French](http://reneefrench.blogspot.com/) та ліцензовано за
[Creative Commons Attribution 3.0 Unported License](http://creativecommons.org/licenses/by/3.0/).

### Інші переклади

Переклади "Go by Example" від волонтерів доступні в наступних версіях:


* [Chinese](https://gobyexample-cn.github.io/) by [gobyexample-cn](https://github.com/gobyexample-cn)
* [Czech](http://gobyexamples.sweb.cz/) by [martinkunc](https://github.com/martinkunc/gobyexample-cz)
* [French](http://le-go-par-l-exemple.keiruaprod.fr) by [keirua](https://github.com/keirua/gobyexample)
* [Italian](http://gobyexample.it) by the [Go Italian community](https://github.com/golangit/gobyexample-it)
* [Japanese](http://spinute.org/go-by-example) by [spinute](https://github.com/spinute)
* [Korean](https://mingrammer.com/gobyexample/) by [mingrammer](https://github.com/mingrammer)
* [Russian](https://gobyexample.com.ru/) by [badkaktus](https://github.com/badkaktus)
* [Spanish](http://goconejemplos.com) by the [Go Mexico community](https://github.com/dabit/gobyexample)
* [Ukrainian](http://butuzov.github.io/gobyexample/) by [butuzov](https://github.com/butuzov/gobyexample)

### Дякуємо

Дякуюємо [Jeremy Ashkenas](https://github.com/jashkenas)
за [Docco](http://jashkenas.github.com/docco/), що надихнули на цей проект.
