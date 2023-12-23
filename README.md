# GoQuiz

A quiz `API` written in `Go` with 
[Fiber](https://github.com/gofiber/fiber).

It is completely open-source, so feel free to download
and build a quiz using this project. Just give credit
to me, please – being nice is greatly appreciated. ;)

## How it operates?

To begin, let me illustrate the folder structure I
used in developing this application.

```
📦 GoQuiz project
    📂 app
        📂 handlers
        📂 repository
        📂 requests
    📂 database
        📂 models
    📂 http
    📂 helpers
    📂 routes
    📜 .env
    📜 main.go
```

I established a "handlers" folder where I defined
various structs and their associated methods. Subsequently,
I registered these handlers in the `routes/api.go` file. Therefore,
when a request is sent to the server, Fiber autonomously
handles the process, invokes the designated handler method,
and returns the corresponding response. That's all.

Middleware, authentication, and session management are defined
in the `http` folder.

For the database, I am utilizing the `gorm` package 
with a MySQL database.

### Credits

[Martin Binder](https://mrtn.vip) - This is my first solo project with Go and Fiber.
