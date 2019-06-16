g:
	git add .
	git commit -m"自动提交 git 代码"
	git push

build:
	protoc -I . --go_out=plugins=micro:. proto/aliyun/aliyun.proto
