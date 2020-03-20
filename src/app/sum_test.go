package main

import "testing"

func TestSum(t *testing.T) {
    total := sum(5, 5)
    if total != 10 {
       t.Errorf("A soma está incorreta, obtivemos: %d, queriamos: %d.", total, 10)
    }
}