#!/bin/bash
for((i=1;i<=1000000;i++))
do
curl -i -H 'Content-Type: application/json' \
    -d '{"Value":"test message"}' http://127.0.0.1:1991/send_msg/32002
done

