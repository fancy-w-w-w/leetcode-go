TARGET = leetcode

PROJECT_PATH = $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

CMD_DIR = $(PROJECT_PATH)/cmd

BUILD_DIR = $(PROJECT_PATH)/build

ifdef CC_GO
CC=$(CC_GO)
endif

ifdef CXX_GO
CXX = $(CXX_GO)
endif

all:
	@echo $(BUILD_DIR)
	@mkdir -p $(BUILD_DIR)
	@cd $(CMD_DIR) && CC=$(CC) CXX=$(CXX) go build --ldflags '-extldflags "-static"' -v -o $(BUILD_DIR)/$(TARGET)
	@echo "build success"

clean:
	rm -f $(BUILD_DIR)/$(TARGET)