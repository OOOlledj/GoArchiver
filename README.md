## In development ##

Simple executable to add all provided files to archieve  &nbsp;

- Directories recognizes as single file (added to archieve with it's content recursivly)  
- Can work with single file (ex. iso images, which are usually large);  
- ~~Does not work with absolute paths, like "/home/user/test.txt"~~;  
- Can work with absolute and relative paths (using library "path")

Options to run in CLI:
- **-m** archieve mod (-m=targz; -m=zip)
- **-o** output filename, can be used with or without format
- **-h** call help
  
Example of usage:

        gfarch -m=targz -o=test ./test go.mod ./go.work

        gfarch -m=zip -o=test.archive ./test go.mod ./go.work
