npm i --save-dev axios


make build-image

docker build -t his-ui .
docker images |grep his-ui
docker run -d -p 8080:8080 --name his-ui his-ui
docker ps
docker rm -f his-ui
docker save his-ui -o his-ui.tar
scp his-ui.tar pascal@192.168.0.100:/tmp



sudo docker load -i his.tar
sudo docker load -i his-ui.tar

make start-hp
docker run -d -p 8000:8000 -p 3000:3000 --name hisv1 his
docker run -d -p 8080:8080 --name his-ui his-ui

check ui with: 'ip':8080/
