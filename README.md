# fileDownloadByGo
Generate a web download link for files in a directory in Linux

docker build -t web_filedownloader . #生成新的镜像 
docker run -d -p 3002:8080 -v ~/downloads:/root/downloads web_filedownloader  #完成部署
