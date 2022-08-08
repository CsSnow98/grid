sudo docker rm -f $(sudo docker ps -aq)
sudo docker network prune
sudo docker volume prune
cd fixtures && ./start.sh
docker-compose up -d
echo "正在等待节点的启动完成，等待10秒"
sleep 10
cd ..
rm grid
go build
./grid
cd application && ./start.sh