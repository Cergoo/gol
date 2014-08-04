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

	t1.Eq(pluralRule1(float64(0)), 0)
	t1.Eq(pluralRule1(float64(0.5)), 0)
	t1.Eq(pluralRule1(float64(100)), 0)

	//_________________	pluralRule2A

	// first form
	t1.Eq(pluralRule2A(float64(-1)), 0)
	t1.Eq(pluralRule2A(float64(1)), 0)

	// second form
	t1.Eq(pluralRule2A(float64(0)), 1)
	t1.Eq(pluralRule2A(float64(0.5)), 1)
	t1.Eq(pluralRule2A(float64(2)), 1)

	//_________________	pluralRule2B

	// first form
	t1.Eq(pluralRule2B(float64(-1)), 0)
	t1.Eq(pluralRule2B(float64(0)), 0)
	t1.Eq(pluralRule2B(float64(1)), 0)

	// second form
	t1.Eq(pluralRule2B(float64(0.5)), 1)
	t1.Eq(pluralRule2B(float64(2)), 1)

	//_________________	pluralRule2C

	// first form
	t1.Eq(pluralRule2C(float64(-1)), 0)
	t1.Eq(pluralRule2C(float64(0)), 0)
	t1.Eq(pluralRule2C(float64(0.5)), 0)
	t1.Eq(pluralRule2C(float64(1)), 0)
	t1.Eq(pluralRule2C(float64(1.5)), 0)

	// second form
	t1.Eq(pluralRule2C(float64(2)), 1)
	t1.Eq(pluralRule2C(float64(2.5)), 1)
	t1.Eq(pluralRule2C(float64(100)), 1)

	//_________________	pluralRule2D

	// first form
	t1.Eq(pluralRule2D(float64(-1)), 0)
	t1.Eq(pluralRule2D(float64(1)), 0)
	t1.Eq(pluralRule2D(float64(21)), 0)

	// second form
	t1.Eq(pluralRule2D(float64(0)), 1)
	t1.Eq(pluralRule2D(float64(0.5)), 1)
	t1.Eq(pluralRule2D(float64(2)), 1)
	t1.Eq(pluralRule2D(float64(11)), 1)

	//_________________	pluralRule2E

	// first form
	t1.Eq(pluralRule2E(float64(-1)), 0)
	t1.Eq(pluralRule2E(float64(0)), 0)
	t1.Eq(pluralRule2E(float64(1)), 0)
	t1.Eq(pluralRule2E(float64(11)), 0)
	t1.Eq(pluralRule2E(float64(12)), 0)
	t1.Eq(pluralRule2E(float64(98)), 0)
	t1.Eq(pluralRule2E(float64(99)), 0)

	// second form
	t1.Eq(pluralRule2E(float64(0.5)), 1)
	t1.Eq(pluralRule2E(float64(2)), 1)
	t1.Eq(pluralRule2E(float64(10)), 1)
	t1.Eq(pluralRule2E(float64(100)), 1)

	//_________________	pluralRule2F

	// first form
	t1.Eq(pluralRule2F(float64(-1)), 0)
	t1.Eq(pluralRule2F(float64(0)), 0)
	t1.Eq(pluralRule2F(float64(1)), 0)
	t1.Eq(pluralRule2F(float64(2)), 0)
	t1.Eq(pluralRule2F(float64(11)), 0)
	t1.Eq(pluralRule2F(float64(12)), 0)
	t1.Eq(pluralRule2F(float64(20)), 0)
	t1.Eq(pluralRule2F(float64(40)), 0)

	// second form
	t1.Eq(pluralRule2F(float64(0.5)), 1)
	t1.Eq(pluralRule2F(float64(3)), 1)
	t1.Eq(pluralRule2F(float64(10)), 1)

	//_________________	pluralRule3A

	// first form
	t1.Eq(pluralRule3A(float64(0)), 0)

	// second form
	t1.Eq(pluralRule3A(float64(-1)), 1)
	t1.Eq(pluralRule3A(float64(1)), 1)
	t1.Eq(pluralRule3A(float64(21)), 1)

	// third form
	t1.Eq(pluralRule3A(float64(0.5)), 2)
	t1.Eq(pluralRule3A(float64(2)), 2)
	t1.Eq(pluralRule3A(float64(10)), 2)
	t1.Eq(pluralRule3A(float64(11)), 2)

	//_________________	pluralRule3B

	// first form
	t1.Eq(pluralRule3B(float64(-1)), 0)
	t1.Eq(pluralRule3B(float64(1)), 0)

	// second form
	t1.Eq(pluralRule3B(float64(-2)), 1)
	t1.Eq(pluralRule3B(float64(2)), 1)

	// third form
	t1.Eq(pluralRule3B(float64(0)), 2)
	t1.Eq(pluralRule3B(float64(0.5)), 2)
	t1.Eq(pluralRule3B(float64(3)), 2)
	t1.Eq(pluralRule3B(float64(11)), 2)

	//_________________	pluralRule3C

	// first form
	t1.Eq(pluralRule3C(float64(-1)), 0)
	t1.Eq(pluralRule3C(float64(1)), 0)

	// second form
	t1.Eq(pluralRule3C(float64(0)), 1)
	t1.Eq(pluralRule3C(float64(-11)), 1)
	t1.Eq(pluralRule3C(float64(11)), 1)
	t1.Eq(pluralRule3C(float64(19)), 1)
	t1.Eq(pluralRule3C(float64(111)), 1)
	t1.Eq(pluralRule3C(float64(119)), 1)

	// third form
	t1.Eq(pluralRule3C(float64(0.5)), 2)
	t1.Eq(pluralRule3C(float64(20)), 2)
	t1.Eq(pluralRule3C(float64(21)), 2)

	//_________________	pluralRule3D

	// first form
	t1.Eq(pluralRule3D(float64(-1)), 0)
	t1.Eq(pluralRule3D(float64(1)), 0)
	t1.Eq(pluralRule3D(float64(21)), 0)

	// second form
	t1.Eq(pluralRule3D(float64(-2)), 1)
	t1.Eq(pluralRule3D(float64(2)), 1)
	t1.Eq(pluralRule3D(float64(9)), 1)
	t1.Eq(pluralRule3D(float64(22)), 1)
	t1.Eq(pluralRule3D(float64(29)), 1)

	// third form
	t1.Eq(pluralRule3D(float64(0)), 2)
	t1.Eq(pluralRule3D(float64(0.5)), 2)
	t1.Eq(pluralRule3D(float64(11)), 2)
	t1.Eq(pluralRule3D(float64(19)), 2)

	//_________________	pluralRule3E

	// first form
	t1.Eq(pluralRule3E(float64(-1)), 0)
	t1.Eq(pluralRule3E(float64(1)), 0)

	// second form
	t1.Eq(pluralRule3E(float64(-2)), 1)
	t1.Eq(pluralRule3E(float64(2)), 1)
	t1.Eq(pluralRule3E(float64(3)), 1)
	t1.Eq(pluralRule3E(float64(4)), 1)

	// third form
	t1.Eq(pluralRule3E(float64(0)), 2)
	t1.Eq(pluralRule3E(float64(0.5)), 2)
	t1.Eq(pluralRule3E(float64(5)), 2)
	t1.Eq(pluralRule3E(float64(9)), 2)
	t1.Eq(pluralRule3E(float64(11)), 2)
	t1.Eq(pluralRule3E(float64(12)), 2)
	t1.Eq(pluralRule3E(float64(14)), 2)

	//_________________	pluralRule3F

	// first form
	t1.Eq(pluralRule3F(float64(0)), 0)

	// second form
	t1.Eq(pluralRule3F(float64(-0.5)), 1)
	t1.Eq(pluralRule3F(float64(0.5)), 1)
	t1.Eq(pluralRule3F(float64(1)), 1)
	t1.Eq(pluralRule3F(float64(1.5)), 1)

	// third form
	t1.Eq(pluralRule3F(float64(-2)), 2)
	t1.Eq(pluralRule3F(float64(2)), 2)
	t1.Eq(pluralRule3F(float64(3)), 2)

	//_________________	pluralRule3G

	// first form
	t1.Eq(pluralRule3G(float64(-0.5)), 0)
	t1.Eq(pluralRule3G(float64(0)), 0)
	t1.Eq(pluralRule3G(float64(0.5)), 0)
	t1.Eq(pluralRule3G(float64(1)), 0)

	// second form
	t1.Eq(pluralRule3G(float64(-2)), 1)
	t1.Eq(pluralRule3G(float64(2)), 1)
	t1.Eq(pluralRule3G(float64(3)), 1)
	t1.Eq(pluralRule3G(float64(9)), 1)
	t1.Eq(pluralRule3G(float64(10)), 1)

	// third
	t1.Eq(pluralRule3G(float64(1.5)), 2)
	t1.Eq(pluralRule3G(float64(11)), 2)
	t1.Eq(pluralRule3G(float64(12)), 2)

	//_________________	pluralRule3H

	// first form
	t1.Eq(pluralRule3H(float64(0)), 0)

	// second form
	t1.Eq(pluralRule3H(float64(-1)), 1)
	t1.Eq(pluralRule3H(float64(1)), 1)

	// third form
	t1.Eq(pluralRule3H(float64(0.5)), 2)
	t1.Eq(pluralRule3H(float64(1.5)), 2)
	t1.Eq(pluralRule3H(float64(2)), 2)
	t1.Eq(pluralRule3H(float64(10)), 2)
	t1.Eq(pluralRule3H(float64(11)), 2)

	//_________________	pluralRule3I

	// first form
	t1.Eq(pluralRule3I(float64(-1)), 0)
	t1.Eq(pluralRule3I(float64(1)), 0)

	// second form
	t1.Eq(pluralRule3I(float64(-2)), 1)
	t1.Eq(pluralRule3I(float64(2)), 1)
	t1.Eq(pluralRule3I(float64(3)), 1)
	t1.Eq(pluralRule3I(float64(4)), 1)
	t1.Eq(pluralRule3I(float64(22)), 1)
	t1.Eq(pluralRule3I(float64(23)), 1)
	t1.Eq(pluralRule3I(float64(24)), 1)

	// third form
	t1.Eq(pluralRule3I(float64(0)), 2)
	t1.Eq(pluralRule3I(float64(0.5)), 2)
	t1.Eq(pluralRule3I(float64(5)), 2)
	t1.Eq(pluralRule3I(float64(9)), 2)
	t1.Eq(pluralRule3I(float64(12)), 2)
	t1.Eq(pluralRule3I(float64(13)), 2)
	t1.Eq(pluralRule3I(float64(14)), 2)
	t1.Eq(pluralRule3I(float64(15)), 2)

	//_________________	pluralRule4A

	// first form
	t1.Eq(pluralRule4A(float64(-1)), 0)
	t1.Eq(pluralRule4A(float64(1)), 0)

	// second form
	t1.Eq(pluralRule4A(float64(-2)), 1)
	t1.Eq(pluralRule4A(float64(2)), 1)

	// third form
	t1.Eq(pluralRule4A(float64(-10)), 2)
	t1.Eq(pluralRule4A(float64(10)), 2)
	t1.Eq(pluralRule4A(float64(20)), 2)
	t1.Eq(pluralRule4A(float64(100)), 2)

	// fourth form
	t1.Eq(pluralRule4A(float64(0)), 3)
	t1.Eq(pluralRule4A(float64(0.5)), 3)
	t1.Eq(pluralRule4A(float64(3)), 3)
	t1.Eq(pluralRule4A(float64(9)), 3)
	t1.Eq(pluralRule4A(float64(11)), 3)

	//_________________	pluralRule4B

	// first form
	t1.Eq(pluralRule4B(float64(-1)), 0)
	t1.Eq(pluralRule4B(float64(1)), 0)
	t1.Eq(pluralRule4B(float64(21)), 0)

	// second form
	t1.Eq(pluralRule4B(float64(-2)), 1)
	t1.Eq(pluralRule4B(float64(2)), 1)
	t1.Eq(pluralRule4B(float64(3)), 1)
	t1.Eq(pluralRule4B(float64(4)), 1)
	t1.Eq(pluralRule4B(float64(22)), 1)
	t1.Eq(pluralRule4B(float64(23)), 1)
	t1.Eq(pluralRule4B(float64(24)), 1)

	// third form
	t1.Eq(pluralRule4B(float64(-5)), 2)
	t1.Eq(pluralRule4B(float64(0)), 2)
	t1.Eq(pluralRule4B(float64(5)), 2)
	t1.Eq(pluralRule4B(float64(6)), 2)
	t1.Eq(pluralRule4B(float64(8)), 2)
	t1.Eq(pluralRule4B(float64(9)), 2)
	t1.Eq(pluralRule4B(float64(11)), 2)
	t1.Eq(pluralRule4B(float64(12)), 2)
	t1.Eq(pluralRule4B(float64(13)), 2)
	t1.Eq(pluralRule4B(float64(14)), 2)

	// fourth form
	t1.Eq(pluralRule4B(float64(0.5)), 3)
	t1.Eq(pluralRule4B(float64(1.5)), 3)

	//_________________	pluralRule4C

	// first form
	t1.Eq(pluralRule4C(float64(-1)), 0)
	t1.Eq(pluralRule4C(float64(1)), 0)

	// second form
	t1.Eq(pluralRule4C(float64(-2)), 1)
	t1.Eq(pluralRule4C(float64(2)), 1)
	t1.Eq(pluralRule4C(float64(3)), 1)
	t1.Eq(pluralRule4C(float64(4)), 1)
	t1.Eq(pluralRule4C(float64(22)), 1)
	t1.Eq(pluralRule4C(float64(23)), 1)
	t1.Eq(pluralRule4C(float64(24)), 1)

	// third form
	t1.Eq(pluralRule4C(float64(-10)), 2)
	t1.Eq(pluralRule4C(float64(10)), 2)
	t1.Eq(pluralRule4C(float64(11)), 2)
	t1.Eq(pluralRule4C(float64(12)), 2)
	t1.Eq(pluralRule4C(float64(13)), 2)
	t1.Eq(pluralRule4C(float64(14)), 2)
	t1.Eq(pluralRule4C(float64(15)), 2)
	t1.Eq(pluralRule4C(float64(16)), 2)
	t1.Eq(pluralRule4C(float64(17)), 2)
	t1.Eq(pluralRule4C(float64(18)), 2)
	t1.Eq(pluralRule4C(float64(19)), 2)
	t1.Eq(pluralRule4C(float64(20)), 2)
	t1.Eq(pluralRule4C(float64(21)), 2)
	t1.Eq(pluralRule4C(float64(25)), 2)
	t1.Eq(pluralRule4C(float64(26)), 2)
	t1.Eq(pluralRule4C(float64(28)), 2)
	t1.Eq(pluralRule4C(float64(29)), 2)

	// fourth form
	t1.Eq(pluralRule4C(float64(0.5)), 3)
	t1.Eq(pluralRule4C(float64(1.5)), 3)

	//_________________	pluralRule4D

	// first form
	t1.Eq(pluralRule4D(float64(-1)), 0)
	t1.Eq(pluralRule4D(float64(1)), 0)
	t1.Eq(pluralRule4D(float64(101)), 0)

	// second form
	t1.Eq(pluralRule4D(float64(-2)), 1)
	t1.Eq(pluralRule4D(float64(2)), 1)
	t1.Eq(pluralRule4D(float64(102)), 1)

	// third form
	t1.Eq(pluralRule4D(float64(-3)), 2)
	t1.Eq(pluralRule4D(float64(3)), 2)
	t1.Eq(pluralRule4D(float64(4)), 2)
	t1.Eq(pluralRule4D(float64(103)), 2)
	t1.Eq(pluralRule4D(float64(104)), 2)

	// fourth form
	t1.Eq(pluralRule4D(float64(0)), 3)
	t1.Eq(pluralRule4D(float64(0.5)), 3)
	t1.Eq(pluralRule4D(float64(5)), 3)
	t1.Eq(pluralRule4D(float64(10)), 3)
	t1.Eq(pluralRule4D(float64(11)), 3)
	t1.Eq(pluralRule4D(float64(12)), 3)
	t1.Eq(pluralRule4D(float64(13)), 3)
	t1.Eq(pluralRule4D(float64(14)), 3)

	//_________________	pluralRule4E

	// first form
	t1.Eq(pluralRule4E(float64(-1)), 0)
	t1.Eq(pluralRule4E(float64(1)), 0)

	// second form
	t1.Eq(pluralRule4E(float64(-2)), 1)
	t1.Eq(pluralRule4E(float64(0)), 1)
	t1.Eq(pluralRule4E(float64(2)), 1)
	t1.Eq(pluralRule4E(float64(10)), 1)
	t1.Eq(pluralRule4E(float64(102)), 1)
	t1.Eq(pluralRule4E(float64(110)), 1)

	// third form
	t1.Eq(pluralRule4E(float64(-11)), 2)
	t1.Eq(pluralRule4E(float64(11)), 2)
	t1.Eq(pluralRule4E(float64(19)), 2)
	t1.Eq(pluralRule4E(float64(111)), 2)
	t1.Eq(pluralRule4E(float64(119)), 2)

	// fourth form
	t1.Eq(pluralRule4E(float64(0.5)), 3)
	t1.Eq(pluralRule4E(float64(20)), 3)
	t1.Eq(pluralRule4E(float64(21)), 3)
	t1.Eq(pluralRule4E(float64(22)), 3)
	t1.Eq(pluralRule4E(float64(29)), 3)

	//_________________	pluralRule4F

	// first form
	t1.Eq(pluralRule4F(float64(-1)), 0)
	t1.Eq(pluralRule4F(float64(1)), 0)
	t1.Eq(pluralRule4F(float64(11)), 0)

	// second form
	t1.Eq(pluralRule4F(float64(-2)), 1)
	t1.Eq(pluralRule4F(float64(2)), 1)
	t1.Eq(pluralRule4F(float64(12)), 1)

	// third form
	t1.Eq(pluralRule4F(float64(-3)), 2)
	t1.Eq(pluralRule4F(float64(3)), 2)
	t1.Eq(pluralRule4F(float64(10)), 2)
	t1.Eq(pluralRule4F(float64(13)), 2)
	t1.Eq(pluralRule4F(float64(19)), 2)

	// fourth form
	t1.Eq(pluralRule4F(float64(0)), 3)
	t1.Eq(pluralRule4F(float64(0.5)), 3)
	t1.Eq(pluralRule4F(float64(20)), 3)
	t1.Eq(pluralRule4F(float64(21)), 3)
	t1.Eq(pluralRule4F(float64(22)), 3)
	t1.Eq(pluralRule4F(float64(23)), 3)
	t1.Eq(pluralRule4F(float64(29)), 3)
	t1.Eq(pluralRule4F(float64(101)), 3)
	t1.Eq(pluralRule4F(float64(102)), 3)
	t1.Eq(pluralRule4F(float64(103)), 3)
	t1.Eq(pluralRule4F(float64(109)), 3)

	//_________________	pluralRule5A

	// first form
	t1.Eq(pluralRule5A(float64(-1)), 0)
	t1.Eq(pluralRule5A(float64(1)), 0)

	// second form
	t1.Eq(pluralRule5A(float64(-2)), 1)
	t1.Eq(pluralRule5A(float64(2)), 1)

	// third form
	t1.Eq(pluralRule5A(float64(-3)), 2)
	t1.Eq(pluralRule5A(float64(3)), 2)
	t1.Eq(pluralRule5A(float64(4)), 2)
	t1.Eq(pluralRule5A(float64(5)), 2)
	t1.Eq(pluralRule5A(float64(6)), 2)

	// fourth form
	t1.Eq(pluralRule5A(float64(-7)), 3)
	t1.Eq(pluralRule5A(float64(7)), 3)
	t1.Eq(pluralRule5A(float64(8)), 3)
	t1.Eq(pluralRule5A(float64(9)), 3)
	t1.Eq(pluralRule5A(float64(10)), 3)

	// fifth form
	t1.Eq(pluralRule5A(float64(0)), 4)
	t1.Eq(pluralRule5A(float64(0.5)), 4)
	t1.Eq(pluralRule5A(float64(11)), 4)
	t1.Eq(pluralRule5A(float64(12)), 4)
	t1.Eq(pluralRule5A(float64(13)), 4)
	t1.Eq(pluralRule5A(float64(14)), 4)
	t1.Eq(pluralRule5A(float64(15)), 4)
	t1.Eq(pluralRule5A(float64(16)), 4)
	t1.Eq(pluralRule5A(float64(17)), 4)
	t1.Eq(pluralRule5A(float64(18)), 4)
	t1.Eq(pluralRule5A(float64(19)), 4)
	t1.Eq(pluralRule5A(float64(20)), 4)

	//_________________	pluralRule5B

	// first form
	t1.Eq(pluralRule5B(float64(-1)), 0)
	t1.Eq(pluralRule5B(float64(1)), 0)
	t1.Eq(pluralRule5B(float64(21)), 0)
	t1.Eq(pluralRule5B(float64(61)), 0)
	t1.Eq(pluralRule5B(float64(81)), 0)
	t1.Eq(pluralRule5B(float64(101)), 0)

	// second form
	t1.Eq(pluralRule5B(float64(-2)), 1)
	t1.Eq(pluralRule5B(float64(2)), 1)
	t1.Eq(pluralRule5B(float64(22)), 1)
	t1.Eq(pluralRule5B(float64(62)), 1)
	t1.Eq(pluralRule5B(float64(82)), 1)
	t1.Eq(pluralRule5B(float64(102)), 1)

	// third form
	t1.Eq(pluralRule5B(float64(-3)), 2)
	t1.Eq(pluralRule5B(float64(3)), 2)
	t1.Eq(pluralRule5B(float64(4)), 2)
	t1.Eq(pluralRule5B(float64(9)), 2)
	t1.Eq(pluralRule5B(float64(23)), 2)
	t1.Eq(pluralRule5B(float64(24)), 2)
	t1.Eq(pluralRule5B(float64(29)), 2)
	t1.Eq(pluralRule5B(float64(63)), 2)
	t1.Eq(pluralRule5B(float64(64)), 2)
	t1.Eq(pluralRule5B(float64(69)), 2)
	t1.Eq(pluralRule5B(float64(83)), 2)
	t1.Eq(pluralRule5B(float64(84)), 2)
	t1.Eq(pluralRule5B(float64(89)), 2)
	t1.Eq(pluralRule5B(float64(103)), 2)
	t1.Eq(pluralRule5B(float64(104)), 2)
	t1.Eq(pluralRule5B(float64(109)), 2)

	// fourth form
	t1.Eq(pluralRule5B(float64(-1000000)), 3)
	t1.Eq(pluralRule5B(float64(1000000)), 3)
	t1.Eq(pluralRule5B(float64(2000000)), 3)
	t1.Eq(pluralRule5B(float64(10000000)), 3)

	// fourth form
	t1.Eq(pluralRule5B(float64(0)), 4)
	t1.Eq(pluralRule5B(float64(0.5)), 4)
	t1.Eq(pluralRule5B(float64(10)), 4)
	t1.Eq(pluralRule5B(float64(11)), 4)
	t1.Eq(pluralRule5B(float64(12)), 4)
	t1.Eq(pluralRule5B(float64(13)), 4)
	t1.Eq(pluralRule5B(float64(14)), 4)
	t1.Eq(pluralRule5B(float64(19)), 4)
	t1.Eq(pluralRule5B(float64(20)), 4)
	t1.Eq(pluralRule5B(float64(71)), 4)
	t1.Eq(pluralRule5B(float64(72)), 4)
	t1.Eq(pluralRule5B(float64(73)), 4)
	t1.Eq(pluralRule5B(float64(74)), 4)
	t1.Eq(pluralRule5B(float64(79)), 4)
	t1.Eq(pluralRule5B(float64(91)), 4)
	t1.Eq(pluralRule5B(float64(92)), 4)
	t1.Eq(pluralRule5B(float64(93)), 4)
	t1.Eq(pluralRule5B(float64(94)), 4)
	t1.Eq(pluralRule5B(float64(99)), 4)
	t1.Eq(pluralRule5B(float64(100)), 4)
	t1.Eq(pluralRule5B(float64(1000)), 4)
	t1.Eq(pluralRule5B(float64(10000)), 4)
	t1.Eq(pluralRule5B(float64(100000)), 4)

	//_________________	pluralRule6A

	// first form
	t1.Eq(pluralRule6A(float64(0)), 0)

	// second form
	t1.Eq(pluralRule6A(float64(-1)), 1)
	t1.Eq(pluralRule6A(float64(1)), 1)

	// third form
	t1.Eq(pluralRule6A(float64(-2)), 2)
	t1.Eq(pluralRule6A(float64(2)), 2)

	// fourth form
	t1.Eq(pluralRule6A(float64(-3)), 3)
	t1.Eq(pluralRule6A(float64(3)), 3)
	t1.Eq(pluralRule6A(float64(4)), 3)
	t1.Eq(pluralRule6A(float64(9)), 3)
	t1.Eq(pluralRule6A(float64(10)), 3)
	t1.Eq(pluralRule6A(float64(103)), 3)
	t1.Eq(pluralRule6A(float64(109)), 3)
	t1.Eq(pluralRule6A(float64(110)), 3)

	// fifth form
	t1.Eq(pluralRule6A(float64(-11)), 4)
	t1.Eq(pluralRule6A(float64(11)), 4)
	t1.Eq(pluralRule6A(float64(12)), 4)
	t1.Eq(pluralRule6A(float64(98)), 4)
	t1.Eq(pluralRule6A(float64(99)), 4)
	t1.Eq(pluralRule6A(float64(111)), 4)
	t1.Eq(pluralRule6A(float64(112)), 4)
	t1.Eq(pluralRule6A(float64(198)), 4)
	t1.Eq(pluralRule6A(float64(199)), 4)

	// sixth form
	t1.Eq(pluralRule6A(float64(0.5)), 5)
	t1.Eq(pluralRule6A(float64(100)), 5)
	t1.Eq(pluralRule6A(float64(102)), 5)
	t1.Eq(pluralRule6A(float64(200)), 5)
	t1.Eq(pluralRule6A(float64(202)), 5)

	//_________________	pluralRule6B

	// first form
	t1.Eq(pluralRule6B(float64(0)), 0)

	// second form
	t1.Eq(pluralRule6B(float64(-1)), 1)
	t1.Eq(pluralRule6B(float64(1)), 1)

	// third form
	t1.Eq(pluralRule6B(float64(-2)), 2)
	t1.Eq(pluralRule6B(float64(2)), 2)

	// fourth form
	t1.Eq(pluralRule6B(float64(-3)), 3)
	t1.Eq(pluralRule6B(float64(3)), 3)

	// fifth form
	t1.Eq(pluralRule6B(float64(-6)), 4)
	t1.Eq(pluralRule6B(float64(6)), 4)

	// sixth form
	t1.Eq(pluralRule6B(float64(0.5)), 5)
	t1.Eq(pluralRule6B(float64(4)), 5)
	t1.Eq(pluralRule6B(float64(5)), 5)
	t1.Eq(pluralRule6B(float64(7)), 5)
	t1.Eq(pluralRule6B(float64(8)), 5)
	t1.Eq(pluralRule6B(float64(9)), 5)
	t1.Eq(pluralRule6B(float64(10)), 5)
	t1.Eq(pluralRule6B(float64(11)), 5)
	t1.Eq(pluralRule6B(float64(12)), 5)
	t1.Eq(pluralRule6B(float64(13)), 5)
	t1.Eq(pluralRule6B(float64(16)), 5)

	//_________________	pluralRuleRu

	// first form
	t1.Eq(pluralRuleRu(float64(-1)), 0)
	t1.Eq(pluralRuleRu(float64(1)), 0)
	t1.Eq(pluralRuleRu(float64(21)), 0)
	t1.Eq(pluralRuleRu(float64(31)), 0)
	t1.Eq(pluralRuleRu(float64(101)), 0)

	// second form
	t1.Eq(pluralRuleRu(float64(0)), 1)
	t1.Eq(pluralRuleRu(float64(-5)), 1)
	t1.Eq(pluralRuleRu(float64(6)), 1)
	t1.Eq(pluralRuleRu(float64(7)), 1)
	t1.Eq(pluralRuleRu(float64(11)), 1)

	// third form
	t1.Eq(pluralRuleRu(float64(-2)), 2)
	t1.Eq(pluralRuleRu(float64(22)), 2)
	t1.Eq(pluralRuleRu(float64(1.5)), 2)
	t1.Eq(pluralRuleRu(float64(101.1)), 2)

}
