#!/bin/sh
configFile="./config/config.yaml"
echo "kill process"
pkill -9 install
pkill -9 comment-app
echo "kill end"
if [ ! -f $configFile ]; then
	echo "start install sever"
	nohup ./install > /dev/null 2>&1 &
	sleep 1
else
	echo "start comment-app sever"
	nohup ./comment-app > /dev/null 2>&1 &
	sleep 1
fi