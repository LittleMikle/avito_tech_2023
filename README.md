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
![Screenshot from 2023-08-30 22-03-26](https://github.com/LittleMikle/avito_tech_2023/assets/101155101/b2362a59-9082-44c7-a2a9-994d7b2b2140)


В соответствии с пунктом 4, где указано что механизм миграций не нужен, SQL файл лежит по пути:

/schema/000001_init.up.sql

### Swagger:

swag init -g cmd/main.go

### Тесты:

go test -v ./...

## Postman коллекция лежит в корневой папке

Создание сегмента:

![Screenshot from 2023-08-30 20-36-21](https://github.com/LittleMikle/avito_tech_2023/assets/101155101/4f41c79b-8cff-4a3e-ad3a-1410f5939de6)

Удаление сегмента:

![Screenshot from 2023-08-30 20-37-06](https://github.com/LittleMikle/avito_tech_2023/assets/101155101/3e10ae95-8824-40a7-bd7f-ce86991043cb)


Добавление сегментов пользователю:

![Screenshot from 2023-08-30 20-39-48](https://github.com/LittleMikle/avito_tech_2023/assets/101155101/0482363c-2f47-4f4a-8cc9-4dad83fdf58a)

Пользователь добавляется, при повторном запросе записи не перетираютя (есть проверка):

![Screenshot from 2023-08-30 20-42-02](https://github.com/LittleMikle/avito_tech_2023/assets/101155101/5417c108-ae5e-4133-9d63-b0e60e0dd30d)



Также появились записи о добавлении в таблицу с историей:

![Screenshot from 2023-08-30 20-41-54](https://github.com/LittleMikle/avito_tech_2023/assets/101155101/b7fa3b44-b1b0-4c33-984e-361895734426)



Удаление нескольких сегментов у пользователя:

![Screenshot from 2023-08-30 20-44-57](https://github.com/LittleMikle/avito_tech_2023/assets/101155101/8bdfb77c-ef87-4ce5-b6ae-f535d48e4306)


Добавились записи об удалении в таблицу с историей:

![Screenshot from 2023-08-30 20-45-50](https://github.com/LittleMikle/avito_tech_2023/assets/101155101/793ebbe1-4a36-43b6-a456-b59d327a336a)


Получение всех активных сегментов пользователя:
![Screenshot from 2023-08-30 20-47-56](https://github.com/LittleMikle/avito_tech_2023/assets/101155101/df446dcf-cab3-4452-be13-19b108b972d4)


## Дополнительные задания:

**1) Экспорт истории пользователя в .csv**

Создается файл папке проекта с названием resultUser и уникальный id пользователя по которому получаем отчет:

![Screenshot from 2023-08-30 20-52-18](https://github.com/LittleMikle/avito_tech_2023/assets/101155101/8a259ec3-f471-4831-87f4-9a5e8995d73f)

Создается файл в папке проекта с названием resultUser и уникальный id пользователя по по которому получаем отчет:

![Screenshot from 2023-08-30 20-53-56](https://github.com/LittleMikle/avito_tech_2023/assets/101155101/6769c522-58c0-41a9-b2b5-292addf7dfe4)

Внутри файла результат селекта по id пользователя:  

![Screenshot from 2023-08-30 20-54-28](https://github.com/LittleMikle/avito_tech_2023/assets/101155101/76cbbd1b-0af4-4136-85ac-40b623a229ec)


**2) Создание и отложенное удаление (pg_cron):**

![Screenshot from 2023-08-30 21-01-48](https://github.com/LittleMikle/avito_tech_2023/assets/101155101/3940c8d0-a3dc-4a33-8b64-63203186072a)


После запроса создается cron.job который удалит запись через n дней:

![Screenshot from 2023-08-30 21-02-46](https://github.com/LittleMikle/avito_tech_2023/assets/101155101/18fcad03-dfe2-45b8-bd26-38f229ca4b6a)


**3) Рандомное добавление пользователей в сегмент:**

Возник вопрос как рассчитывать % для добавления, поскольку неизвестно откуда брать количество юзеров. Для реализации создал таблицу users. В реальном проекте количество юзеров может приходить извне. 

![Screenshot from 2023-08-30 21-41-07](https://github.com/LittleMikle/avito_tech_2023/assets/101155101/2b5aa574-a2bb-42e8-9aec-9a9be4bca4a0)


Для проверки очистим таблицу с сегментами и историю пользователей.

![Screenshot from 2023-08-30 21-42-43](https://github.com/LittleMikle/avito_tech_2023/assets/101155101/7865b0ce-3e38-49ee-8c08-f93715a135e6)


Получаем записи в таблицу с сегментами:

![Screenshot from 2023-08-30 21-43-18](https://github.com/LittleMikle/avito_tech_2023/assets/101155101/300ae8a7-9a8f-4552-b95c-347b08ade8d0)


Также появились записи в истории. Добавились в том же порядке. Можно сравнить по user_id: 
![Screenshot from 2023-08-30 21-43-24](https://github.com/LittleMikle/avito_tech_2023/assets/101155101/e79f9557-2d86-4b63-9a55-b70fe908593d)




