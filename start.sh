#!/bin/sh

./appritstoreback &
BACK_PID=$!

cd ./web 
npm start &
FRONT_PID=$!

while sleep 60; do
  ps -fp $BACK_PID 
  BACK_PROCESS_STATUS=$?
  if [ $BACK_PROCESS_STATUS -ne 0 ]; then
    echo "backend process has already exited."
    exit 1
  fi
  
  ps -fp $FRONT_PID 
  FRONT_PROCESS_STATUS=$?
  if [ $FRONT_PROCESS_STATUS -ne 0 ]; then
    echo "fontend process has already exited."
    exit 1
  fi
done
