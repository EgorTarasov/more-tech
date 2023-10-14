FROM node:18-alpine

WORKDIR /app

COPY . /app

RUN npm i

RUN npm run build

RUN npm add serve -g

EXPOSE 3000

CMD ["npm", "run", "serve"]