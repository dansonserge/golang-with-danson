package main

import (
	"fmt"
	"sync"
	"time"
)

//.6 declare mutex
//var m = sync.Mutex{}
var m = sync.RWMutex{}

//1. declare the wait group
var wg = sync.WaitGroup{}


var dbData = []string{"id1","id2","id3","id4","id5"}
var results = []string{}
func dbCall(i int){
	//var delay float32 = rand.Float32()*2000
	var delay float32 = 2000
	time.Sleep(time.Duration(delay)*time.Microsecond)
	fmt.Println("the result from DB is: ", dbData[i])

	m.Lock()
	//results = append(results, dbData[i])
	save(dbData[i])
	log()
	m.Unlock()

	//3. remove the execution count
	wg.Done()
}
func main() {

	t0 := time.Now()

	for i:=0; i<len(dbData); i++{
		//2. add an execution count
		wg.Add(1)
		go dbCall(i)
	}

	//4. wait until all the executions are count to 0, the continue
	wg.Wait()
	fmt.Printf("\n Total execution time: %v", time.Since(t0))
	fmt.Printf("\n Results: %v", results)

	//.5 we should use mutex to control writing in our slice in a sage way when using concurency

}

func save(result string){
	m.Lock()
	results= append(results, result)
	m.Unlock()
}

func log(){
	m.RLock()
	fmt.Println("\nThe current results are: ", results)
	m.RUnlock()
}