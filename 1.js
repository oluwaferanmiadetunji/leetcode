// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */

var twoSum = function (nums, target) {
	// create a map object
	let numberIndex = new Map();
	// initialize the result array
	let result = [];

	for (let i = 0; i < nums.length; i++) {
		// current number
		let num = nums[i];
		// difference
		let difference = target - num;

		// check if diiference is in map
		if (numberIndex.has(difference)) {
			result[0] = numberIndex.get(difference);
			result[1] = i;

			// return result
			return result;
		}

		numberIndex.set(num, i);
	}

	return result;
};
