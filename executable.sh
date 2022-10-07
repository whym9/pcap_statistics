sudo docker build -t statistics . 
sudo docker run --name=stat --env-file .env -p 6006:6006 -d statistics
