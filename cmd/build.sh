GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o cloud-companion-agent .
docker build -t yohansutanto5/cloud-companion-agent .
docker run -d --name cloudagent -p 8078:8078 yohansutanto5/cloud-companion-agent