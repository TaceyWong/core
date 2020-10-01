
## <font color=red> Just to practice basic system programming, not to use in real world </font>

# core


![License](https://img.shields.io/badge/license-GPL3.0-brightgreen.svg?style=flat-square)
![Go Report Card](https://goreportcard.com/badge/github.com/TaceyWong/core)

Golang Implement of Coreutils (git://git.sv.gnu.org/coreutils)



## Main

**Output of entire files**

+ [ ] `cat`: concatenate files and print on the standard output
+ [ ] `tac`: concatenate and print files in reverse 
+ [ ] `nl`:  number lines of files
+ [ ] `od`: dump files in octal and other formats
+ [ ] `base32`: base32 encode/decode data and print to standard output
+ [ ] `base64`: base64 encode/decode data and print to standard output
+ [ ] `basenc`: 

**Formatting file contents**

+ [ ] `fm`: `sudo apt install fmtools` : simple optimal text formatter
+ [ ] `pr`: convert text files for printing
+ [ ] `fold`: wrap each input line to fit in specified width

**Output of parts of files**

+ [ ] `head`: output the first part of files
+ [ ] `tail`: output the last part of files
+ [ ] `split`: split a file into pieces
+ [ ] `csplit`: split a file into sections determined by context lines

**Summarizing files**

+ [ ] `cksum`: checksum and count the bytes in a file
+ [ ] `b2sum`: compute and check BLAKE2 message digest
+ [ ] `md5sum`: compute and check MD5 message digest
+ [ ] `sha1sum`: compute and check SHA1 message digest
+ [ ] `sha2 utilties`: 

**Operating on sorted files**

+ [ ] `sort`: sort lines of text files
+ [ ] `shuf`: generate random permutations
+ [ ] `uniq`: report or omit repeated lines
+ [ ] `comm`: compare two sorted files line by line
+ [ ] `ptx`: produce a permuted index of file contents
+ [ ] `tsort`: perform topological sort

**Operating on fields**

+ [ ] `cut`: remove sections from each line of files
+ [ ] `paste`: merge lines of files
+ [ ] `join`: join lines of two files on a common field

**Operating on characters**

+ [ ] `tr`: translate or delete characters
+ [ ] `expand`: convert tabs to spaces
+ [ ] `unexpand`:  convert spaces to tabs

**Directory listing**

+ [ ] `ls`: list directory contents
+ [ ] `dir`: list directory contents
+ [ ] `vdir`:  list directory contents
+ [ ] `dircolors`: color setup for ls

**Basic operations**

+ [ ] `cp`: copy files and directories
+ [ ] `dd`: convert and copy a file
+ [ ] `install`:  copy files and set attributes
+ [ ] `mv`:  move (rename) files
+ [ ] `rm`: remove files or directories
+ [ ] `shred`:  overwrite a file to hide its contents, and optionally delete it

**Special file types**

+ [ ] `link`: call the link function to create a link to a file
+ [ ] `ln`:  make links between files
+ [ ] `mkdir`:  make directories
+ [ ] `mkfifo`: make FIFOs (named pipes)
+ [ ] `mknod`:  make block or character special files
+ [ ] `readlink`: print resolved symbolic links or canonical file names
+ [ ] `rmdir`: remove empty directories
+ [ ] `unlink`: call the unlink function to remove the specified file

**Changing file attributes**

+ [ ] `chown`: change file owner and group
+ [ ] `chgrp`: change group ownership
+ [ ] `chmod`:  change file mode bits
+ [ ] `touch`: change file timestamps

**Disk usage**

+ [ ] `df`: report file system disk space usage
+ [ ] `du`:  estimate file space usage
+ [ ] `stat`: display file or file system status
+ [ ] `sync`: Synchronize cached writes to persistent storage
+ [ ] `truncate`: shrink or extend the size of a file to the specified size 

**Print Text**

+ [ ] `echo`:  display a line of text
+ [ ] `printf`: format and print data
+ [ ] `yes`: output a string repeatedly until killed

**Conditions**

+ [ ] `false`: do nothing, unsuccessfully
+ [ ] `true`: do nothing, successfully
+ [ ] `test`:  check file types and compare values
+ [ ] `expr`:  evaluate expressions

**Redirection**

+ [ ] `tee`:  read from standard input and write to standard output and files

**File name manipulation**

+ [ ] `basename`: strip directory and suffix from filenames
+ [ ] `dirname`: strip last component from file name
+ [ ] `pathchk`: check whether file names are valid or portable
+ [ ] `mktemp`: create a temporary file or directory
+ [ ] `realpath`: print the resolved path

**Working context**

+ [ ] `pwd`:  print name of current/working directory
+ [ ] `stty`:  change and print terminal line settings
+ [ ] `printenv`:  print all or part of environment
+ [ ] `tty`:  print the file name of the terminal connected to standard input

**User information**

+ [ ] `id`: print real and effective user and group IDs
+ [ ] `logname`: print user´s login name
+ [ ] `whoami`:  print effective userid
+ [ ] `groups`: print the groups a user is in
+ [ ] `users`: - print the user names of users currently logged in to the current host
+ [ ] `who`: show who is logged on

**System context**

+ [ ] `date`: print or set the system date and time
+ [ ] `arch`:  print machine hardware name (same as uname -m)
+ [ ] `nproc`: print the number of processing units available
+ [ ] `uname`: print system information
+ [ ] `hostname`: show or set the system's host name
+ [ ] `hostid`:  print the numeric identifier for the current host
+ [ ] `uptime`: Tell how long the system has been running


**SELinux context**

+ [ ] `chcon`: change file security context
+ [ ] `runcon`: run command with specified security context

**Modified command ivocation**

+ [ ] `chroot`:run command or interactive shell with special root directory
+ [ ] `env`: run a program in a modified environment
+ [ ] `nice`: run a program with modified scheduling priority
+ [ ] `nohup`: run a command immune to hangups, with output to a non-tty
+ [ ] `stdbuf`: Run COMMAND, with modified buffering operations for its standard streams.
+ [ ] `timeout`:  run a command with a time limit

**Process control**

+ [ ] `kill`: send a signal to a process

**Delaying**

+ [ ] `sleep`:  delay for a specified amount of time


**Numeric operation**

+ [ ] `factor`: factor numbers 
+ [ ] `numfmt`: Convert numbers from/to human-readable strings
+ [ ] `seq`: print a sequence of numbers


## References

+ https://www.gnu.org/software/coreutils/
+ http://www.maizure.org/projects/decoded-gnu-coreutils/
+ https://en.wikipedia.org/wiki/List_of_GNU_Core_Utilities_commands




















































































