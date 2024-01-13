REPO_NAME="suroi"
TAG_NAME="1"

git clone https://github.com/HasangerGames/suroi.git ./internal/app

docker build -t "0x152a/$REPO_NAME:$TAG_NAME" ./internal
docker push "0x152a/$REPO_NAME:$TAG_NAME"
