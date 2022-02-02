package vertexId

import (
	"math/rand"
	"time"
)

const letters = "abcdefghijklmnopqrstuvwxyz"
const numbers = "1234567890"
const coupletLength = 4
const vertexIdCoupletQuantity = 16

func init() {
	rand.Seed(time.Now().Unix())
}

func VertexId() string {
	b64 := VertexIdB64()
	return string(b64[:])
}

func VertexIdB64() (vertexId [64]byte) {
	for i := 0; i < vertexIdCoupletQuantity; i++ {
		for j, b := range couplet() {
			vertexId[(i*coupletLength)+j] = b
		}
	}
	return vertexId
}

func letterComponent() (byte, byte) {
	return letters[rand.Intn(26)], letters[rand.Intn(26)]
}

func numberComponent() (byte, byte) {
	return numbers[rand.Intn(10)], numbers[rand.Intn(10)]
}

func couplet() (doubleCouplet [4]byte) {
	lc0, lc1 := letterComponent()
	nc0, nc1 := numberComponent()
	return [4]byte{lc0, lc1, nc0, nc1}
}
