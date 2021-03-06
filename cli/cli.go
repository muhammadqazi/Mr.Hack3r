package cli

import (
	"fmt"
	"os"

	"github.com/muhammadqazi/Mr.Hack3r/controllers"
	"github.com/muhammadqazi/Mr.Hack3r/helpers"
	Colors "github.com/muhammadqazi/Mr.Hack3r/utils"
)

func Cli() {
	const Head = `
	
	███╗   ███╗██████╗    ██╗  ██╗ █████╗  ██████╗██╗  ██╗██████╗ ██████╗ 
	████╗ ████║██╔══██╗   ██║  ██║██╔══██╗██╔════╝██║ ██╔╝╚════██╗██╔══██╗
	██╔████╔██║██████╔╝   ███████║███████║██║     █████╔╝  █████╔╝██████╔╝
	██║╚██╔╝██║██╔══██╗   ██╔══██║██╔══██║██║     ██╔═██╗  ╚═══██╗██╔══██╗
	██║ ╚═╝ ██║██║  ██║██╗██║  ██║██║  ██║╚██████╗██║  ██╗██████╔╝██║  ██║
	╚═╝     ╚═╝╚═╝  ╚═╝╚═╝╚═╝  ╚═╝╚═╝  ╚═╝ ╚═════╝╚═╝  ╚═╝╚═════╝ ╚═╝  ╚═╝
																		  
	`

	fmt.Println(string(Colors.Red()), Head)

	fmt.Println(Colors.Red(), "\t[", Colors.White(), " * ", Colors.Red(), "]", Colors.White(), "Modern Cracking Tools With Advanced Functionality", Colors.Red(), "\t[", Colors.White(), " * ", Colors.Red(), "]")
	fmt.Println(Colors.Red(), "\t[", Colors.White(), " * ", Colors.Red(), "]", Colors.White(), "  Cracking , Scanning , Information Collector ", Colors.Red(), "\t[", Colors.White(), " * ", Colors.Red(), "]")
	fmt.Println(Colors.Red(), "\t[", Colors.White(), " * ", Colors.Red(), "]", Colors.White(), "\t  https://github.com/muhammadqazi", Colors.Red(), "\t\t[", Colors.White(), " * ", Colors.Red(), "]")
	fmt.Println(Colors.Red(), "\t[", Colors.White(), " * ", Colors.Red(), "]", Colors.White(), "\t\t  By Muhammad Qazi", Colors.Red(), "\t\t\t[", Colors.White(), " * ", Colors.Red(), "]")

	fmt.Println(Colors.Red(), "\n\n--------------------------------------------------------------------------------------------------")
	fmt.Println(Colors.Red(), "\nCracking Modules:")

	fmt.Println(Colors.Red(), "\n\n["+Colors.White()+"1"+Colors.Red()+"]", Colors.White(), " Zip Cracker",
		Colors.Red(), "\t\t["+Colors.White()+"2"+Colors.Red()+"]", Colors.White(), " TCP Scanner",
		Colors.Red(), "\t\t["+Colors.White()+"3"+Colors.Red()+"]", Colors.White(), " SSH Cracker")

	fmt.Println(Colors.Red(), "\n\n["+Colors.White()+"4"+Colors.Red()+"]", Colors.White(), " Net Scan")

	fmt.Println(Colors.Red(), "\n\n["+Colors.White()+"0"+Colors.Red()+"]", " Exit")

	fmt.Print(Colors.Red(), "\n\n┌─[ "+Colors.Green()+"Mr.Hack3r"+Colors.Blue()+"~"+Colors.White()+"@HOME"+Colors.Red()+" ]"+"\n└──╼ "+Colors.White())

	var validated bool = false

	for !validated {

		var input int
		fmt.Scanf("%d", &input)

		switch input {

		case 1:

			helpers.ClearScreen()

			var (
				zipFile     string
				dictonary   string
				concurrency int
			)

			fmt.Println(Colors.Red(), "\n\nEnter the name of the zip file: ", Colors.White())
			fmt.Scanf("%s", &zipFile)

			fmt.Println(Colors.Red(), "\nEnter the name of the dictonary file: ", Colors.White())
			fmt.Scanf("%s", &dictonary)

			fmt.Println(Colors.Red(), "\nEnter the number of concurrency (default 8)", Colors.White())
			fmt.Scanf("%d", &concurrency)

			if zipFile == "" || dictonary == "" || concurrency == 0 {

				fmt.Println("Failed: Missing file name")

				validated = true

			} else {

				controllers.ZipCracker(zipFile, dictonary, concurrency)

				validated = true
			}

		case 2:
			helpers.ClearScreen()

			fmt.Println(Colors.Red(), "\n\nEnter the target IP: ", Colors.White())
			var target string
			fmt.Scanf("%s", &target)

			if target == "" {

				fmt.Println("Failed: Missing Target IP")

				validated = true

			} else {

				controllers.TCPScanner(target)

				validated = true
			}

		case 3:
			helpers.ClearScreen()

			var (
				host string
				port string
				user string
				pass string
			)

			fmt.Println(Colors.Red(), "\n\nEnter the Host: ", Colors.White())
			fmt.Scanf("%s", &host)

			fmt.Println(Colors.Red(), "\nEnter the Port: ", Colors.White())
			fmt.Scanf("%s", &port)

			fmt.Println(Colors.Red(), "\nEnter the User: ", Colors.White())
			fmt.Scanf("%s", &user)

			fmt.Println(Colors.Red(), "\nEnter the location for password dictonary: ", Colors.White())
			fmt.Scanf("%s", &pass)

			if host == "" || port == "" || user == "" || pass == "" {

				fmt.Println("Failed: Missing Information")

				validated = true

			} else {

				controllers.SSHCracker(host, port, user, pass)

				validated = true
			}

		case 4:
			helpers.ClearScreen()

			fmt.Println(Colors.Red(), "\n\nEnter the First three octets. Ex: 192.168.0", Colors.White())

			var ip string
			fmt.Scanf("%s", &ip)

			if ip == "" {

				fmt.Println("Failed: Missing IP")

				validated = true

			} else {

				controllers.NetScan(ip)

				validated = true
			}

		case 0:
			fmt.Println("Mr.Hack3r is out")
			os.Exit(0)

		default:
			fmt.Println("Invalid Choice ! Try Again")
			fmt.Print(Colors.Red(), "\n\n┌─[ "+Colors.Green()+"Mr.Hack3r"+Colors.Blue()+"~"+Colors.White()+"@HOME"+Colors.Red()+" ]"+"\n└──╼ "+Colors.White())
			validated = false
		}

	}

}
