package gpio

import (
	"fmt"
	"os"
	"reflect"
	"syscall"
	"unsafe"
)

var (
	base    int64
	GPIO_DR []uint32
)

func main() {
	f, err := os.OpenFile("/dev/mem", os.O_RDWR, 0)
	if err != nil {
		fmt.Println("Open Mem Error")
		return
	}
	defer f.Close()

	base = (0xFF7F0000 & 0xFFFFF000)
	p, err := syscall.Mmap(int(f.Fd()), base, 4096, syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
	if err != nil {
		fmt.Println("mmap Error")
		return
	}
	defer syscall.Munmap(p)

	s := *(*reflect.SliceHeader)(unsafe.Pointer(&p))
	s.Len /= 4
	s.Cap = 4
	GPIO_DR = *(*[]uint32)(unsafe.Pointer(&s))
	GPIO_DR[0] |= (0x1 << 1)
	fmt.Println(GPIO_DR)
}
