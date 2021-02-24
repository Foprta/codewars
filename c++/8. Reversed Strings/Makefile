TEST_FILE_NAME=tests

all: clean build run_tests

build: 
	g++ ${TEST_FILE_NAME}.cpp -o ${TEST_FILE_NAME} -lgtest_main -lgtest

clean:
	rm ${TEST_FILE_NAME}

run_tests:
	./${TEST_FILE_NAME}