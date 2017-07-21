#!/bin/bash

TARGET=targets/develop

cd $TARGET && APP_ENV=development GOMAXPROCS=2 ./massliking
