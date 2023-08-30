# Тестовое задание для стажера Golang
## Сервис динамического сегментирования пользователей

**Выполнены все дополнительные задания:**

**1) Экспорт истории в .csv**

**2) Создание и отложенное удаление сегментов (pg_cron)**

**3) Рандомное добавление пользователей в сегмент**

"""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""

Сервис выполнен в соответствии с чистой архитектурой.

### Запуск:

docker-compose-up

Все стартует:
![Screenshot from 2023-08-30 22-03-26](https://github.com/LittleMikle/avito_tech_2023/assets/101155101/f718e64b-0db6-4d8e-bed8-cd172aa739a6)

В соответствии с пунктом 4, где указано что механизм миграций не нужен, SQL файл лежит по пути:

/schema/000001_init.up.sql

### Swagger:

swag init -g cmd/main.go

### Тесты:

go test -v ./...

## Postman коллекция лежит в корневой папке

Создание сегмента:

![Screenshot from 2023-08-30 20-36-21](https://github.com/LittleMikle/avito_tech_2023/assets/101155101/62abc9d9-4bd4-48f4-af4b-9807f9bddb47)

Удаление сегмента:

![Screenshot from 2023-08-30 20-37-06](https://github.com/LittleMikle/avito_tech_2023/assets/101155101/d62b0848-bde1-41c3-b62d-8b6f9d9b7e43)

Добавление сегментов пользователю:

![Screenshot from 2023-08-30 20-39-48](https://github.com/LittleMikle/avito_tech_2023/assets/101155101/42960f85-ce2b-44c5-8b65-8415c26d0572)

Пользователь добавляется, при повторном запросе записи не перетираютя (есть проверка):
![Screenshot from 2023-08-30 20-41-54](https://github.com/LittleMikle/avito_tech_2023/assets/101155101/f0aafb7d-4efb-42fa-92e3-a36c09ecd594)

Также появились записи о добавлении в таблицу с историей:
![Screenshot from 2023-08-30 20-42-02](https://github.com/LittleMikle/avito_tech_2023/assets/101155101/d8754873-d6dc-4b5c-a371-7a430e9ce39e)

Удаление нескольких сегментов у пользователя:
![Screenshot from 2023-08-30 20-44-57](https://github.com/LittleMikle/avito_tech_2023/assets/101155101/15dd9473-0ff0-424c-8fee-b7a70df94563)

Добавились записи об удалении в таблицу с историей:
![Screenshot from 2023-08-30 20-45-50](https://github.com/LittleMikle/avito_tech_2023/assets/101155101/270e92d6-5762-43c7-8887-eeae467f115b)

Получение всех активных сегментов пользователя:
![Screenshot from 2023-08-30 20-47-56](https://github.com/LittleMikle/avito_tech_2023/assets/101155101/67468325-c398-4f24-91d1-bdd9e27b5aaf)

## Дополнительные задания:

**1) Экспорт истории пользователя в .csv**

Создается файл папке проекта с названием resultUser и уникальный id пользователя по которому получаем отчет:

![Screenshot from 2023-08-30 20-52-18](https://github.com/LittleMikle/avito_tech_2023/assets/101155101/307332e9-4dba-4d58-b273-b042d7df96b3)

Создается файл в папке проекта с названием resultUser и уникальный id пользователя по по которому получаем отчет:

![Screenshot from 2023-08-30 20-53-56](https://github.com/LittleMikle/avito_tech_2023/assets/101155101/33d262d9-c4c3-4506-9286-3117d9a7255f)

Внутри файла результат селекта по id пользователя:  
![Screenshot from 2023-08-30 20-54-28](https://github.com/LittleMikle/avito_tech_2023/assets/101155101/5b8a2ba4-aeae-4cfe-94c7-1b921a417d8c)

**2) Создание и отложенное удаление (pg_cron):**

![Screenshot from 2023-08-30 21-01-48](https://github.com/LittleMikle/avito_tech_2023/assets/101155101/52706c57-1bbb-4a42-8993-e2a7a3689272)

После запроса создается cron.job который удалит запись через n дней:

![Screenshot from 2023-08-30 21-02-46](https://github.com/LittleMikle/avito_tech_2023/assets/101155101/17b68d4c-af73-4e3a-a003-6e31447b21ee)

**3) Рандомное добавление пользователей в сегмент:**

Возник вопрос как рассчитывать % для добавления, поскольку неизвестно откуда брать количество юзеров. Для реализации создал таблицу users. В реальном проекте количество юзеров может приходить извне. 

![Screenshot from 2023-08-30 21-41-07](https://github.com/LittleMikle/avito_tech_2023/assets/101155101/1ec25f22-7be1-464c-9ffd-6175df0a1fae)

Для проверки очистим таблицу с сегментами и историю пользователей.

![Screenshot from 2023-08-30 21-42-43](https://github.com/LittleMikle/avito_tech_2023/assets/101155101/2bffb003-88ac-4fd4-b6e8-df64a9fa19c9)

Получаем записи в таблицу с сегментами:

![Screenshot from 2023-08-30 21-43-18](https://github.com/LittleMikle/avito_tech_2023/assets/101155101/478e0214-8223-4767-a7a7-18e6247692d4)

Также появились записи в истории. Добавились в том же порядке. Можно сравнить по user_id: 
![Screenshot from 2023-08-30 21-43-24](https://github.com/LittleMikle/avito_tech_2023/assets/101155101/c51e464b-67c7-401c-bd19-6a33f5cfe3c7)




