// Copyright 2018 Bob Clinton <wildbearbob@gmail.com>
// All rights reserved
// https://github.com/wildbearbob/golab

package calc

//import "log"

type Float64Slice []float64

func calculate(a Float64Slice) float64 {
        n := len(a)
        slice := make([]Float64Slice, n)
        for i := 0; i < n; i++ {
                slice[i] = make(Float64Slice, n)
        }
        copy(slice[0], a)

        var k int
        for m := 1; m < n; m++ {
                k = n - m
                for q := 0; q < k; q++{
                        if m % 2 == 0 {
                                if slice[m-1][q+1] == slice[m-1][q] {
                                        slice[m][q] = 1.0
                                } else if slice[m-1][q] == 0.0 {
                                        slice[m][q] = slice[m-1][q]
                                } else {
                                        slice[m][q] = slice[m-1][q+1] / slice[m-1][q]
                                }
                                //slice[m][q] = slice[m-1][q+1] / slice[m-1][q]
                        } else {
                                slice[m][q] = slice[m-1][q+1] - slice[m-1][q]
                        }
                }
        }

        answer := slice[n-1][0]
        for w := n - 2; w > -1; w-- {
                if w % 2 == 0 {
                        answer = answer + slice[w][n-w-1]
                } else {
                        answer = answer * slice[w][n-w-1]
                }
        }

        return answer
}

func (data Float64Slice) Forecast(solution int) Float64Slice {
        l := len(data)

        // fix bug
        //if l + solution > 5 {
        //        log.Fatal()
        //}

        arr := make(Float64Slice, l, l + solution)
        copy(arr, data)

        for s := 0; s < solution; s++ {
                arr = append(arr, calculate(arr))
        }

        return arr
}
