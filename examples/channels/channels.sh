# Коли ми запускаємо программу повідомлення `ping`
# успішно передається з однієї горутини до головної
# через канал повідомлень.
$ go run channels.go
ping

# Відповідно стандартних налаштувань _надсилання_
# та _отримання_ блокуватимуть виконання програми
# допоки _відправкник_ або _отримувач_ не будуть готові.
# Ця властивість дає нам змогу зачекати
# на повідомлення `ping` невикористовуючи інші
# типи синхронізації.
