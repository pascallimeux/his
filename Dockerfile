FROM debian:latest

RUN apt-get update
RUN apt-get install -y libltdl-dev nano && rm -rf /var/lib/apt/lists/*

RUN mkdir -p /var/his/keys
RUN mkdir -p /var/log/his
COPY ./build/bin/his /var/his/his
COPY his_prod.toml /var/his/his.toml
COPY ./keys/server.key /var/his/keys/server.key
COPY ./keys/server.crt /var/his/keys/server.crt
ADD ./fixtures /var/his/fixtures

# Set binary as entrypoint
ENTRYPOINT cd /var/his && ./his

# Expose port (8000)
EXPOSE 8000