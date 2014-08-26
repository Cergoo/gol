package plural

import (
	"github.com/Cergoo/gol/test"
	"testing"
)

var t1 *test.TT

func Test_main(t *testing.T) {
	t1 = test.New(t)
	//_________________	isInt

	// positive test cases
	t1.Eq(isInt(float64(0)), true)
	t1.Eq(isInt(float64(1)), true)
	t1.Eq(isInt(float64(0.0)), true)
	t1.Eq(isInt(float64(1.0000)), true)
	t1.Eq(isInt(float64(-50)), true)

	// negative
	t1.Eq(isInt(float64(0.1)), false)
	t1.Eq(isInt(float64(-0.1)), false)
	t1.Eq(isInt(float64(0.00000000000001)), false)

	//_________________	pluralRule1

	t1.Eq(pluralRule1(float64(0)), uint8(0))
	t1.Eq(pluralRule1(float64(0.5)), uint8(0))
	t1.Eq(pluralRule1(float64(100)), uint8(0))

	//_________________	pluralRule2A

	// first form
	t1.Eq(pluralRule2A(float64(-1)), uint8(0))
	t1.Eq(pluralRule2A(float64(1)), uint8(0))

	// second form
	t1.Eq(pluralRule2A(float64(0)), uint8(1))
	t1.Eq(pluralRule2A(float64(0.5)), uint8(1))
	t1.Eq(pluralRule2A(float64(2)), uint8(1))

	//_________________	pluralRule2B

	// first form
	t1.Eq(pluralRule2B(float64(-1)), uint8(0))
	t1.Eq(pluralRule2B(float64(0)), uint8(0))
	t1.Eq(pluralRule2B(float64(1)), uint8(0))

	// second form
	t1.Eq(pluralRule2B(float64(0.5)), uint8(1))
	t1.Eq(pluralRule2B(float64(2)), uint8(1))

	//_________________	pluralRule2C

	// first form
	t1.Eq(pluralRule2C(float64(-1)), uint8(0))
	t1.Eq(pluralRule2C(float64(0)), uint8(0))
	t1.Eq(pluralRule2C(float64(0.5)), uint8(0))
	t1.Eq(pluralRule2C(float64(1)), uint8(0))
	t1.Eq(pluralRule2C(float64(1.5)), uint8(0))

	// second form
	t1.Eq(pluralRule2C(float64(2)), uint8(1))
	t1.Eq(pluralRule2C(float64(2.5)), uint8(1))
	t1.Eq(pluralRule2C(float64(100)), uint8(1))

	//_________________	pluralRule2D

	// first form
	t1.Eq(pluralRule2D(float64(-1)), uint8(0))
	t1.Eq(pluralRule2D(float64(1)), uint8(0))
	t1.Eq(pluralRule2D(float64(21)), uint8(0))

	// second form
	t1.Eq(pluralRule2D(float64(0)), uint8(1))
	t1.Eq(pluralRule2D(float64(0.5)), uint8(1))
	t1.Eq(pluralRule2D(float64(2)), uint8(1))
	t1.Eq(pluralRule2D(float64(11)), uint8(1))

	//_________________	pluralRule2E

	// first form
	t1.Eq(pluralRule2E(float64(-1)), uint8(0))
	t1.Eq(pluralRule2E(float64(0)), uint8(0))
	t1.Eq(pluralRule2E(float64(1)), uint8(0))
	t1.Eq(pluralRule2E(float64(11)), uint8(0))
	t1.Eq(pluralRule2E(float64(12)), uint8(0))
	t1.Eq(pluralRule2E(float64(98)), uint8(0))
	t1.Eq(pluralRule2E(float64(99)), uint8(0))

	// second form
	t1.Eq(pluralRule2E(float64(0.5)), uint8(1))
	t1.Eq(pluralRule2E(float64(2)), uint8(1))
	t1.Eq(pluralRule2E(float64(10)), uint8(1))
	t1.Eq(pluralRule2E(float64(100)), uint8(1))

	//_________________	pluralRule2F

	// first form
	t1.Eq(pluralRule2F(float64(-1)), uint8(0))
	t1.Eq(pluralRule2F(float64(0)), uint8(0))
	t1.Eq(pluralRule2F(float64(1)), uint8(0))
	t1.Eq(pluralRule2F(float64(2)), uint8(0))
	t1.Eq(pluralRule2F(float64(11)), uint8(0))
	t1.Eq(pluralRule2F(float64(12)), uint8(0))
	t1.Eq(pluralRule2F(float64(20)), uint8(0))
	t1.Eq(pluralRule2F(float64(40)), uint8(0))

	// second form
	t1.Eq(pluralRule2F(float64(0.5)), uint8(1))
	t1.Eq(pluralRule2F(float64(3)), uint8(1))
	t1.Eq(pluralRule2F(float64(10)), uint8(1))

	//_________________	pluralRule3A

	// first form
	t1.Eq(pluralRule3A(float64(0)), uint8(0))

	// second form
	t1.Eq(pluralRule3A(float64(-1)), uint8(1))
	t1.Eq(pluralRule3A(float64(1)), uint8(1))
	t1.Eq(pluralRule3A(float64(21)), uint8(1))

	// third form
	t1.Eq(pluralRule3A(float64(0.5)), uint8(2))
	t1.Eq(pluralRule3A(float64(2)), uint8(2))
	t1.Eq(pluralRule3A(float64(10)), uint8(2))
	t1.Eq(pluralRule3A(float64(11)), uint8(2))

	//_________________	pluralRule3B

	// first form
	t1.Eq(pluralRule3B(float64(-1)), uint8(0))
	t1.Eq(pluralRule3B(float64(1)), uint8(0))

	// second form
	t1.Eq(pluralRule3B(float64(-2)), uint8(1))
	t1.Eq(pluralRule3B(float64(2)), uint8(1))

	// third form
	t1.Eq(pluralRule3B(float64(0)), uint8(2))
	t1.Eq(pluralRule3B(float64(0.5)), uint8(2))
	t1.Eq(pluralRule3B(float64(3)), uint8(2))
	t1.Eq(pluralRule3B(float64(11)), uint8(2))

	//_________________	pluralRule3C

	// first form
	t1.Eq(pluralRule3C(float64(-1)), uint8(0))
	t1.Eq(pluralRule3C(float64(1)), uint8(0))

	// second form
	t1.Eq(pluralRule3C(float64(0)), uint8(1))
	t1.Eq(pluralRule3C(float64(-11)), uint8(1))
	t1.Eq(pluralRule3C(float64(11)), uint8(1))
	t1.Eq(pluralRule3C(float64(19)), uint8(1))
	t1.Eq(pluralRule3C(float64(111)), uint8(1))
	t1.Eq(pluralRule3C(float64(119)), uint8(1))

	// third form
	t1.Eq(pluralRule3C(float64(0.5)), uint8(2))
	t1.Eq(pluralRule3C(float64(20)), uint8(2))
	t1.Eq(pluralRule3C(float64(21)), uint8(2))

	//_________________	pluralRule3D

	// first form
	t1.Eq(pluralRule3D(float64(-1)), uint8(0))
	t1.Eq(pluralRule3D(float64(1)), uint8(0))
	t1.Eq(pluralRule3D(float64(21)), uint8(0))

	// second form
	t1.Eq(pluralRule3D(float64(-2)), uint8(1))
	t1.Eq(pluralRule3D(float64(2)), uint8(1))
	t1.Eq(pluralRule3D(float64(9)), uint8(1))
	t1.Eq(pluralRule3D(float64(22)), uint8(1))
	t1.Eq(pluralRule3D(float64(29)), uint8(1))

	// third form
	t1.Eq(pluralRule3D(float64(0)), uint8(2))
	t1.Eq(pluralRule3D(float64(0.5)), uint8(2))
	t1.Eq(pluralRule3D(float64(11)), uint8(2))
	t1.Eq(pluralRule3D(float64(19)), uint8(2))

	//_________________	pluralRule3E

	// first form
	t1.Eq(pluralRule3E(float64(-1)), uint8(0))
	t1.Eq(pluralRule3E(float64(1)), uint8(0))

	// second form
	t1.Eq(pluralRule3E(float64(-2)), uint8(1))
	t1.Eq(pluralRule3E(float64(2)), uint8(1))
	t1.Eq(pluralRule3E(float64(3)), uint8(1))
	t1.Eq(pluralRule3E(float64(4)), uint8(1))

	// third form
	t1.Eq(pluralRule3E(float64(0)), uint8(2))
	t1.Eq(pluralRule3E(float64(0.5)), uint8(2))
	t1.Eq(pluralRule3E(float64(5)), uint8(2))
	t1.Eq(pluralRule3E(float64(9)), uint8(2))
	t1.Eq(pluralRule3E(float64(11)), uint8(2))
	t1.Eq(pluralRule3E(float64(12)), uint8(2))
	t1.Eq(pluralRule3E(float64(14)), uint8(2))

	//_________________	pluralRule3F

	// first form
	t1.Eq(pluralRule3F(float64(0)), uint8(0))

	// second form
	t1.Eq(pluralRule3F(float64(-0.5)), uint8(1))
	t1.Eq(pluralRule3F(float64(0.5)), uint8(1))
	t1.Eq(pluralRule3F(float64(1)), uint8(1))
	t1.Eq(pluralRule3F(float64(1.5)), uint8(1))

	// third form
	t1.Eq(pluralRule3F(float64(-2)), uint8(2))
	t1.Eq(pluralRule3F(float64(2)), uint8(2))
	t1.Eq(pluralRule3F(float64(3)), uint8(2))

	//_________________	pluralRule3G

	// first form
	t1.Eq(pluralRule3G(float64(-0.5)), uint8(0))
	t1.Eq(pluralRule3G(float64(0)), uint8(0))
	t1.Eq(pluralRule3G(float64(0.5)), uint8(0))
	t1.Eq(pluralRule3G(float64(1)), uint8(0))

	// second form
	t1.Eq(pluralRule3G(float64(-2)), uint8(1))
	t1.Eq(pluralRule3G(float64(2)), uint8(1))
	t1.Eq(pluralRule3G(float64(3)), uint8(1))
	t1.Eq(pluralRule3G(float64(9)), uint8(1))
	t1.Eq(pluralRule3G(float64(10)), uint8(1))

	// third
	t1.Eq(pluralRule3G(float64(1.5)), uint8(2))
	t1.Eq(pluralRule3G(float64(11)), uint8(2))
	t1.Eq(pluralRule3G(float64(12)), uint8(2))

	//_________________	pluralRule3H

	// first form
	t1.Eq(pluralRule3H(float64(0)), uint8(0))

	// second form
	t1.Eq(pluralRule3H(float64(-1)), uint8(1))
	t1.Eq(pluralRule3H(float64(1)), uint8(1))

	// third form
	t1.Eq(pluralRule3H(float64(0.5)), uint8(2))
	t1.Eq(pluralRule3H(float64(1.5)), uint8(2))
	t1.Eq(pluralRule3H(float64(2)), uint8(2))
	t1.Eq(pluralRule3H(float64(10)), uint8(2))
	t1.Eq(pluralRule3H(float64(11)), uint8(2))

	//_________________	pluralRule3I

	// first form
	t1.Eq(pluralRule3I(float64(-1)), uint8(0))
	t1.Eq(pluralRule3I(float64(1)), uint8(0))

	// second form
	t1.Eq(pluralRule3I(float64(-2)), uint8(1))
	t1.Eq(pluralRule3I(float64(2)), uint8(1))
	t1.Eq(pluralRule3I(float64(3)), uint8(1))
	t1.Eq(pluralRule3I(float64(4)), uint8(1))
	t1.Eq(pluralRule3I(float64(22)), uint8(1))
	t1.Eq(pluralRule3I(float64(23)), uint8(1))
	t1.Eq(pluralRule3I(float64(24)), uint8(1))

	// third form
	t1.Eq(pluralRule3I(float64(0)), uint8(2))
	t1.Eq(pluralRule3I(float64(0.5)), uint8(2))
	t1.Eq(pluralRule3I(float64(5)), uint8(2))
	t1.Eq(pluralRule3I(float64(9)), uint8(2))
	t1.Eq(pluralRule3I(float64(12)), uint8(2))
	t1.Eq(pluralRule3I(float64(13)), uint8(2))
	t1.Eq(pluralRule3I(float64(14)), uint8(2))
	t1.Eq(pluralRule3I(float64(15)), uint8(2))

	//_________________	pluralRule4A

	// first form
	t1.Eq(pluralRule4A(float64(-1)), uint8(0))
	t1.Eq(pluralRule4A(float64(1)), uint8(0))

	// second form
	t1.Eq(pluralRule4A(float64(-2)), uint8(1))
	t1.Eq(pluralRule4A(float64(2)), uint8(1))

	// third form
	t1.Eq(pluralRule4A(float64(-10)), uint8(2))
	t1.Eq(pluralRule4A(float64(10)), uint8(2))
	t1.Eq(pluralRule4A(float64(20)), uint8(2))
	t1.Eq(pluralRule4A(float64(100)), uint8(2))

	// fourth form
	t1.Eq(pluralRule4A(float64(0)), uint8(3))
	t1.Eq(pluralRule4A(float64(0.5)), uint8(3))
	t1.Eq(pluralRule4A(float64(3)), uint8(3))
	t1.Eq(pluralRule4A(float64(9)), uint8(3))
	t1.Eq(pluralRule4A(float64(11)), uint8(3))

	//_________________	pluralRule4B

	// first form
	t1.Eq(pluralRule4B(float64(-1)), uint8(0))
	t1.Eq(pluralRule4B(float64(1)), uint8(0))
	t1.Eq(pluralRule4B(float64(21)), uint8(0))

	// second form
	t1.Eq(pluralRule4B(float64(-2)), uint8(1))
	t1.Eq(pluralRule4B(float64(2)), uint8(1))
	t1.Eq(pluralRule4B(float64(3)), uint8(1))
	t1.Eq(pluralRule4B(float64(4)), uint8(1))
	t1.Eq(pluralRule4B(float64(22)), uint8(1))
	t1.Eq(pluralRule4B(float64(23)), uint8(1))
	t1.Eq(pluralRule4B(float64(24)), uint8(1))

	// third form
	t1.Eq(pluralRule4B(float64(-5)), uint8(2))
	t1.Eq(pluralRule4B(float64(0)), uint8(2))
	t1.Eq(pluralRule4B(float64(5)), uint8(2))
	t1.Eq(pluralRule4B(float64(6)), uint8(2))
	t1.Eq(pluralRule4B(float64(8)), uint8(2))
	t1.Eq(pluralRule4B(float64(9)), uint8(2))
	t1.Eq(pluralRule4B(float64(11)), uint8(2))
	t1.Eq(pluralRule4B(float64(12)), uint8(2))
	t1.Eq(pluralRule4B(float64(13)), uint8(2))
	t1.Eq(pluralRule4B(float64(14)), uint8(2))

	// fourth form
	t1.Eq(pluralRule4B(float64(0.5)), uint8(3))
	t1.Eq(pluralRule4B(float64(1.5)), uint8(3))

	//_________________	pluralRule4C

	// first form
	t1.Eq(pluralRule4C(float64(-1)), uint8(0))
	t1.Eq(pluralRule4C(float64(1)), uint8(0))

	// second form
	t1.Eq(pluralRule4C(float64(-2)), uint8(1))
	t1.Eq(pluralRule4C(float64(2)), uint8(1))
	t1.Eq(pluralRule4C(float64(3)), uint8(1))
	t1.Eq(pluralRule4C(float64(4)), uint8(1))
	t1.Eq(pluralRule4C(float64(22)), uint8(1))
	t1.Eq(pluralRule4C(float64(23)), uint8(1))
	t1.Eq(pluralRule4C(float64(24)), uint8(1))

	// third form
	t1.Eq(pluralRule4C(float64(-10)), uint8(2))
	t1.Eq(pluralRule4C(float64(10)), uint8(2))
	t1.Eq(pluralRule4C(float64(11)), uint8(2))
	t1.Eq(pluralRule4C(float64(12)), uint8(2))
	t1.Eq(pluralRule4C(float64(13)), uint8(2))
	t1.Eq(pluralRule4C(float64(14)), uint8(2))
	t1.Eq(pluralRule4C(float64(15)), uint8(2))
	t1.Eq(pluralRule4C(float64(16)), uint8(2))
	t1.Eq(pluralRule4C(float64(17)), uint8(2))
	t1.Eq(pluralRule4C(float64(18)), uint8(2))
	t1.Eq(pluralRule4C(float64(19)), uint8(2))
	t1.Eq(pluralRule4C(float64(20)), uint8(2))
	t1.Eq(pluralRule4C(float64(21)), uint8(2))
	t1.Eq(pluralRule4C(float64(25)), uint8(2))
	t1.Eq(pluralRule4C(float64(26)), uint8(2))
	t1.Eq(pluralRule4C(float64(28)), uint8(2))
	t1.Eq(pluralRule4C(float64(29)), uint8(2))

	// fourth form
	t1.Eq(pluralRule4C(float64(0.5)), uint8(3))
	t1.Eq(pluralRule4C(float64(1.5)), uint8(3))

	//_________________	pluralRule4D

	// first form
	t1.Eq(pluralRule4D(float64(-1)), uint8(0))
	t1.Eq(pluralRule4D(float64(1)), uint8(0))
	t1.Eq(pluralRule4D(float64(101)), uint8(0))

	// second form
	t1.Eq(pluralRule4D(float64(-2)), uint8(1))
	t1.Eq(pluralRule4D(float64(2)), uint8(1))
	t1.Eq(pluralRule4D(float64(102)), uint8(1))

	// third form
	t1.Eq(pluralRule4D(float64(-3)), uint8(2))
	t1.Eq(pluralRule4D(float64(3)), uint8(2))
	t1.Eq(pluralRule4D(float64(4)), uint8(2))
	t1.Eq(pluralRule4D(float64(103)), uint8(2))
	t1.Eq(pluralRule4D(float64(104)), uint8(2))

	// fourth form
	t1.Eq(pluralRule4D(float64(0)), uint8(3))
	t1.Eq(pluralRule4D(float64(0.5)), uint8(3))
	t1.Eq(pluralRule4D(float64(5)), uint8(3))
	t1.Eq(pluralRule4D(float64(10)), uint8(3))
	t1.Eq(pluralRule4D(float64(11)), uint8(3))
	t1.Eq(pluralRule4D(float64(12)), uint8(3))
	t1.Eq(pluralRule4D(float64(13)), uint8(3))
	t1.Eq(pluralRule4D(float64(14)), uint8(3))

	//_________________	pluralRule4E

	// first form
	t1.Eq(pluralRule4E(float64(-1)), uint8(0))
	t1.Eq(pluralRule4E(float64(1)), uint8(0))

	// second form
	t1.Eq(pluralRule4E(float64(-2)), uint8(1))
	t1.Eq(pluralRule4E(float64(0)), uint8(1))
	t1.Eq(pluralRule4E(float64(2)), uint8(1))
	t1.Eq(pluralRule4E(float64(10)), uint8(1))
	t1.Eq(pluralRule4E(float64(102)), uint8(1))
	t1.Eq(pluralRule4E(float64(110)), uint8(1))

	// third form
	t1.Eq(pluralRule4E(float64(-11)), uint8(2))
	t1.Eq(pluralRule4E(float64(11)), uint8(2))
	t1.Eq(pluralRule4E(float64(19)), uint8(2))
	t1.Eq(pluralRule4E(float64(111)), uint8(2))
	t1.Eq(pluralRule4E(float64(119)), uint8(2))

	// fourth form
	t1.Eq(pluralRule4E(float64(0.5)), uint8(3))
	t1.Eq(pluralRule4E(float64(20)), uint8(3))
	t1.Eq(pluralRule4E(float64(21)), uint8(3))
	t1.Eq(pluralRule4E(float64(22)), uint8(3))
	t1.Eq(pluralRule4E(float64(29)), uint8(3))

	//_________________	pluralRule4F

	// first form
	t1.Eq(pluralRule4F(float64(-1)), uint8(0))
	t1.Eq(pluralRule4F(float64(1)), uint8(0))
	t1.Eq(pluralRule4F(float64(11)), uint8(0))

	// second form
	t1.Eq(pluralRule4F(float64(-2)), uint8(1))
	t1.Eq(pluralRule4F(float64(2)), uint8(1))
	t1.Eq(pluralRule4F(float64(12)), uint8(1))

	// third form
	t1.Eq(pluralRule4F(float64(-3)), uint8(2))
	t1.Eq(pluralRule4F(float64(3)), uint8(2))
	t1.Eq(pluralRule4F(float64(10)), uint8(2))
	t1.Eq(pluralRule4F(float64(13)), uint8(2))
	t1.Eq(pluralRule4F(float64(19)), uint8(2))

	// fourth form
	t1.Eq(pluralRule4F(float64(0)), uint8(3))
	t1.Eq(pluralRule4F(float64(0.5)), uint8(3))
	t1.Eq(pluralRule4F(float64(20)), uint8(3))
	t1.Eq(pluralRule4F(float64(21)), uint8(3))
	t1.Eq(pluralRule4F(float64(22)), uint8(3))
	t1.Eq(pluralRule4F(float64(23)), uint8(3))
	t1.Eq(pluralRule4F(float64(29)), uint8(3))
	t1.Eq(pluralRule4F(float64(101)), uint8(3))
	t1.Eq(pluralRule4F(float64(102)), uint8(3))
	t1.Eq(pluralRule4F(float64(103)), uint8(3))
	t1.Eq(pluralRule4F(float64(109)), uint8(3))

	//_________________	pluralRule5A

	// first form
	t1.Eq(pluralRule5A(float64(-1)), uint8(0))
	t1.Eq(pluralRule5A(float64(1)), uint8(0))

	// second form
	t1.Eq(pluralRule5A(float64(-2)), uint8(1))
	t1.Eq(pluralRule5A(float64(2)), uint8(1))

	// third form
	t1.Eq(pluralRule5A(float64(-3)), uint8(2))
	t1.Eq(pluralRule5A(float64(3)), uint8(2))
	t1.Eq(pluralRule5A(float64(4)), uint8(2))
	t1.Eq(pluralRule5A(float64(5)), uint8(2))
	t1.Eq(pluralRule5A(float64(6)), uint8(2))

	// fourth form
	t1.Eq(pluralRule5A(float64(-7)), uint8(3))
	t1.Eq(pluralRule5A(float64(7)), uint8(3))
	t1.Eq(pluralRule5A(float64(8)), uint8(3))
	t1.Eq(pluralRule5A(float64(9)), uint8(3))
	t1.Eq(pluralRule5A(float64(10)), uint8(3))

	// fifth form
	t1.Eq(pluralRule5A(float64(0)), uint8(4))
	t1.Eq(pluralRule5A(float64(0.5)), uint8(4))
	t1.Eq(pluralRule5A(float64(11)), uint8(4))
	t1.Eq(pluralRule5A(float64(12)), uint8(4))
	t1.Eq(pluralRule5A(float64(13)), uint8(4))
	t1.Eq(pluralRule5A(float64(14)), uint8(4))
	t1.Eq(pluralRule5A(float64(15)), uint8(4))
	t1.Eq(pluralRule5A(float64(16)), uint8(4))
	t1.Eq(pluralRule5A(float64(17)), uint8(4))
	t1.Eq(pluralRule5A(float64(18)), uint8(4))
	t1.Eq(pluralRule5A(float64(19)), uint8(4))
	t1.Eq(pluralRule5A(float64(20)), uint8(4))

	//_________________	pluralRule5B

	// first form
	t1.Eq(pluralRule5B(float64(-1)), uint8(0))
	t1.Eq(pluralRule5B(float64(1)), uint8(0))
	t1.Eq(pluralRule5B(float64(21)), uint8(0))
	t1.Eq(pluralRule5B(float64(61)), uint8(0))
	t1.Eq(pluralRule5B(float64(81)), uint8(0))
	t1.Eq(pluralRule5B(float64(101)), uint8(0))

	// second form
	t1.Eq(pluralRule5B(float64(-2)), uint8(1))
	t1.Eq(pluralRule5B(float64(2)), uint8(1))
	t1.Eq(pluralRule5B(float64(22)), uint8(1))
	t1.Eq(pluralRule5B(float64(62)), uint8(1))
	t1.Eq(pluralRule5B(float64(82)), uint8(1))
	t1.Eq(pluralRule5B(float64(102)), uint8(1))

	// third form
	t1.Eq(pluralRule5B(float64(-3)), uint8(2))
	t1.Eq(pluralRule5B(float64(3)), uint8(2))
	t1.Eq(pluralRule5B(float64(4)), uint8(2))
	t1.Eq(pluralRule5B(float64(9)), uint8(2))
	t1.Eq(pluralRule5B(float64(23)), uint8(2))
	t1.Eq(pluralRule5B(float64(24)), uint8(2))
	t1.Eq(pluralRule5B(float64(29)), uint8(2))
	t1.Eq(pluralRule5B(float64(63)), uint8(2))
	t1.Eq(pluralRule5B(float64(64)), uint8(2))
	t1.Eq(pluralRule5B(float64(69)), uint8(2))
	t1.Eq(pluralRule5B(float64(83)), uint8(2))
	t1.Eq(pluralRule5B(float64(84)), uint8(2))
	t1.Eq(pluralRule5B(float64(89)), uint8(2))
	t1.Eq(pluralRule5B(float64(103)), uint8(2))
	t1.Eq(pluralRule5B(float64(104)), uint8(2))
	t1.Eq(pluralRule5B(float64(109)), uint8(2))

	// fourth form
	t1.Eq(pluralRule5B(float64(-1000000)), uint8(3))
	t1.Eq(pluralRule5B(float64(1000000)), uint8(3))
	t1.Eq(pluralRule5B(float64(2000000)), uint8(3))
	t1.Eq(pluralRule5B(float64(10000000)), uint8(3))

	// fourth form
	t1.Eq(pluralRule5B(float64(0)), uint8(4))
	t1.Eq(pluralRule5B(float64(0.5)), uint8(4))
	t1.Eq(pluralRule5B(float64(10)), uint8(4))
	t1.Eq(pluralRule5B(float64(11)), uint8(4))
	t1.Eq(pluralRule5B(float64(12)), uint8(4))
	t1.Eq(pluralRule5B(float64(13)), uint8(4))
	t1.Eq(pluralRule5B(float64(14)), uint8(4))
	t1.Eq(pluralRule5B(float64(19)), uint8(4))
	t1.Eq(pluralRule5B(float64(20)), uint8(4))
	t1.Eq(pluralRule5B(float64(71)), uint8(4))
	t1.Eq(pluralRule5B(float64(72)), uint8(4))
	t1.Eq(pluralRule5B(float64(73)), uint8(4))
	t1.Eq(pluralRule5B(float64(74)), uint8(4))
	t1.Eq(pluralRule5B(float64(79)), uint8(4))
	t1.Eq(pluralRule5B(float64(91)), uint8(4))
	t1.Eq(pluralRule5B(float64(92)), uint8(4))
	t1.Eq(pluralRule5B(float64(93)), uint8(4))
	t1.Eq(pluralRule5B(float64(94)), uint8(4))
	t1.Eq(pluralRule5B(float64(99)), uint8(4))
	t1.Eq(pluralRule5B(float64(100)), uint8(4))
	t1.Eq(pluralRule5B(float64(1000)), uint8(4))
	t1.Eq(pluralRule5B(float64(10000)), uint8(4))
	t1.Eq(pluralRule5B(float64(100000)), uint8(4))

	//_________________	pluralRule6A

	// first form
	t1.Eq(pluralRule6A(float64(0)), uint8(0))

	// second form
	t1.Eq(pluralRule6A(float64(-1)), uint8(1))
	t1.Eq(pluralRule6A(float64(1)), uint8(1))

	// third form
	t1.Eq(pluralRule6A(float64(-2)), uint8(2))
	t1.Eq(pluralRule6A(float64(2)), uint8(2))

	// fourth form
	t1.Eq(pluralRule6A(float64(-3)), uint8(3))
	t1.Eq(pluralRule6A(float64(3)), uint8(3))
	t1.Eq(pluralRule6A(float64(4)), uint8(3))
	t1.Eq(pluralRule6A(float64(9)), uint8(3))
	t1.Eq(pluralRule6A(float64(10)), uint8(3))
	t1.Eq(pluralRule6A(float64(103)), uint8(3))
	t1.Eq(pluralRule6A(float64(109)), uint8(3))
	t1.Eq(pluralRule6A(float64(110)), uint8(3))

	// fifth form
	t1.Eq(pluralRule6A(float64(-11)), uint8(4))
	t1.Eq(pluralRule6A(float64(11)), uint8(4))
	t1.Eq(pluralRule6A(float64(12)), uint8(4))
	t1.Eq(pluralRule6A(float64(98)), uint8(4))
	t1.Eq(pluralRule6A(float64(99)), uint8(4))
	t1.Eq(pluralRule6A(float64(111)), uint8(4))
	t1.Eq(pluralRule6A(float64(112)), uint8(4))
	t1.Eq(pluralRule6A(float64(198)), uint8(4))
	t1.Eq(pluralRule6A(float64(199)), uint8(4))

	// sixth form
	t1.Eq(pluralRule6A(float64(0.5)), uint8(5))
	t1.Eq(pluralRule6A(float64(100)), uint8(5))
	t1.Eq(pluralRule6A(float64(102)), uint8(5))
	t1.Eq(pluralRule6A(float64(200)), uint8(5))
	t1.Eq(pluralRule6A(float64(202)), uint8(5))

	//_________________	pluralRule6B

	// first form
	t1.Eq(pluralRule6B(float64(0)), uint8(0))

	// second form
	t1.Eq(pluralRule6B(float64(-1)), uint8(1))
	t1.Eq(pluralRule6B(float64(1)), uint8(1))

	// third form
	t1.Eq(pluralRule6B(float64(-2)), uint8(2))
	t1.Eq(pluralRule6B(float64(2)), uint8(2))

	// fourth form
	t1.Eq(pluralRule6B(float64(-3)), uint8(3))
	t1.Eq(pluralRule6B(float64(3)), uint8(3))

	// fifth form
	t1.Eq(pluralRule6B(float64(-6)), uint8(4))
	t1.Eq(pluralRule6B(float64(6)), uint8(4))

	// sixth form
	t1.Eq(pluralRule6B(float64(0.5)), uint8(5))
	t1.Eq(pluralRule6B(float64(4)), uint8(5))
	t1.Eq(pluralRule6B(float64(5)), uint8(5))
	t1.Eq(pluralRule6B(float64(7)), uint8(5))
	t1.Eq(pluralRule6B(float64(8)), uint8(5))
	t1.Eq(pluralRule6B(float64(9)), uint8(5))
	t1.Eq(pluralRule6B(float64(10)), uint8(5))
	t1.Eq(pluralRule6B(float64(11)), uint8(5))
	t1.Eq(pluralRule6B(float64(12)), uint8(5))
	t1.Eq(pluralRule6B(float64(13)), uint8(5))
	t1.Eq(pluralRule6B(float64(16)), uint8(5))

	//_________________	pluralRuleRu

	// first form
	t1.Eq(pluralRuleRu(float64(-1)), uint8(0))
	t1.Eq(pluralRuleRu(float64(1)), uint8(0))
	t1.Eq(pluralRuleRu(float64(21)), uint8(0))
	t1.Eq(pluralRuleRu(float64(31)), uint8(0))
	t1.Eq(pluralRuleRu(float64(101)), uint8(0))

	// second form
	t1.Eq(pluralRuleRu(float64(0)), uint8(1))
	t1.Eq(pluralRuleRu(float64(-5)), uint8(1))
	t1.Eq(pluralRuleRu(float64(6)), uint8(1))
	t1.Eq(pluralRuleRu(float64(7)), uint8(1))
	t1.Eq(pluralRuleRu(float64(11)), uint8(1))

	// third form
	t1.Eq(pluralRuleRu(float64(-2)), uint8(2))
	t1.Eq(pluralRuleRu(float64(22)), uint8(2))
	t1.Eq(pluralRuleRu(float64(1.5)), uint8(2))
	t1.Eq(pluralRuleRu(float64(101.1)), uint8(2))

}
