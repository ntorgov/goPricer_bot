FROM golang:latest
LABEL authors="ntorg"

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

ENV TOKEN=my_value

# Копируем файлы проекта в рабочую директорию
COPY . .

# Собираем проект
RUN go build -o main .

# Указываем порт, который будет слушать приложение
EXPOSE 8080

# Запускаем приложение при запуске контейнера
CMD ["./main"]