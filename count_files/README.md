# Count files and directories
This binary count the files and directories of a path.
### Build

```bash
git clone https://github.com/carht/gomisc.git
cd count_files
go build .
```

### Usage with examples

The parameter **-v** is for verbose mode, return the list of files and directories. The position of this parameter don't change the result, see 3 and 4. The example 2 is a synonymous for 3 and 4.

The parameter **-p** is for add a path, if the path don't exists return an error. The default value is "."

-----

* 1
```bash
./count_files
Files: 4, Directories: 1
```

* 2
```bash
./count_files -v
/home/carht/src/gomisc/count_files
/home/carht/src/gomisc/count_files/README.md
/home/carht/src/gomisc/count_files/count_files
/home/carht/src/gomisc/count_files/go.mod
/home/carht/src/gomisc/count_files/main.go
Files: 4, Directories: 1
```

* 3
```bash
./count_files -p "." -v
/home/carht/src/gomisc/count_files
/home/carht/src/gomisc/count_files/README.md
/home/carht/src/gomisc/count_files/count_files
/home/carht/src/gomisc/count_files/go.mod
/home/carht/src/gomisc/count_files/main.go
Files: 4, Directories: 1
```

* 4
```bash
./count_files -v -p "."
/home/carht/src/gomisc/count_files
/home/carht/src/gomisc/count_files/README.md
/home/carht/src/gomisc/count_files/count_files
/home/carht/src/gomisc/count_files/go.mod
/home/carht/src/gomisc/count_files/main.go
Files: 4, Directories: 1
```

* 5
```bash
./count_files -p "/home/carht/src/gomisc"
Files: 68, Directories: 37
```
