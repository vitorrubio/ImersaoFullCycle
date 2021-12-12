FROM node:14.15.4-alpine3.12

RUN cp -r . /home/node/app

RUN npm install -g @nestjs/cli@8.0.0

USER node

WORKDIR /home/node/app

CMD npm start:dev