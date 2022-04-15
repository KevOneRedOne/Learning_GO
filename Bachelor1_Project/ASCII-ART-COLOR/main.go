package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

//EntreeUtilisateur allows the user to choose which character he wants to be colored
func EntreeUtilisateur(color string) []int {
	var colorplace []int

	fmt.Print("\033[H\033[2J")

	fmt.Println("Comment voulez-vous colorier le mot ?\n")
	fmt.Println("Si vous voulez tout colorier tapez \"Entrée\" sinon : ")
	fmt.Println("- Taper les lettres que vous voulez colorier")
	fmt.Println("- Laissez des espaces sur celles que vous ne voulez pas colorier (cf. Exemple ci dessous)\n")
	fmt.Println("Exemple : ")
	fmt.Println("Couleur")
	fmt.Println("Cou eur")
	fmt.Println(colors[color] + "Cou" + "\033[0m" + "l" + colors[color] + "eur" + "\033[0m" + "\n")
	fmt.Println("Vous aviez originellent entré :", flag.Args()[0])
	fmt.Print("Vous voulez colorier :		")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	acolorier := scanner.Text()

	if acolorier == "" {
		for index := range []rune(flag.Args()[0]) {
			colorplace = append(colorplace, index)
		}
		return colorplace
	}

	var (
		reset bool
		cmpt  int
	)

	for index := 0; index < len(acolorier); index++ {
		if reset == true {
			index = 0
			reset = false
		}

		if len(acolorier) > len(flag.Args()[0]) {
			fmt.Print("\033[H\033[2J")
			fmt.Println("MAUVAIS CHOIX")
			fmt.Println("Vous devez coloriez : 	", flag.Args()[0])
			fmt.Println("Vous avez entré : 	", acolorier, "\n")
			fmt.Print("Recommencer : ")
			scanner = bufio.NewScanner(os.Stdin)
			scanner.Scan()
			acolorier = scanner.Text()
			reset = true
		}

		if acolorier[index] == flag.Args()[0][index] {
			colorplace = append(colorplace, index)
		} else if acolorier[index] == ' ' {
			continue
		} else {
			if cmpt >= 4 {
				fmt.Print("\033[H\033[2J")
				fmt.Println("MAUVAIS CHOIX")
				fmt.Println("Vous devez coloriez : 	", flag.Args()[0])
				fmt.Println("Vous avez entré : 	", acolorier, "\n")
				fmt.Println("Pour rappel vous devez :")
				fmt.Println(" - Taper les lettres que vous voulez colorier")
				fmt.Println(" - Laissez des espaces pour les lettres que vous ne voulez pas colorier\n")
				fmt.Println("Exemple : ")
				fmt.Println("Couleur")
				fmt.Println("Cou eur")
				fmt.Println(colors[color] + "C" + "\033[0m" + "o" + colors[color] + "u" + "\033[0m" + "l" + colors[color] + "e" + "\033[0m" + colors[color] + "u" + "\033[0m" + colors[color] + "r" + "\033[0m" + "\n")
				fmt.Println("Si vous voulez tout colorier tapez \"Entrée\"")
				fmt.Println("Votre mot originel :   ", flag.Args()[0])
				fmt.Print("Recommencer :		")
			} else if cmpt >= 1 {
				fmt.Print("\033[H\033[2J")
				fmt.Println("MAUVAIS CHOIX")
				fmt.Println("Vous devez coloriez : 	", flag.Args()[0])
				fmt.Println("Vous avez entré : 	", acolorier, "\n")
				fmt.Println("Pour rappel vous devez :")
				fmt.Println(" - Taper les lettres que vous voulez colorier")
				fmt.Println(" - Laissez des espaces pour les lettres que vous ne voulez pas colorier\n")
				fmt.Print("Recommencer : ")
			} else {
				fmt.Print("\033[H\033[2J")
				fmt.Println("MAUVAIS CHOIX")
				fmt.Println("Vous devez coloriez : 	", flag.Args()[0])
				fmt.Println("Vous avez entré : 	", acolorier, "\n")
				fmt.Print("Recommencer : ")
			}
			scanner = bufio.NewScanner(os.Stdin)
			scanner.Scan()
			acolorier = scanner.Text()
			reset = true
			if acolorier == "" {
				for index := range []rune(flag.Args()[0]) {
					colorplace = append(colorplace, index)
				}
				return colorplace
			}
			colorplace = colorplace[:0]
			cmpt++
		}
	}
	return colorplace
}

//ReadFile return an array of string which is the same as the file (line = line)
func ReadFile(StylizedFile string) []string {
	var source []string
	file, _ := os.Open(StylizedFile)  // open the .txt
	scanner := bufio.NewScanner(file) // scanner scans the file
	scanner.Split(bufio.ScanLines)    // set-up scanner preference to read the file line-by-line
	for scanner.Scan() {              // loop that performs a line-by-line scan on each new iteration
		if scanner.Text() != "" {
			source = append(source, scanner.Text()) // adds the value of scanner (that contains the characters from StylizedFile) to source
		}
	}
	file.Close() // closes the file
	return source
}

//Affichage prints the stylized characters
func Affichage(arguments []rune, txtfile []string, color string, colorplace []int) {
	for ligne := 0; ligne < 8; ligne++ { // Each character is composed of 8 lines
		for index, char := range arguments {
			for i := range colorplace {
				if colorplace[i] == index {
					fmt.Print(colors[color])
				}
			}
			fmt.Print(txtfile[ligne+(int(char)-32)*8]) //Prints the stylized character
			fmt.Print("\033[0m")
			if index == len(arguments)-1 && ligne != 7 { //Executes a newline after every stylized character
				fmt.Println()
			}
		}
	}
}

func main() {
	var colorplace []int

	colorPtr := flag.String("color", "white", "a string to define color")
	flag.Parse()
	color := strings.ToLower(*colorPtr)

	_, ok := colors[color]
	if !ok { // check si bonne couleur ou non
		fmt.Println("Error")
		os.Exit(0)
	}

	if len(flag.Args()) == 1 { // lenght of arguments is os.Args[0,1] = 2 => "main.go" "word"
		arguments := []rune(flag.Args()[0])

		if color != "white" {
			fmt.Print("ok")
			colorplace = EntreeUtilisateur(color)
		}

		for index := range arguments {
			if arguments[index] < 32 || arguments[index] > 126 { //If char isn't between space and tilde
				fmt.Println("Error")
				os.Exit(0) //Stops the program from running
			}
		}
		txtfile := ReadFile("standard.txt")
		var start int
		for index := range arguments {
			if arguments[index] == '\\' && index != len(arguments)-1 { // condition to check if at the index of "arguments" the value is strictly equal to '\\' (which means a simple `\`)
				// and if the index is different from the length of "arguments"
				if arguments[index+1] == 'n' { // condition to verify if the character following '\' is a 'n'
					Affichage(arguments[start:index], txtfile, color, colorplace)
					fmt.Println()     // Acts as "\n" and prints a newline
					start = index + 2 // Jumps the newline, actualizes the value start and keeps reading the argument
				}
			} else if index == len(arguments)-1 { //If the end of os.Args[1] is reached
				Affichage(arguments[start:], txtfile, color, colorplace)
			}
		}
	} else {
		fmt.Println("Error")
	}
}

var colors = map[string]string{
	"black":             "\033[38;5;0m",
	"maroon":            "\033[38;5;1m",
	"green":             "\033[38;5;2m",
	"olive":             "\033[38;5;3m",
	"navy":              "\033[38;5;4m",
	"purple":            "\033[38;5;5m",
	"teal":              "\033[38;5;6m",
	"silver":            "\033[38;5;7m",
	"grey":              "\033[38;5;8m",
	"red":               "\033[38;5;9m",
	"lime":              "\033[38;5;10m",
	"yellow":            "\033[38;5;11m",
	"blue":              "\033[38;5;12m",
	"fushia":            "\033[38;5;13m",
	"aqua":              "\033[38;5;14m",
	"white":             "\033[0m",
	"grey0":             "\033[38;5;16m",
	"navyblue":          "\033[38;5;17m",
	"darkblue":          "\033[38;5;18m",
	"blue3":             "\033[38;5;19m",
	"blue31":            "\033[38;5;20m",
	"blue1":             "\033[38;5;21m",
	"darkgreen":         "\033[38;5;22m",
	"deepskyblue4":      "\033[38;5;23m",
	"deepskyblue41":     "\033[38;5;24m",
	"deepskyblue42":     "\033[38;5;25m",
	"dodgerblue3":       "\033[38;5;26m",
	"dodgerblue2":       "\033[38;5;27m",
	"green4":            "\033[38;5;28m",
	"springgreen4":      "\033[38;5;29m",
	"turquoise4":        "\033[38;5;30m",
	"deepskyblue3":      "\033[38;5;31m",
	"deepskyblue31":     "\033[38;5;32m",
	"dodgerblue1":       "\033[38;5;33m",
	"green3":            "\033[38;5;34m",
	"springgreen3":      "\033[38;5;35m",
	"darkcyan":          "\033[38;5;36m",
	"lightseagreen":     "\033[38;5;37m",
	"deepskyblue2":      "\033[38;5;38m",
	"deepskyblue1":      "\033[38;5;39m",
	"green31":           "\033[38;5;40m",
	"springgreen31":     "\033[38;5;41m",
	"springgreen2":      "\033[38;5;42m",
	"cyan3":             "\033[38;5;43m",
	"darktutquoise":     "\033[38;5;44m",
	"turquoise2":        "\033[38;5;45m",
	"green1":            "\033[38;5;46m",
	"springgreen21":     "\033[38;5;47m",
	"springgreen1":      "\033[38;5;48m",
	"medianspringgreen": "\033[38;5;49m",
	"cyan2":             "\033[38;5;50m",
	"cyan1":             "\033[38;5;51m",
	"darkred":           "\033[38;5;52m",
	"deeppink4":         "\033[38;5;53m",
	"purple4":           "\033[38;5;54m",
	"purple41":          "\033[38;5;55m",
	"purple3":           "\033[38;5;56m",
	"blueviolet":        "\033[38;5;57m",
	"orange4":           "\033[38;5;58m",
	"grey37":            "\033[38;5;59m",
	"mediumpurple4":     "\033[38;5;60m",
	"slateblue3":        "\033[38;5;61m",
	"slateblue31":       "\033[38;5;62m",
	"royalblue1":        "\033[38;5;63m",
	"chartreuse4":       "\033[38;5;64m",
	"darkseagreen4":     "\033[38;5;65m",
	"paleturquoise4":    "\033[38;5;66m",
	"steelblue":         "\033[38;5;67m",
	"steelblue3":        "\033[38;5;68m",
	"cornflowerblue":    "\033[38;5;69m",
	"chartreuse3":       "\033[38;5;70m",
	"darkseagreen41":    "\033[38;5;71m",
	"cadetblue":         "\033[38;5;72m",
	"cadetblue1":        "\033[38;5;73m",
	"skyblue3":          "\033[38;5;74m",
	"steelblue1":        "\033[38;5;75m",
	"chartreuse31":      "\033[38;5;76m",
	"palegreen3":        "\033[38;5;77m",
	"seagreen3":         "\033[38;5;78m",
	"aquamarine3":       "\033[38;5;79m",
	"mediumturquoise":   "\033[38;5;80m",
	"steelblue11":       "\033[38;5;81m",
	"chartreuse2":       "\033[38;5;82m",
	"seagreen2":         "\033[38;5;83m",
	"seagreen1":         "\033[38;5;84m",
	"seagreen11":        "\033[38;5;85m",
	"aquamarine1":       "\033[38;5;86m",
	"darkslategray2":    "\033[38;5;87m",
	"darkred1":          "\033[38;5;88m",
	"deeppink41":        "\033[38;5;89m",
	"darkmagenta":       "\033[38;5;90m",
	"darkmagenta1":      "\033[38;5;91m",
	"darkviolet":        "\033[38;5;92m",
	"purple1":           "\033[38;5;93m",
	"orange41":          "\033[38;5;94m",
	"lightpink4":        "\033[38;5;95m",
	"plum4":             "\033[38;5;96m",
	"mediumpurple3":     "\033[38;5;97m",
	"mediumpurple31":    "\033[38;5;98m",
	"slateblue1":        "\033[38;5;99m",
	"yellow4":           "\033[38;5;100m",
	"wheat4":            "\033[38;5;101m",
	"grey53":            "\033[38;5;102m",
	"lightslategrey":    "\033[38;5;103m",
	"mediumpurple":      "\033[38;5;104m",
	"lightslateblue":    "\033[38;5;105m",
	"yellow41":          "\033[38;5;106m",
	"darkolivegreen3":   "\033[38;5;107m",
	"darkseagreen":      "\033[38;5;108m",
	"lightskyblue3":     "\033[38;5;109m",
	"lightskyblue31":    "\033[38;5;110m",
	"skyblue2":          "\033[38;5;111m",
	"chartreuse21":      "\033[38;5;112m",
	"darkolivegreen31":  "\033[38;5;113m",
	"palegreen31":       "\033[38;5;114m",
	"darkslategreen3":   "\033[38;5;115m",
	"darkslategrey3":    "\033[38;5;116m",
	"skyblue1":          "\033[38;5;117m",
	"chartreuse1":       "\033[38;5;118m",
	"lightgreen":        "\033[38;5;119m",
	"lightgreen1":       "\033[38;5;120m",
	"palegreean1":       "\033[38;5;121m",
	"aquamarine11":      "\033[38;5;122m",
	"darkslategrey1":    "\033[38;5;123m",
	"red3":              "\033[38;5;124m",
	"deeppink42":        "\033[38;5;125m",
	"mediumvioletred":   "\033[38;5;126m",
	"magenta3":          "\033[38;5;127m",
	"darkviolet1":       "\033[38;5;128m",
	"purple2":           "\033[38;5;129m",
	"darkorange3":       "\033[38;5;130m",
	"indianred":         "\033[38;5;131m",
	"hotpink3":          "\033[38;5;132m",
	"mediumorchid3":     "\033[38;5;133m",
	"mediumorchid":      "\033[38;5;134m",
	"mediumpurple2":     "\033[38;5;135m",
	"darkgolderod":      "\033[38;5;136m",
	"lightsalmon3":      "\033[38;5;137m",
	"rosybrown":         "\033[38;5;138m",
	"grey63":            "\033[38;5;139m",
	"mediumpurple21":    "\033[38;5;140m",
	"mediumpurple1":     "\033[38;5;141m",
	"gold3":             "\033[38;5;142m",
	"darkkhak":          "\033[38;5;143m",
	"navajowhite3":      "\033[38;5;144m",
	"grey69":            "\033[38;5;145m",
	"lightsteelblue3":   "\033[38;5;146m",
	"lightslyleblue":    "\033[38;5;147m",
	"yellow3":           "\033[38;5;148m",
	"darkolivegreen32":  "\033[38;5;149m",
	"darkseagreen3":     "\033[38;5;150m",
	"darkseagreen2":     "\033[38;5;151m",
	"lightcyan3":        "\033[38;5;152m",
	"lightskyblue1":     "\033[38;5;153m",
	"greenyellow":       "\033[38;5;154m",
	"darkolivegreen2":   "\033[38;5;155m",
	"palegreen1":        "\033[38;5;156m",
	"darkseagreen21":    "\033[38;5;157m",
	"darkseagreen1":     "\033[38;5;158m",
	"paleturquoise1":    "\033[38;5;159m",
	"red31":             "\033[38;5;160m",
	"deeppink31":        "\033[38;5;161m",
	"deeppink3":         "\033[38;5;162m",
	"magenta31":         "\033[38;5;163m",
	"magenta32":         "\033[38;5;164m",
	"magenta2":          "\033[38;5;165m",
	"darkorange31":      "\033[38;5;166m",
	"indianred1":        "\033[38;5;167m",
	"hotpink31":         "\033[38;5;168m",
	"hotpink2":          "\033[38;5;169m",
	"orchid":            "\033[38;5;170m",
	"mediumorchid1":     "\033[38;5;171m",
	"orange3":           "\033[38;5;172m",
	"lightsalmon31":     "\033[38;5;173m",
	"lightpink3":        "\033[38;5;174m",
	"pink3":             "\033[38;5;175m",
	"plum3":             "\033[38;5;176m",
	"violet":            "\033[38;5;177m",
	"gold31":            "\033[38;5;178m",
	"lightgoldenrod3":   "\033[38;5;179m",
	"tan":               "\033[38;5;180m",
	"mistyrose3":        "\033[38;5;181m",
	"thistle3":          "\033[38;5;182m",
	"plum2":             "\033[38;5;183m",
	"yellow31":          "\033[38;5;184m",
	"khaki3":            "\033[38;5;185m",
	"lightgoldenrod2":   "\033[38;5;186m",
	"lightyellow3":      "\033[38;5;187m",
	"grey84":            "\033[38;5;188m",
	"lightsteelblue1":   "\033[38;5;189m",
	"yellow2":           "\033[38;5;190m",
	"darkolivegreen1":   "\033[38;5;191m",
	"darkolivegreen11":  "\033[38;5;192m",
	"darkseagreen11":    "\033[38;5;193m",
	"honeydew2":         "\033[38;5;194m",
	"lightcyan1":        "\033[38;5;195m",
	"red1":              "\033[38;5;196m",
	"deeppink2":         "\033[38;5;197m",
	"deeppink1":         "\033[38;5;198m",
	"deeppink11":        "\033[38;5;199m",
	"magenta1":          "\033[38;5;200m",
	"magenta11":         "\033[38;5;201m",
	"orangered1":        "\033[38;5;202m",
	"indianred11":       "\033[38;5;203m",
	"indianred12":       "\033[38;5;204m",
	"hotpink":           "\033[38;5;205m",
	"hotpink1":          "\033[38;5;206m",
	"mediumorchid11":    "\033[38;5;207m",
	"darkorange":        "\033[38;5;208m",
	"salmon1":           "\033[38;5;209m",
	"lightcoral":        "\033[38;5;210m",
	"palevioletred":     "\033[38;5;211m",
	"orchid2":           "\033[38;5;212m",
	"orchid1":           "\033[38;5;213m",
	"orange":            "\033[38;5;214m",
	"sandybrown":        "\033[38;5;215m",
	"lightsalmon1":      "\033[38;5;216m",
	"lightpink":         "\033[38;5;217m",
	"pink":              "\033[38;5;218m",
	"plum1":             "\033[38;5;219m",
	"gold1":             "\033[38;5;220m",
	"lightgoldenrod21":  "\033[38;5;221m",
	"lightgoldenrod22":  "\033[38;5;222m",
	"navajowhite1":      "\033[38;5;223m",
	"mistyrose1":        "\033[38;5;224m",
	"thistle1":          "\033[38;5;225m",
	"yellow1":           "\033[38;5;226m",
	"lightgoldenrod1":   "\033[38;5;227m",
	"khakil1":           "\033[38;5;228m",
	"wheat1":            "\033[38;5;229m",
	"cornsilk1":         "\033[38;5;230m",
	"grey100":           "\033[38;5;231m",
	"grey3":             "\033[38;5;232m",
	"grey7":             "\033[38;5;233m",
	"grey11":            "\033[38;5;234m",
	"grey15":            "\033[38;5;235m",
	"grey19":            "\033[38;5;236m",
	"grey23":            "\033[38;5;237m",
	"grey27":            "\033[38;5;238m",
	"grey30":            "\033[38;5;239m",
	"grey35":            "\033[38;5;240m",
	"grey39":            "\033[38;5;241m",
	"grey42":            "\033[38;5;242m",
	"grey46":            "\033[38;5;243m",
	"grey50":            "\033[38;5;244m",
	"grey54":            "\033[38;5;245m",
	"grey58":            "\033[38;5;246m",
	"grey62":            "\033[38;5;247m",
	"grey66":            "\033[38;5;248m",
	"grey70":            "\033[38;5;249m",
	"grey74":            "\033[38;5;250m",
	"grey78":            "\033[38;5;251m",
	"grey82":            "\033[38;5;252m",
	"grey85":            "\033[38;5;253m",
	"grey93":            "\033[38;5;255m",
	"grey89":            "\033[38;5;254m",
}
