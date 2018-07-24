package application

import (
	"fmt"
	"testing"
)

func TestUnmarshal(t *testing.T) {

	test := `schema: "1"
name: basicapp
owner: "clint@getweave.com"
repo: "weavelab/basicapp"
namespace: default
ports:
- name: main-port
  protocol: TCP
  number: 9000
- name: diagnostics-port
  protocol: TCP
  number: 7777

resources:
  ram: 100m
  cpu: .5u
  minReplicas: 3
  maxReplicas: 10

env:
- name: THIS_PORT
  value: 9000
- name: PROPERTY_TWO
  configMapName: myConfigMap
  key: propertytwo
  

expose:
- name: http
  type: ingress
  domain: clint.getweave.com
  paths:
  - "/hi"
  - "/hello"
  destination: main-port
- name: diagnostics
  type: nodeport
  port: 30211
  destination: diagnostics-port
`
	a, err := Unmarshal([]byte(test))
	if err != nil {
		fmt.Print(err)
		t.Fail()
	}
	if a.Name != "basicapp" {
		t.Fail()
	}
	//fmt.Printf("%#v", a)

}
