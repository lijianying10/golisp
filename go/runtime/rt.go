package runtime

import (
	"bufio"
	"fmt"
	"os"
)

func READ() string{
	fmt.Print("golisp> ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return text
}

func EVAL(source string) string{
	return source
}

func PRINT(source string){
	fmt.Print(source)
}

func GOLISP(){
	for {
		PRINT(EVAL(READ()))
	}
}
