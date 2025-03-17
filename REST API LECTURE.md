---
tags:
  - golang
  - rest-api
  - rest
---
# REST API
REST API (**Representational State Transfer API**) — это **архитектурный стиль** взаимодействия клиента и сервера через HTTP, в котором данные передаются в виде ресурсов, а операции над ними выполняются стандартными методами HTTP (**GET, POST, PUT, PATCH, DELETE**).
# Gin легковесный фреймворк для создания api

**1. Что такое Gin и зачем он нужен?**

Gin — это легковесный и высокопроизводительный HTTP-фреймворк на Go, созданный для удобной разработки веб-приложений и API.

**Преимущества Gin:**

- Высокая скорость (использует net/http и эффективную маршрутизацию)
- Удобный API и middleware
- Автоматическая обработка JSON

---

**2. Установка и первый запуск**

**Установка Gin**
Перед началом работы необходимо установить Gin с помощью go get:
```sh

go get -u github.com/gin-gonic/gin

```
**Минимальный пример сервера**

Создадим main.go с простым веб-сервером:
```go
package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default() // Создание роутера с логгером и обработчиком ошибок

    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "Hello, Gin!"}) // JSON-ответ
    })

    r.Run(":8080") // Запуск сервера на порту 8080
}
```


---

**3. Маршрутизация в Gin**

Маршруты в Gin позволяют обрабатывать HTTP-запросы разных типов (GET, POST, PUT, DELETE).

**Примеры маршрутов**
```go
r.GET("/hello", func(c *gin.Context) {
    c.String(200, "Hello, World!")
})

r.POST("/submit", func(c *gin.Context) {
    name := c.PostForm("name")
    c.JSON(200, gin.H{"name": name})
})

r.PUT("/update", func(c *gin.Context) {
    c.JSON(200, gin.H{"status": "updated"})
})

r.DELETE("/delete", func(c *gin.Context) {
    c.JSON(200, gin.H{"status": "deleted"})
})
```

---

**4. URL-параметры и Query-параметры**

Gin позволяет передавать параметры через URL и query-строку.

**URL-параметры**
```go
r.GET("/user/:id", func(c *gin.Context) {
    id := c.Param("id")
    c.JSON(200, gin.H{"user_id": id})
})
```

Пример запроса
```
GET /user/123
```

**Query-параметры**
```go
r.GET("/search", func(c *gin.Context) {
    query := c.Query("id")//?id=123.
    c.JSON(200, gin.H{"search": query})
})
```

Пример запроса: 
```
GET /search?id=123
```

**5. Работа с JSON в Gin**

Gin умеет автоматически парсить JSON-запросы и отправлять JSON-ответы.

**Приём JSON-запросов**
```go
type User struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}

r.POST("/user", func(c *gin.Context) {
    var user User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, gin.H{"user": user})
})
```

# **Введение в GORM**

[GORM](https://gorm.io/) — это **ORM (Object-Relational Mapping) для Go**, позволяющая работать с базами данных удобным и декларативным способом. Она поддерживает **PostgreSQL, MySQL, SQLite, SQL Server** и другие СУБД.

---

**1. Установка GORM**

Для начала установите GORM и драйвер для нужной базы данных (например, PostgreSQL):

```sh

go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres

```

---

**2. Подключение к базе данных**

Пример подключения к **PostgreSQL**:

```go
package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	dsn := "host=localhost user=postgres password=secret dbname=mydb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Ошибка подключения к БД:", err)
	}

	log.Println("Подключение успешно!")
}
```

Замените dsn на свои реальные параметры подключения.

---

**3. Определение модели**

В GORM модель — это структура, соответствующая таблице в БД.

```go
type User struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string
	Email string `gorm:"unique"`
	Age   int
}
```

GORM автоматически создаст таблицу users с колонками id, name, email и age.

---

**4. Миграции (Создание таблиц)**

Чтобы создать таблицу, используйте:

```go
db.AutoMigrate(&User{})
```

**Важно:** AutoMigrate только добавляет новые поля, но **не удаляет и не изменяет существующие**.

---

**5. CRUD-операции в GORM**

**Создание записи**

```go
user := User{Name: "Beksultan", Email: "beks123@example.com", Age: 20}
db.Create(&user)
```

**Чтение данных**

```go
var user User
db.First(&user, 1) // Найти пользователя с ID=1
db.First(&user, "email = ?", "ivan@example.com") // Найти по email
```

**Обновление записи**

```go
db.Model(&user).Update("Age", 31)               // Обновить одно поле
db.Model(&user).Updates(User{Name: "Иван Петров", Age: 32}) // Обновить несколько полей
```

**Удаление записи**

```go
db.Delete(&user)
```

---

**6. Фильтрация и сортировка**

```go
var users []User
db.Where("age > ?", 25).Find(&users) // Выбрать всех старше 25
db.Order("age desc").Find(&users)    // Сортировать по возрасту по убыванию
```

---

**7. Работа с связями (1 к 1, 1 ко многим, многие ко многим)**

**Связь “Один ко многим” (User → Posts)**

```go
type Post struct {
	ID     uint
	Title  string
	UserID uint
}

type User struct {
	ID    uint
	Name  string
	Posts []Post `gorm:"foreignKey:UserID"`
}
```

