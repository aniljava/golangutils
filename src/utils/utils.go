package utils

import "encoding/csv"
import "os"
import "bufio"
import "io"

func Csv_as_channel(file string) chan map[string]string {

	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	reader := csv.NewReader(bufio.NewReader(f))

	head, err := reader.Read()
	if err != nil {
		panic(err)
	}

	channel := make(chan map[string]string, 1024)

	go func() {
		for {
			record, err := reader.Read()
			if err == io.EOF {
				close(channel)
				return
			} else {
				m := map[string]string{}
				for i := 0; i < len(head); i++ {
					key := head[i]
					value := record[i]
					m[key] = value
				}
				channel <- m
			}
		}
	}()

	return channel
}
