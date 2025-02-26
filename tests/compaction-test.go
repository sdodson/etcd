package main

import (
    "context"
    "fmt"
    "log"
    "time"

    clientv3 "go.etcd.io/etcd/client/v3"
)

func main() {
    client, err := clientv3.New(clientv3.Config{
        Endpoints: []string{"http://localhost:2400"},
    })
    if err != nil {
        log.Fatal(err)
    }
    defer client.Close()

    ctx := context.Background()

    // Generate 400 unique keys with 800 revisions each, approx 256b in size
    for i := 0; i < 400; i++ {
        key := fmt.Sprintf("key%d", i)
        for j := 0; j < 800; j++ {
            _, err := client.Put(ctx, key, fmt.Sprintf("bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb%d", j))
            if err != nil {
                log.Fatal(err)
            }
        }
        if i%100 == 0 {
            fmt.Printf("Generated %d keys\n", i)
        }
    }

    // Get the current revision
    status, err := client.Status(ctx, "http://localhost:2400")
    if err != nil {
        log.Fatal(err)
    }

    currentRevision := status.Header.Revision
 
    // Perform compaction
    start := time.Now()
    _, err = client.Compact(ctx, currentRevision)
    if err != nil {
        log.Fatal(err)
    }
    duration := time.Since(start)

    fmt.Printf("Compaction completed in %v\n", duration)

    // Defragment to reclaim space
//    _, err = client.Defragment(ctx, "http://localhost:2400")
//    if err != nil {
//#        log.Fatal(err)
//#    }
}
