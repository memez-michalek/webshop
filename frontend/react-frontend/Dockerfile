FROM node:15.5.1

WORKDIR /app

COPY package.json /app/

RUN yarn install

COPY . .

EXPOSE 3000

CMD ["npm", "start"]