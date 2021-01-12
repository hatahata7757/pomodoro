# 構成
## API
- 使用言語：Go
- フレームワーク：gin
- ORM：gorm
- アーキテクチャ：Clean Architecture
## Reverse proxy
- Nginx
## Database
- mysql

# 起動
apiとmysql, nginxを起動する
```
docker-compose up -d
```
## API
### ルーティング
```
// 一覧
curl -i -H 'Content-Type:application/json' localhost:1234/api/v1/users
```

```
// 詳細
curl -i -H 'Content-Type:application/json' localhost:1234/api/v1/users/:id
```

```
// 登録
curl -i -H "Accept: application/json" -H "Content-type: application/json" -X POST -d '{"Name": "XX"}' localhost:1234/api/v1/users
```

```
// 削除
curl -i -H "Accept: application/json" -H "Content-type: application/json" -X DELETE localhost:1234/api/v1/users/:id
```
