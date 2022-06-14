# Тестовое задание для "constanta"

## Для запуска приложения написать `make run`

### Что было реализовано
- для реализации задачи использовался Go 1.18
- использовались только компоненты стандартной библиотеки Go
- сервер не принимает запрос если количество url в нем больше 20
- таймаут на запрос одного url - одна секунда
- сервис запускается одной командой на 8080 порту
- TODO пожалуйста, добавь README с описанием того, как сервис собирать и запускать (с использованием docker) и какого формата запрос он принимает, какой ответ нам ожидать
- по возможности добавлял комментарии

Со звездочкой:

- для каждого входящего запроса не больше 4 **одновременных** исходящих
- сервер не обслуживает больше чем 100 одновременных входящих http-запросов (ошибка с кодом 503)
- обработка запроса может быть отменена клиентом в любой момент, это должно повлечь за собой остановку всех операций связанных с этим запросом - **не сделал** (реализовал бы добавив id_request и по нему делал бы отмену контекста)
- сервис должен поддерживать 'graceful shutdown' - **не сделал**

### Формат запроса и ответа

Сервис ждет POST запрос вида

- `localhost:8080/send`
- `{
"urls": [
"https://www.google.com",
"https://www.google.com",
"https://www.google.com",
"https://www.google.com",
"https://www.google.com",
"https://www.google.com",
"https://www.google.com",
"https://www.google.com",
"https://www.google.com",
"https://www.google.com",
"https://www.google.com"
]
}`

В ответе возвращается массив тел ответов