# Copyright 2020 Jan Šimůnek github.com/hanysek
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

APP_NAME = c1-utils
LDFLAGS := -ldflags="-s -w"

deps:
	mkdir -p $(shell go env GOPATH)/bin
.PHONY: deps

build:
	go build -o ./bin/$(APP_NAME) ./cmd/c1-utils.go

install:
	go install ./cmd/c1-utils.go
