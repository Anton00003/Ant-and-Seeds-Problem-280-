package main

import (
	"fmt"
	"math"
	"time"
)

func main() {

	//karts_1:= make([]float64,2*l)
	var karts_1 [][]float64
	var karts_2 [][][]m
	var y [][]float64
	var x [][]u
	var h int
	var l int
	var n_karts_00 int
	var rez float64
	var rez_ float64
	h = 5
	l = 5

	//karts_2:=make([][][]m,1)

	var k [2]int
	k[0] = 2
	k[1] = 2

	karts_1 = add_kart1(karts_1, l)
	karts_2 = add_kart2(karts_2, karts_1, h, l)
	karts_1, karts_2, n_karts_00 = add_karts(karts_2, karts_1, h, l, 0, n_karts_00)

	//karts_1[1][1]=5

	//x=append(x,ttt)
	x, y = zapolnenie(h, l)
	z(x, h, l)
	kvad(karts_2, y, k)
	//tik (karts_2, x, h,l,0)

	for i := 1; i < 4000; i++ {
		_, _, rez = tiks(karts_2, karts_1, x, h, l, n_karts_00, rez)
		rez_ = rez_ + rez*float64(i)
	}

	//for _,v:= range x{
	//	fmt.Println(v)}

	//	fmt.Println(len (x))
	//	fmt.Println(len (x[0]))
	//	fmt.Println(n_karts_00)
	fmt.Println(rez_)
	//	fmt.Println(karts_1)
	//	fmt.Println(len(karts_1))
	//	fmt.Println(karts_2)
	//	for _,v:= range karts_2{
	//	fmt.Println()
	//	for _,t:= range v{
	//	fmt.Println(t)}}
	time.Sleep(time.Second)
}

//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
type u struct {
	z float64
	u float64
	d float64
	l float64
	r float64
}

type m struct {
	z float64
	k float64
	p int
}

//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

func zapolnenie(h int, l int) ([][]u, [][]float64) {

	x := make([][]u, h)
	y := make([][]float64, h)

	for i := 0; i < h; i++ {

		x[i] = make([]u, l)
		y[i] = make([]float64, l)
	}
	return x, y

}

func kvad(karts_2 [][][]m, y [][]float64, k [2]int) {
	karts_2[0][k[0]][k[1]].z = 1

}

func z(x [][]u, h int, l int) {
	for i := 0; i < h; i++ {
		for j := 0; j < l; j++ {
			if i-1 >= 0 {
				x[i][j].u = 1
			}
			if i+1 <= len(x)-1 {
				x[i][j].d = 1
			}
			if j-1 >= 0 {
				x[i][j].l = 1
			}
			if j+1 <= len(x[0])-1 {
				x[i][j].r = 1
			}
		}
	}
}

func tiks(karts_2 [][][]m, karts_1 [][]float64, x [][]u, h int, l int, n_karts_00 int, rez float64) ([][]float64, [][][]m, float64) {
	var n_riad int
	rez = 0

	for i := 0; i < len(karts_2); i++ {
		tik(karts_2, x, h, l, i)
	}

	for i := 0; i < len(karts_2); i++ {
		for j := 0; j < l; j++ {
			n_riad = int((karts_1[i][2*l+1] - 1) * float64(h-1))
			if karts_2[i][n_riad][j].p > 0 {
				karts_2[karts_2[i][n_riad][j].p][n_riad][j].z = karts_2[karts_2[i][n_riad][j].p][n_riad][j].z + karts_2[i][n_riad][j].z
				karts_2[i][n_riad][j].z = 0
			}
		}
	}

	for j := 0; j < l; j++ {
		rez = rez + karts_2[n_karts_00][h-1][j].z
		karts_2[n_karts_00][h-1][j].z = 0
	}
	return karts_1, karts_2, rez
}

func tik(karts_2 [][][]m, x [][]u, h int, l int, g int) {
	var znach float64

	y := make([][]float64, h)
	for i := 0; i < h; i++ {
		y[i] = make([]float64, l)
	}

	for i := 0; i < h; i++ {
		for j := 0; j < l; j++ {
			znach = karts_2[g][i][j].z / (x[i][j].u + x[i][j].d + x[i][j].l + x[i][j].r)

			if i-1 >= 0 {
				y[i-1][j] = y[i-1][j] + znach
			}
			if i+1 <= len(x)-1 {
				y[i+1][j] = y[i+1][j] + znach
			}
			if j-1 >= 0 {
				y[i][j-1] = y[i][j-1] + znach
			}
			if j+1 <= len(x[0])-1 {
				y[i][j+1] = y[i][j+1] + znach
			}

			//y[i][j] = y[i][j] - x[i][j].z

		}
	}

	for i := 0; i < h; i++ {
		for j := 0; j < l; j++ {

			karts_2[g][i][j].z = y[i][j]

		}
	}

}

func add_kart1(karts_1 [][]float64, l int) [][]float64 {

	b := make([]float64, 2*l+2)
	karts_1 = append(karts_1, b)
	for i := 0; i < l; i++ {
		karts_1[0][i] = 1
		karts_1[0][2*l] = karts_1[0][2*l]*10 + karts_1[0][i]
	}
	for i := l; i < 2*l; i++ {
		karts_1[0][i] = 2
		karts_1[0][2*l] = karts_1[0][2*l]*10 + karts_1[0][i]
	}
	karts_1[0][2*l+1] = 1
	return karts_1

}

func add_karts(karts_2 [][][]m, karts_1 [][]float64, h int, l int, n int, n_karts_00 int) ([][]float64, [][][]m, int) {

	var num float64
	var id int

	b := make([]float64, 2*l+2)
	for i := 0; i < 2*l+2; i++ {
		b[i] = karts_1[n][i]
	}

	for i := 0; i < 2*l; i++ {
		if b[i] == b[2*l+1] {

			num = b[2*l] - b[i]*math.Pow(10, float64(2*l-i-1))

			id = 0

			for j := 0; j < len(karts_1); j++ {
				if num == karts_1[j][2*l] {
					id = 1
					karts_2[n][int((karts_1[n][2*l+1]-1)*float64(h-1))][i%l].p = j
				}
			}
			if id == 0 {

				karts_1 = add_kart(karts_1, b, l)
				karts_1[len(karts_1)-1][i] = 0
				karts_1[len(karts_1)-1][2*l] = num
				karts_1[len(karts_1)-1][2*l+1] = math.Abs(b[2*l+1]-2) + 1

				if num == 0 {
					n_karts_00 = len(karts_1) - 1
				}

				karts_2 = add_kart2(karts_2, karts_1, h, l)

				karts_2[n][int((karts_1[n][2*l+1]-1)*float64(h-1))][i%l].p = len(karts_2) - 1

				karts_1, karts_2, n_karts_00 = add_karts(karts_2, karts_1, h, l, len(karts_1)-1, n_karts_00)

			}
		}
	}
	return karts_1, karts_2, n_karts_00

}

func add_kart(karts_1 [][]float64, b []float64, l int) [][]float64 {
	c := make([]float64, 2*l+2)

	for i := 0; i < 2*l+2; i++ {
		c[i] = b[i]
	}
	karts_1 = append(karts_1, c)
	return karts_1
}

func add_kart2(karts_2 [][][]m, karts_1 [][]float64, h int, l int) [][][]m {

	c := make([][]m, h)
	for i := 0; i < h; i++ {
		c[i] = make([]m, l)
	}

	//c[1][1].k = 2

	for i := 0; i < l; i++ {
		c[0][i].k = karts_1[len(karts_1)-1][i]
		c[h-1][i].k = karts_1[len(karts_1)-1][l+i]
	}

	karts_2 = append(karts_2, c)

	return karts_2
}
