FROM debian:latest

RUN apt-get update
RUN apt-get install -y

RUN apt-get install -y nginx

RUN mkdir -p /var/his-ui
ADD ./build /var/his-ui/

RUN rm -v /etc/nginx/nginx.conf
ADD ./nginx/nginx.conf /etc/nginx/

RUN echo "daemon off;" >> /etc/nginx/nginx.conf

EXPOSE 8080

CMD service nginx start
