// Given a 32-bit signed integer, reverse digits of an integer

/**
 * @param {number} x
 * @return {number}
 */
var reverse = function (x) {
	const reversedNum = parseInt(x.toString().split('').reverse().join('')) * Math.sign(x);
	if (reversedNum >= Math.pow(2, -31) * -1 && reversedNum <= Math.pow(2, 31) - 1) {
		return reversedNum;
	} else {
		return 0;
	}
};
