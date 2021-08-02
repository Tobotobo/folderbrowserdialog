# folderbrowserdialog
golang folderbrowserdialog

``` go
package main

import (
	"fmt"

	"github.com/Tobotobo/folderbrowserdialog"
)

func main() {
	if ok, selectedPath := folderbrowserdialog.Show(); ok {
		fmt.Println(selectedPath)
	} else {
		fmt.Println("Cancel")
	}
}
```
![image](https://user-images.githubusercontent.com/46508541/127917084-5ce39f24-1c48-4ced-a64b-c1d0856895f3.png)

