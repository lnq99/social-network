FROM node:alpine as build-stage

WORKDIR /app

COPY package.json ./

RUN npm install --verbose

COPY . .

RUN npm run build


FROM nginx:alpine as production-stage

RUN mkdir /app

COPY --from=build-stage /app/dist /app

COPY nginx.conf /etc/nginx/nginx.conf
