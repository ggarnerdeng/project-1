package main

import (
	"fmt"
	"net/http"

	//"text/html"

	"github.com/sfreiberg/simplessh"
)

func main() { /*
		var hostname, userName, password string
		fmt.Printf("Enter a hostname(IP):  ")
		fmt.Scanln(&hostname)
		fmt.Printf("Enter a username:  ")
		fmt.Scanln(&userName)
		fmt.Printf("Enter a password:  ")
		fmt.Scanln(&password)*/
	hostname := "192.168.1.33"
	userName := "garner"
	password := "password"

	//PSTREE: view all processes as a tree
	http.HandleFunc("/top", func(a http.ResponseWriter, b *http.Request) {
		var top = terminalCommand(hostname, userName, password, "top -bn1")
		fmt.Fprintf(a, "%s\n", "Viewing all processes")
		fmt.Fprintln(a, " ")
		fmt.Fprintln(a, " ")
		fmt.Fprintf(a, "%s", top)

	})

	//TREE: view all files as a tree
	//a = response
	//b = request
	http.HandleFunc("/tree", func(a http.ResponseWriter, b *http.Request) {
		var tree = terminalCommand(hostname, userName, password, "tree")
		fmt.Fprintf(a, "%s\n", "Viewing all files")
		fmt.Fprintln(a, " ")
		fmt.Fprintln(a, " ")
		fmt.Fprintf(a, "%s", tree)
	})
	//HISTORY: view all history
	http.HandleFunc("/history", func(a http.ResponseWriter, b *http.Request) {
		var history = terminalCommand(hostname, userName, password, "cat ~/.bash_history | nl")
		fmt.Fprintf(a, "%s\n", "Viewing user command history")
		fmt.Fprintln(a, " ")
		fmt.Fprintln(a, " ")
		fmt.Fprintf(a, "%s", history)
	})
	//FINGER: list user information
	http.HandleFunc("/finger", func(a http.ResponseWriter, b *http.Request) {
		var finger = terminalCommand(hostname, userName, password, "finger")
		fmt.Fprintf(a, "%s\n", "Viewing user information")
		fmt.Fprintln(a, " ")
		fmt.Fprintln(a, " ")
		fmt.Fprintf(a, "%s", finger)
	})
	//SYSLOG
	http.HandleFunc("/syslog", func(a http.ResponseWriter, b *http.Request) {
		var syslog = terminalCommand(hostname, userName, password, "cat /var/log/syslog")
		fmt.Fprintf(a, "%s\n", "Viewing system log")
		fmt.Fprintln(a, " ")
		fmt.Fprintln(a, " ")
		fmt.Fprintf(a, "%s", syslog)
	})
	//AUTHLOG
	http.HandleFunc("/authlog", func(a http.ResponseWriter, b *http.Request) {
		var authlog = terminalCommand(hostname, userName, password, "cat /var/log/auth.log")
		fmt.Fprintf(a, "%s\n", "Viewing login records")
		fmt.Fprintln(a, " ")
		fmt.Fprintln(a, " ")
		fmt.Fprintf(a, "%s", authlog)
	})

	http.HandleFunc("/", func(a http.ResponseWriter, b *http.Request) {
		var output = terminalCommand(hostname, userName, password, "ps")
		var html = `<html>
		<head>
		
		<style>
		ul {
		list-style-type: none;
		margin: 0;
		padding: 0;
		overflow: hidden;
		background-color: #333333;
	  }
	  
	  li {
		float: center;
	  }
	  
	  li a {
		display: block;
		color: red;
		text-align: left;
		padding: 16px;
		text-decoration: none;
	  }
	  
	  li a:hover {
		background-color: #111111;
	  }
		* {
		 box-sizing: border-box; 
		}
		
		body {
		  margin: 0;
		}
		#main {
		  display: flex;
		  min-height: calc(100vh - 40vh);
		}
		#main > article {
		  flex: 1;
		}
		
		#main > nav, 
		#main > aside {
		  flex: 0 0 20vw;
		  background: beige;
		}
		#main > nav {
		  order: -1;
		}
		header, footer, article, nav, aside {
		  padding: 1em;
		}
		header, footer {
		  background: yellowgreen;
		  height: 20vh;
		}
	  </style>
			</head>
	  <body>
		<header>Logged in as garner@192.168.1.33</header>
		<div id="main">
		  <article>Command Line Options
		  <ol>
		 
			  <li><a href="/top">View all processes</a></li>
			  <li><a href="/tree">View all files</a></li>
			  <li><a href="/history">View all history</a></li>
			  <li><a href="/finger">View user information</a></li>
			  <li><a href="/syslog">View System Log</a></li>
			  <li><a href="/authlog">View login history</a></li>
		  </ol></article>
		  <nav></nav>
		  <aside></aside>
		</div>
		<footer></footer>
	  </body>
					</html>
			`
		fmt.Fprintf(a, html, output)
	})
	fmt.Println()
	fmt.Println()
	fmt.Println("Open Localhost:12345")
	http.ListenAndServe(":12345", nil)
}

func terminalCommand(hostname string, userName string, password string, command string) []byte {
	var client *simplessh.Client
	var err error
	if client, err = simplessh.ConnectWithPassword(hostname, userName, password); err != nil {
		fmt.Print(err)
	}
	if err != nil {
		panic(err)
	}
	defer client.Close()
	output, err := client.Exec(command)
	if err != nil {
		panic(err)
	}
	return output
}
