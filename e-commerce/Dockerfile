### STAGE 1: Build ###
FROM node:12.18.2-alpine AS buildxyz
WORKDIR /usr/src/app
COPY package.json package-lock.json ./
RUN npm install
COPY . .
RUN npm run build

### STAGE 2: Run ###
FROM nginx:1.19-alpine
COPY --from=buildxyz /usr/src/app/dist/e-commerce /usr/share/nginx/html