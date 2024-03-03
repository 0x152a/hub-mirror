REPO_NAME="curlconverter.github.io"
TAG_NAME="latest"

git clone https://github.com/curlconverter/curlconverter.github.io.git ./internal/app
cd ./internal/app 
# git checkout "$TAG_NAME"
cd ../..

docker build -t "0x152a/$REPO_NAME:$TAG_NAME" ./internal
docker push "0x152a/$REPO_NAME:$TAG_NAME"
