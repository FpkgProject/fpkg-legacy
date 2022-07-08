all:
		rm -rf build
		mkdir build
		go build "./src/main.go"
		mv main "build/fpkg"
