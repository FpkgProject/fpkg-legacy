# Fpkg
A package manager for unix-like system built with golang.

The Purpose of fpkg is to create a universal package manager for unix-like system.

## How it works
Fpkg uses `make` to build and install packages. It works like ports (FreeBSD), Portage (Gentoo) and AUR(Arch Linux).

1. In search you lock a package url, if exists in repoList.

2. In install it clones a repo in a temporary directory and run "**./configure**", "**make**" and "**makefile**" case exists to generate a native binary, before this he install in machine and remove the temporary directory. 
