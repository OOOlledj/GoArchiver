## In development##

Simple executable file to add all files in the directory to archieve  &nbsp;

- Directories recognizes as single file (added to archieve with it's content recursivly)  
- Can work with single file (ex. iso images, which are usually large);  
- ~~Does not work with absolute paths, like "/home/user/test.txt"~~;  
- Can work with absolute and relative paths (using library "path")

Options to run (**options are unavailable now, but will be released in future with CLI usage**):
- -d <directory_path> set up diretory to work with; if not specified - work with current directory
- -s (separately) each of listed files will be added to separate archieve. if not specified - put all directory content to a single archieve
- -tar.gz
- -zip
- -rar