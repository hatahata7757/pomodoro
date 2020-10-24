FROM node:12.16.1-alpine

RUN apk update && \
  apk upgrade && \
  yarn global add @vue/cli

RUN yarn install

EXPOSE 3000
CMD ["yarn", "run", "dev"]
