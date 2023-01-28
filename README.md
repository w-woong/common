# Common

- error
- http
- logger

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
DOMAIN_NAME=woong
cd $WOONG_HOME/$DOMAIN_NAME/cmd/http
nohup go run $WOONG_HOME/$DOMAIN_NAME/cmd/http/main.go --autoMigrate >> $WOONG_HOME/$DOMAIN_NAME/cmd/http/logs/agent.log 2>&1 &

DOMAIN_NAME=auth
cd ${WOONG_HOME}/${DOMAIN_NAME}/cmd
nohup go run ${WOONG_HOME}/${DOMAIN_NAME}/cmd/main.go --autoMigrate >> ${WOONG_HOME}/${DOMAIN_NAME}/cmd/logs/agent.log 2>&1 &

DOMAIN_NAME=payment
cd $WOONG_HOME/$DOMAIN_NAME/cmd
nohup go run $WOONG_HOME/$DOMAIN_NAME/cmd/main.go --autoMigrate >> $WOONG_HOME/$DOMAIN_NAME/cmd/logs/agent.log 2>&1 &

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