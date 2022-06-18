
if ($IsLinux) {
    # First time run
    # sudo apt-get install golang gcc libgl1-mesa-dev xorg-dev
    go build -o dist/go-desktop-app
}


if ($IsWindows) {

    # Used references
    # https://www.msys2.org/docs/ci/
    # https://www.msys2.org/wiki/Launchers/
    C:\tools\msys64\usr\bin\env.exe `
        MSYSTEM=MINGW64 `
        CHERE_INVOKING=1 `
        /usr/bin/bash -lc "go build -o dist/go-desktop-app.exe ."

}
