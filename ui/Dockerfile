# Stage 1: Build Angular app

FROM node:18 AS builder

WORKDIR /app

COPY package*.json ./
RUN npm install

COPY . .

ARG API_URL=http://localhost:8080
ARG AUTH0_DOMAIN=dev-dq1p4t4guh2d2oj3.au.auth0.com
ARG AUTH0_CLIENT_ID=CkHz9nnnyx4ZMrag1ZCBejluZNiMeniz
ARG AUTH0_AUDIENCE=https://api.dev.gainline.io

RUN sed -i "s|http://localhost:8080|${API_URL}|g" src/environments/environment.prod.ts && \
    sed -i "s|AUTH0_DOMAIN_PLACEHOLDER|${AUTH0_DOMAIN}|g" src/environments/environment.prod.ts && \
    sed -i "s|AUTH0_CLIENT_ID_PLACEHOLDER|${AUTH0_CLIENT_ID}|g" src/environments/environment.prod.ts && \
    sed -i "s|AUTH0_AUDIENCE_PLACEHOLDER|${AUTH0_AUDIENCE}|g" src/environments/environment.prod.ts

RUN npm run build

# Stage 2: Serve with Nginx

FROM nginx:alpine

WORKDIR /app

COPY --from=builder /app/dist/gainline-ui/browser /usr/share/nginx/html/
COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]