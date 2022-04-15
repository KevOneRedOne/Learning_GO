# ASCII-ART-COLOR-FS


## Presentation

> Project Ascii-color-fs is the 2nd project in Ascii-Art

The team is composed of:
 - Kyllian (as KGROLLEMUND) (Captain)
 - Marc (as MRAZAFINDRIANTSOA)
 - Yanis (as YDJOUDI)
 - Kévin (as KALVES)
 - Antoine (as ABERNARD3)

<hr>

## Explanation

The project Ascii-color-fs consists on receiving three `string` as arguments :
 - The first `string` is a flag which allows to choose the color of the characters.
 - The second `string` is the text to be changed into art.
 - The third `string` is the choosen graphical representation of ASCII characters (standard, thinkertoy, shadow).

The program can work with :
 - Only the text to be changed ;
 - Only the text and the chosen color ;
 - Only the text and the chosen graphical representation of ASCII characters (standard, thinkertoy, shadow).
 - Only the text, the color and the chosen graphical representation of ASCII characters (standard, thinkertoy, shadow).

<hr>

### How to use the program ?

If you want only your text in ASCII :
 - writes the program and the text on the console. 

```console
PS C:\VSCODEProjet\Ytrack\B1\ASCII-ART-COLOR\ASCII-COLOR-FS> go run main.go "Hello There"
 _    _          _   _                 _______   _
| |  | |        | | | |               |__   __| | |
| |__| |   ___  | | | |   ___            | |    | |__     ___   _ __    ___
|  __  |  / _ \ | | | |  / _ \           | |    |  _ \   / _ \ | '__|  / _ \
| |  | | |  __/ | | | | | (_) |          | |    | | | | |  __/ | |    |  __/
|_|  |_|  \___| |_| |_|  \___/           |_|    |_| |_|  \___| |_|     \___|

```

If you want to chose an other graphical representation of the Ascii for your text :
 - writes the program, the text and the graphical representation (standard, thinkertoy or shadow)


```console
PS C:\VSCODEProjet\Ytrack\B1\ASCII-ART-COLOR\ASCII-COLOR-FS> go run main.go "Hello There" shadow

_|    _|          _| _|                _|_|_|_|_| _|
_|    _|   _|_|   _| _|   _|_|             _|     _|_|_|     _|_|   _|  _|_|   _|_|
_|_|_|_| _|_|_|_| _| _| _|    _|           _|     _|    _| _|_|_|_| _|_|     _|_|_|_|
_|    _| _|       _| _| _|    _|           _|     _|    _| _|       _|       _|
_|    _|   _|_|_| _| _|   _|_|             _|     _|    _|   _|_|_| _|         _|_|_|


PS C:\VSCODEProjet\Ytrack\B1\ASCII-ART-COLOR\ASCII-COLOR-FS> go run main.go "Hello There" thinkertoy
     
o  o     o o           o-O-o o
|  |     | |             |   |
O--O o-o | | o-o         |   O--o o-o o-o o-o
|  | |-' | | | |         |   |  | |-' |   |-'
o  o o-o o o o-o         o   o  o o-o o   o-o

```


If you want to add color on your ascii text :
 - Writes the program name, the flag of you color (example : "--color=red") and you text.
 - Then, the program will ask you to choose the letters you want to color :
    - Press enter if you want to color all of the text 
    - Or, writes only the letters to be colored (Don't forget to put spaces for the letters that you don't want to color )

```console 
PS C:\VSCODEProjet\Ytrack\B1\ASCII-ART-COLOR\ASCII-COLOR-FS> go run main.go --color=red "Hello There"  

Comment voulez-vous colorier le mot ?

Si vous voulez tout colorier tapez "Entrée" sinon :
- Taper les lettres que vous voulez colorier
- Laissez des espaces sur celles que vous ne voulez pas colorier (cf. Exemple ci dessous)

Exemple :
Couleur
Cou eur
Couleur

Vous aviez originellent entré : Hello There
Vous voulez colorier :          H l o T e e

 _    _          _   _                 _______   _                           
| |  | |        | | | |               |__   __| | |                          
| |__| |   ___  | | | |   ___            | |    | |__     ___   _ __    ___  
|  __  |  / _ \ | | | |  / _ \           | |    |  _ \   / _ \ | '__|  / _ \ 
| |  | | |  __/ | | | | | (_) |          | |    | | | | |  __/ | |    |  __/ 
|_|  |_|  \___| |_| |_|  \___/           |_|    |_| |_|  \___| |_|     \___| 

```

If you want to add color and to chose a graphical representation of the Ascii for your text :
 - writes the program, the flag of you color (example : "--color=red"),the text and the graphical representation (standard, thinkertoy or shadow)
 - Then, the program will ask you to choose the letters you want to color :
    - Press enter if you want to color all of the text 
    - Or, writes only the letters to be colored (Don't forget to put spaces for the letters that you don't want to color )

```console 
PS C:\VSCODEProjet\Ytrack\B1\ASCII-ART-COLOR\ASCII-COLOR-FS> go run main.go --color=red "Hello There" shadow 

Comment voulez-vous colorier le mot ?

Si vous voulez tout colorier tapez "Entrée" sinon :
- Taper les lettres que vous voulez colorier
- Laissez des espaces sur celles que vous ne voulez pas colorier (cf. Exemple ci dessous)

Exemple :
Couleur
Cou eur
Couleur

Vous aviez originellent entré : Hello There
Vous voulez colorier :          Hello      

_|    _|          _| _|                _|_|_|_|_| _|
_|    _|   _|_|   _| _|   _|_|             _|     _|_|_|     _|_|   _|  _|_|   _|_|
_|_|_|_| _|_|_|_| _| _| _|    _|           _|     _|    _| _|_|_|_| _|_|     _|_|_|_|
_|    _| _|       _| _| _|    _|           _|     _|    _| _|       _|       _|
_|    _|   _|_|_| _| _|   _|_|             _|     _|    _|   _|_|_| _|         _|_|_|


```

<hr>                                                               
            