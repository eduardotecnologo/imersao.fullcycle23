# imersao.fullcycle23
# Setup NodeWSL
# https://cloudbytes.dev/snippets/how-to-install-nodejs-and-npm-on-wsl2
# Setup NestJS
$npx @nestjs/cli new nest-api
$cd nest-api
$npm run start
$docker-compose exec app bash

# Generate Routes
$nest g resource routes

# Mongo
$npm i mongoose @nestjs/mongoose --save
$npm i @nestjs/config --save

# Kafka (Connect)
$npm install kafkajs @nestjs/microservices --save