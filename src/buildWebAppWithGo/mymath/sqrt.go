package mymath
//建议包名和目录名一致

func Sqrt(x float64) float64 {
	//函数名大写，在其他包就可以调用
	z := 0.0
	for i:=0;i<1000;i++{
		z -= (z*z -x )/ (2*x)
	}
	return z
}
