package lt278

import "sort"

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */
//二分查找
func firstBadVersion(n int) int {
    return sort.Search(n, func(version int) bool { return isBadVersion(version) })
}
func isBadVersion(version int) bool