package main

import (
	"fmt"
	"log"

	"goexcersise/award/batch"
	"goexcersise/award/conf"
	logger "goexcersise/award/log"
)

func main() {
	conf.Parse()
	logger.InitLogConf()
	batch.InitAwardPool()

	//for {
	//	go WinPrize("diuge")
	//	time.Sleep(time.Duration(200)*time.Millisecond)
	//}

	batch.Start()
}

func drawPrize(awardBatches []batch.AwardBatch) {
	awardBatch := batch.WinPrize("diuge")

	if awardBatch == nil {
		fmt.Println("很遗憾，您未能抽中奖品")
	}

	log.Println("恭喜您抽中:", awardBatch.GetName())

	//var mutex sync.Mutex
	//
	//defer mutex.Unlock()
	//
	//mutex.Lock()
	//
	//for i, award := range awardBatches {
	//	if strings.Compare(award.GetName(),awardBatch.GetName()) == 0 || awardBatches[i].totalBalance > 0 {
	//		awardBatches[i].totalBalance -= 1
	//		fmt.Println("恭喜您抽中奖品 : " + award.GetName())
	//	}
	//}

}
