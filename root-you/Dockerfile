FROM tomyj/my-env:node
ENV API_HOSTNAME=api:3232
WORKDIR /app
COPY package-lock.json .
COPY package.json .
COPY . .
RUN npm ci
CMD [ "npx", "next", "dev", "-p", "80" ]
