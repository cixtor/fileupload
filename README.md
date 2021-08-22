# File Upload [![GoReport](https://goreportcard.com/badge/github.com/cixtor/fileupload)](https://goreportcard.com/report/github.com/cixtor/fileupload) [![GoDoc](https://godoc.org/github.com/cixtor/fileupload?status.svg)](https://godoc.org/github.com/cixtor/fileupload)

> Uploading refers to transmitting data from one computer system to another through means of a network.
>
> Uploading directly contrasts with downloading, where data is received over a network.
> 
> — https://en.wikipedia.org/wiki/Upload

## Installation

```
go get -u github.com/cixtor/fileupload
```

## Usage

Run `fileupload` to start a file server in the current working directory, listening on a random ephemeral port number. The web server returns a web page with a form to allow you, or anyone in your local network, to upload one or more files at once to the machine running the program. Then, you can leverage additional tools like [ngrok](https://ngrok.com) to allow users outside your local network to access the web page and offer them an easy way to transfer one or more files to you without the risks that public file transfer services pose.

```
$ fileupload
2021/08/22 13:50:53 localhost:3000 [::1]:51483 "GET  /              HTTP/1.1" 200 420 "-" 1.571861ms
2021/08/22 13:51:20 localhost:3000 [::1]:51483 "POST /upload        HTTP/1.1" 302   0 "-" 983.471µs
2021/08/22 13:51:20 localhost:3000 [::1]:51483 "GET  /?success=true HTTP/1.1" 200 420 "-" 131.556µs
^C
Server stopped
```

Use option `-m` to control the maximum amount of bytes per file the server will allow users to upload, default is 20 MiB. Use option `-mp` to control the maximum (accumulated) amount of bytes to hold in memory before all files are uploaded, default is 200 MiB.
