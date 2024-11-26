FROM golang:1.22

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy the source code.
#COPY *.go ./
COPY . .

# Build
RUN GCO_ENABLED=0 GOOS=linux go build -o /CustomerMS
#RUN go build -o /CustomerMS

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/reference/dockerfile/#expose
EXPOSE 8382

# Run
CMD [ "/CustomerMS" ]