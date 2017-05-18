# Start from a base image
FROM golang

# Copy local package files to the container
ADD . /go/src/github.com/richard8thday/xlsx2json
ADD ./input.xlsx /go/bin/xlsx2json

# Build the project inside the container
RUN go get github.com/tealeg/xlsx
RUN go install github.com/richard8thday/xlsx2json

# Run the command when the container starts
ENTRYPOINT /go/bin/xlsx2json