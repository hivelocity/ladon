/*
 * Copyright © 2016-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * @author		Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @copyright 	2015-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @license 	Apache-2.0
 */

package ladon_test

import (
	"fmt"
	"os"
	"sync"
	"testing"

	. "gitlab.host1plus.com/linas/ladon"
	. "gitlab.host1plus.com/linas/ladon/manager/memory"
)

var managers = map[string]Manager{}
var migrators = map[string]ManagerMigrator{}

func TestMain(m *testing.M) {
	var wg sync.WaitGroup
	wg.Add(1)
	connectMEM(&wg)
	wg.Wait()

	s := m.Run()
	os.Exit(s)
}

func connectMEM(wg *sync.WaitGroup) {
	defer wg.Done()
	managers["memory"] = NewMemoryManager()
}

func TestManagers(t *testing.T) {
	t.Run("type=get errors", func(t *testing.T) {
		for k, s := range managers {
			t.Run("manager="+k, TestHelperGetErrors(s))
		}
	})

	t.Run("type=CRUD", func(t *testing.T) {
		for k, s := range managers {
			t.Run(fmt.Sprintf("manager=%s", k), TestHelperCreateGetDelete(s))
		}
	})

	t.Run("type=find", func(t *testing.T) {
		for k, s := range map[string]Manager{
			"postgres": managers["postgres"],
			"mysql":    managers["mysql"],
		} {
			t.Run(fmt.Sprintf("manager=%s", k), TestHelperFindPoliciesForSubject(k, s))
			t.Run(fmt.Sprintf("manager=%s", k), TestHelperFindPoliciesForResource(k, s))
		}
	})
}
