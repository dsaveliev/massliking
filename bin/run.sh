#!/bin/bash

TARGET=targets/develop

cd $TARGET && APP_ENV=development GOMAXPROCS=8 ./massliking
