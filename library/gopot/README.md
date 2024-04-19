## GOPOT

该库主要实现了对于`Pot Server`的调用封装，使用该库可以方便的对于远端`Pot`进行调用。

### Usage

```
package main

import (
	"fmt"
	"time"

	"github.com/swxctx/pot/library/gopot"
)

func main() {
	// 初始化客户端
	potClient, err := gopot.NewClient(&gopot.Config{
		Address: "localhost:9577",
	})
	if err != nil {
		fmt.Printf("err-> %v", err)
		return
	}
	time.Sleep(time.Duration(1) * time.Second)

	key := "pot:client:key"
	value := "Hello Pot!"
	fmt.Printf("key-> %s, value-> %s\n", key, value)

	// set
	setCmd := potClient.Set(key, value, time.Duration(20)*time.Second)
	if setCmd.Err() != nil {
		fmt.Printf("set err-> %v", err)
		return
	}
	fmt.Println("set cache success...")

	// get
	getCmd := potClient.Get(key)
	if getCmd.Err() != nil {
		fmt.Printf("get err-> %v", err)
		return
	}
	fmt.Printf("get cache, key-> %s, value-> %s\n", key, getCmd.String())

	// exists cmd
	existsCmd := potClient.Exists(key)
	if existsCmd.Err() != nil {
		fmt.Printf("exists err-> %v", err)
		return
	}
	if existsCmd.Result() != gopot.EXPIRATION_IS_EXPIRED {
		fmt.Sprintf("key: %s exists\n", key)
	}

	// ttl
	ttlCmd := potClient.TTL(key)
	if ttlCmd.Err() != nil {
		fmt.Printf("ttl err-> %v", err)
		return
	}
	fmt.Printf("key: %s, ttl-> %d\n", key, ttlCmd.Result())

	// expire
	expireCmd := potClient.Expire(key, time.Duration(5)*time.Second)
	if expireCmd.Err() != nil {
		fmt.Printf("expire err-> %v", err)
		return
	}
	fmt.Println("set cache expire success...")

	time.Sleep(time.Duration(10) * time.Second)

	// exists cmd
	e := potClient.Exists(key)
	if e.Result() < 1 {
		fmt.Printf("key: %s not exists\n", key)
	}

	// set
	s := potClient.Set(key, value, time.Duration(20)*time.Second)
	fmt.Printf("again set key-> %s, success-> %v\n", key, s.Success())

	// exists
	if potClient.Exists(key).Result() > 0 {
		fmt.Printf("key: %s exists\n", key)
	}

	// del
	potClient.Del(key)

	e1 := potClient.Exists(key)
	if e1.Result() < 1 {
		fmt.Printf("key: %s not exists\n", key)
	}
}
```