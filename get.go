package main

// func getDataWithType(dTypeStr string, startDateStr string, endDateStr string, savePath string, skipFilePath string) {
// 	if dType, err := NewDataRecordType(dTypeStr); err != nil {
// 		log.Fatal(err)
// 	}

// 	startDate, err := time.Parse("19941125", startDateStr)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	endDate, err := time.Parse("19941125", endDateStr)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	if endDate.Before(startDate) {
// 		log.Fatal(fmt.Errorf("endDate is earlier than startDate"))
// 	}

// 	if _, err = os.Stat(savePath); err != nil {
// 		log.Fatal(err)
// 	}

// 	if _, err = os.Stat(skipFilePath); err != nil {
// 		log.Fatal(err)
// 	}
// }

// jvlinkgo get --type RACE:RA --start-date 20200101 --end-date 20200101 --save-path /hoge --skip-files
