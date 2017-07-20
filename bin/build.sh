#!/bin/bash
export APP_ENV=development
export NODE_ENV=development
export GOARCH=amd64
export GOOS=linux

TARGET=targets/develop

rm -rf ./targets;

mkdir -p ./targets;
mkdir -p ./$TARGET;
mkdir -p ./$TARGET/static;
mkdir -p ./$TARGET/config;
mkdir -p ./$TARGET/log;
mkdir -p ./$TARGET/pids;

cp -r ./config/dev.env.js ./frontend/config/;
cp -r ./config/development.yml ./$TARGET/config/;
cd ./frontend && npm run build && cd ../;
cd ./backend && go build -o ../$TARGET/massliking *go && cd ../;
