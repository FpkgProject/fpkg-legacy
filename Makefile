PROJ_NAME=fpkg
all:
		rm -rf build
		mkdir build
		cd src/; \
		go build -o ../main
		mv main "build/$(PROJ_NAME)"
