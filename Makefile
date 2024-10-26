# 设置输出二进制文件的目录
BIN_DIR=./bin

# 项目的主目录
PROJECT_DIR=~/projects/blockchain

# 设置二进制文件的名称
BINARY_NAME=blockchain

# 默认目标
all: build

# 构建项目
build:
	mkdir -p $(BIN_DIR)
	cd $(PROJECT_DIR) && go build -o $(BIN_DIR)/$(BINARY_NAME)

# 运行构建的程序
run: build
	$(BIN_DIR)/$(BINARY_NAME)	

# 测试项目
test:
	cd $(PROJECT_DIR) && go test -v ./...

# 清理构建文件
clean:
	rm -f $(BIN_DIR)/$(BINARY_NAME)

.PHONY: all build run test clean