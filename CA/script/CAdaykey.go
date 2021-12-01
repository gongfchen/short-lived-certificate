package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"time"
)

func Date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func (ca Ca) gen_day_key(domain_name string, count int) (encrypt_key [32]uint8) {
	// DO NOT CHANGE
	// year, month, day := time.Now().Date()
	// fmt.Println(year, month, day)
	// t_now := Date(year, int(month), day)
	// t_start := ca.first_start_time
	// days := t_now.Sub(t_start).Hours() / 24
	// fmt.Println(days)

	// salt := domain_name + time.Time.String(ca.first_start_time.AddDate(0, 0, int(days)))
	salt := domain_name + time.Time.String(ca.first_start_time.AddDate(0, 0, count))
	encrypt_key = kdf(ca.master_key, salt)
	return encrypt_key
}

func (ca Ca) main_daykey(domain_name string, num_of_cert int) {
	fmt.Println("domain name: " + domain_name)
	// count := 0
	keypath := "../../CA-middle-daemon-storage/Daily Keys/" + domain_name
	if _, err := os.Stat(keypath); os.IsNotExist(err) {
		os.MkdirAll(keypath, 0700)
	}
	// for count < num_of_cert {
	key := ca.gen_day_key(domain_name, num_of_cert)
	daily_key := hex.EncodeToString(key[:])

	f, err := os.Create(keypath + "/daily_key.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString(daily_key)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written successfully --- Daily key written to shared storage")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	// count += 1
	// }

}
