# Common

- error
- http
- logger

## protoc

```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative *.proto
```

## Woong server
```bash
WOONG_HOME=$HOME/Documents/projects/woong
```

```bash
# ps -ef | grep $WOONG_HOME | grep -v 'grep' | awk '{print $2}'
# kill `ps -ef | grep $WOONG_HOME | grep -v 'grep' | awk '{print $2}'`

# processes run with `go run` command should be killed like this
kill `ps -ef | grep '\-\-autoMigrate' | grep -v 'grep' | awk '{print $2}'`
```

```bash
DOMAIN_NAME=oauth2
cd ${WOONG_HOME}/${DOMAIN_NAME}/cmd/http
nohup go run ${WOONG_HOME}/${DOMAIN_NAME}/cmd/http/main.go --autoMigrate --key='./certs/key.pem' --pem='./certs/cert.pem' --config='./configs/server-woong.yml' >> ${WOONG_HOME}/${DOMAIN_NAME}/cmd/http/logs/agent.log 2>&1 &

DOMAIN_NAME=woong
cd $WOONG_HOME/$DOMAIN_NAME/cmd/http
nohup go run $WOONG_HOME/$DOMAIN_NAME/cmd/http/main.go --autoMigrate >> $WOONG_HOME/$DOMAIN_NAME/cmd/http/logs/agent.log 2>&1 &

DOMAIN_NAME=payment
cd $WOONG_HOME/$DOMAIN_NAME/cmd/http
nohup go run $WOONG_HOME/$DOMAIN_NAME/cmd/http/main.go --autoMigrate >> $WOONG_HOME/$DOMAIN_NAME/cmd/http/logs/agent.log 2>&1 &

DOMAIN_NAME=user
cd $WOONG_HOME/$DOMAIN_NAME/cmd/grpc
nohup go run $WOONG_HOME/$DOMAIN_NAME/cmd/grpc/main.go --autoMigrate >> $WOONG_HOME/$DOMAIN_NAME/cmd/grpc/logs/agent.log 2>&1 &

cd $WOONG_HOME/$DOMAIN_NAME/cmd/http
nohup go run $WOONG_HOME/$DOMAIN_NAME/cmd/http/main.go --autoMigrate >> $WOONG_HOME/$DOMAIN_NAME/cmd/http/logs/agent.log 2>&1 &

DOMAIN_NAME=product
cd $WOONG_HOME/$DOMAIN_NAME/cmd/http
nohup go run $WOONG_HOME/$DOMAIN_NAME/cmd/http/main.go --autoMigrate >> $WOONG_HOME/$DOMAIN_NAME/cmd/http/logs/agent.log 2>&1 &

DOMAIN_NAME=order
cd $WOONG_HOME/$DOMAIN_NAME/cmd/http
nohup go run $WOONG_HOME/$DOMAIN_NAME/cmd/http/main.go --autoMigrate >> $WOONG_HOME/$DOMAIN_NAME/cmd/http/logs/agent.log 2>&1 &

DOMAIN_NAME=partner
cd $WOONG_HOME/$DOMAIN_NAME/cmd/http
nohup go run $WOONG_HOME/$DOMAIN_NAME/cmd/http/main.go --autoMigrate >> $WOONG_HOME/$DOMAIN_NAME/cmd/http/logs/agent.log 2>&1 &

DOMAIN_NAME=resource
cd $WOONG_HOME/$DOMAIN_NAME/cmd/http
nohup go run $WOONG_HOME/$DOMAIN_NAME/cmd/http/main.go --autoMigrate >> $WOONG_HOME/$DOMAIN_NAME/cmd/http/logs/agent.log 2>&1 &


```

## Git
```bash
DOMAIN_NAME=woong
cd $WOONG_HOME/$DOMAIN_NAME
git fetch origin main:main

DOMAIN_NAME=auth
cd ${WOONG_HOME}/${DOMAIN_NAME}
git fetch origin main:main

DOMAIN_NAME=payment
cd $WOONG_HOME/$DOMAIN_NAME
git fetch origin main:main

DOMAIN_NAME=user
cd $WOONG_HOME/$DOMAIN_NAME
git fetch origin main:main

cd $WOONG_HOME/$DOMAIN_NAME
git fetch origin main:main

DOMAIN_NAME=product
cd $WOONG_HOME/$DOMAIN_NAME
git fetch origin main:main

DOMAIN_NAME=order
cd $WOONG_HOME/$DOMAIN_NAME
git fetch origin main:main

DOMAIN_NAME=partner
cd $WOONG_HOME/$DOMAIN_NAME
git fetch origin main:main

DOMAIN_NAME=resource
cd $WOONG_HOME/$DOMAIN_NAME
git fetch origin main:main

```

## Drop tables
### woong DB
```sql
drop table home_group_products;
drop table short_notices;
drop table main_promotions;
drop table homes;
drop table app_configs;
drop table tags;
```

### woong_auth db
```sql
drop table auth_requests;
drop table auth_states;
drop table tokens;
```

### woong_order
```sql
drop table cart_products;
drop table carts;
```

### woong_product
```sql
drop table group_products;
drop table groups;
drop table products;
```

```sql
delete from group_products;
delete from groups;
delete from products;
```

### woong_user
```sql
drop table emails;
drop table passwords;
drop table personals;
drop table users;
```
### woong_partner
```sql

```