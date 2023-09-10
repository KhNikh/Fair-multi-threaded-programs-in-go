package main
import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

const (
    NUM_RECORDS = 3
    NUM_CONN    = 6
)

type RecordData struct {
    Attrs [3]int
}

type Record struct {
    Data RecordData
    Lock sync.Mutex
}

type Database struct {
    Records [NUM_RECORDS]Record
}

var db Database

func acquireLock(txn rune, recIdx int) {
    fmt.Printf("txn %c: wants to acquire lock on record: %d\n", txn, recIdx)
    db.Records[recIdx].Lock.Lock()
    fmt.Printf("txn %c: acquired lock on record: %d\n", txn, recIdx)
}

func releaseLock(txn rune, recIdx int) {
    db.Records[recIdx].Lock.Unlock()
    fmt.Printf("txn %c: released lock on record: %d\n", txn, recIdx)
}

func initDB() {
    for i := 0; i < NUM_RECORDS; i++ {
        db.Records[i].Data.Attrs[0] = i // id
        db.Records[i].Data.Attrs[1] = rand.Intn(20) + 10 // age
    }
}

func mimicLoad(tname rune) {
    for {
        rec1 := rand.Intn(NUM_RECORDS)
        rec2 := rand.Intn(NUM_RECORDS)

        if rec1 == rec2 {
            continue
        }

        if rec1 > rec2 {
            rec1, rec2 = rec2, rec1
        }

        acquireLock(tname, rec1)
        acquireLock(tname, rec2)

        time.Sleep(2 * time.Second)

        releaseLock(tname, rec2)
        releaseLock(tname, rec1)

        time.Sleep(1 * time.Second)
    }
}

func main() {
    rand.Seed(time.Now().UnixNano())
    initDB()

    var wg sync.WaitGroup

    for i := 0; i < NUM_CONN; i++ {
        tname := rune('A' + i)
        wg.Add(1)
        go func() {
            defer wg.Done()
            mimicLoad(tname)
        }()
    }

    wg.Wait()

    for i := 0; i < NUM_RECORDS; i++ {
        db.Records[i].Lock = sync.Mutex{}
    }
}
