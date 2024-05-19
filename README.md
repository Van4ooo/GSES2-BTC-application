# `GSES2 BTC application`

## Опис
### Цей сервіс дозволяє:

 - Отримувати поточний курс біткоіна (BTC) у гривнях (UAH).
 - Підписувати електронні адреси на отримання інформації про зміну курсу.
 - Відправляти поточний курс на всі підписані електронні адреси.
Дані зберігаються у файловій системі, що виключає необхідність у використанні бази даних.

### Структура проекту
 - **main.go**: Головний файл для запуску сервісу.
 - **handlers.go**: Містить обробники для всіх HTTP-запитів.
 - **storage.go**: Містить функції для роботи з файловою системою.
 - **email.go**: Містить функції для відправки електронних листів.
 - **getRate.go**: Функції для взаємодії з api third party сервіса.
 - **api_test.go**: Файл з тестами.

## Запуск
Попередньо має бути **обов'язково** заповнений [.env](.env) файл відповідними параметрами
### Піднімаємо сервер
```shell
docker-compose up -d --build app-service
```
### Піднімаємо сервер з тестуванням api
```shell
docker-compose up -d --build
```
### Запускаємо тести окремо
```shell
docker-compose up -d --build test-api
```

## API Endpoints
Для прикладу буде використаний **python** з бібліотекою **requests**

### GET /api/rate

Цей запит повертає поточний курс BTC до UAH, використовуючи **публічний API** сервісу курсу валют.

``` python
>>> response = requests.get("http://127.0.0.1:8080/api/rate")
>>> response, response.text
(<Response [200]>, '2604430\n')
```

### POST /api/subscribe

Цей запит підписує e-mail на розсилку з інформацією про курс.

#### Параметри:

**email**: електронна адреса

```python
>>> response = requests.post("http://127.0.0.1:8080/api/subscribe", data={"email":"example1@gmail.com"})
>>> response
<Response [200]>
>>> response = requests.post("http://127.0.0.1:8080/api/subscribe", data={"email":"example1@gmail.com"})
>>> response
<Response [409]>
```

### POST /api/sendEmails

Цей запит відправляє поточний курс BTC до UAH на всі підписані електронні адреси.

```python
>>> response = requests.post("http://127.0.0.1:8080/api/sendEmails")
>>> response
<Response [200]>
```

## :)