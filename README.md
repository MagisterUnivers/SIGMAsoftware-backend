# SIGMAsoftware-backend

![Tech Titans1](https://github.com/MagisterUnivers/SIGMAsoftware-backend/assets/36455862/630cb0cc-a7a4-4793-ab0b-2a2db3c9e4bc)

Привет! Меня зовут MagisterUnivers, я из группы TechTitans. Этот проект представляет собой простой сервер HTTP CRUD (Create, Read, Update, Delete) на языке программирования Go Lang, созданный в рамках тестового задания для компании SIGMA Software.

## Инструкции по запуску

1. **Установка Go:**
   Убедитесь, что на вашем компьютере установлен Go. Если нет, вы можете скачать и установить его с [официального сайта Go](https://golang.org/dl/).

2. **Клонирование репозитория:**
   В терминале выполните команду:

   ```bash
   git clone https://github.com/MagisterUnivers/SIGMAsoftware-backend.git  
   ```

3. **Запуск проекта:**
   В терминале выполните команду

   ```bash
   go run main.go
   ```

   **Использование Docker**

   <h3>Вы можете использовать образ из открытого репозитория на Docker Hub (Ссылка в описании)</h2>


   *Создание образа*

    ```bash
    docker build -t sigmago:init .
    ```

    *Использование образа*

    ```bash
    docker run -p 8080:8080 -d sigmago:init
    ```
