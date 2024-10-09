FROM golang:1.22-alpine3.20

ENV NAME=gomicro-pinger

# Set working directory
WORKDIR /go/src/${NAME}

# Add build dependencies
RUN set -eux; apk update; apk add --no-cache make ca-certificates tzdata

# Copy source code
COPY . .

# Install dependencies
RUN go mod download -x

# Build the application
RUN make build

RUN cp ./build/${NAME} /service

ENTRYPOINT ["/service"]
