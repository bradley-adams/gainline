# Stage 1: Build Angular app
FROM node:18 AS builder

WORKDIR /app

COPY package*.json ./
RUN npm install

COPY . .

ARG API_URL=http://localhost:8080
RUN sed -i "s|http://localhost:8080|${API_URL}|g" src/environments/environment.prod.ts

RUN npm run build

# Stage 2: Serve with Nginx
FROM nginx:alpine

WORKDIR /app

COPY --from=builder /app/dist/gainline-ui/browser /usr/share/nginx/html/
COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]