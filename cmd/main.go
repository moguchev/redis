package main

import (
	"context"
	"log"
	"time"
	"unsafe"

	"github.com/go-redis/redis/v8"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func StringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}

func main() {
	ctx := context.Background()

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81",
		DB:       0, // use default DB
	})
	defer client.Close()

	cmd := client.Ping(ctx)
	if err := cmd.Err(); err != nil {
		log.Fatal(err)
	} else {
		log.Println(cmd.Val()) // PONG
	}

	Get(ctx, client)
	Mget(ctx, client)
}

func Get(ctx context.Context, client *redis.Client) {
	raw, err := client.Get(ctx, "key").Bytes()
	if err != nil {
		if err == redis.Nil {
			log.Println("key does not exist")
		} else {
			log.Println("Get failed", err)
		}
	}

	s := new(emptypb.Empty) // just for example
	if err = proto.Unmarshal(raw, s); err != nil {
		log.Println("proto.Unmarshal failed", err)
	}
	// encoding.BinaryMarshaler
}

func Mget(ctx context.Context, client *redis.Client) {
	t := timestamppb.Now()
	bytes, err := proto.Marshal(t)
	if err != nil {
		log.Println("proto.Marshal failed", err)
	}

	log.Println("set key-1 with ex 2 sec")
	if err = client.SetEX(ctx, "key-1", bytes, 2*time.Second).Err(); err != nil {
		log.Println("SetEX failed", err)
	}

	log.Println("set key-2 with ex 5 sec")
	if err = client.SetEX(ctx, "key-2", bytes, 5*time.Second).Err(); err != nil {
		log.Println("SetEX failed", err)
	}

	log.Println("MGet key-1, key-2, key-3")
	values, err := client.MGet(ctx, "key-1", "key-2", "key-3").Result()
	if err != nil {
		log.Println("MGet failed", err)
	} else {
	Loop:
		for _, value := range values {
			log.Printf("value: %#v", value)

			var bytes []byte

			switch v := value.(type) {
			case nil:
				log.Println("is nil")
				continue Loop
			case string:
				log.Println("is string")
				bytes = StringToBytes(v)
			case []byte:
				bytes = v
			}

			obj := new(timestamppb.Timestamp)
			if err = proto.Unmarshal(bytes, obj); err != nil {
				log.Println("proto.Unmarshal failed", err)
			}
			log.Println(obj)
		}
	}

	time.Sleep(3 * time.Second)

	log.Println("MGet key-1 key-2 key-3")
	values, err = client.MGet(ctx, "key-1", "key-2", "key-3").Result()
	if err != nil {
		log.Println("MGet failed", err)
	} else {
	Loop2:
		for _, value := range values {
			log.Printf("value: %#v", value)

			var bytes []byte

			switch v := value.(type) {
			case nil:
				log.Println("is nil")
				continue Loop2
			case string:
				log.Println("is string")
				bytes = StringToBytes(v)
			case []byte:
				bytes = v
			}

			obj := new(timestamppb.Timestamp)
			if err = proto.Unmarshal(bytes, obj); err != nil {
				log.Println("proto.Unmarshal failed", err)
			}
			log.Println(obj)
		}
	}

	time.Sleep(3 * time.Second)

	log.Println("MGet key-1 key-2 key-3")
	values, err = client.MGet(ctx, "key-1", "key-2", "key-3").Result()
	if err != nil {
		log.Println("MGet failed", err)
	} else {
	Loop3:
		for _, value := range values {
			log.Printf("value: %#v", value)

			var bytes []byte

			switch v := value.(type) {
			case nil:
				log.Println("is nil")
				continue Loop3
			case string:
				log.Println("is string")
				bytes = StringToBytes(v)
			case []byte:
				bytes = v
			}

			obj := new(timestamppb.Timestamp)
			if err = proto.Unmarshal(bytes, obj); err != nil {
				log.Println("proto.Unmarshal failed", err)
			}
			log.Println(obj)
		}
	}
}
