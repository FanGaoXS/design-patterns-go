package main

import "fmt"

func main() {
	appleBuilder, _ := GetPhoneBuilder("Apple")
	appleBuilder.BuildCPU("Apple A15")
	appleBuilder.BuildScreen("samsung")
	appleBuilder.BuildBattery("宁德时代")
	apple := appleBuilder.GetPhone()
	fmt.Printf("apple = %#v\n", apple)

	huaweiBuilder, _ := GetPhoneBuilder("HUAWEI")
	huaweiBuilder.BuildCPU("麒麟9800")
	huaweiBuilder.BuildScreen("京东方")
	huaweiBuilder.BuildBattery("宁德时代")
	huawei := huaweiBuilder.GetPhone()
	fmt.Printf("huawei = %#v\n", huawei)
}
