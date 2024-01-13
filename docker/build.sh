REPO_NAME="tldraw"
TAG_NAME="v2.0.0-beta.2"

git clone https://github.com/tldraw/tldraw.git ./internal/app
cd ./internal/app 
git checkout v2.0.0-beta.2
cd ../..

docker build -t "0x152a/$REPO_NAME:$TAG_NAME" ./internal
docker push "0x152a/$REPO_NAME:$TAG_NAME"
