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

func (v Vec2) magSqr() float64 {
	return v.X*v.X + v.Y*v.Y
}

func (v Vec2) dot(target Vec2) float64 {
	return v.X*target.X + v.Y*target.Y
}

func (v Vec2) Angle(target Vec2) float64 {
	magSqr1 := v.magSqr()
	magSqr2 := target.magSqr()
	if magSqr1 == 0 || magSqr2 == 0 {
		return 0.0
	}

	dot := v.dot(target)
	theta := dot / (math.Sqrt(magSqr1 * magSqr2))
	return math.Acos(theta)
}

/**
叉积
*/
func (v Vec2) cross(target Vec2) float64 {
	return v.X*target.Y - v.Y*target.X
}

/**
带方向的夹角的弧度
*/
func (v Vec2) SignAngle(target Vec2) float64 {
	angle := v.Angle(target)
	c := v.cross(target)
	if c < 0 {
		return -angle
	}
	return angle
}

func (v Vec2) NormalizeSelf() Vec2 {
	magSqr := v.X*v.X + v.Y*v.Y
	if magSqr == 1.0 || magSqr == 0.0 {
		return v
	}

	invSqrt := 1.0 / math.Sqrt(magSqr)
	v.X *= invSqrt
	v.Y *= invSqrt
	return v
}
