#!/bin/sh
buildPath="build"
app="./"
program="$buildPath/$app"
printf "\nStart app: $app\n"
# Set all ENV vars for the program to run
#export $(cat .env | xargs)
time make automigrate
time make migrate
time ./$program
# This should unset all the ENV vars, just in case.
#unset $(cat .env | sed -E 's/(.*)=.*/\1/' | xargs)
printf "\nStopped app: $app\n\n"
