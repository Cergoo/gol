package plural

import (
	"gol/test"
	"testing"
)

var t1 *test.TT

func Test_main(t *testing.T) {
	t1 = test.New(t)
	//_________________	isInt

	// positive test cases
	t1.Eq("test1_isInt", isInt(float64(0)), true)
	t1.Eq("test2_isInt", isInt(float64(1)), true)
	t1.Eq("test3_isInt", isInt(float64(0.0)), true)
	t1.Eq("test4_isInt", isInt(float64(1.0000)), true)
	t1.Eq("test5_isInt", isInt(float64(-50)), true)

	// negative
	t1.Eq("test6_isInt", isInt(float64(0.1)), false)
	t1.Eq("test7_isInt", isInt(float64(-0.1)), false)
	t1.Eq("test8_isInt", isInt(float64(0.00000000000001)), false)

	//_________________	pluralRule1

	t1.Eq("test1_pluralRule1", pluralRule1(float64(0)), 0)
	t1.Eq("test2_pluralRule1", pluralRule1(float64(0.5)), 0)
	t1.Eq("test3_pluralRule1", pluralRule1(float64(100)), 0)

	//_________________	pluralRule2A

	// first form
	t1.Eq("test1_pluralRule2A", pluralRule2A(float64(-1)), 0)
	t1.Eq("test2_pluralRule2A", pluralRule2A(float64(1)), 0)

	// second form
	t1.Eq("test3_pluralRule2A", pluralRule2A(float64(0)), 1)
	t1.Eq("test4_pluralRule2A", pluralRule2A(float64(0.5)), 1)
	t1.Eq("test5_pluralRule2A", pluralRule2A(float64(2)), 1)

	//_________________	pluralRule2B

	// first form
	t1.Eq("test1_pluralRule2B", pluralRule2B(float64(-1)), 0)
	t1.Eq("test2_pluralRule2B", pluralRule2B(float64(0)), 0)
	t1.Eq("test3_pluralRule2B", pluralRule2B(float64(1)), 0)

	// second form
	t1.Eq("test4_pluralRule2B", pluralRule2B(float64(0.5)), 1)
	t1.Eq("test5_pluralRule2B", pluralRule2B(float64(2)), 1)

	//_________________	pluralRule2C

	// first form
	t1.Eq("test1_pluralRule2C", pluralRule2C(float64(-1)), 0)
	t1.Eq("test2_pluralRule2C", pluralRule2C(float64(0)), 0)
	t1.Eq("test3_pluralRule2C", pluralRule2C(float64(0.5)), 0)
	t1.Eq("test4_pluralRule2C", pluralRule2C(float64(1)), 0)
	t1.Eq("test5_pluralRule2C", pluralRule2C(float64(1.5)), 0)

	// second form
	t1.Eq("test6_pluralRule2C", pluralRule2C(float64(2)), 1)
	t1.Eq("test7_pluralRule2C", pluralRule2C(float64(2.5)), 1)
	t1.Eq("test8_pluralRule2C", pluralRule2C(float64(100)), 1)

	//_________________	pluralRule2D

	// first form
	t1.Eq("test1_pluralRule2D", pluralRule2D(float64(-1)), 0)
	t1.Eq("test2_pluralRule2D", pluralRule2D(float64(1)), 0)
	t1.Eq("test3_pluralRule2D", pluralRule2D(float64(21)), 0)

	// second form
	t1.Eq("test4_pluralRule2D", pluralRule2D(float64(0)), 1)
	t1.Eq("test5_pluralRule2D", pluralRule2D(float64(0.5)), 1)
	t1.Eq("test6_pluralRule2D", pluralRule2D(float64(2)), 1)
	t1.Eq("test7_pluralRule2D", pluralRule2D(float64(11)), 1)

	//_________________	pluralRule2E

	// first form
	t1.Eq("test1_pluralRule2E", pluralRule2E(float64(-1)), 0)
	t1.Eq("test2_pluralRule2E", pluralRule2E(float64(0)), 0)
	t1.Eq("test3_pluralRule2E", pluralRule2E(float64(1)), 0)
	t1.Eq("test4_pluralRule2E", pluralRule2E(float64(11)), 0)
	t1.Eq("test5_pluralRule2E", pluralRule2E(float64(12)), 0)
	t1.Eq("test6_pluralRule2E", pluralRule2E(float64(98)), 0)
	t1.Eq("test7_pluralRule2E", pluralRule2E(float64(99)), 0)

	// second form
	t1.Eq("test8_pluralRule2E", pluralRule2E(float64(0.5)), 1)
	t1.Eq("test9_pluralRule2E", pluralRule2E(float64(2)), 1)
	t1.Eq("test10_pluralRule2E", pluralRule2E(float64(10)), 1)
	t1.Eq("test11_pluralRule2E", pluralRule2E(float64(100)), 1)

	//_________________	pluralRule2F

	// first form
	t1.Eq("test1_pluralRule2F", pluralRule2F(float64(-1)), 0)
	t1.Eq("test2_pluralRule2F", pluralRule2F(float64(0)), 0)
	t1.Eq("test3_pluralRule2F", pluralRule2F(float64(1)), 0)
	t1.Eq("test4_pluralRule2F", pluralRule2F(float64(2)), 0)
	t1.Eq("test5_pluralRule2F", pluralRule2F(float64(11)), 0)
	t1.Eq("test6_pluralRule2F", pluralRule2F(float64(12)), 0)
	t1.Eq("test7_pluralRule2F", pluralRule2F(float64(20)), 0)
	t1.Eq("test8_pluralRule2F", pluralRule2F(float64(40)), 0)

	// second form
	t1.Eq("test9_pluralRule2F", pluralRule2F(float64(0.5)), 1)
	t1.Eq("test10_pluralRule2F", pluralRule2F(float64(3)), 1)
	t1.Eq("test11_pluralRule2F", pluralRule2F(float64(10)), 1)

	//_________________	pluralRule3A

	// first form
	t1.Eq("test1_pluralRule3A", pluralRule3A(float64(0)), 0)

	// second form
	t1.Eq("test2_pluralRule3A", pluralRule3A(float64(-1)), 1)
	t1.Eq("test3_pluralRule3A", pluralRule3A(float64(1)), 1)
	t1.Eq("test4_pluralRule3A", pluralRule3A(float64(21)), 1)

	// third form
	t1.Eq("test5_pluralRule3A", pluralRule3A(float64(0.5)), 2)
	t1.Eq("test6_pluralRule3A", pluralRule3A(float64(2)), 2)
	t1.Eq("test7_pluralRule3A", pluralRule3A(float64(10)), 2)
	t1.Eq("test8_pluralRule3A", pluralRule3A(float64(11)), 2)

	//_________________	pluralRule3B

	// first form
	t1.Eq("test1_pluralRule3B", pluralRule3B(float64(-1)), 0)
	t1.Eq("test2_pluralRule3B", pluralRule3B(float64(1)), 0)

	// second form
	t1.Eq("test3_pluralRule3B", pluralRule3B(float64(-2)), 1)
	t1.Eq("test4_pluralRule3B", pluralRule3B(float64(2)), 1)

	// third form
	t1.Eq("test5_pluralRule3B", pluralRule3B(float64(0)), 2)
	t1.Eq("test6_pluralRule3B", pluralRule3B(float64(0.5)), 2)
	t1.Eq("test7_pluralRule3B", pluralRule3B(float64(3)), 2)
	t1.Eq("test8_pluralRule3B", pluralRule3B(float64(11)), 2)

	//_________________	pluralRule3C

	// first form
	t1.Eq("test1_pluralRule3C", pluralRule3C(float64(-1)), 0)
	t1.Eq("test2_pluralRule3C", pluralRule3C(float64(1)), 0)

	// second form
	t1.Eq("test3_pluralRule3C", pluralRule3C(float64(0)), 1)
	t1.Eq("test4_pluralRule3C", pluralRule3C(float64(-11)), 1)
	t1.Eq("test5_pluralRule3C", pluralRule3C(float64(11)), 1)
	t1.Eq("test6_pluralRule3C", pluralRule3C(float64(19)), 1)
	t1.Eq("test7_pluralRule3C", pluralRule3C(float64(111)), 1)
	t1.Eq("test8_pluralRule3C", pluralRule3C(float64(119)), 1)

	// third form
	t1.Eq("test9_pluralRule3C", pluralRule3C(float64(0.5)), 2)
	t1.Eq("test10_pluralRule3C", pluralRule3C(float64(20)), 2)
	t1.Eq("test11_pluralRule3C", pluralRule3C(float64(21)), 2)

	//_________________	pluralRule3D

	// first form
	t1.Eq("test1_pluralRule3D", pluralRule3D(float64(-1)), 0)
	t1.Eq("test2_pluralRule3D", pluralRule3D(float64(1)), 0)
	t1.Eq("test3_pluralRule3D", pluralRule3D(float64(21)), 0)

	// second form
	t1.Eq("test4_pluralRule3D", pluralRule3D(float64(-2)), 1)
	t1.Eq("test5_pluralRule3D", pluralRule3D(float64(2)), 1)
	t1.Eq("test6_pluralRule3D", pluralRule3D(float64(9)), 1)
	t1.Eq("test7_pluralRule3D", pluralRule3D(float64(22)), 1)
	t1.Eq("test8_pluralRule3D", pluralRule3D(float64(29)), 1)

	// third form
	t1.Eq("test9_pluralRule3D", pluralRule3D(float64(0)), 2)
	t1.Eq("test10_pluralRule3D", pluralRule3D(float64(0.5)), 2)
	t1.Eq("test11_pluralRule3D", pluralRule3D(float64(11)), 2)
	t1.Eq("test12_pluralRule3D", pluralRule3D(float64(19)), 2)

	//_________________	pluralRule3E

	// first form
	t1.Eq("test1_pluralRule3E", pluralRule3E(float64(-1)), 0)
	t1.Eq("test2_pluralRule3E", pluralRule3E(float64(1)), 0)

	// second form
	t1.Eq("test3_pluralRule3E", pluralRule3E(float64(-2)), 1)
	t1.Eq("test4_pluralRule3E", pluralRule3E(float64(2)), 1)
	t1.Eq("test5_pluralRule3E", pluralRule3E(float64(3)), 1)
	t1.Eq("test6_pluralRule3E", pluralRule3E(float64(4)), 1)

	// third form
	t1.Eq("test7_pluralRule3E", pluralRule3E(float64(0)), 2)
	t1.Eq("test8_pluralRule3E", pluralRule3E(float64(0.5)), 2)
	t1.Eq("test9_pluralRule3E", pluralRule3E(float64(5)), 2)
	t1.Eq("test10_pluralRule3E", pluralRule3E(float64(9)), 2)
	t1.Eq("test11_pluralRule3E", pluralRule3E(float64(11)), 2)
	t1.Eq("test12_pluralRule3E", pluralRule3E(float64(12)), 2)
	t1.Eq("test13_pluralRule3E", pluralRule3E(float64(14)), 2)

	//_________________	pluralRule3F

	// first form
	t1.Eq("test1_pluralRule3F", pluralRule3F(float64(0)), 0)

	// second form
	t1.Eq("test2_pluralRule3F", pluralRule3F(float64(-0.5)), 1)
	t1.Eq("test3_pluralRule3F", pluralRule3F(float64(0.5)), 1)
	t1.Eq("test4_pluralRule3F", pluralRule3F(float64(1)), 1)
	t1.Eq("test5_pluralRule3F", pluralRule3F(float64(1.5)), 1)

	// third form
	t1.Eq("test6_pluralRule3F", pluralRule3F(float64(-2)), 2)
	t1.Eq("test7_pluralRule3F", pluralRule3F(float64(2)), 2)
	t1.Eq("test8_pluralRule3F", pluralRule3F(float64(3)), 2)

	//_________________	pluralRule3G

	// first form
	t1.Eq("test1_pluralRule3G", pluralRule3G(float64(-0.5)), 0)
	t1.Eq("test2_pluralRule3G", pluralRule3G(float64(0)), 0)
	t1.Eq("test3_pluralRule3G", pluralRule3G(float64(0.5)), 0)
	t1.Eq("test4_pluralRule3G", pluralRule3G(float64(1)), 0)

	// second form
	t1.Eq("test5_pluralRule3G", pluralRule3G(float64(-2)), 1)
	t1.Eq("test6_pluralRule3G", pluralRule3G(float64(2)), 1)
	t1.Eq("test7_pluralRule3G", pluralRule3G(float64(3)), 1)
	t1.Eq("test8_pluralRule3G", pluralRule3G(float64(9)), 1)
	t1.Eq("test9_pluralRule3G", pluralRule3G(float64(10)), 1)

	// third
	t1.Eq("test10_pluralRule3G", pluralRule3G(float64(1.5)), 2)
	t1.Eq("test11_pluralRule3G", pluralRule3G(float64(11)), 2)
	t1.Eq("test12_pluralRule3G", pluralRule3G(float64(12)), 2)

	//_________________	pluralRule3H

	// first form
	t1.Eq("test1_pluralRule3H", pluralRule3H(float64(0)), 0)

	// second form
	t1.Eq("test2_pluralRule3H", pluralRule3H(float64(-1)), 1)
	t1.Eq("test3_pluralRule3H", pluralRule3H(float64(1)), 1)

	// third form
	t1.Eq("test3_pluralRule3H", pluralRule3H(float64(0.5)), 2)
	t1.Eq("test4_pluralRule3H", pluralRule3H(float64(1.5)), 2)
	t1.Eq("test5_pluralRule3H", pluralRule3H(float64(2)), 2)
	t1.Eq("test6_pluralRule3H", pluralRule3H(float64(10)), 2)
	t1.Eq("test7_pluralRule3H", pluralRule3H(float64(11)), 2)

	//_________________	pluralRule3I

	// first form
	t1.Eq("test1_pluralRule3I", pluralRule3I(float64(-1)), 0)
	t1.Eq("test2_pluralRule3I", pluralRule3I(float64(1)), 0)

	// second form
	t1.Eq("test3_pluralRule3I", pluralRule3I(float64(-2)), 1)
	t1.Eq("test4_pluralRule3I", pluralRule3I(float64(2)), 1)
	t1.Eq("test5_pluralRule3I", pluralRule3I(float64(3)), 1)
	t1.Eq("test6_pluralRule3I", pluralRule3I(float64(4)), 1)
	t1.Eq("test7_pluralRule3I", pluralRule3I(float64(22)), 1)
	t1.Eq("test8_pluralRule3I", pluralRule3I(float64(23)), 1)
	t1.Eq("test9_pluralRule3I", pluralRule3I(float64(24)), 1)

	// third form
	t1.Eq("test10_pluralRule3I", pluralRule3I(float64(0)), 2)
	t1.Eq("test11_pluralRule3I", pluralRule3I(float64(0.5)), 2)
	t1.Eq("test12_pluralRule3I", pluralRule3I(float64(5)), 2)
	t1.Eq("test13_pluralRule3I", pluralRule3I(float64(9)), 2)
	t1.Eq("test14_pluralRule3I", pluralRule3I(float64(12)), 2)
	t1.Eq("test15_pluralRule3I", pluralRule3I(float64(13)), 2)
	t1.Eq("test16_pluralRule3I", pluralRule3I(float64(14)), 2)
	t1.Eq("test17_pluralRule3I", pluralRule3I(float64(15)), 2)

	//_________________	pluralRule4A

	// first form
	t1.Eq("test1_pluralRule4A", pluralRule4A(float64(-1)), 0)
	t1.Eq("test2_pluralRule4A", pluralRule4A(float64(1)), 0)

	// second form
	t1.Eq("test3_pluralRule4A", pluralRule4A(float64(-2)), 1)
	t1.Eq("test4_pluralRule4A", pluralRule4A(float64(2)), 1)

	// third form
	t1.Eq("test5_pluralRule4A", pluralRule4A(float64(-10)), 2)
	t1.Eq("test6_pluralRule4A", pluralRule4A(float64(10)), 2)
	t1.Eq("test7_pluralRule4A", pluralRule4A(float64(20)), 2)
	t1.Eq("test8_pluralRule4A", pluralRule4A(float64(100)), 2)

	// fourth form
	t1.Eq("test9_pluralRule4A", pluralRule4A(float64(0)), 3)
	t1.Eq("test10_pluralRule4A", pluralRule4A(float64(0.5)), 3)
	t1.Eq("test11_pluralRule4A", pluralRule4A(float64(3)), 3)
	t1.Eq("test12_pluralRule4A", pluralRule4A(float64(9)), 3)
	t1.Eq("test13_pluralRule4A", pluralRule4A(float64(11)), 3)

	//_________________	pluralRule4B

	// first form
	t1.Eq("test1_pluralRule4B", pluralRule4B(float64(-1)), 0)
	t1.Eq("test2_pluralRule4B", pluralRule4B(float64(1)), 0)
	t1.Eq("test3_pluralRule4B", pluralRule4B(float64(21)), 0)

	// second form
	t1.Eq("test4_pluralRule4B", pluralRule4B(float64(-2)), 1)
	t1.Eq("test5_pluralRule4B", pluralRule4B(float64(2)), 1)
	t1.Eq("test6_pluralRule4B", pluralRule4B(float64(3)), 1)
	t1.Eq("test7_pluralRule4B", pluralRule4B(float64(4)), 1)
	t1.Eq("test8_pluralRule4B", pluralRule4B(float64(22)), 1)
	t1.Eq("test9_pluralRule4B", pluralRule4B(float64(23)), 1)
	t1.Eq("test10_pluralRule4B", pluralRule4B(float64(24)), 1)

	// third form
	t1.Eq("test11_pluralRule4B", pluralRule4B(float64(-5)), 2)
	t1.Eq("test12_pluralRule4B", pluralRule4B(float64(0)), 2)
	t1.Eq("test13_pluralRule4B", pluralRule4B(float64(5)), 2)
	t1.Eq("test14_pluralRule4B", pluralRule4B(float64(6)), 2)
	t1.Eq("test15_pluralRule4B", pluralRule4B(float64(8)), 2)
	t1.Eq("test16_pluralRule4B", pluralRule4B(float64(9)), 2)
	t1.Eq("test17_pluralRule4B", pluralRule4B(float64(11)), 2)
	t1.Eq("test18_pluralRule4B", pluralRule4B(float64(12)), 2)
	t1.Eq("test19_pluralRule4B", pluralRule4B(float64(13)), 2)
	t1.Eq("test20_pluralRule4B", pluralRule4B(float64(14)), 2)

	// fourth form
	t1.Eq("test21_pluralRule4B", pluralRule4B(float64(0.5)), 3)
	t1.Eq("test22_pluralRule4B", pluralRule4B(float64(1.5)), 3)

	//_________________	pluralRule4C

	// first form
	t1.Eq("test1_pluralRule4C", pluralRule4C(float64(-1)), 0)
	t1.Eq("test2_pluralRule4C", pluralRule4C(float64(1)), 0)

	// second form
	t1.Eq("test3_pluralRule4C", pluralRule4C(float64(-2)), 1)
	t1.Eq("test4_pluralRule4C", pluralRule4C(float64(2)), 1)
	t1.Eq("test5_pluralRule4C", pluralRule4C(float64(3)), 1)
	t1.Eq("test6_pluralRule4C", pluralRule4C(float64(4)), 1)
	t1.Eq("test7_pluralRule4C", pluralRule4C(float64(22)), 1)
	t1.Eq("test8_pluralRule4C", pluralRule4C(float64(23)), 1)
	t1.Eq("test9_pluralRule4C", pluralRule4C(float64(24)), 1)

	// third form
	t1.Eq("test10_pluralRule4C", pluralRule4C(float64(-10)), 2)
	t1.Eq("test11_pluralRule4C", pluralRule4C(float64(10)), 2)
	t1.Eq("test12_pluralRule4C", pluralRule4C(float64(11)), 2)
	t1.Eq("test13_pluralRule4C", pluralRule4C(float64(12)), 2)
	t1.Eq("test14_pluralRule4C", pluralRule4C(float64(13)), 2)
	t1.Eq("test15_pluralRule4C", pluralRule4C(float64(14)), 2)
	t1.Eq("test16_pluralRule4C", pluralRule4C(float64(15)), 2)
	t1.Eq("test17_pluralRule4C", pluralRule4C(float64(16)), 2)
	t1.Eq("test18_pluralRule4C", pluralRule4C(float64(17)), 2)
	t1.Eq("test19_pluralRule4C", pluralRule4C(float64(18)), 2)
	t1.Eq("test20_pluralRule4C", pluralRule4C(float64(19)), 2)
	t1.Eq("test21_pluralRule4C", pluralRule4C(float64(20)), 2)
	t1.Eq("test22_pluralRule4C", pluralRule4C(float64(21)), 2)
	t1.Eq("test24_pluralRule4C", pluralRule4C(float64(25)), 2)
	t1.Eq("test25_pluralRule4C", pluralRule4C(float64(26)), 2)
	t1.Eq("test26_pluralRule4C", pluralRule4C(float64(28)), 2)
	t1.Eq("test27_pluralRule4C", pluralRule4C(float64(29)), 2)

	// fourth form
	t1.Eq("test28_pluralRule4C", pluralRule4C(float64(0.5)), 3)
	t1.Eq("test29_pluralRule4C", pluralRule4C(float64(1.5)), 3)

	//_________________	pluralRule4D

	// first form
	t1.Eq("test1_pluralRule4D", pluralRule4D(float64(-1)), 0)
	t1.Eq("test2_pluralRule4D", pluralRule4D(float64(1)), 0)
	t1.Eq("test3_pluralRule4D", pluralRule4D(float64(101)), 0)

	// second form
	t1.Eq("test4_pluralRule4D", pluralRule4D(float64(-2)), 1)
	t1.Eq("test5_pluralRule4D", pluralRule4D(float64(2)), 1)
	t1.Eq("test6_pluralRule4D", pluralRule4D(float64(102)), 1)

	// third form
	t1.Eq("test7_pluralRule4D", pluralRule4D(float64(-3)), 2)
	t1.Eq("test8_pluralRule4D", pluralRule4D(float64(3)), 2)
	t1.Eq("test9_pluralRule4D", pluralRule4D(float64(4)), 2)
	t1.Eq("test10_pluralRule4D", pluralRule4D(float64(103)), 2)
	t1.Eq("test11_pluralRule4D", pluralRule4D(float64(104)), 2)

	// fourth form
	t1.Eq("test12_pluralRule4D", pluralRule4D(float64(0)), 3)
	t1.Eq("test13_pluralRule4D", pluralRule4D(float64(0.5)), 3)
	t1.Eq("test14_pluralRule4D", pluralRule4D(float64(5)), 3)
	t1.Eq("test15_pluralRule4D", pluralRule4D(float64(10)), 3)
	t1.Eq("test16_pluralRule4D", pluralRule4D(float64(11)), 3)
	t1.Eq("test17_pluralRule4D", pluralRule4D(float64(12)), 3)
	t1.Eq("test18_pluralRule4D", pluralRule4D(float64(13)), 3)
	t1.Eq("test19_pluralRule4D", pluralRule4D(float64(14)), 3)

	//_________________	pluralRule4E

	// first form
	t1.Eq("test1_pluralRule4E", pluralRule4E(float64(-1)), 0)
	t1.Eq("test2_pluralRule4E", pluralRule4E(float64(1)), 0)

	// second form
	t1.Eq("test3_pluralRule4E", pluralRule4E(float64(-2)), 1)
	t1.Eq("test4_pluralRule4E", pluralRule4E(float64(0)), 1)
	t1.Eq("test5_pluralRule4E", pluralRule4E(float64(2)), 1)
	t1.Eq("test6_pluralRule4E", pluralRule4E(float64(10)), 1)
	t1.Eq("test7_pluralRule4E", pluralRule4E(float64(102)), 1)
	t1.Eq("test8_pluralRule4E", pluralRule4E(float64(110)), 1)

	// third form
	t1.Eq("test9_pluralRule4E", pluralRule4E(float64(-11)), 2)
	t1.Eq("test10_pluralRule4E", pluralRule4E(float64(11)), 2)
	t1.Eq("test11_pluralRule4E", pluralRule4E(float64(19)), 2)
	t1.Eq("test12_pluralRule4E", pluralRule4E(float64(111)), 2)
	t1.Eq("test13_pluralRule4E", pluralRule4E(float64(119)), 2)

	// fourth form
	t1.Eq("test14_pluralRule4E", pluralRule4E(float64(0.5)), 3)
	t1.Eq("test15_pluralRule4E", pluralRule4E(float64(20)), 3)
	t1.Eq("test16_pluralRule4E", pluralRule4E(float64(21)), 3)
	t1.Eq("test17_pluralRule4E", pluralRule4E(float64(22)), 3)
	t1.Eq("test18_pluralRule4E", pluralRule4E(float64(29)), 3)

	//_________________	pluralRule4F

	// first form
	t1.Eq("test1_pluralRule4F", pluralRule4F(float64(-1)), 0)
	t1.Eq("test2_pluralRule4F", pluralRule4F(float64(1)), 0)
	t1.Eq("test3_pluralRule4F", pluralRule4F(float64(11)), 0)

	// second form
	t1.Eq("test4_pluralRule4F", pluralRule4F(float64(-2)), 1)
	t1.Eq("test5_pluralRule4F", pluralRule4F(float64(2)), 1)
	t1.Eq("test6_pluralRule4F", pluralRule4F(float64(12)), 1)

	// third form
	t1.Eq("test7_pluralRule4F", pluralRule4F(float64(-3)), 2)
	t1.Eq("test8_pluralRule4F", pluralRule4F(float64(3)), 2)
	t1.Eq("test9_pluralRule4F", pluralRule4F(float64(10)), 2)
	t1.Eq("test10_pluralRule4F", pluralRule4F(float64(13)), 2)
	t1.Eq("test11_pluralRule4F", pluralRule4F(float64(19)), 2)

	// fourth form
	t1.Eq("test12_pluralRule4F", pluralRule4F(float64(0)), 3)
	t1.Eq("test13_pluralRule4F", pluralRule4F(float64(0.5)), 3)
	t1.Eq("test14_pluralRule4F", pluralRule4F(float64(20)), 3)
	t1.Eq("test15_pluralRule4F", pluralRule4F(float64(21)), 3)
	t1.Eq("test16_pluralRule4F", pluralRule4F(float64(22)), 3)
	t1.Eq("test17_pluralRule4F", pluralRule4F(float64(23)), 3)
	t1.Eq("test18_pluralRule4F", pluralRule4F(float64(29)), 3)
	t1.Eq("test19_pluralRule4F", pluralRule4F(float64(101)), 3)
	t1.Eq("test20_pluralRule4F", pluralRule4F(float64(102)), 3)
	t1.Eq("test21_pluralRule4F", pluralRule4F(float64(103)), 3)
	t1.Eq("test22_pluralRule4F", pluralRule4F(float64(109)), 3)

	//_________________	pluralRule5A

	// first form
	t1.Eq("test1_pluralRule5A", pluralRule5A(float64(-1)), 0)
	t1.Eq("test2_pluralRule5A", pluralRule5A(float64(1)), 0)

	// second form
	t1.Eq("test3_pluralRule5A", pluralRule5A(float64(-2)), 1)
	t1.Eq("test4_pluralRule5A", pluralRule5A(float64(2)), 1)

	// third form
	t1.Eq("test5_pluralRule5A", pluralRule5A(float64(-3)), 2)
	t1.Eq("test6_pluralRule5A", pluralRule5A(float64(3)), 2)
	t1.Eq("test7_pluralRule5A", pluralRule5A(float64(4)), 2)
	t1.Eq("test8_pluralRule5A", pluralRule5A(float64(5)), 2)
	t1.Eq("test9_pluralRule5A", pluralRule5A(float64(6)), 2)

	// fourth form
	t1.Eq("test10_pluralRule5A", pluralRule5A(float64(-7)), 3)
	t1.Eq("test11_pluralRule5A", pluralRule5A(float64(7)), 3)
	t1.Eq("test12_pluralRule5A", pluralRule5A(float64(8)), 3)
	t1.Eq("test13_pluralRule5A", pluralRule5A(float64(9)), 3)
	t1.Eq("test14_pluralRule5A", pluralRule5A(float64(10)), 3)

	// fifth form
	t1.Eq("test15_pluralRule5A", pluralRule5A(float64(0)), 4)
	t1.Eq("test16_pluralRule5A", pluralRule5A(float64(0.5)), 4)
	t1.Eq("test17_pluralRule5A", pluralRule5A(float64(11)), 4)
	t1.Eq("test18_pluralRule5A", pluralRule5A(float64(12)), 4)
	t1.Eq("test19_pluralRule5A", pluralRule5A(float64(13)), 4)
	t1.Eq("test20_pluralRule5A", pluralRule5A(float64(14)), 4)
	t1.Eq("test21_pluralRule5A", pluralRule5A(float64(15)), 4)
	t1.Eq("test22_pluralRule5A", pluralRule5A(float64(16)), 4)
	t1.Eq("test23_pluralRule5A", pluralRule5A(float64(17)), 4)
	t1.Eq("test24_pluralRule5A", pluralRule5A(float64(18)), 4)
	t1.Eq("test25_pluralRule5A", pluralRule5A(float64(19)), 4)
	t1.Eq("test26_pluralRule5A", pluralRule5A(float64(20)), 4)

	//_________________	pluralRule5B

	// first form
	t1.Eq("test1_pluralRule5B", pluralRule5B(float64(-1)), 0)
	t1.Eq("test2_pluralRule5B", pluralRule5B(float64(1)), 0)
	t1.Eq("test3_pluralRule5B", pluralRule5B(float64(21)), 0)
	t1.Eq("test4_pluralRule5B", pluralRule5B(float64(61)), 0)
	t1.Eq("test5_pluralRule5B", pluralRule5B(float64(81)), 0)
	t1.Eq("test6_pluralRule5B", pluralRule5B(float64(101)), 0)

	// second form
	t1.Eq("test7_pluralRule5B", pluralRule5B(float64(-2)), 1)
	t1.Eq("test8_pluralRule5B", pluralRule5B(float64(2)), 1)
	t1.Eq("test9_pluralRule5B", pluralRule5B(float64(22)), 1)
	t1.Eq("test10_pluralRule5B", pluralRule5B(float64(62)), 1)
	t1.Eq("test11_pluralRule5B", pluralRule5B(float64(82)), 1)
	t1.Eq("test12_pluralRule5B", pluralRule5B(float64(102)), 1)

	// third form
	t1.Eq("test13_pluralRule5B", pluralRule5B(float64(-3)), 2)
	t1.Eq("test14_pluralRule5B", pluralRule5B(float64(3)), 2)
	t1.Eq("test15_pluralRule5B", pluralRule5B(float64(4)), 2)
	t1.Eq("test16_pluralRule5B", pluralRule5B(float64(9)), 2)
	t1.Eq("test17_pluralRule5B", pluralRule5B(float64(23)), 2)
	t1.Eq("test18_pluralRule5B", pluralRule5B(float64(24)), 2)
	t1.Eq("test19_pluralRule5B", pluralRule5B(float64(29)), 2)
	t1.Eq("test20_pluralRule5B", pluralRule5B(float64(63)), 2)
	t1.Eq("test21_pluralRule5B", pluralRule5B(float64(64)), 2)
	t1.Eq("test22_pluralRule5B", pluralRule5B(float64(69)), 2)
	t1.Eq("test23_pluralRule5B", pluralRule5B(float64(83)), 2)
	t1.Eq("test24_pluralRule5B", pluralRule5B(float64(84)), 2)
	t1.Eq("test25_pluralRule5B", pluralRule5B(float64(89)), 2)
	t1.Eq("test26_pluralRule5B", pluralRule5B(float64(103)), 2)
	t1.Eq("test27_pluralRule5B", pluralRule5B(float64(104)), 2)
	t1.Eq("test28_pluralRule5B", pluralRule5B(float64(109)), 2)

	// fourth form
	t1.Eq("test29_pluralRule5B", pluralRule5B(float64(-1000000)), 3)
	t1.Eq("test30_pluralRule5B", pluralRule5B(float64(1000000)), 3)
	t1.Eq("test31_pluralRule5B", pluralRule5B(float64(2000000)), 3)
	t1.Eq("test32_pluralRule5B", pluralRule5B(float64(10000000)), 3)

	// fourth form
	t1.Eq("test33_pluralRule5B", pluralRule5B(float64(0)), 4)
	t1.Eq("test34_pluralRule5B", pluralRule5B(float64(0.5)), 4)
	t1.Eq("test35_pluralRule5B", pluralRule5B(float64(10)), 4)
	t1.Eq("test36_pluralRule5B", pluralRule5B(float64(11)), 4)
	t1.Eq("test37_pluralRule5B", pluralRule5B(float64(12)), 4)
	t1.Eq("test38_pluralRule5B", pluralRule5B(float64(13)), 4)
	t1.Eq("test39_pluralRule5B", pluralRule5B(float64(14)), 4)
	t1.Eq("test40_pluralRule5B", pluralRule5B(float64(19)), 4)
	t1.Eq("test41_pluralRule5B", pluralRule5B(float64(20)), 4)
	t1.Eq("test42_pluralRule5B", pluralRule5B(float64(71)), 4)
	t1.Eq("test43_pluralRule5B", pluralRule5B(float64(72)), 4)
	t1.Eq("test44_pluralRule5B", pluralRule5B(float64(73)), 4)
	t1.Eq("test45_pluralRule5B", pluralRule5B(float64(74)), 4)
	t1.Eq("test46_pluralRule5B", pluralRule5B(float64(79)), 4)
	t1.Eq("test47_pluralRule5B", pluralRule5B(float64(91)), 4)
	t1.Eq("test48_pluralRule5B", pluralRule5B(float64(92)), 4)
	t1.Eq("test49_pluralRule5B", pluralRule5B(float64(93)), 4)
	t1.Eq("test50_pluralRule5B", pluralRule5B(float64(94)), 4)
	t1.Eq("test51_pluralRule5B", pluralRule5B(float64(99)), 4)
	t1.Eq("test52_pluralRule5B", pluralRule5B(float64(100)), 4)
	t1.Eq("test53_pluralRule5B", pluralRule5B(float64(1000)), 4)
	t1.Eq("test54_pluralRule5B", pluralRule5B(float64(10000)), 4)
	t1.Eq("test55_pluralRule5B", pluralRule5B(float64(100000)), 4)

	//_________________	pluralRule6A

	// first form
	t1.Eq("test1_pluralRule6A", pluralRule6A(float64(0)), 0)

	// second form
	t1.Eq("test2_pluralRule6A", pluralRule6A(float64(-1)), 1)
	t1.Eq("test3_pluralRule6A", pluralRule6A(float64(1)), 1)

	// third form
	t1.Eq("test4_pluralRule6A", pluralRule6A(float64(-2)), 2)
	t1.Eq("test5_pluralRule6A", pluralRule6A(float64(2)), 2)

	// fourth form
	t1.Eq("test6_pluralRule6A", pluralRule6A(float64(-3)), 3)
	t1.Eq("test7_pluralRule6A", pluralRule6A(float64(3)), 3)
	t1.Eq("test8_pluralRule6A", pluralRule6A(float64(4)), 3)
	t1.Eq("test9_pluralRule6A", pluralRule6A(float64(9)), 3)
	t1.Eq("test10_pluralRule6A", pluralRule6A(float64(10)), 3)
	t1.Eq("test11_pluralRule6A", pluralRule6A(float64(103)), 3)
	t1.Eq("test12_pluralRule6A", pluralRule6A(float64(109)), 3)
	t1.Eq("test13_pluralRule6A", pluralRule6A(float64(110)), 3)

	// fifth form
	t1.Eq("test14_pluralRule6A", pluralRule6A(float64(-11)), 4)
	t1.Eq("test15_pluralRule6A", pluralRule6A(float64(11)), 4)
	t1.Eq("test16_pluralRule6A", pluralRule6A(float64(12)), 4)
	t1.Eq("test17_pluralRule6A", pluralRule6A(float64(98)), 4)
	t1.Eq("test18_pluralRule6A", pluralRule6A(float64(99)), 4)
	t1.Eq("test19_pluralRule6A", pluralRule6A(float64(111)), 4)
	t1.Eq("test20_pluralRule6A", pluralRule6A(float64(112)), 4)
	t1.Eq("test21_pluralRule6A", pluralRule6A(float64(198)), 4)
	t1.Eq("test22_pluralRule6A", pluralRule6A(float64(199)), 4)

	// sixth form
	t1.Eq("test23_pluralRule6A", pluralRule6A(float64(0.5)), 5)
	t1.Eq("test24_pluralRule6A", pluralRule6A(float64(100)), 5)
	t1.Eq("test25_pluralRule6A", pluralRule6A(float64(102)), 5)
	t1.Eq("test26_pluralRule6A", pluralRule6A(float64(200)), 5)
	t1.Eq("test27_pluralRule6A", pluralRule6A(float64(202)), 5)

	//_________________	pluralRule6B

	// first form
	t1.Eq("test1_pluralRule6B", pluralRule6B(float64(0)), 0)

	// second form
	t1.Eq("test2_pluralRule6B", pluralRule6B(float64(-1)), 1)
	t1.Eq("test3_pluralRule6B", pluralRule6B(float64(1)), 1)

	// third form
	t1.Eq("test4_pluralRule6B", pluralRule6B(float64(-2)), 2)
	t1.Eq("test5_pluralRule6B", pluralRule6B(float64(2)), 2)

	// fourth form
	t1.Eq("test6_pluralRule6B", pluralRule6B(float64(-3)), 3)
	t1.Eq("test7_pluralRule6B", pluralRule6B(float64(3)), 3)

	// fifth form
	t1.Eq("test8_pluralRule6B", pluralRule6B(float64(-6)), 4)
	t1.Eq("test9_pluralRule6B", pluralRule6B(float64(6)), 4)

	// sixth form
	t1.Eq("test10_pluralRule6B", pluralRule6B(float64(0.5)), 5)
	t1.Eq("test11_pluralRule6B", pluralRule6B(float64(4)), 5)
	t1.Eq("test12_pluralRule6B", pluralRule6B(float64(5)), 5)
	t1.Eq("test13_pluralRule6B", pluralRule6B(float64(7)), 5)
	t1.Eq("test14_pluralRule6B", pluralRule6B(float64(8)), 5)
	t1.Eq("test15_pluralRule6B", pluralRule6B(float64(9)), 5)
	t1.Eq("test16_pluralRule6B", pluralRule6B(float64(10)), 5)
	t1.Eq("test17_pluralRule6B", pluralRule6B(float64(11)), 5)
	t1.Eq("test18_pluralRule6B", pluralRule6B(float64(12)), 5)
	t1.Eq("test19_pluralRule6B", pluralRule6B(float64(13)), 5)
	t1.Eq("test20_pluralRule6B", pluralRule6B(float64(16)), 5)

	//_________________	pluralRuleRu

	// first form
	t1.Eq("test1_pluralRuleRu", pluralRuleRu(float64(-1)), 0)
	t1.Eq("test2_pluralRuleRu", pluralRuleRu(float64(1)), 0)
	t1.Eq("test3_pluralRuleRu", pluralRuleRu(float64(21)), 0)
	t1.Eq("test4_pluralRuleRu", pluralRuleRu(float64(31)), 0)
	t1.Eq("test5_pluralRuleRu", pluralRuleRu(float64(101)), 0)

	// second form
	t1.Eq("test6_pluralRuleRu", pluralRuleRu(float64(0)), 1)
	t1.Eq("test7_pluralRuleRu", pluralRuleRu(float64(-5)), 1)
	t1.Eq("test8_pluralRuleRu", pluralRuleRu(float64(6)), 1)
	t1.Eq("test9_pluralRuleRu", pluralRuleRu(float64(7)), 1)
	t1.Eq("test10_pluralRuleRu", pluralRuleRu(float64(11)), 1)

	// third form
	t1.Eq("test11_pluralRuleRu", pluralRuleRu(float64(-2)), 2)
	t1.Eq("test12_pluralRuleRu", pluralRuleRu(float64(22)), 2)
	t1.Eq("test13_pluralRuleRu", pluralRuleRu(float64(1.5)), 2)
	t1.Eq("test14_pluralRuleRu", pluralRuleRu(float64(101.1)), 2)

}
