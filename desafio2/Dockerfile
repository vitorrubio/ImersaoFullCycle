FROM node:14.15.4-alpine3.12

RUN npm install -g @nestjs/cli@8.0.0

WORKDIR /home/node/app

COPY . .

RUN mkdir -p /home/node/app/dist

RUN chmod -R 777 /home/node/app/dist

RUN npm install

USER node

CMD npm run start:dev