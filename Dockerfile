FROM golang:1.23-alpine AS builder

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux go build -o seonaut cmd/server/main.go

FROM node:18-alpine3.18 AS front
WORKDIR /home/node
COPY --from=builder /app ./app/
RUN npm install --save-exact esbuild && ./node_modules/esbuild/bin/esbuild ./app/web/css/style.css \
	--bundle \
	--minify \
	--outdir=./app/web/static \
	--public-path=/resources \
	--loader:.woff=file \
	--loader:.woff2=file

FROM alpine:latest AS production
COPY --from=front /home/node/app /app/

ENV WAIT_VERSION 2.9.0
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/$WAIT_VERSION/wait /bin/wait
RUN chmod +x /bin/wait

WORKDIR /app

# Make the start script executable
RUN chmod +x start.sh

# Expose the port
EXPOSE 10000

# Set default config file
ENV CONFIG_FILE=config.prod

# Run the application
CMD ["./start.sh"]