package main

// #include <unistd.h>
import "C"

func main() {
	println(C.sysconf(C._SC_PHYS_PAGES)*C.sysconf(C._SC_PAGE_SIZE), " bytes")
}
