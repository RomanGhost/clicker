package clickwebsocket

import (
	"crypto/sha256"
	"fmt"
	"log"
	"sort"
	"time"
)

func clicksPerSecondValid(batch *ClickBatch) bool {
	second := 0
	countClick := 0
	for _, click := range batch.ClicksInfo {
		if second != time.Unix(click.ClickTime, 0).Second() {
			if countClick > 14 {
				return false
			}
			countClick = 0
		} else if second == time.Unix(click.ClickTime, 0).Second() {
			countClick++
		}
	}
	return true
}

func autoClickValid(batch *ClickBatch) bool {
	var diff []int64
	var sum int64
	for i := 1; i < len(batch.ClicksInfo); i++ {
		resDiff := batch.ClicksInfo[i].ClickTime - batch.ClicksInfo[i-1].ClickTime
		diff = append(diff, resDiff)
		sum += resDiff
	}
	mean := float64(float64(sum) / float64(len(batch.ClicksInfo)))

	var median float64
	sort.Slice(diff, func(i, j int) bool {
		return i > j
	})
	mNumber := len(diff) / 2

	if len(diff)%2 == 0 {
		median = float64(diff[mNumber])
	} else {
		median = float64(diff[mNumber-1]+diff[mNumber]) / 2.0
	}

	return mean != median
}

func ValidateBatch(batch *ClickBatch) float64 {
	batchTime := time.Unix(batch.SendTime, 0)
	serverTime := time.Now().UTC()
	differenceTime := serverTime.Sub(batchTime)

	if differenceTime.Abs() > time.Second*45 {
		log.Println("the batchtime arrived is too late - reject")
		return 0
	}

	cps := clicksPerSecondValid(batch)
	if !cps {
		log.Println("the message batch clicks is too fast")
		return 0
	}
	ac := autoClickValid(batch)
	if !ac {
		log.Println("the message batch clicks have same timing")
		return 0
	}

	var resultClicks float64
	for _, click := range batch.ClicksInfo {
		resultClicks += click.ClickValue

	}
	return resultClicks
}

func ValidateMessageValid(message Validate, userLogin string) error {
	valid := message.Valid
	nonce := message.Nonce
	// message format: "login_valid_nonce"
	sum := sha256.Sum256([]byte(fmt.Sprintf("%v_%v_%v", userLogin, valid, nonce)))
	log.Printf("Res of sum: %x\n", sum)
	if sum[0] != 0 && sum[1] < 128 {
		return fmt.Errorf("sha256 sum is not valid")
	}

	return nil
}
