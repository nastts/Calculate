# Калькулятор на HTTP сервере Golang

данный проект может считать арифмитические выражения,переданные через HTTP-запрос, с использованием операций:
___
### ➕ сложение
### ➖ вычитание
### ✖️ умножение
### ➗ деление
### (...) использование скобок

> [!NOTE]
> ### в данном калькуляторе работает приоритет операций:
>  1. умножение(✖️) и деление(➗)
>  2. сложение(➕) и вычитание(➖)
> то, что находится в скобках, решается раньше того, что вне скобок.
> в скобках приоритет операций такой же

___
> [!TIP]
>данный проект был написан в **VScode**, поэтому советую тестировать проект именно там, чтобы >не было проблем с запуском и отправкой запроса

# Запуск проекта
### 1. **установите [Golang](https://go.dev/dl/)**
### 2. **сохраните [проект](https://github.com/nastts/calculate/archive/refs/heads/main.zip)**
### 3. **откройте терминал и пропишите команду для запуска сервера**
```powershell
go run ./cmd/main.go
```
### 4. **после того, как вы получили сообщение:**
```
2024/12/19 19:17:58 сервер запущен
```
>[!IMPORTANT]
>после вы должны запустить новый терминал, чтобы отправить запрос на сервер
### 5. после того, как вы открыли новый терминал, вы должны прописать:
```
Invoke-WebRequest -Method 'POST' -Uri 'http://localhost:8080/api/v1/calculate' -ContentType 'application/json' -Body '{"expression":"2+2*2"}' | Select-Object -Expand Content
```
>[!TIP]
>это очень похожее на curl, но используется для VScode

>[!IMPORTANT]
>метод должен быть POST




# Статусы сервера

### 200✅ - StatusOK 
>всё работает корректно и вы получите ответ

### 422❌ - StatusUnprocessableEntity 
>входные данные не соответствуют требованиям — например, кроме цифр и разрешённых операций присутствует символ английского алфавита. убедитесь, что ваше выражение было написано корректно

### 405❌ - StatusMethodNotAllowed
> Метод, используемый в запросе, не соответствует для данного сервера. в прошлом important сказано, какой стоит использовать. убедитесь, что используете метод POST

### 500❌ - StatusInternalServerError
>в случае какой-либо иной ошибки («Что-то пошло не так»)

# Пример ответов сервера

### 200✅ (StatusOK)

```
Invoke-WebRequest -Method 'POST' -Uri 'http://localhost:8080/api/v1/calculate' -ContentType 'application/json' -Body '{"expression":"2+2*2"}' | Select-Object -Expand Content
```
### результат
```
{"result":"6.00"}
```
___
### 422❌ (StatusMethodNotAllowed)
```
Invoke-WebRequest -Method 'POST' -Uri 'http://localhost:8080/api/v1/calculate' -ContentType 'application/json' -Body '{"expression":"2+2*"}' | Select-Object -Expand Content
```
### результат
```
Invoke-WebRequest : {"error":"expression is not valid"}
строка:1 знак:1
+ Invoke-WebRequest -Method 'POST' -Uri 'http://localhost:8080/api/v1/c ...
+ ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
    + CategoryInfo          : InvalidOperation: (System.Net.HttpWebRequest:HttpWebRequest) [Invoke-WebRequest], WebException
    + FullyQualifiedErrorId : WebCmdletWebResponseException,Microsoft.PowerShell.Commands.InvokeWebRequestCommand
```
___
### 405❌ (StatusMethodNotAllowed)
```
Invoke-WebRequest -Method 'GET' -Uri 'http://localhost:8080/api/v1/calculate' -ContentType 'application/json' -Body '{"expression":"2+2*2"}' | Select-Object -Expand Content
```

### результат
```
Invoke-WebRequest : {"error":"method not allowed"}
строка:1 знак:1
+ Invoke-WebRequest -Method 'GET' -Uri 'http://localhost:8080/api/v1/ca ...
+ ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
    + CategoryInfo          : NotSpecified: (:) [Invoke-WebRequest], ProtocolViolationException
    + FullyQualifiedErrorId : System.Net.ProtocolViolationException,Microsoft.PowerShell.Commands.InvokeWebRequestCommand
```

>[!TIP]
>### локальная **[ссылка](http://localhost:8080/api/v1/calculate)** сервера 
