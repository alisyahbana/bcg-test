# catalog-management

## deploy

1. create file log at "/var/log/bcg-test/bcg-test.log"
  ```bash
sudo mkdir /var/log/bcg-test && sudo chmod -R 777 /var/log/bcg-test
```
2. run file with command
  ```bash
go run cmd/app/bcg-test/main.go
```
3. Run mysql migration from sql/init.sql
4. visit url http://localhost:8081