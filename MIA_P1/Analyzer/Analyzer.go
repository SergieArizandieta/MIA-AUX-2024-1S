package Analyzer

import (
	"flag"
	"fmt"
	"bufio"
	"os"
	"regexp"
	"strings"
	"MIA_P1/DiskManagement"
	"MIA_P1/FileSystem"
	"MIA_P1/User"
)

var re = regexp.MustCompile(`-(\w+)=("[^"]+"|\S+)`)

func getCommandAndParams(input string) (string, string) {
	parts := strings.Fields(input) 
	if len(parts) > 0 {
		command := strings.ToLower(parts[0])                   
		params := strings.Join(parts[1:], " ") 
		return command, params
	}
	return "", input
}

func Analyze(){

	for true {
		var input string
		fmt.Println("Enter command: ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan() 
		input = scanner.Text()

		command, params := getCommandAndParams(input)

		fmt.Println("Command: ", command, "Params: ", params)

		AnalyzeCommnad(command, params)

		//mkdisk -size=3000 -unit=K -fit=BF
		//fdisk -size=300 -driveletter=A -name=Particion1
		//mount -driveletter=A -name=Particion1
		//mkfs -type=full -id=A119
		//login -user=root -pass=123 -id=A119

	}
}

func AnalyzeCommnad(command string, params string){

	if strings.Contains(command, "mkdisk") {
		fn_mkdisk(params)
	}else if strings.Contains(command, "fdisk") {
		fn_fdisk(params)
	}else if strings.Contains(command, "mount") {
		fn_mount(params)
	}else if strings.Contains(command, "mkfs") {
		fn_mkfs(params)
	}else if strings.Contains(command, "login") {
		fn_login(params)
	}else{
		fmt.Println("Error: Command not found")
	}

}

func fn_login(input string){
	// Define flags
	fs := flag.NewFlagSet("login", flag.ExitOnError)
	user := fs.String("user", "", "Usuario")
	pass := fs.String("pass", "", "Contraseña")
	id := fs.String("id", "", "Id")

	// Parse the flags
	fs.Parse(os.Args[1:])

	// find the flags in the input
	matches := re.FindAllStringSubmatch(input, -1)

	// Process the input
	for _, match := range matches {
		flagName := match[1]
		flagValue := match[2]

		flagValue = strings.Trim(flagValue, "\"")

		switch flagName {
		case "user", "pass", "id":
			fs.Set(flagName, flagValue)
		default:
			fmt.Println("Error: Flag not found")
		}
	}

	// Call the function
	User.Login(*user, *pass, *id)

}

func fn_mkfs(input string){
	// Define flags
	fs := flag.NewFlagSet("mkfs", flag.ExitOnError)
	id := fs.String("id", "", "Id")
	type_ := fs.String("type", "", "Tipo")
	fs_ := fs.String("fs", "2fs", "Fs")

	// Parse the flags
	fs.Parse(os.Args[1:])

	// find the flags in the input
	matches := re.FindAllStringSubmatch(input, -1)

	// Process the input
	for _, match := range matches {
		flagName := match[1]
		flagValue := match[2]

		flagValue = strings.Trim(flagValue, "\"")

		switch flagName {
		case "id", "type", "fs":
			fs.Set(flagName, flagValue)
		default:
			fmt.Println("Error: Flag not found")
		}
	}

	// Call the function
	FileSystem.Mkfs(*id, *type_, *fs_)
	

}

func fn_mount(input string){
	// Define flags
	fs := flag.NewFlagSet("mount", flag.ExitOnError)
	driveletter := fs.String("driveletter", "", "Letra")
	name := fs.String("name", "", "Nombre")

	// Parse the flags
	fs.Parse(os.Args[1:])

	// find the flags in the input
	matches := re.FindAllStringSubmatch(input, -1)

	// Process the input
	for _, match := range matches {
		flagName := match[1]
		flagValue := strings.ToLower(match[2])

		flagValue = strings.Trim(flagValue, "\"")

		switch flagName {
		case "driveletter", "name":
			fs.Set(flagName, flagValue)
		default:
			fmt.Println("Error: Flag not found")
		}
	}

	// Call the function
	DiskManagement.Mount(*driveletter, *name)
}

func fn_fdisk(input string) {
	// Define flags
	fs := flag.NewFlagSet("fdisk", flag.ExitOnError)
	size := fs.Int("size", 0, "Tamaño")
	driveletter := fs.String("driveletter", "", "Letra")
	name := fs.String("name", "", "Nombre")
	unit := fs.String("unit", "m", "Unidad")
	type_ := fs.String("type", "p", "Tipo")
	fit := fs.String("fit", "f", "Ajuste")

	// Parse the flags
	fs.Parse(os.Args[1:])

	// find the flags in the input
	matches := re.FindAllStringSubmatch(input, -1)

	// Process the input
	for _, match := range matches {
		flagName := match[1]
		flagValue := strings.ToLower(match[2])

		flagValue = strings.Trim(flagValue, "\"")

		switch flagName {
		case "size", "fit", "unit", "driveletter", "name", "type":
			fs.Set(flagName, flagValue)
		default:
			fmt.Println("Error: Flag not found")
		}
	}

	// Call the function
	DiskManagement.Fdisk(*size, *driveletter, *name, *unit, *type_, *fit)
}

func fn_mkdisk(params string){
	// Define flags
	fs := flag.NewFlagSet("mkdisk", flag.ExitOnError)
	size := fs.Int("size", 0, "Tamaño")
	fit := fs.String("fit", "f", "Ajuste")
	unit := fs.String("unit", "m", "Unidad")

	// Parse the flags
	fs.Parse(os.Args[1:])

	// find the flags in the input
	matches := re.FindAllStringSubmatch(params, -1)

	// Process the input
	for _, match := range matches {
		flagName := match[1]
		flagValue := strings.ToLower(match[2])

		flagValue = strings.Trim(flagValue, "\"")

		switch flagName {
			case "size", "fit", "unit":
				fs.Set(flagName, flagValue)
			default:
				fmt.Println("Error: Flag not found")
		}
	}

	// Call the function
	DiskManagement.Mkdisk(*size, *fit, *unit)

} 

