# Go desktop app



## Project setup
Follow the installation guide from https://developer.fyne.io/started/

The main steps are:
1. Install https://www.msys2.org/
2. Run the following commands
```
pacman -Syu
pacman -S git mingw-w64-x86_64-toolchain

echo "export PATH=$PATH:/c/Program\ Files/Go/bin:~/Go/bin" >> ~/.bashrc
```

Open the `~/.bashrc` and move the PATH variable BEFORE the interactive mode.

Example:
```
export PATH=/mingw64/bin:/usr/local/bin:/usr/bin:/bin:/c/Windows/System32:/c/Windows:/c/Windows/System32/Wbem:/c/Windows/System32/WindowsPowerShell/v1.0/:/usr/bin/site_perl:/usr/bin/vendor_perl:/usr/bin/core_perl:/c/Program\ Files/Go/bin:~/Go/bin


# If not running interactively, don't do anything
[[ "$-" != *i* ]] && return
```

## Run the application
Once all the tools are installed you can run the application with:
```bash
# In powershell
run_app.ps1
```


## Debug from VSCode
0. Ensure that you are using the latest version of GO and dlv
    `go install github.com/go-delve/delve/cmd/dlv@latest`
1. Open the MSYS2 terminal
2. Launch vscode from that MSYS2 terminal with the `code` command
3. Open the project folder and from Debug section run the `Launch Package` command

Do not debug an app that was generated using go run.

Alternatively you can use the `.\run_build_app.ps1 open` command.
It will build the executable file with the debug information included.
At any point in time you can attach to the running process.

