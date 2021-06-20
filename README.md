# chan-map
A safe map use channel, very simple .

# useage 
import this package
```
import "github.com/caoyuewen/chan-map"
```

```
go get "github.com/caoyuewen/chan-map"
```

# example 
```

package main

import (
	"fmt"
	"github.com/caoyuewen/chan-map"
)

func main() {

	// create a map
	cMap := chan_map.NewChanMap()

	// set string type key
	cMap.Set("name", "kerry")
	// set int type key
	cMap.Set(8, "anything")

	// get value
	v1 := cMap.Get("name")
	v2 := cMap.Get(8)
	fmt.Println(v1.(string), v2.(string))

	// delete key
	cMap.Delete(8)

	// show the map current size
	fmt.Println(cMap.Size())

	fn := func(k, v interface{}) bool {
		// just show
		fmt.Println(k, v)
		//switch k.(type) {
		//case string:
		//	//...
		//	return true
		//case int:
		//	//...
		//	return false
		//}
		return true
	}
	cMap.Range(fn)
}
```
