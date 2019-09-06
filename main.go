package main

func main() {
    outChan := make(chan int64)
    timer := erpTimer(outChan)
    writer := writer(outChan)
    go timer.run()
    go writer.run()
    for {
    }
}