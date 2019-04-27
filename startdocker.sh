docker container stop dtdataserv
docker container rm dtdataserv
docker run -d --name dtdataserv \
    -v $PWD/cfg:/app/dtdataserv/cfg \
    -v $PWD/logs:/app/dtdataserv/logs \
    -v $PWD/dat:/app/dtdataserv/dat \
    -p 7053:7053 \
    -p 7788:7788 \
    dtdataserv