# Build Vite application
FROM node:22.7.0 AS build

WORKDIR /build

COPY . .

RUN npm install
RUN npm run build

# Run server
FROM golang:1.23.0

WORKDIR /app

ARG DB_CONNECTION_STRING
ENV DB_CONNECTION_STRING=${DB_CONNECTION_STRING}

COPY --from=build /build/dist ./dist/

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
COPY src/server ./src/server/
EXPOSE 8000

CMD ["go", "run", "."]
