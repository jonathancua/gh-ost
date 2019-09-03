package logic

import (
    "bufio"
    "fmt"
    "os"
)

func pause() {
    buf := bufio.NewReader(os.Stdin)
    fmt.Println("\n...pausing... press any key to continue\n")
    buf.ReadBytes('\n')
}
