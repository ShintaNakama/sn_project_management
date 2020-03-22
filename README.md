# sn_project_management
- プロジェクト管理アプリ

- develop-env Start and Create DataBase
- Hot Reload ... realize
  - ./.realize.yaml
```
$ docker-compose up --build
```
- vscodeでのdebugの場合
  - docker経由だとvscodeでのdebugが難しいので、mysqlだけコンテナ起動し、realize startをローカルで実行する。
  - db接続のためのDNSはなども変更も必要。
  - vscodeで、main.goの処理終了のあたりでブレークポイントを設置し、debugしたい箇所にもブレークポイントを置くことでdebugができる。

- DataBase Migration ... sql-migrate
  - ./dbconfig.yml
```
$ sql-migrate status
$ sql-migrate new_create_`table_name` ... Create Migration File
$ sql-migrate up Create Tables
$ sql-migrate down RollBack
```
