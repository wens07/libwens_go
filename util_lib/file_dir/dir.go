/**
 * Author: wengqiang (email: wens.wq@gmail.com  site: qiangweng.site)
 *
 * Copyright © 2015--2018 . All rights reserved.
 *
 * File: dir.go, Date: 2018-11-17
 *
 *
 * This library is free software under the terms of the GNU General Public License
 * as published by the Free Software Foundation; either version 3 of the License,
 * or (at your option) any later version.
 *
 */

package file_dir

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func ListDir(path string, recursive bool) []string {

	var res []string

	//list dir recursive
	if recursive {
		filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			res = append(res, info.Name())
			return nil
		})
	} else {
		fileinfo, err := ioutil.ReadDir(path)
		if err != nil {
			panic(err)
		}

		for _, v := range fileinfo {
			res = append(res, v.Name())
		}

	}

	return res
}
