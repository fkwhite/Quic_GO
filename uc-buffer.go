package quic

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"sync"
	"sort"
	"time"
)

type StreamBuffer struct {
	mtxs        []sync.Mutex
	buffers     []float64
	registerIn  []map[int64]float64           // maybe hacer una struct de esto
	registerOut []map[int64]map[int64]float64 // para no repetir esto

	// len int
}

var globalBuffers StreamBuffer

func GlobalBuffersInit(numStreams int) {
	globalBuffers.mtxs = make([]sync.Mutex, numStreams)
	globalBuffers.buffers = make([]float64, numStreams)
	globalBuffers.registerIn = make([]map[int64]float64, numStreams)
	globalBuffers.registerOut = make([]map[int64]map[int64]float64, numStreams)
	// globalBuffers.len = numStreams
}

func GlobalBuffersLog(streamIdx int) {
	globalBuffers.mtxs[streamIdx].Lock()
	defer globalBuffers.mtxs[streamIdx].Unlock()
	configFile := "conf_Scheduler.json"
	file, err := os.Open(configFile)
	if err != nil {
		fmt.Println("An error has ocurred -- Opening configuration file")
		panic(err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err = decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}

	nameFile := fmt.Sprint("tmp/", configuration.Scheduler_name, "logSchedulerDelay", "_", streamIdx, ".log") //logSchedulerXX.log
	f, err := os.OpenFile(nameFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)                                  //os.O_TRUNC
	if err != nil {
		fmt.Println("An error has ocurred -- Opening file")
		panic(err)
	}
	defer f.Close()

	keys := make([]float64, 0)
	for k, _ := range globalBuffers.registerOut[streamIdx] {
		keys = append(keys, float64(k))
	}
	// fmt.Println("+++++++++ Timestamps ordered +++++++++++")
	sort.Float64s((keys))

	for _, key1 :=  range keys {
		for key2, element := range globalBuffers.registerOut[streamIdx][int64(key1)] {
			logMessage := fmt.Sprint(key1, " ", key2, " ", element, "\n")
			//fmt.Println(logMessage) //por pantalla
			_, err = f.WriteString(logMessage) //fichero log
			if err != nil {
				fmt.Println("An error has ocurred -- Writing file")
				panic(err)
			}
		}
	}
	// for key1, map_ := range globalBuffers.registerOut[streamIdx] {
	// 	for key2, element := range map_ {
	// 		logMessage := fmt.Sprint(key1, " ", key2, " ", element, "\n")
	// 		//fmt.Println(logMessage) //por pantalla
	// 		_, err = f.WriteString(logMessage) //fichero log
	// 		if err != nil {
	// 			fmt.Println("An error has ocurred -- Writing file")
	// 			panic(err)
	// 		}
	// 	}

	// }

}

// delta can be positive or negative
func GlobalBuffersIncr(streamIdx int, delta float64) {
	globalBuffers.mtxs[streamIdx].Lock()
	defer globalBuffers.mtxs[streamIdx].Unlock()
	keys := make([]float64, 0)
	globalBuffers.buffers[streamIdx] += delta
	//fmt.Println("El tamaño del buffer es: ", globalBuffers.buffers[streamIdx])
	//fmt.Println("Lenght ", len(globalBuffers.registerIn[streamIdx]))
	if delta != 0 {
		if math.Signbit(delta) { // Delta negativo -> true
			for k, _ := range globalBuffers.registerIn[streamIdx] {
				keys = append(keys, float64(k))
			}
			// fmt.Println("+++++++++ Timestamps ordered +++++++++++")
			sort.Float64s((keys))
			// timestamp := time.Now().UnixMicro()
			// fmt.Printf("New try - Stream %d\n",streamIdx)
			// for _, k := range keys {
			// 	fmt.Printf("Pkt. delay time %d - Pkt. size %d\n", int64(k),globalBuffers.registerIn[streamIdx][int64(k)])
			// }

			var aux = math.Abs(delta)
			for _, key :=  range keys {
				element := globalBuffers.registerIn[streamIdx][int64(key)]
				// fmt.Printf("Pkt. delay %d - Pkt. size %f\n",key,element)
				timeOut := time.Now().UnixMicro()
				if element > aux {
					globalBuffers.registerIn[streamIdx][int64(key)] = element - aux
					globalBuffers.registerOut[streamIdx][timeOut] = make(map[int64]float64)
					globalBuffers.registerOut[streamIdx][timeOut][int64(key)] = aux
					break
				} else if element == aux {
					globalBuffers.registerOut[streamIdx][timeOut] = make(map[int64]float64)
					globalBuffers.registerOut[streamIdx][timeOut][int64(key)] = aux
					delete(globalBuffers.registerIn[streamIdx], int64(key))
					break
				} else {
					aux = aux - element
					globalBuffers.registerOut[streamIdx][timeOut] = make(map[int64]float64)
					globalBuffers.registerOut[streamIdx][timeOut][int64(key)] = element
					delete(globalBuffers.registerIn[streamIdx], int64(key))
				}
			}

		} else {
			timeIn := time.Now().UnixMicro()
			globalBuffers.registerIn[streamIdx][timeIn] = delta

		}

	}

}


func GlobalBuffersRead(streamIdx int) float64 {
	// if (streamIdx >= globalBuffers.len){
	// 	fmt.Printf("Error in stream idx %d, max len %d \n", streamIdx, globalBuffers.len)
	// }
	globalBuffers.mtxs[streamIdx].Lock()
	defer globalBuffers.mtxs[streamIdx].Unlock()
	return globalBuffers.buffers[streamIdx]

}

func GlobalBuffersWrite(streamIdx int, newVal float64) {
	globalBuffers.mtxs[streamIdx].Lock()
	defer globalBuffers.mtxs[streamIdx].Unlock()
	globalBuffers.buffers[streamIdx] = newVal
	globalBuffers.registerIn[streamIdx] = make(map[int64]float64)
	globalBuffers.registerOut[streamIdx] = make(map[int64]map[int64]float64)
}

func GlobalBuffersSojournTimeLog(scheduler string, timestamp int64, id int, sum int64) {
	globalBuffers.mtxs[id].Lock()
	defer globalBuffers.mtxs[id].Unlock()
	nameFile := fmt.Sprint("tmp/logSchedulerSojournTime_", scheduler, ".log") //logSchedulerXX.log
	f, err := os.OpenFile(nameFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)  //os.O_TRUNC
	if err != nil {
		fmt.Println("An error has ocurred -- Opening file")
		panic(err)
	}
	defer f.Close()

	logMessage := fmt.Sprint(timestamp, "	", id, "	", sum, "\n")
	//fmt.Println(logMessage) //por pantalla
	_, err = f.WriteString(logMessage) //fichero log
	if err != nil {
		fmt.Println("An error has ocurred -- Writing file")
		panic(err)
	}

}

func GlobalBuffersPktDelay(streamIdx int) int64{
	var  min int64
	min = 999999999999999999 
	globalBuffers.mtxs[streamIdx].Lock()
	defer globalBuffers.mtxs[streamIdx].Unlock()

	// timestamp := time.Now().UnixMicro()
	// fmt.Println("+++++++++Timestamp++++++++")
	if(len(globalBuffers.registerIn[streamIdx]) == 0){
		// fmt.Println("--------Empty--------")
		min = 0
	}else{
		// for val, _ := range globalBuffers.registerIn[streamIdx] {
		// 	fmt.Printf("delay Pkt. %d  and buffer len %d\n",timestamp-val,len(globalBuffers.registerIn[streamIdx]))
			
		// }

		for val, _ := range globalBuffers.registerIn[streamIdx] {
			if (val < min) {
				// fmt.Printf("min delay Pkt. %d \n",val)
				min = val
			}
		}
	}
	return min
	
}


func GlobalBuffersPktDelayFirst(streamIdx int) int64{
	var  min int64
	min = 999999999999999999 
	globalBuffers.mtxs[streamIdx].Lock()
	defer globalBuffers.mtxs[streamIdx].Unlock()

	
	// fmt.Println("+++++++++Timestamp++++++++")
	if(len(globalBuffers.registerIn[streamIdx]) == 0){
		// fmt.Println("--------Empty--------")
		min = 0
	}else{
	

		for val, _ := range globalBuffers.registerIn[streamIdx] {
			min = val
			break
		}
	}
	return min
	
}

func GlobalBuffersTotalDelay(streamIdx int) int64 {
	var sum int64
	sum = 0
	globalBuffers.mtxs[streamIdx].Lock()
	defer globalBuffers.mtxs[streamIdx].Unlock()
	for val, _ := range globalBuffers.registerIn[streamIdx] {
		sum += val
	}
	return sum
}



func GlobalBuffersSojournTime(streamIdx int) (int64,int) {
	var sum int64
	sum = 0
	globalBuffers.mtxs[streamIdx].Lock()
	defer globalBuffers.mtxs[streamIdx].Unlock()
	num_data := len(globalBuffers.registerIn[streamIdx])
	if (num_data!=0){
		for val, _ := range globalBuffers.registerIn[streamIdx] {
			sum += val
		}
	}
	
	return sum,num_data
}