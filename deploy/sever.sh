#!/bin/sh
configFile="./config/config.yaml"
pkill -9 install
pkill -9 comment-app

if [ ! -f "$configFile" ]; then
	echo "run install"
	chmod 774 ./install
	nohup ./install >> log.out & sleep 1
else
	echo "run main"
	chmod 774 ./comment-app
	nohup ./comment-app >> log.out & sleep 1
fi