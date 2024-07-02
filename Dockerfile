# Latest version
FROM golang:alpine
# Create a folder and copy all the contents of the project into it
ADD . /my_template
# Define the working directory
WORKDIR /my_template

# Metadata
LABEL version="1.0"
LABEL description="my_template"

# Delete caches after installing
RUN apk --no-cache add bash

# Default port
EXPOSE 1111
# The instruction to be executed when the container is created
CMD go run ./cmd/web/.

# sudo docker image build -f Dockerfile -t my_template .
# sudo docker run --detach -p 1111:1111 my_template

