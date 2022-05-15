package test

import "time"

func PreCreateOrder() {
	time.Sleep(time.Duration(1) * time.Second)
}

func CoreCreateOrder() {

}

func AfterCreateOrder() {

}
