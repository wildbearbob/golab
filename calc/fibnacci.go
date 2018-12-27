package calc

func Fibnacci(n int) int {
        var m1 = 1
        var m2, m3 int
        for i := 1; i < n; i++ {
                m2 = m1 + m3
                m3 = m1
                m1 = m2
        }
        return m1
}
