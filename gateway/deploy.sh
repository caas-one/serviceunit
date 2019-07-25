#!/bin/zsh

# clean
rm -fr /Users/mac/goWorkSpace/pkg/darwin_amd64/gitlab.yeepay.com/yce/nodeport/*

# build a new binary
GOOS=linux GOARCH=amd64 go build .

# kill the process
ssh root@10.151.30.217 "pkill nodeport"
ssh root@10.151.30.218 "pkill nodeport"
ssh root@10.151.30.223 "pkill nodeport"

# sync file
scp nodeport root@10.151.30.217:~
scp nodeport root@10.151.30.218:~
scp nodeport root@10.151.30.223:~

# start the process
START_QA="nodeport --apiserver=10.151.32.61:6443 --ca=/etc/controller/qa/ca.crt --cert=/etc/controller/qa/client.crt --key=/etc/controller/qa/client.key &"
START_CICD="nodeport --apiserver=10.151.33.87:6443 --ca=/etc/controller/cicd/ca.crt --cert=/etc/controller/cicd/client.crt --key=/etc/controller/cicd/client.key &"
# ssh root@10.151.30.217 ${START_QA} 
# ssh root@10.151.30.218 ${START_QA} 
# ssh root@10.151.30.223 ${START_QA} 

# md5
md5 nodeport
ssh root@10.151.30.217 "md5sum nodeport"
ssh root@10.151.30.218 "md5sum nodeport"
ssh root@10.151.30.223 "md5sum nodeport"





