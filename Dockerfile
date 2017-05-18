# Start from a base image
FROM golang

# Copy local package files to the container
ADD . /go/src/github.com/richard8thday/xls-to-json
ADD ./input.xlsx /go/bin/xls-to-json

# Build the project inside the container
RUN go get github.com/tealeg/xlsx
RUN go install github.com/richard8thday/xls-to-json

# Run the command when the container starts
ENTRYPOINT /go/bin/xls-to-json