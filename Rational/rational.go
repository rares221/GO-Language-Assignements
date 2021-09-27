package main

import (
	"fmt"
)

type Rational struct {
	TopNum, DownNum int
}

func (r Rational) getTop() int{
	return r.TopNum
}

func (r Rational) getDown() int{
	return r.DownNum
}

func (r Rational) newInstance(x, y int) Rational{
	t:=Rational{x, y}
	return t
}


func (r Rational) add(t Rational) Rational{
	if t.TopNum == 0 {
		return r
	} else if r.DownNum ==t.DownNum {
		r.TopNum += t.TopNum
	} else {
		cmmmc := (r.DownNum * t.DownNum)/ cmmdc(r.DownNum,t.DownNum)
		r.TopNum = r.TopNum*t.DownNum + t.TopNum*r.DownNum
		r.DownNum *= cmmmc
	}
	return r
}

func (r Rational) substract(t Rational) Rational{
	if t.TopNum == 0 {
		return r
	} else if r.DownNum ==t.DownNum {
		r.TopNum -= t.TopNum
	} else {
		cmmmc := (r.DownNum * t.DownNum)/ cmmdc(r.DownNum,t.DownNum)
		r.TopNum = r.TopNum*t.DownNum - t.TopNum*r.DownNum
		r.DownNum *= cmmmc
	}

	return r
}

func (r Rational) multiply(t Rational) Rational{

	r.TopNum = r.TopNum *t.TopNum
	r.DownNum = r.DownNum *t.DownNum
	return r
}

func (r Rational) multiplyInt(t int) Rational{
	r.TopNum = r.TopNum *t
	return r
}

func (r Rational) divideInt(t int) Rational{
	r.DownNum = r.DownNum /t
	return r
}

func (r Rational) substractInt(t int) Rational{
	r.TopNum -= r.DownNum *t
	return r
}

func (r Rational) addInt(t int) Rational{
	r.TopNum += r.DownNum *t
	return r
}

func (r Rational) addTopNum(t int) Rational{
	r.TopNum +=t
	return r
}

func (r Rational) addDownNum(t int) Rational{
	r.DownNum +=t
	return r
}

func (r Rational) addBoth(t int) Rational{
	r.TopNum +=t
	r.DownNum +=t
	return r
}

func (r Rational) substractTopNum(t int) Rational{
	r.TopNum -=t
	return r
}

func (r Rational) substractDownNum(t int) Rational{
	r.DownNum -=t
	return r
}

func (r Rational) substractBoth(t int) Rational{
	r.TopNum -=t
	r.DownNum -=t
	return r
}

func (r Rational) isNull() bool{
	return r.TopNum == 0
}

func (r Rational) getRealValue() float32{
	return float32(r.TopNum)/ float32(r.DownNum)
}

func (r Rational) getAbsValue() Rational{
	if r.TopNum <0 {
		r.TopNum = -r.TopNum
	}
	if  r.DownNum < 0{
		r.DownNum = -r.DownNum
	}
	return r
}

func (r Rational) divide(t Rational) Rational{
	aux := t.TopNum
	t.TopNum = t.DownNum
	t.DownNum =aux
	r= r.multiply(t)
	return r
}

func (r Rational) simplify() Rational{
	c:=cmmdc(r.TopNum, r.DownNum)
	r.TopNum /= c
	r.DownNum /= c
	return r
}

func (r Rational) pow(n int) Rational{

	if n==0 {
		r.TopNum =1
		r.DownNum =1
	}
	numa := r.TopNum
	numi := r.DownNum

	for i:=0; i < n; i++ {
		r.TopNum = r.TopNum * numa
		r.DownNum = r.DownNum * numi
	}
	return r
}


func (r Rational) bigerThan(t Rational) bool{
	if r.DownNum == t.DownNum {
		return r.TopNum > t.TopNum
	}
	r.TopNum *= t.DownNum
	t.TopNum *= r.DownNum
	return r.TopNum > t.TopNum
}

func (r Rational) smallerThan(t Rational) bool{
	if r.DownNum == t.DownNum {
		return r.TopNum < t.TopNum
	}
	r.TopNum *= t.DownNum
	t.TopNum *= r.DownNum
	return r.TopNum < t.TopNum
}

func (r Rational) equals(t Rational) bool{
	rSimplificat := r.simplify()
	tSimplificat := t.simplify()

	return rSimplificat.TopNum == tSimplificat.TopNum && rSimplificat.DownNum == tSimplificat.DownNum
}

func (r Rational) inverse() Rational{
	aux :=r.TopNum
	r.TopNum =r.DownNum
	r.DownNum =aux
	return r
}

func (r Rational) isNatural() bool{

	return r.DownNum == 1
}

func cmmdc(a, b int) int{
	for a % b != 0{
		r:=a % b
		a = b
		b = r
	}
	return b
}


func main(){
	a:=Rational{TopNum: 3, DownNum: 2}
	b:=Rational{TopNum: 3, DownNum: 2}
	c:= a.add(b)
	fmt.Println(c)
	c=a.divideInt(2)
	fmt.Println(c)
	c=a.addTopNum(4)
	fmt.Println(c)


}
