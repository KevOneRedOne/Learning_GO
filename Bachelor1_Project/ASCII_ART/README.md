# ASCII-ART-COLOR-FS


## Presentation

> Project Ascii-color-fs is our 2nd project in Ascii-Art

The team is composed of:
 - Kyllian (as KGROLLEMUND) (Captain)
 - Marc (as MRAZAFINDRIANTSOA)
 - Yanis (as YDJOUDI)
 - Kévin (as KALVES)
 - Antoine (as ABERNARD3)

<hr>

## Explanation

The project Ascii-color-fs consists on receiving three `string` as arguments entered by the user in the terminal:
 - The first `string` is a flag which allows the user to choose the color of the characters.
 - The second `string` is the text that will be changed into ASCII characters. (using the `standard` graphical representation by default)
 - The third `string` is the chosen graphical representation for ASCII characters (`standard`, `thinkertoy`, `shadow`).

The program can do the following :
 - Convert the text into ASCII characters ;
 - Convert the text and add a color to the selected characters ;
 - Convert the text using one of the graphical representation for ASCII characters (standard, thinkertoy, shadow) ;
 - Convert the text using one of the graphical representation for ASCII characters (standard, thinkertoy, shadow) and add a color to the selected characters.

<hr>

### How to use the program ?

It is recommanded to first, build the program and then to run it:

```
go build main.go
```



If you want to convert your text in ASCII characters:
 - Type the name of the program and the text between double quotes on the terminal.

```
PS C:\VSCODEProjet\Ytrack\B1\ASCII-ART-COLOR\ASCII-COLOR-FS> go run main.go "Hello There"
 _    _          _   _                 _______   _
| |  | |        | | | |               |__   __| | |
| |__| |   ___  | | | |   ___            | |    | |__     ___   _ __    ___
|  __  |  / _ \ | | | |  / _ \           | |    |  _ \   / _ \ | '__|  / _ \
| |  | | |  __/ | | | | | (_) |          | |    | | | | |  __/ | |    |  __/
|_|  |_|  \___| |_| |_|  \___/           |_|    |_| |_|  \___| |_|     \___|

```

If you want to choose another graphical representation for ASCII characters for your text :
 - Type the name of the program, the text between double quotes followed by one of the graphical representation (`standard`, `thinkertoy` or `shadow`)


```
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


If you want to add color to your ASCII characters :
 - Type the name of the program, the flag of your color (example : "--color=red") and your text between double quotes.
 - Then, the program will ask you to choose the letters you want to color :
    - Press enter if you want to color all of the text 
    - Or, type only the letters you want to be colored (Don't forget to put spaces for the letters that you don't want to color ), the other will stay white by default

```
PS C:\VSCODEProjet\Ytrack\B1\ASCII-ART-COLOR\ASCII-COLOR-FS> go run main.go --color=red "Hello There"  

Comment voulez-vous colorier le mot ?

Si vous voulez tout colorier tapez "Entrée" sinon :
- Taper les lettres que vous voulez colorier
- Laissez des espaces sur celles que vous ne voulez pas colorier (cf. Exemple ci dessous)

Exemple :
Couleur
Cou eur
Couleur

Vous aviez originellement entré : Hello There
Vous voulez colorier :            H l o T e e

 _    _          _   _                 _______   _                           
| |  | |        | | | |               |__   __| | |                          
| |__| |   ___  | | | |   ___            | |    | |__     ___   _ __    ___  
|  __  |  / _ \ | | | |  / _ \           | |    |  _ \   / _ \ | '__|  / _ \ 
| |  | | |  __/ | | | | | (_) |          | |    | | | | |  __/ | |    |  __/ 
|_|  |_|  \___| |_| |_|  \___/           |_|    |_| |_|  \___| |_|     \___| 

```

If you want to add a color and choose a graphical representation of ASCII characters for your text :
 - Type the name of the program, the flag of your color (example : "--color=red"),the text between double quotes followed by the graphical representation (standard, thinkertoy or shadow)
 - Then, the program will ask you to choose the letters you want to color :
    - Press enter if you want to color all of the text 
    - Or, type only the letters you want to be colored (Don't forget to put spaces for the letters that you don't want to color )

```
PS C:\VSCODEProjet\Ytrack\B1\ASCII-ART-COLOR\ASCII-COLOR-FS> go run main.go --color=red "Hello There" shadow 

Comment voulez-vous colorier le mot ?

Si vous voulez tout colorier tapez "Entrée" sinon :
- Taper les lettres que vous voulez colorier
- Laissez des espaces sur celles que vous ne voulez pas colorier (cf. Exemple ci dessous)

Exemple :
Couleur
Cou eur
Couleur

Vous aviez originellement entré : Hello There
Vous voulez colorier :            Hello      

_|    _|          _| _|                _|_|_|_|_| _|
_|    _|   _|_|   _| _|   _|_|             _|     _|_|_|     _|_|   _|  _|_|   _|_|
_|_|_|_| _|_|_|_| _| _| _|    _|           _|     _|    _| _|_|_|_| _|_|     _|_|_|_|
_|    _| _|       _| _| _|    _|           _|     _|    _| _|       _|       _|
_|    _|   _|_|_| _| _|   _|_|             _|     _|    _|   _|_|_| _|         _|_|_|


```

<hr>                                                               
            
