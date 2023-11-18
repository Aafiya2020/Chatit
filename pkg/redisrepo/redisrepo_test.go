package redisrepo

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load("../../.env")
}

func TestUpdateContactList(t *testing.T) {
	InitialiseRedis()
	defer redisClient.Close()

	for i := 4; i < 20; i++ {
		UpdateContactList("user1", fmt.Sprintf("user%d", i))
		time.Sleep(time.Second * 2)
	}
}

func TestFetchContactList(t *testing.T) {
	InitialiseRedis()
	defer redisClient.Close()

	res, err := FetchContactList("user1")
	if err != nil {
		t.Error("error in fetch", err)
		return
	}

	t.Log("success", res)
}

func TestFetchChatBetween(t *testing.T) {
	InitialiseRedis()
	defer redisClient.Close()

	res, err := FetchChatBetween("user1", "user2", "0", "+inf")

	if err != nil {
		t.Error("error in fetch", err)
		return
	}

	t.Log("success", res)
}

func TestIndexExist(t *testing.T) {
	InitialiseRedis()
	defer redisClient.Close()
	res, err := redisClient.Do(context.Background(),
		"FT._LIST",
	).Result()

	fmt.Printf("%T\n", res.([]interface{})[0])
	t.Log(res, err)
	fmt.Println(res.([]interface{})[0])
}

func TestCreateSortableIndex(t *testing.T) {
	InitialiseRedis()
	defer redisClient.Close()

	res, err := redisClient.Do(context.Background(),
		"FT.CREATE",
		"idx#chats",
		"ON", "JSON",
		"PREFIX", "1", "chat#",
		"SCHEMA", "$.from", "AS", "from", "TAG",
		"$.to", "AS", "to", "TAG",
		"$.timestamp", "AS", "timestamp", "NUMERIC", "SORTABLE",
	).Result()

	fmt.Println(res, err)
}
