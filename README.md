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


