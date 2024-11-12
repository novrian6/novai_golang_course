// login_test.go
package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/tebeka/selenium"
)

const (
	seleniumPath = "/path/to/selenium-server-standalone.jar"
	port         = 8080
	webDriverURL = "http://localhost:4444/wd/hub" // Default Selenium WebDriver URL
)

func TestLoginSuccess(t *testing.T) {
	// Start a Selenium WebDriver session with Chrome
	caps := selenium.Capabilities{"browserName": "chrome"}
	wd, err := selenium.NewRemote(caps, webDriverURL)
	if err != nil {
		t.Fatalf("Failed to open session: %s", err)
	}
	defer wd.Quit()

	// Open the login page
	if err := wd.Get("http://localhost:8080/login"); err != nil {
		t.Fatalf("Failed to load page: %s", err)
	}

	// Find the username, password fields, and login button
	usernameElem, err := wd.FindElement(selenium.ByCSSSelector, "input[name='username']")
	if err != nil {
		t.Fatalf("Failed to find username input: %s", err)
	}
	passwordElem, err := wd.FindElement(selenium.ByCSSSelector, "input[name='password']")
	if err != nil {
		t.Fatalf("Failed to find password input: %s", err)
	}
	loginButton, err := wd.FindElement(selenium.ByCSSSelector, "input[type='submit']")
	if err != nil {
		t.Fatalf("Failed to find submit button: %s", err)
	}

	// Perform login with correct credentials
	if err := usernameElem.SendKeys("user"); err != nil {
		t.Fatalf("Failed to enter username: %s", err)
	}
	if err := passwordElem.SendKeys("pass"); err != nil {
		t.Fatalf("Failed to enter password: %s", err)
	}
	if err := loginButton.Click(); err != nil {
		t.Fatalf("Failed to click login button: %s", err)
	}

	// Wait for the response and check the outcome
	time.Sleep(1 * time.Second) // Small delay for page load

	body, err := wd.FindElement(selenium.ByTagName, "body")
	if err != nil {
		t.Fatalf("Failed to find body: %s", err)
	}
	text, err := body.Text()
	if err != nil {
		t.Fatalf("Failed to retrieve body text: %s", err)
	}

	// Verify login success
	expected := "Login Successful"
	if text != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, text)
	} else {
		fmt.Println("TestLoginSuccess passed.")
	}
}
