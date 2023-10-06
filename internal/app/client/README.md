# keeper client CLI documentation

### Invoke findUser

> go run cmd/client/main.go registerUser --name=test --email=test@test.ru --password=test

> go run cmd/client/main.go getUser

> go run cmd/client/main.go loginUser --email=test@test.ru --password=test


> go run cmd/client/main.go createAccount --title=test --login=test@test.ru --password=test --comment=test --url=test.ru

> go run cmd/client/main.go updateAccount --id=uuid --title=test --login=test@test.ru --password=test --comment=test --url=test.ru

> go run cmd/client/main.go deleteAccount --id=uuid

> go run cmd/client/main.go listAccount --limit=<10> --offset=<10>

> go run cmd/client/main.go createText --title=test --data='текст который нужно сохранить' --comment=test

> go run cmd/client/main.go updateText --id=uuid --title=test --title=test --data=текст который нужно сохранить --comment=test

> go run cmd/client/main.go deleteText --id=uuid

> go run cmd/client/main.go listText --limit=<10> --offset=<10>

> go run cmd/client/main.go createBinary --title=binary --data='текст который нужно сохранить' --comment=test

> go run cmd/client/main.go updateBinary --id=uuid --title=binary --title=binary --data=текст который нужно сохранить --comment=binary

> go run cmd/client/main.go deleteBinary --id=uuid

> go run cmd/client/main.go listBinary --limit=<10> --offset=<10>

> go run cmd/client/main.go createCard --name=test --number='2204-1201-0110-1398' --date=08/29 --cvc=988 --comment=test

> go run cmd/client/main.go updateCard --id=uuid --name=test --number='2204-1201-0110-1398' --date=08/29 --cvc=988 --comment=test

> go run cmd/client/main.go deleteCard --id=uuid

> go run cmd/client/main.go listCard --limit=<10> --offset=<10>