#!/bin/sh
app="graphql-server"
printf "\nStart running: $app\n"
# Set all ENV vars for the server to run
# export $(cat .env | xargs)
time /$GOPATH/bin/air
# This should unset all the ENV vars, just in case.
# unset $(cat .env | sed -E 's/(.*)=.*/\1/''' | xargs)
printf "\nStopped running: $app\n\n"
