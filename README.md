# Fpkg
A package manager for unix-like system built with golang.

The Purpose of fpkg is to create a universal package manager for unix-like system.

## How Works

The Fpkg uses the GNU/make + curl to build and install a new package. In system similar to ports (freeBSD), portage (gentoo) and AUR(archLinux).

1. In search you lock a package url, and call a curl to run make to generate a file, and a make install to install in machine. 
