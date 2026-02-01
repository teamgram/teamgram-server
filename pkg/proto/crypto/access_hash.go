// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package crypto

import (
	"encoding/binary"
	"math/rand"
	"strconv"
	"time"

	"golang.org/x/crypto/blake2b"
)

// User represents a Telegram user with an ID, base hash, timestamp, and salt
type User struct {
	ID        int
	BaseHash  int64
	Timestamp int64
	Salt      int
}

// Database simulates a simple in-memory user database
var Database = make(map[int]User)

// GenerateBaseHash generates a base hash for a user and returns it as int64
func GenerateBaseHash(userID int, timestamp int64, salt int) int64 {
	// Combine userID, timestamp, and salt to create a unique string
	uniqueString := strconv.Itoa(userID) + strconv.FormatInt(timestamp, 10) + strconv.Itoa(salt)
	// Generate BLAKE2b-256 hash of the unique string
	hash := blake2b.Sum256([]byte(uniqueString))
	// Convert the first 8 bytes of the hash to int64
	return int64(binary.BigEndian.Uint64(hash[:8]))
}

// GenerateAccessHash generates a unique access hash for a user accessed by a specific visitor and returns it as int64
func GenerateAccessHash(user User, visitorID int) int64 {
	// Combine the user's base hash with the visitor's ID to create a unique string
	uniqueString := strconv.FormatInt(user.BaseHash, 10) + strconv.Itoa(visitorID)
	// Generate BLAKE2b-256 hash of the unique string
	hash := blake2b.Sum256([]byte(uniqueString))
	// Convert the first 8 bytes of the hash to int64
	return int64(binary.BigEndian.Uint64(hash[:8]))
}

// CreateUser creates a new user and stores it in the database
func CreateUser(userID int) User {
	timestamp := time.Now().Unix()
	salt := rand.Int()
	baseHash := GenerateBaseHash(userID, timestamp, salt)
	user := User{
		ID:        userID,
		BaseHash:  baseHash,
		Timestamp: timestamp,
		Salt:      salt,
	}
	Database[userID] = user
	return user
}

// VerifyAccessHash verifies if the provided access hash matches the expected hash for a visitor
func VerifyAccessHash(userID int, visitorID int, providedAccessHash int64) bool {
	user, exists := Database[userID]
	if !exists {
		return false
	}
	expectedHash := GenerateAccessHash(user, visitorID)
	return expectedHash == providedAccessHash
}

//func main() {
//	// Seed the random number generator
//	rand.Seed(time.Now().UnixNano())
//
//	// Simulate user creation
//	userID := 12345
//	user := CreateUser(userID)
//
//	// Print user details
//	fmt.Printf("User ID: %d\n", user.ID)
//	fmt.Printf("Base Hash: %d\n", user.BaseHash)
//
//	// Simulate API request from a visitor
//	visitorID := 54321
//	providedAccessHash := GenerateAccessHash(user, visitorID)
//	fmt.Printf("Provided Access Hash for Visitor ID %d: %d\n", visitorID, providedAccessHash)
//
//	// Verify access hash
//	if VerifyAccessHash(userID, visitorID, providedAccessHash) {
//		fmt.Println("Access hash verified successfully!")
//	} else {
//		fmt.Println("Access hash verification failed.")
//	}
//}
//
