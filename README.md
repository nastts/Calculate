# Калькулятор на HTTP сервере Golang

#### данный проект может считать арифмитические выражения,переданные через HTTP-запрос, с использованием операций:
### ➕ сложение
### ➖ вычитание
### ✖️ умножение
### ➗ деление
### (...) использование скобок

> [!NOTE]
> ### в данном калькуляторе работает приоритет операций:
>  1. умножение(✖️) и деление(➗)
>  2. сложение(➕) и вычитание(➖)

>[!TIP]
> ### то, что находится в скобках, решается раньше того, что вне скобок. в скобках приоритет операций такой же

> [!IMPORTANT]
>### данный проект был написан на **Windows**, поэтому запуск проекта ниже написан для **Windows**

>[!TIP]
>### локальная **[ссылка](http://localhost:8080/api/v1/calculate)** сервера 

>[!TIP]
>### end-point сервера
>/api/v1/calculate



# Запуск проекта
### 1. **установите [Golang](https://go.dev/dl/)**
### 2. **сохраните [проект](https://github.com/nastts/Calculate/archive/refs/heads/main.zip)**
### 3. **откройте терминал и пропишите команду для запуска сервера**
```powershell
go run ./cmd/main.go
```
### 4. **если вы получили сообщение:**
```Go
2024/12/19 19:17:58 сервер запущен
```
### **значит сервер запустился корректно**
>[!IMPORTANT]
>после того, как вы запустили сервер, создайте новый терминал, что отправить запрос
### 5. после того, как вы открыли новый терминал, вы должны прописать:
```
Invoke-WebRequest -Method 'POST' -Uri 'http://localhost:8080/api/v1/calculate' -ContentType 'application/json' -Body '{"expression":"2+2*2"}' | Select-Object -Expand Content
```
>[!TIP]
>это очень похожее на curl, но используется на **Windows**

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
### 422❌ (StatusUnprocessableEntity) ошибка пользователя в написании выражения
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
### 405❌ (StatusMethodNotAllowed) не правильный метод
```
Invoke-WebRequest -Method 'GET' -Uri 'http://localhost:8080/api/v1/calculate' -ContentType 'application/json' -Body '{"expression":"2+2*2"}' | Select-Object -Expand Content
```

### результат
```
Invoke-WebRequest : Невозможно отправить тело содержимого с данным типом предиката.
строка:1 знак:1
+ Invoke-WebRequest -Method 'GET' -Uri 'http://localhost:8080/api/v1/ca ...
+ ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
    + CategoryInfo          : NotSpecified: (:) [Invoke-WebRequest], ProtocolViolationException
    + FullyQualifiedErrorId : System.Net.ProtocolViolationException,Microsoft.PowerShell.Commands.InvokeWebRequestCommand
```
