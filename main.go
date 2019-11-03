package main

import (
	"fmt"
	"net/http"

	"github.com/sfreiberg/simplessh"
)

func main() {
	var hostname, userName, password string
	fmt.Printf("Enter a hostname(IP):  ")
	fmt.Scanln(&hostname)
	fmt.Printf("Enter a username:  ")
	fmt.Scanln(&userName)
	fmt.Printf("Enter a password:  ")
	fmt.Scanln(&password)
	http.HandleFunc("/top", func(w http.ResponseWriter, r *http.Request) {
		var top = terminalCommand(hostname, userName, password, "top -bn1")
		fmt.Fprintf(w, "%s\n", "Ouput for top command")
		fmt.Fprintf(w, "%s", top)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var output = terminalCommand(hostname, userName, password, "ps")
		var html = `<html>
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
	  <body>
		<header>Logged in as garner@192.168.1.33</header>
		<div id="main">
		  <article>Command Line Options
		  <ol>
			  
			  <li><a href="/top">Ru top</a></li>
		  </ol></article>
		  <nav></nav>
		  <aside></aside>
		</div>
		<footer></footer>
	  </body>
					</html>
			`
		fmt.Fprintf(w, html, output)
	})
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
