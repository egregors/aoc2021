package days

import (
	"testing"

	"github.com/egregors/aoc2021/utils"
	"github.com/stretchr/testify/assert"
)

func Test_packetDecoder(t *testing.T) {
	// tests
	assert.Equal(t, 16, packetDecoderP1("8A004A801A8002F478"))
	assert.Equal(t, 12, packetDecoderP1("620080001611562C8802118E34"))
	assert.Equal(t, 23, packetDecoderP1("C0015000016115A2E0802F182340"))
	assert.Equal(t, 31, packetDecoderP1("A0016C880162017C3686B18A3D4780"))

	// solve
	xs, _ := utils.ReadLines("d16_input.txt")
	assert.Equal(t, 913, packetDecoderP1(xs[0]))

}

func Test_packetDEcoderP2(t *testing.T) {
	// tests
	assert.Equal(t, 3, packetDecoderP2("C200B40A82"))
	assert.Equal(t, 54, packetDecoderP2("04005AC33890"))
	assert.Equal(t, 7, packetDecoderP2("880086C3E88112"))
	assert.Equal(t, 9, packetDecoderP2("CE00C43D881120"))
	assert.Equal(t, 1, packetDecoderP2("D8005AC2A8F0"))
	assert.Equal(t, 0, packetDecoderP2("F600BC2D8F"))
	assert.Equal(t, 0, packetDecoderP2("9C005AC2F8F0"))
	assert.Equal(t, 1, packetDecoderP2("9C0141080250320F1802104A08"))

	// solve
	xs, _ := utils.ReadLines("d16_input.txt")
	assert.Equal(t, 1510977819698, packetDecoderP2(xs[0]))
}
