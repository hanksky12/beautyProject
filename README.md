## 後端(生產者)
api server (produce pc status)
## 消費者
consumer (record pc status)
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

## 6. 註冊使用者
//todo 


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

主要套件
logrus: log
gorm: orm
gin: web
segmentio/kafka-go: kafka
golang-migrate/migrate: sql migrate



