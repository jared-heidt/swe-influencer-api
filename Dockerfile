# syntax=docker/dockerfile:1

FROM golang:1.16-alpine as builder

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy the source code.
COPY . ./

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/main .

FROM scratch

# Copy the Pre-built binary file
COPY --from=builder /app/bin/main .

# Run executable
CMD ["./main"]

#OLD CODE syntax=docker/dockerfile:1 FROM golang:1.16-alpine WORKDIR /app COPY go.mod ./ COPY go.sum ./ RUN go mod download COPY . ./ RUN go build -o /swe-influencer-api EXPOSE 3200 CMD [ "/swe-influencer-api" ]