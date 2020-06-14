
## gss 'go serve static' 
### Info:
GSS serve's static files from a given directory(recursivly) via http on a defined port. This tool is meant for temporary usage to serve files for development or file movement purposes. It is not meant to replace a webserver.
### Install:
##### Build it yourself
The current makefile supports systems that include /usr/bin in their PATH. If your system doesn't include that you either add it or compile it manually with "go build -o gss" and copy it into your supported PATH. 
```sh 
$ git clone https://github.com/voodooEntity/gss 
$ cd gss 
$ make && make build && make install 
```
### Args:
* -port INT  * Specifys the port number used to server the files, default '8091'*
* -dir STRING *Path to directory to server the files from, default './' * 

### Usage:
##### Serving the current directory on default port 8091
 ```sh 
 $ gss
 ```
##### Serving a given directory '/srv/foo/bar/' on default port 8091
 ```sh 
 $ gss -dir /srv/foo/bar 
 ```
##### Serving a given directory '/srv/foo/bar/' on given port 8008
 ```sh 
 $ gss -dir /srv/foo/bar -port 8008
 ```
##### Serving current directory on a given port 8008
 ```sh 
 $ gss -port 8008 
 ```

### FAQ

##### Why did you build gss?
I just thought there could be a tool with just simplistic input doing the job, and since golang is very good at simple http actions i created it. It's meant to be easy and simple.
##### What are the future plans of threader?
I have no speficic future plans for the tool. The base funcionality is given and i gonne focus on making sure this tool always stays stable and working. Im still open to suggestions on how to improve the tool, but i can't promise anything .) just feel free to comment.

