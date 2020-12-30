/**
* @Author: Jason
* @Date: 2020/9/29 14:15
* @File : io
* @Software: GoLand
**/

package dcore

import (
	"io"
	"net"
)

/**
判断是否末尾
*/
func IsEOFOrNetReadError(err error) bool {
	if err == io.EOF {
		return true
	}
	n, ok := err.(*net.OpError)
	return ok && n.Op == "read"
}
