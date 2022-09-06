# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy the source code.
COPY . ./

# Build
RUN go build -o /swe-influencer-api

EXPOSE 3200

# Run
CMD [ "/swe-influencer-api" ]