# Stage 1: Build the Go application
FROM golang:1.23.4 AS builder

# Set the working directory for the container
WORKDIR /auth

RUN go install github.com/air-verse/air@latest

# Копируем только go.mod и go.sum для кэширования зависимостей
COPY go.mod go.sum /auth/

# Загружаем зависимости
RUN go mod download

# Устанавливаем рабочую директорию для исходников
WORKDIR /auth/cmd

COPY . /auth

# Запускаем air для live reload
CMD ["air"]
