REM 编译32位的dll

set GOARCH=386
set GOROOT=C:\go386

set str=%path%
set path=C:\mingw32\MinGW\bin;%str%

C:\go386\bin\go.exe build -ldflags "-s -w" -v -x -buildmode=c-archive -o c\go.a 

set path = %str%
set GOROOT=C:\Go
set GOARCH=amd64

pause