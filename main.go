package main

import (
	"flag"
	"fmt"
	"log"
	"net/smtp"
	"os"
	"os/exec"
	"strings"
)

func main() {

	var repository string
	flag.StringVar(&repository, "repository", "", "repository to be cloned")

	var branch string
	flag.StringVar(&branch, "branch", "master", "branch to be cloned")

	var browser string
	flag.StringVar(&browser, "browser", "electron", "branch to be cloned")

	flag.Parse()

	if repository == "" {
		log.Fatalln("repository flag is not given")
	}

	sl := strings.Split(repository, "/")
	name := sl[len(sl)-1]
	name = strings.ReplaceAll(name, "\"", "")

	err := cloneRepository(repository, branch)

	if err != nil {
		log.Fatal(err)
	}

	err = runCypress(name, browser)

	var testsPassed bool
	if err != nil {
		log.Print("tests failed")
		log.Print(err)
		testsPassed = false
	} else {
		log.Print("tests passed")
		testsPassed = true
	}

	err = deleteRepository(name)
	if err != nil {
		log.Fatal(err)
	}

	if !testsPassed {
		os.Exit(1)
	}
}

func cloneRepository(repo string, branch string) error {
	log.Printf("going to clone the repository")
	out, err := exec.Command("sh", "-c", "git clone "+repo+" -b "+branch).CombinedOutput()
	if out != nil {
		log.Print(string(out))
	}
	return err
}

func npmInstall(name string) error {
	log.Printf("going to npm i")
	cmd := exec.Command("sh", "-c", "cd "+name+" && npm i --save-dev")
	log.Print("cmd: " + cmd.String())
	out, err := cmd.CombinedOutput()
	if out != nil {
		log.Print(string(out))
	}
	return err
}

func runCypress(name string, browser string) error {
	log.Printf("going to run cypress")
	cmd := exec.Command("sh", "-c", "cd "+name+" && npx cypress run --browser "+browser)
	log.Print("cmd: " + cmd.String())
	out, err := cmd.CombinedOutput()
	if out != nil {
		log.Print(string(out))
	}
	if err != nil {
		sendMail(string(out))
	}
	return err
}

func deleteRepository(name string) error {
	log.Printf("going to delete repo")
	cmd := exec.Command("rm", "-rf", name)
	log.Print("cmd: " + cmd.String())
	out, err := cmd.CombinedOutput()
	if out != nil {
		log.Print(string(out))
	}
	return err
}

func sendMail(logs string) {
	// Sender data.
	from := ""
	password := ""

	// Receiver email address.
	to := []string{}

	// smtp server configuration.
	smtpHost := ""
	smtpPort := ""

	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Message.
	message := []byte("Tests failed see this for more infos \n\n " + logs)

	// Authentication.

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}
