/*
Copyright 2020 Jan Šimůnek github.com/hanysek
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package command

import (
	"fmt"
	"path"
	"path/filepath"
	"strings"
)

func fileNameWithoutExtension(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}

func addSuffixToFilePath(fpath string, suffix string) string {
	dir := filepath.Dir(fpath)
	fileName := filepath.Base(fpath)
	fileWithoutExt := fileNameWithoutExtension(fileName)
	ext := filepath.Ext(fpath)
	newFileName := path.Join(dir, fmt.Sprintf("%s%s%s", fileWithoutExt, suffix, ext))
	return newFileName
}
