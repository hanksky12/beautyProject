## 前端
mouse action (web socket)
## gin後端(生產者)
api server (produce server status,mouse action)(query server status,mouse action)
## kafka消費者
consumer (record server status,mouse action)
## 定期
cron (calculate pc status)


# 步驟:

## 1. run docker
docker-compose -f compose-elk.yaml up

docker-compose -f compose-kafka.yaml up

docker-compose -f compose-db.yaml up

## 2. kafka-ui(http://localhost:8080/) 設定Kafka Cluster
host.docker.internal 29092

## 3. kibana(http://localhost:5601/) 設定DataView 
Management -> Index Patterns -> Create Index Pattern -> logstash-* -> @timestamp

## 4. sql migrate
go run ./cmd/migrate/cmd.go -c Run -p up -p 0

## 5. run service
go run ./cmd/backend/main.go (http://localhost:8070/)

go run ./cmd/consumer/main.go

go run ./cmd/cron/main.go

## 6. run 前端 
cd frontend
npm run serve

## 7.註冊使用者
http://localhost:8081/personal-register

## 8.登入
http://localhost:8081/login


Log 查詢keyword
log.file.path: 從檔案下手
request_id: 從請求request_id下手
go_id: 從goroutine下手


---
### Technical-Framework
#### vue.js(front) , bootstrap(front) , gin(back)

---
# Frontend(SPA)

view: composed of multiple components

components: base components not related to logic

store: vuex for global variable

mixins: js for reuse &  api

---
# Backend

## 分層架構

cmd/app/controller: route

internal/pkg/model: table schema

internal/services : logic

---

# 套件

| 功能       | 套件名稱                   |
|----------|------------------------|
| log      | logrus                 |
| orm      | gorm                   |
| backend  | gin                    |
| frontend | vue.js                 |
| kafka    | segmentio/kafka-go     |
| migrate  | golang-migrate/migrate |
| test     | stretchr/testify       |
| mock     | vektra/mockery         |




