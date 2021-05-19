# cpubalancer

**cpubalancer** Позволяет автоматически изменять значение переменной типа int в зависимости от загрузки процессора в заданном интервале

## Install (with GOPATH set on your machine)
----------

* Get the `cpubalancer` package
```
go get github.com/rurick/balancer

```

##Usage
----------
```
package main

import (
	cpubalancer "github.com/rurick/balancer"
)

// init and run
	balancer := cpubalancer.New(10)
	go balancer.Run(ctx, 30, 70, 10)

// For using value
	val := balancer.Value()
```

