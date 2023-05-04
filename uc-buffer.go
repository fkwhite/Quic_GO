package quic

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"sync"
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

	for key1, map_ := range globalBuffers.registerOut[streamIdx] {
		for key2, element := range map_ {
			logMessage := fmt.Sprint(key1, " ", key2, " ", element, "\n")
			//fmt.Println(logMessage) //por pantalla
			_, err = f.WriteString(logMessage) //fichero log
			if err != nil {
				fmt.Println("An error has ocurred -- Writing file")
				panic(err)
			}
		}

	}

}

// delta can be positive or negative
func GlobalBuffersIncr(streamIdx int, delta float64) {
	globalBuffers.mtxs[streamIdx].Lock()
	defer globalBuffers.mtxs[streamIdx].Unlock()

	globalBuffers.buffers[streamIdx] += delta
	//fmt.Println("El tamaño del buffer es: ", globalBuffers.buffers[streamIdx])
	//fmt.Println("Lenght ", len(globalBuffers.registerIn[streamIdx]))
	if delta != 0 {
		if math.Signbit(delta) { // Delta negativo -> true
			var aux = math.Abs(delta)
			for key, element := range globalBuffers.registerIn[streamIdx] {
				timeOut := time.Now().UnixMicro()
				if element > aux {
					globalBuffers.registerIn[streamIdx][key] = element - aux
					globalBuffers.registerOut[streamIdx][timeOut] = make(map[int64]float64)
					globalBuffers.registerOut[streamIdx][timeOut][key] = aux

					break
				} else if element == aux {
					globalBuffers.registerOut[streamIdx][timeOut] = make(map[int64]float64)
					globalBuffers.registerOut[streamIdx][timeOut][key] = aux
					delete(globalBuffers.registerIn[streamIdx], key)
					break
				} else {
					aux = aux - element
					globalBuffers.registerOut[streamIdx][timeOut] = make(map[int64]float64)
					globalBuffers.registerOut[streamIdx][timeOut][key] = element
					delete(globalBuffers.registerIn[streamIdx], key)
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

	globalBuffers.mtxs[streamIdx].Lock()
	defer globalBuffers.mtxs[streamIdx].Unlock()
	retrun globalBuffers.registerIn[streamIdx]
		
	
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
