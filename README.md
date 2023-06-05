
# File Server

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

File Server is a small and lightweight Go application that allows you to share a directory over the local network. It provides a simple and convenient way to serve files from a specified directory, making it easy for users on the same network to access and download files.

## Features

* Share a directory over the local network.
* Customize the port and directory path through configuration or command line flags.
* Easy setup and usage.

## Installation

To use the File Server, make sure you have Go installed on your system. Then, follow these steps:

```
go install github.com/VinukaThejana/file-server@latest
```

## Usage

To run the File Server, use the following command:

```
file-server [--port PORT] [--path PATH]
```

The server will start and listen for incoming requests on the specified port. By default, if no port is provided, it will use port 8080. If no path is provided, the default home directory will be used.

## Configuration

Alternatively, you can define the port and path in a config.yml file located in the ~/.config/file-server directory. The format of the config.yml file should be as follows:

```
port: 8080
path: /path/to/directory
```

## Piping the path

You can also pipe the path to the File Server using a command like pwd to serve your current working directory. For example:

```
pwd | file-server
```

The server will use the piped path as the directory to be served.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

