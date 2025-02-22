FROM node:20-alpine AS build

# Set the working directory inside the container
WORKDIR /app

# Copy package.json and package-lock.json before installing dependencies
COPY package.json package-lock.json ./

RUN npm install

# # Generate OpenApi files
RUN openapi-typescript-codegen --input http://localhost:8080/swagger/doc.json --output ./src/app/api

# Copy the entire Angular app into the container
COPY . .
RUN npm run build -- --configuration=production

FROM nginx:alpine

COPY --from=build /app/dist/frontend /usr/share/nginx/html

# COPY nginx.conf /etc/nginx/nginx.conf

EXPOSE 4200

CMD ["nginx", "-g", "daemon off;"]
