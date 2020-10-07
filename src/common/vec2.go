/**
* @Author: Jason
* @Date: 2020/9/30 10:17
* @File : vec2
* @Software: GoLand
**/

package common

import "math"

type Position struct {
	X int32
	Y int32
}

type Vec2 struct {
	X float64
	Y float64
}

func (self Vec2) magSqr() float64 {
	return self.X*self.X + self.Y*self.Y
}

func (self Vec2) dot(target Vec2) float64 {
	return self.X*target.X + self.Y*target.Y
}

func (self Vec2) Angle(target Vec2) float64 {
	magSqr1 := self.magSqr()
	magSqr2 := target.magSqr()
	if magSqr1 == 0 || magSqr2 == 0 {
		return 0.0
	}

	dot := self.dot(target)
	theta := dot / (math.Sqrt(magSqr1 * magSqr2))
	return math.Acos(theta)
}

/**
叉积
*/
func (self Vec2) cross(target Vec2) float64 {
	return self.X*target.Y - self.Y*target.X
}

/**
带方向的夹角的弧度
*/
func (self Vec2) SignAngle(target Vec2) float64 {
	angle := self.Angle(target)
	c := self.cross(target)
	if c < 0 {
		return -angle
	}
	return angle
}

func (self Vec2) NormalizeSelf() Vec2 {
	magSqr := self.X*self.X + self.Y*self.Y
	if magSqr == 1.0 || magSqr == 0.0 {
		return self
	}

	invSqrt := 1.0 / math.Sqrt(magSqr)
	self.X *= invSqrt
	self.Y *= invSqrt
	return self
}
