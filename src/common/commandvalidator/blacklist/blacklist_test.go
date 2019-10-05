/*******************************************************************************
* Copyright 2019 Samsung Electronics All Rights Reserved.
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
* http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
*
*******************************************************************************/

package blacklist

import (
	"testing"
)

func TestIsBlack(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		command := "test"
		if ok := IsBlack(command); ok != false {
			t.Error("unexpected error")
		}
	})
	t.Run("Error", func(t *testing.T) {
		t.Run("IsBlack", func(t *testing.T) {
			for _, black := range blackList {
				if ok := IsBlack(black); ok == false {
					t.Error("unexpected succeed, " + black)
				}
			}
		})
	})
}
