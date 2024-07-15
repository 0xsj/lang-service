# pull from base image
FROM kakao:nx-base AS builder

# Extend from base Dockerfile
ARG APP_NAME=gateway
ARG NODE_ENV
WORKDIR /app/builder
COPY . .
# COPY --from=deps /app/node_modules ./node_modules
RUN npx nx build ${APP_NAME}

FROM kakao:nx-base

WORKDIR /app
# COPY --from=builder /app/builder/dist/apps/${APP_NAME} ./dist/apps/${APP_NAME}/
COPY --from=builder /app/builder ./
ENV NODE_ENV=$NODE_ENV

CMD ["node", "./dist/apps/gateway/main.js"]
