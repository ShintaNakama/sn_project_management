# sn_project_management
- プロジェクト管理
- 個人の勉強用
- Go 1.13
- echo
- gorp
- terraform
- GitHubActions
- CRUD

### develop-env Start and Create DataBase
- Hot Reload ... realize
- docker
```
docker-compose up --build
```

### DataBase Migration ... sql-migrate
  - ./dbconfig.yml
```
$ sql-migrate status
$ sql-migrate new_create_`table_name` ... Create Migration File
$ sql-migrate up Create Tables
$ sql-migrate down RollBack
```

### docker コンテナ内でのdelve debug
```
# コンテナ内に入り、attachするプロセスを確認してattach
ps aux
dlv attach PID
# あとはdlvの使い方通りにdebug
## ブレークポイント設置
b パッケージ名.関数名
## 確認したい挙動を行い、そこで止まっているはずなので
ls # delveのlsコマンド実行
c  # ブレークポイントまで移動
n  # next
locals # local変数一覧確認
p 変数 # 変数の値確認
q # debug終了 processをkillするか聞かれるのY or n
```
