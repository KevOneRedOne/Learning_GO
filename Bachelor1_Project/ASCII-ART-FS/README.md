# ASCII-ART-FS


## Presentation

> Project ascii-art-fs is the 2nd project in Go-Reloaded

The team is composed of:
 - Kyllian (as KGROLLEMUND) (Captain)
 - Marc (as MRAZAFINDRIANTSOA)
 - Yanis (as YDJOUDI)
 - Kévin (as KALVES)
 - Antoine (as ABERNARD3)

Date of launch is Monday, 14th of December 2020.

Language used is Golang.

Every built-in function is allowed.

<hr>

## Explanation

ascii-art-fs consists on receiving two `string` as arguments :
The first `string` is the text to be changed into art.
The second `string` is the choosen graphical representation of ASCII characters (standard, thinkertoy, shadow).

### Example

**Command prompt**

```console
student@ubuntu:~/ascii-art$ go build
student@ubuntu:~/ascii-art$ ./ascii-art "hello" standard
  _                _    _
 | |              | |  | |
 | |__      ___   | |  | |    ___
 |  _ \    / _ \  | |  | |   / _ \
 | | | |  |  __/  | |  | |  | (_) |
 |_| |_|   \___|  |_|  |_|   \___/



student@ubuntu:~/ascii-art$ ./ascii-art "Hello There!" shadow

_|    _|          _| _|                _|_|_|_|_| _|                                  _|
_|    _|   _|_|   _| _|   _|_|             _|     _|_|_|     _|_|   _|  _|_|   _|_|   _|
_|_|_|_| _|_|_|_| _| _| _|    _|           _|     _|    _| _|_|_|_| _|_|     _|_|_|_| _|
_|    _| _|       _| _| _|    _|           _|     _|    _| _|       _|       _|
_|    _|   _|_|_| _| _|   _|_|             _|     _|    _|   _|_|_| _|         _|_|_| _|



student@ubuntu:~/ascii-art$ ./ascii-art "Hello There!" thinkertoy

o  o     o o           o-O-o o
|  |     | |             |   |                o
O--O o-o | | o-o         |   O--o o-o o-o o-o |
|  | |-' | | | |         |   |  | |-' |   |-' o
o  o o-o o o o-o         o   o  o o-o o   o-o
                                              O


student@ubuntu:~/ascii-art$
```

<hr>