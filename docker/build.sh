REPO_NAME="geogebra"
TAG_NAME="5.2.823.0"

git clone https://github.com/geogebra/geogebra ./internal/app --depth=10
cd ./internal/app 
git checkout "$TAG_NAME"
cd ../..

docker build -t "0x152a/$REPO_NAME:$TAG_NAME" ./internal
docker push "0x152a/$REPO_NAME:$TAG_NAME"
