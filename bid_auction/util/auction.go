package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"greedyGame/bid_auction/model"
	"io/ioutil"
	"net/http"
	"sort"
	"sync"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

//ValidateGetBid ...
func ValidateGetBid(c *gin.Context) (string, error) {
	var (
		err           error
		reqPayload    []byte
		adRequestBody model.GetBidStruct
		adSlotID      string
		isExists      bool
	)
	if reqPayload, err = ioutil.ReadAll(c.Request.Body); err == nil {
		if err = json.Unmarshal(reqPayload, &adRequestBody); err == nil {
			if _, err = govalidator.ValidateStruct(adRequestBody); err == nil {
				if isExists, err = isSlotExists(adRequestBody.AdPlacementID); err == nil && isExists {
					adSlotID = adRequestBody.AdPlacementID
				} else {
					if isExists {
						err = errors.New("isSlotExists: " + err.Error())
					} else {
						err = errors.New("ad slot does not exist")
					}
				}
			}
		}
	}
	return adSlotID, err
}

//GetBid ...
func GetBid(adSlotID string) (int, error) {
	var (
		err        error
		highestBid int
	)
	bids := getBids(adSlotID)
	if len(bids) > 0 {
		sort.Ints(bids)
		highestBid = bids[len(bids)-1]
		if highestBid == 0 {
			err = errors.New(" no valid bid")
		}
		fmt.Println(err, highestBid)
	} else {
		fmt.Println("boom")
	}
	return highestBid, err
}

func getBids(adSlotID string) []int {
	var wg sync.WaitGroup
	bids := make([]int, 0)

	//bidding services
	//these are url of multiple bidding service
	//we will store
	var urls = []string{
		`http://localhost:8090/bid/v1/getbidbyadplacementid?adplacementid=`,
		`http://localhost:8090/bid/v1/getbidbyadplacementid1?adplacementid=`,
	}
	for _, url := range urls {
		url = url + adSlotID
		wg.Add(1)
		// Launch a goroutine to fetch the URL.
		go func(url string) {
			// Fetch the URL.
			timeout := time.Duration(200 * time.Millisecond)
			c := http.Client{
				Timeout: timeout,
			}
			resp, err := c.Get(url)
			if err != nil {
				fmt.Printf("got an error")
			} else {
				defer resp.Body.Close()
				body, err := ioutil.ReadAll(resp.Body)
				if err == nil {
					var resp model.BidResp
					if err := json.Unmarshal(body, &resp); err == nil {
						bids = append(bids, resp.Price)
					} else {
						fmt.Println(err, "------")
					}
				} else {
					fmt.Println("boom")
				}
			}
			defer wg.Done()
		}(url)
	}
	wg.Wait()
	fmt.Println(bids)
	return bids
}
