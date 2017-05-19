FROM debian:latest

RUN apt-get update
RUN apt-get install -y libltdl-dev nano && rm -rf /var/lib/apt/lists/*

RUN mkdir -p /var/his
RUN mkdir -p /var/log/his
COPY his /var/his/his
COPY his_prod.toml /var/his/his.toml
COPY server.key /var/his/server.key
COPY server.crt /var/his/server.crt
ADD ./fixtures /var/his/fixtures

# Set binary as entrypoint
ENTRYPOINT cd /var/his && ./his

# Expose port (8000)
EXPOSE 8000