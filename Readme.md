# audigo
3D Led CubeのPCレス(Raspberry pi)音響サービス  

<!-- toc -->  
* [💊  Requirements](#-requirements)
* [📌 Installing](#-installing)
* [🎧  Usage](#-usage)
* [🌏  Api](#-api)
* [🎃  Notes](#-notes)
<!-- tocstop -->  

# Getting Started
## 💊 Requirements

* git
* Go 1.11 or later

## 📌 Installing

1. Goto GOPATH  
    **WIndows**
    ```sh
    $ cd %GOPATH%/go
    ```

    **Others**
    ```sh
    $ cd $HOME/go
    ```

2. Get src
    ```sh
    $ git clone https://github.com/code560/audigo.git ./src/github.com/code560/audigo
    $ cd ./src/github.com/code560/audigo
    $ dep ensure
    ```

# 🎧 Usage
Start audio service.
```sh
$ go run audigo.go 80
```

## 🔨 Commands

### port
add port number. default port 8080


# 🌏️ Api
| REST | URI                             | note                          | arguments     |
|------|---------------------------------|-------------------------------|---------------|
| GET  | /audio/v1/ping                  | I Can Fly !                   | none          |
| POST | /audio/v1/init/\<content id>    | init players in memory        | none          |
| POST | /audio/v1/play/\<content id>    | play sound                    | src: "bgm_wave.wav" (file name in ./asset/audio/) <br>loop: true or false (loop play or single play) <br>stop: true or false (start and pause or normal play)        |
| POST | /audio/v1/stop/\<content id>    | stop content player sound     | none          |
| POST | /audio/v1/pause/\<content id>   | pause content player sound    | none          |
| POST | /audio/v1/resume/\<content id>  | resume content player sound   | none          |
| POST | /audio/v1/volume/\<content id>  | change volume                 | vol: 2 (0 - n, 0 is silent)          |


# 🎃 Notes

| Platform / Architecture        | x86 | x86_64 |
|--------------------------------|-----|--------|
| Windows (7, 10 or Later)       |     | ✓     |
| Rasbian (?)                    |     |        |
| OSX (?)                        |     |        |


以上  