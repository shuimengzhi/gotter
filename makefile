.PHONY: all build clean run check cover lint docker help

gotter ="   _____    ____    _______   _______   ______   _____\n"
gotter +="  / ____|  / __ \  |__   __| |__   __| |  ____| |  __ \ \n"
gotter +=" | |  __  | |  | |    | |       | |    | |__    | |__) |\n"
gotter +=" | | |_ | | |  | |    | |       | |    |  __|   |  _  /\n"
gotter +=" | |__| | | |__| |    | |       | |    | |____  | | \ \\n"
gotter +="  \_____|  \____/     |_|       |_|    |______| |_|  \_\\n"

#记录的时间
time = $(shell date "+%Y-%m-%d %H:%M:%S")
# 应用生成的位置
OUT_FILE_PATH=./bin/app

run:build logo
	${OUT_FILE_PATH}_mac
build:
	@echo "==============编译开始:"$(time)
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${OUT_FILE_PATH}_linux main.go
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ${OUT_FILE_PATH}_win main.go
	@CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ${OUT_FILE_PATH}_mac main.go
	@echo "==============编译结束:"$(time)

clean:
	rm -f ${OUT_FILE_PATH}_linux
	rm -f ${OUT_FILE_PATH}_win
	rm -f ${OUT_FILE_PATH}_mac
logo:
	@echo  $(gotter)


