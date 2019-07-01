package runtime

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
)

func READ() string{
	fmt.Print("golisp> ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	if len(text)==0{
		fmt.Println("READ EOF exit!")
		os.Exit(0)
	}
	return text
}

func EVAL(source string) string{
	return source
}

func PRINT(source string){
	fmt.Print(source)
}

func SIGNALHandler(){
	c := make(chan os.Signal, 1)

	// Passing no signals to Notify means that
	// all signals will be sent to the channel.
	signal.Notify(c)

	// Block until any signal is received.
	s := <-c
	fmt.Println("Got signal:", s)
	os.Exit(0)
}

func GOLISP(){
	go SIGNALHandler()
	for {
		PRINT(EVAL(READ()))
	}
}
