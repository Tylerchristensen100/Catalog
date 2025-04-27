FROM node:22-alpine AS build-react

WORKDIR /app

COPY admin/package*.json ./

RUN rm -rf node_modules package-lock.json && npm install

COPY admin/ ./
COPY core/ ./core/

RUN npm run build


FROM nginx:alpine AS reverse-proxy

COPY nginx.conf /etc/nginx/conf.d/default.conf

COPY --from=build-react /app/dist /etc/nginx/html 

COPY /core /etc/nginx/html/core


EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]