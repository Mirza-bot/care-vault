# Use Node.js as the base image
FROM node:18

# Set the working directory inside the container
WORKDIR /app

# Copy package.json and package-lock.json before installing dependencies
COPY package.json package-lock.json ./

# Install dependencies
RUN npm install

# Copy the entire Angular app into the container
COPY . .

# Expose the Angular development server port
EXPOSE 4200

# Start the Angular app in development mode
CMD ["npm", "run", "start"]
