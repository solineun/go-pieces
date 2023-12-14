package main

func main() {
	
}

func turnOffSelectBranch(in, in2 chan int, done chan struct{}) {
	for {
		select {
		case v, ok := <- in:
			if !ok {
				in = nil
				continue
			}
			v += 1 //some work
		case v, ok := <- in2:
			if !ok {
				in2 = nil
				continue
			}
			v += 2 //some work
		case <- done:
			return
		}
	}
}