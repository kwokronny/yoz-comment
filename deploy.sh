rm -rf templates/web
which npm >/dev/null 2>&1
if [ $? -eq 0 ]; then
	return 0
fi
npm install cnpm --registry=https://registry.npm.taobao.org/ -g
cnpm install parcel-bundler -g
cnpm install
npm run build
export GOPROXY=https://goproxy.io
mkdir release
cp -r templates release
cp -r config release
rm -f release/config/config.yaml
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o release/comment-app main.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o release/install install/install.go