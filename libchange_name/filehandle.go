/**
 * Author: wengqiang (email: wens.wq@gmail.com  site: qiangweng.site)
 *
 * Copyright © 2015--2017 . All rights reserved.
 *
 * This library is free software under the terms of the GNU General Public License
 * as published by the Free Software Foundation; either version 3 of the License,
 * or (at your option) any later version.
 *
 */

package libchange_name

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func IsFileExist(filename string) bool {

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

func RenameFiles(files []string, rps map[string]string) {

	for _, file := range files {
		exist := IsFileExist(file)
		if !exist {
			fmt.Println("file not exist: ", file)
			continue
		}

		readFile, rErr := os.OpenFile(file, os.O_RDWR, 0666)
		writeFile, wErr := os.OpenFile(file+"_temp", os.O_CREATE|os.O_RDWR, 0666)
		if rErr != nil || wErr != nil {
			fmt.Println("replace file error: ", file, rErr.Error(), wErr.Error())
			continue
		}

		rd := bufio.NewReader(readFile)
		wd := bufio.NewWriter(writeFile)

		for {

			line, err := rd.ReadBytes('\n')

			if err != nil && err != io.EOF {
				panic("rename file error!")
			}

			for old, new := range rps {
				line = []byte(strings.Replace(string(line), old, new, -1))
			}

			wd.Write(line)
			wd.Flush()

			if io.EOF == err {
				break
			}

		}

		readFile.Close()
		os.Remove(file)

		writeFile.Close()
		os.Rename(file+"_temp", file)

	}
}

func RenameFileName(files []string, rps map[string]string) []string {

	res := files
	for _, file := range files {
		exist := IsFileExist(file)
		if !exist {
			fmt.Println("file not exist: ", file)
			continue
		}

		oldFile := file

		for old, new := range rps {
			if strings.Contains(file, old) {
				file = strings.Replace(file, old, new, -1)
			}

		}

		err := os.Rename(oldFile, file)
		if err != nil {
			fmt.Printf("rename file (%s) error!\n", file)
			fmt.Println(err.Error())
			continue
		}
	}

	return res
}
