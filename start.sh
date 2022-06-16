#!/bin/sh

./appritstoreback &

cd ./web 
npm start &

wait

exit $?
