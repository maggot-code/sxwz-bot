/*
 * @FilePath: /sxwz-bot/pkg/fans/fans_test.go
 * @Author: maggot-code
 * @Date: 2023-09-19 08:39:28
 * @LastEditors: maggot-code
 * @LastEditTime: 2023-09-19 10:12:01
 * @Description:
 */
package fans

import (
	"fmt"
	"testing"
)

func TestNewFansRepo(t *testing.T) {
	fans, err := NewFansRepo(1660392980)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(fans)
}

func TestToUserTotal(t *testing.T) {
	fans, err := NewFansRepo(1660392980)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(fans.ToUserTotal())
}

func TestToUserTrue(t *testing.T) {
	fans, err := NewFansRepo(1660392980)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(fans.ToUserTrue())
}
