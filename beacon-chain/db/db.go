// Package db defines the ability to create a new database
// for an eth2 beacon node.
package db

import (
	"context"

	"github.com/prysmaticlabs/prysm/beacon-chain/cache"
	"github.com/prysmaticlabs/prysm/beacon-chain/db/kv"
)

// NewDB initializes a new DB.
func NewDB(ctx context.Context, dirPath string, stateSummaryCache *cache.StateSummaryCache) (Database, error) {
	return kv.NewKVStore(ctx, dirPath, stateSummaryCache)
}
