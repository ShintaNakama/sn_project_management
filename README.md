# sn_project_management
- プロジェクト管理アプリ

- develop-env Start and Create DataBase
- Hot Reload ... realize
  - ./.realize.yaml
```
$ docker-compose up --build
```

- DataBase Migration ... gore
  - ./dbconfig.yml
```
$ sql-migrate status
$ sql-migrate new_create_`table_name` ... Create Migration File
$ sql-migrate up Create Tables
$ sql-migrate down RollBack
```
