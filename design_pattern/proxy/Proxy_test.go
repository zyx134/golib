package proxy

import (
	"fmt"
	"testing"
)

func TestProxy(t *testing.T) {
	var d Device
	d = &HardDiskProxy{
		OpId: "admin",
		hd:   &HardDisk{},
	}
	d.Write([]byte("hello"))
	data, _ := d.Read()
	fmt.Println(string(data))
}

func TestProxy1(t *testing.T) {
	var d Device
	d = &HardDiskProxy{
		OpId: "reader",
		hd: &HardDisk{
			storage: []byte("123"),
		},
	}
	err := d.Write([]byte("hello"))
	if err == nil {
		data, _ := d.Read()
		fmt.Println(string(data))
	} else {
		fmt.Println(err)
	}
}
func TestProxy2(t *testing.T) {
	var d Device
	d = &HardDiskProxy{
		OpId: "writer",
		hd: &HardDisk{
			storage: []byte("123"),
		},
	}
	err := d.Write([]byte("hello"))
	if err == nil {
		data, err := d.Read()
		if err == nil {
			fmt.Println(string(data))
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}
}
