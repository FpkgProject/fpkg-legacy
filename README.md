## :warning: No longer maintained :warning:
This is no longer maintained, please consider using [RedsonBr140/gpkg](https://github.com/RedsonBr140/gpkg)

# Fpkg
A package manager for Unix-like Operating systems, written in Golang.

The Purpose of fpkg is to create a universal package manager for Unix-like operating systems.

## How it works
fpkg is source-code based, it checks if a repository with the specified name exists and provides a list for you to choose from. It then runs `./configure` to... Yes, configure the program, then runs `make` to compile the software and `make install` to install.

The main idea is to act similarly to ArchLinux AUR, Gentoo Portage/Emerge and BSD Ports, but universally.

## Contributing
 - [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) - We are trying to do organized commits, so if you have the intention of contributing, would be cool if you use Conventional Commits.
 
 - Branches
 
 We use two principal branches, `master` and `dev`. ALL new features are first made in the **dev** branch, and after a few testing, merged to the master.
 To doing a Pull Request, you need to:
 1. [Fork](https://github.com/FpkgProject/Fpkg/fork) the repository
 2. Create a new branch from the dev branch:
 ```
 git checkout -b my-new-feature dev
 ```
 3. Add your files, commit and push:
```
git add .
git commit -m "feat: A awesome feature"
git push
```
5. After these steps, open a [Pull Request](https://github.com/FpkgProject/Fpkg/pulls) to merge your branch with our dev branch.

## License
BSD-3-Clause license

---
❤ `Keep it simple, stupid!`
