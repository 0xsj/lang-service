ARG NODE_VERSION=20.15.0

FROM node:${NODE_VERSION}-alpine AS deps

ARG NODE_ENV
WORKDIR /app

COPY package*.json ./

RUN npm ci