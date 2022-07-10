# Fpkg
A package manager for unix-like system built with golang.

The Purpose of fpkg is to create a universal package manager for unix-like system.

## How Works

The Fpkg uses the make to build and install a new package. In system similar to ports (freeBSD), portage (gentoo) and AUR(archLinux).

1. In search you lock a package url, if exists in repoList.

2. In install he clone a repo in a temporary directory and run "**./configure**", "**make**" and "**makefile**" case exists to generate a native binary, before this he install in machine and remove the temporary directory. 
