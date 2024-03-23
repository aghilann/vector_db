.PHONY: all clean run

BIN_DIR := bin
EXECUTABLE := $(BIN_DIR)/main

# Target 'all' depends on 'run' now, to separate building and running
all: run

# New build target for just compiling the code
build: $(EXECUTABLE)

$(EXECUTABLE): main.go IndexFlatL2.go | $(BIN_DIR)
	go build -o $(EXECUTABLE) .

$(BIN_DIR):
	mkdir -p $(BIN_DIR)

# Separate run target to execute the binary, depends on 'build' to ensure it's up to date
run: build
	./$(EXECUTABLE)

clean:
	rm -rf $(BIN_DIR)
