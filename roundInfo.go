/*
Copyright 2017 Mosaic Networks Ltd

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
package hashgraph

type RoundInfo struct {
	Witnesses map[string]bool //witness => famous
	Events    []string
}

func (r *RoundInfo) AddEvent(x string, witness bool) {
	r.Events = append(r.Events, x)
	if witness {
		r.AddWitness(x)
	}
}

func (r *RoundInfo) AddWitness(x string) {
	if r.Witnesses == nil {
		r.Witnesses = make(map[string]bool)
	}
	r.Witnesses[x] = false
}

func (r *RoundInfo) SetFame(x string) {
	if r.Witnesses == nil {
		r.Witnesses = make(map[string]bool)
	}
	r.Witnesses[x] = true
}