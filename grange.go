/*  Grange - simple go library for pythonic range
    Copyright (C) 2020 Stanislav Arnaudov

This program is free software; you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation; either version 2 of the License, or
(at your option) any prior version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License along
with this program; if not, write to the Free Software Foundation, Inc.,
51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA. */

package grange

func grange(nums ...int) []int {

	if len(nums) == 1 {
		return make([]int, nums[0])
	}

	if len(nums) == 2 {

		if nums[0] >= nums[1] {
			return make([]int, 0)
		}

		list := make([]int, 0)

		for i := nums[0]; i < nums[1]; i++ {
			list = append(list, i)
		}
		return list
	}

	if len(nums) == 3 {

		if nums[0] >= nums[1] {
			return make([]int, 0)
		}

		list := make([]int, 0)

		for i := nums[0]; i < nums[1]; i += nums[2] {
			list = append(list, i)
		}
		return list
	}

	return make([]int, 0)
}
