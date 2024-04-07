// api.go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// createAd handles POST requests to create a new ad.
func createAd(c *gin.Context) {

	var newAd Ad
	if err := c.ShouldBindJSON(&newAd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Assign an ID to the new ad.
	newAdID, err := rdb.Incr(context.Background(), "ad:id").Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate ad ID"})
		return
	}
	newAd.ID = int(newAdID)

	// Serialize the ad to JSON for storage.
	adJSON, err := json.Marshal(newAd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to serialize ad"})
		return
	}

	// Store the ad in Redis.
	if err := rdb.Set(context.Background(), fmt.Sprintf("ad:%d", newAd.ID), adJSON, 0).Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store ad in Redis"})
		return
	}

	c.JSON(http.StatusCreated, newAd)
}

// filterAds returns active ads.
func filterAds(c *gin.Context) {
	keys, err := rdb.Keys(context.Background(), "ad:*").Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch ad keys from Redis"})
		return
	}

	var activeAds []Ad
	for _, key := range keys {
		adJSON, err := rdb.Get(context.Background(), key).Result()
		if err != nil {
			// Log error, skip this ad.
			continue
		}

		var ad Ad
		if err := json.Unmarshal([]byte(adJSON), &ad); err != nil {
			// Log error, skip this ad.
			continue
		}

		// Check if the ad is active.
		if time.Now().After(ad.StartAt) && time.Now().Before(ad.EndAt) {
			activeAds = append(activeAds, ad)
		}
	}

	c.JSON(http.StatusOK, gin.H{"active_ads": activeAds})
}
