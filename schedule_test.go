/**
Copyright (c) 2013, Ryan Veach
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:
    * Redistributions of source code must retain the above copyright
      notice, this list of conditions and the following disclaimer.
    * Redistributions in binary form must reproduce the above copyright
      notice, this list of conditions and the following disclaimer in the
      documentation and/or other materials provided with the distribution.
    * Neither the name of the <organization> nor the
      names of its contributors may be used to endorse or promote products
      derived from this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL <COPYRIGHT HOLDER> BE LIABLE FOR ANY
DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
(INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
**/

package scheduler

import (
	"fmt"
	"testing"

	"time"
)

func TestID(t *testing.T) {
	var id ID
	if id.GetID() != 0 {
		t.Errorf("new default ID = %v, should be %v", id, 0)
	}

	var in, out uint = 1, 1
	id = ID(in)
	if id.GetID() != out {
		t.Errorf("ID(%v) = %v, want %v", in, id, out)
	}

	id = 3
	if id.GetID() != 3 {
		t.Errorf("id = %v", id.GetID())
	}
}

var cronTests = []struct {
	cronStr string
	time    int
	result  bool
}{
	{"1-3", 2, true},
	{"1-3", 3, true},
	{"1-3", 4, false},
	{"1-3,4", 4, true},
	{"1,3,4", 3, true},
	{"1,3,4", 5, false},
	{"1-3,5-8", 4, false},
	{"1-3/2", 3, true},
	{"1-3/2", 2, false},
}

func TestisMatch(t *testing.T) {
	for _, ct := range cronTests {
		check := timeMatchesCron(ct.cronStr, ct.time)
		if check != ct.result {
			t.Errorf("%s with time %d, expected result %v, got result %v",
				ct.cronStr, ct.time, ct.result, check)
		}
	}
}

func TestCronTime(t *testing.T) {
	cron, _ := NewCron("* 30 2 * * 1")
	fmt.Println(cron)
	nextTime := cron.NextRunTime(time.Now())
	fmt.Println(nextTime)
}

func TestRoundDown(t *testing.T) {
	fmt.Println(roundDownToSecond(time.Now()))
}
