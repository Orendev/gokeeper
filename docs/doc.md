# Домены
1. Пользователь
   1. имя
   2. логин
   3. пароль
   4. email
   5. дата создания 
   6. дата изменения 
   7. дата удаления
2. Аккаунты
   1. заголовок
   2. логин 
   3. пароль 
   4. адрес сайта 
   5. комментарий 
   6. дата удаления 
   7. дата изменения 
   8. дата создания 
   9. пользователь создавший 
   10. версия
3. Текстовые данные 
   1. заголовок
   2. комментарий
   3. основные данные
   4. версия 
   5. дата удаления
   6. дата изменения
   7. дата создания 
   8. пользователь создавший
4. Бинарные данные
   1. заголовок
   2. комментарий
   3. основные данные
   4. версия
   5. дата удаления
   6. дата изменения
   7. дата создания
   6. пользователь создавший
5. Данные банковских карт
   1. заголовок
   2. номер карты
   3. владелец карты
   4. срок действия месяц/год
   5. CVC2 / CVV2
   6. комментарий 
   7. дата удаления
   8. дата создания
   9. дата обновления 
   10. пользователь создавший
   11. версия

# Usecase
### keeper client
1. UC-1 получает информацию о версии и дате сборке бинарного файла клиента
2. UC-2 проходит процедуру первичную регистрацию
3. UC-3 проходит процедуру аутентификации
4. UC-4 синхронизует данные с сервером
5. UC-5 отображает данные для пользователя
6. UC-6 запрашивает пары логин пароль 
7. UC-7 запрашивает произвольные текстовые данные 
8. UC-8 запрашивает произвольные бинарные данные 
9. UC-9 запрашивает банковские карты 
10. UC-10 добавляет пары логин пароль 
11. UC-11 добавляет произвольные текстовые данные 
12. UC-12 добавляет произвольные бинарные данные 
13. UC-13 добавляет банковские карты 
14. UC-14 редактирует пары логин пароль 
15. UC-15 редактирует произвольные текстовые данные 
16. UC-16 редактирует произвольные бинарные данные 
17. UC-17 редактирует банковские карты 
18. UC-18 удаляет пары логин пароль 
19. UC-19 удаляет произвольные текстовые данные 
20. UC-20 удаляет произвольные бинарные данные 
21. UC-21 удаляет банковские карты

### keeper server
1. UC-1 регистрация пользователя
2. UC-2 аутентификации
3. UC-3 синхронизация данных с клиентом
4. UC-4 добавляет  пары логин пароль 
5. UC-5 добавляет произвольные текстовые данные 
6. UC-6 добавляет произвольные бинарные данные 
7. UC-7 добавляет банковские карты 
8. UC-8 редактирование пары логин пароль. 
9. UC-9 редактирование произвольные текстовые данные. 
10. UC-10 редактирование произвольные бинарные данные. 
11. UC-11 редактирование банковские карты. 
12. UC-12 чтение пары логин пароль. 
13. UC-13 чтение произвольные текстовые данные. 
14. UC-14 чтение произвольные бинарные данные. 
15. UC-15 чтение банковские карты. 
16. UC-16 удаление пары логин пароль. 
17. UC-17 удаление произвольные текстовые данные. 
18. UC-18 удаление произвольные бинарные данные. 
19. UC-19 удаление банковские карты.